.PHONY: build
build: out
	go build -o out/panoptichain cmd/main.go

out:
	mkdir out

metrics.md:
	go run cmd/read-observers/main.go --md > metrics.md


