// Package heimdall provides helper functions for decoding Heimdall v2 vote extensions.
package heimdall

import (
	"google.golang.org/protobuf/proto"
)

// BlockIDFlag constants for convenient access.
const (
	BlockIDFlagUnknown = int(BlockIDFlag_BLOCK_ID_FLAG_UNKNOWN)
	BlockIDFlagAbsent  = int(BlockIDFlag_BLOCK_ID_FLAG_ABSENT)
	BlockIDFlagCommit  = int(BlockIDFlag_BLOCK_ID_FLAG_COMMIT)
	BlockIDFlagNil     = int(BlockIDFlag_BLOCK_ID_FLAG_NIL)
)

// UnmarshalExtendedCommitInfo decodes a protobuf-encoded ExtendedCommitInfo.
func UnmarshalExtendedCommitInfo(data []byte) (*ExtendedCommitInfo, error) {
	info := &ExtendedCommitInfo{}
	if err := proto.Unmarshal(data, info); err != nil {
		return nil, err
	}
	return info, nil
}

// UnmarshalVoteExtension decodes a protobuf-encoded VoteExtension.
func UnmarshalVoteExtension(data []byte) (*VoteExtension, error) {
	ve := &VoteExtension{}
	if err := proto.Unmarshal(data, ve); err != nil {
		return nil, err
	}
	return ve, nil
}
