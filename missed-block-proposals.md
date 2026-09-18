# `panoptichain_heimdall_missed_block_proposal`: how it works, and where it breaks

Notes from an 11-hour investigation on 2026-09-17, kept so the reasoning does
not have to be re-derived. Two bugs were found and fixed; the main one is still
open.

---

## What the metric actually counts

For each Heimdall block `h`, `refreshMissedBlockProposal` fetches the validator
set at `h-1`, sorts it by `proposer_priority` descending, and counts **every
validator ranked above the block's actual proposer** as having missed.

```go
proposer := block.ProposerAddress()
validators := <set at h-1>, sorted by proposer_priority desc
for _, v := range validators {
    if v.Address == proposer { break }
    proposers = append(proposers, v.Address)   // "missed"
}
```

So the value per block is **the proposer's index in that sorted list**. Zero is
healthy.

### The assumption, and why it is only approximately true

CometBFT does not pick the highest current priority. It **increments every
validator's priority by its voting power, then takes the max**, then subtracts
total voting power from the winner. The real rule is:

```
proposer(h) = argmax( priority(h-1) + voting_power )
```

The code omits `+ voting_power`. That works while priorities are spread wide
relative to voting power, and silently stops working when they are not.

---

## The 2026-09-17 Amoy incident

A validator joined the Amoy set at height 46,549,275 (15:30:27Z) with:

```
voting_power       = 1
proposer_priority  = -520,048,305     (total voting power: 481,535,557)
```

That is CometBFT's new-validator penalty (`≈ -1.125 × totalVotingPower`). With
voting power 1 it would need ~520 million blocks to ever propose — harmless in
itself, but it sat far outside the normal priority range and skewed the
ordering.

Six minutes later the metric stepped from ~30/min to ~620/min across all 24
pre-existing validators and stayed there for 11 hours.

**Measured by replaying the algorithm against the live API:**

| window | mean proposer index | blocks badly wrong |
|---|---|---|
| before the change | 1.00 | 0 of 12 |
| during (~18:50Z) | **5.86** | **5 of 14** (indices 7–20) |
| after recovery | 0.50 | 0 of 14 |

The error was **bursty — roughly a third of blocks — not a constant offset**. A
sparse sample of 9 heights returned mostly zeros and was actively misleading;
dense sampling is required to see it.

It recovered on its own at ~02:25Z the next day, when the chain re-centred that
validator's priority. Note this was a **step, not gradual drift**:

```
15:31Z  -520,058,320      (first 6.7h: only voting-power-1 accumulation)
22:14Z  -520,033,927
02:44Z  -259,988,560      <-- step, and the metric recovers at ~02:25Z
```

**The chain was healthy throughout.** Amoy's block interval held at 0.99s —
faster than mainnet's 1.21s.

---

## Ruled out, and how

Recording these because several were plausible and cost real time.

| hypothesis | verdict | what killed it |
|---|---|---|
| The chain degraded | No | Block interval flat at 0.99s across the whole window. |
| A Heimdall deploy changed the API | No | Metric flat at 28–43/min for four days, straight through the 09-16 Heimdall v2 rollout. |
| `proposer_priority` changed type string → int | No | That errors on every unmarshal; `"Failed to get Heimdall validators"` appeared **once in 3 hours**. |
| Empty or zero priorities collapsing the sort | No | Live API returns 25 distinct valid signed integers. |
| The recurring buffered-checkpoint unmarshal error | No | Chronic and flat — 60–290/hour for days, no step at the incident. Real bug, unrelated. |
| A stale cached validator set | No | There is no cache. `api.GetJSON` sets `Cache-Control: no-cache, no-store`; only the HTTP connection is pooled. |
| The API returning non-historical priorities | No | CometBFT's recurrence `priority(h) = priority(h-1) + vp − total(if proposer)` holds **exactly**, 24/24 and 25/25, before and after. |
| The pagination bug (below) | No | Amoy's 25 validators fit in one page. Both builds byte-identical on Amoy. |

---

## Fixed

### #108 → v7.0.2 — pagination, and a swallowed parse error

`getValidators` sent `/validators` with no page parameters and never read
`Result.Total`, so CometBFT capped it at its default page size of 30:

```
mainnet  /validators (no params):  count=30  total=105  returned=30
amoy     /validators (no params):  count=25  total=25   returned=25
```

**Mainnet has 105 validators and this saw 30** — hence only ever 30 signer
series there. Amoy fit inside the page, which is why the bug hid for so long.
Switched to `getValidatorsAtHeight`, which already paginated correctly and sat
directly below it; the unpaginated version is removed.

Verified by running both builds against live mainnet for ~100s:

| build | signers | off page 1 |
|---|---|---|
| v7.0.1 | **30** | **0** |
| v7.0.2 | **36** | **13** |

Confirmed in production after deploy: the mainnet signer count was pinned at
exactly 30, then climbed 81 → 94 → toward 105. Mainnet's rate also roughly
halved, which is *correct* — with the list truncated, any block whose proposer
was not in the first 30 flagged all 30.

Also stopped discarding the `strconv.Atoi` error in the sort. An unparseable
priority silently became `0`, which makes every validator compare equal and
`sort.Slice` unstable. Not hypothetical: this API already returns `""` for
numeric fields elsewhere, which is what the buffered-checkpoint warnings are.

### #109 → v8.0.0 — consistent address labels

`signer_address` is now `0x`-prefixed lowercase everywhere. Previously the
`heimdall_*` family emitted bare uppercase and the `rpc_*` family emitted
0x-prefixed lowercase, so the same validator appeared under two spellings and
cross-metric joins silently missed.

> **Breaking for consumers.** Any Grafana rule grouping by `signer_address` has a
> new deduplication key, so alerts open under the old spelling will not be
> matched by their resolve.

---

## Still open

**The algorithm.** Adding `+ voting_power` to the sort is measurably better —
perfect on 6/6 pre-change samples versus 5/6 for the current code — but it was
still wrong during the incident (mean index 6.25), so it is not the whole fix.

The likely remaining piece is **multi-round consensus**. When a proposal round
fails, CometBFT increments priorities again, so `priority(h-1) + vp` no longer
predicts the proposer. Amoy had occasional multi-round blocks during the
incident and none before:

```
before: rounds [0,0,0,0,0,0,0,0,0,0,0,0]   0/12 non-zero
during: rounds [0,0,0,0,1,0,0,0,0,0,0,0]   1/12 non-zero
```

A failed round is *exactly* what this metric ought to be counting, so some of
the increase was real signal. The magnitude never reconciled — ~8% multi-round
blocks cannot produce ~36% wrong — and that gap is the open question.

**Suggested next steps**

1. Add `+ voting_power` to the sort. Strictly closer to CometBFT.
2. **Record the consensus round** alongside the metric. Without it there is no
   way to separate genuine failed rounds from algorithm error, which is the
   single reason this took hours.
3. Consider whether the whole approach should be replaced by reading the commit
   signatures directly rather than inferring misses from priority ordering.

---

## Reproducing it

The over-count reproduces in a fresh local process against the public
endpoints — it is deterministic, not production state.

```yaml
# config.yml
http: { address: 127.0.0.1, port: 9599, pprof_port: 6599, path: /metrics }
logs: { pretty: false, verbosity: info }
namespace: panoptichain
networks: [ { name: Polygon Amoy } ]
providers:
  heimdall:
    - heimdall_url: https://heimdall-api-amoy.polygon.technology
      interval: 5s
      label: polygon.technology
      name: Polygon Amoy
      tendermint_url: https://tendermint-api-amoy.polygon.technology
```

```sh
go build -o pano ./cmd && ./pano config.yml
# after ~5 min:
curl -s localhost:9599/metrics \
  | sed -n 's/^panoptichain_heimdall_missed_block_proposal{.*signer_address="\([^"]*\)"} \(.*\)$/\1 \2/p'
```

Divide the total by `panoptichain_heimdall_block_interval_count` to get misses
per block. Healthy mainnet is ~1.5; Amoy during the incident was ~10.4.

To check the algorithm directly at a given height, without running anything —
this is what settled the investigation:

```sh
H=<heimdall height>
B=https://tendermint-api-amoy.polygon.technology
curl -s "$B/block?height=$H" | jq -r .result.block.header.proposer_address
curl -s "$B/validators?height=$((H-1))" \
  | jq -r '.result.validators | sort_by(-(.proposer_priority|tonumber)) | .[].address'
```

The proposer's position in that list is what the metric counts. Near the top is
healthy; double digits means the ordering has stopped predicting the proposer.
