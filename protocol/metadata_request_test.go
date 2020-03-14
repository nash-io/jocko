package protocol

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMetadataRequest(t *testing.T) {
	req := require.New(t)
	exp := &MetadataRequest{
		APIVersion: 1,
	}
	b, err := Encode(exp)
	req.NoError(err)
	var act MetadataRequest
	err = Decode(b, &act, exp.Version())
	req.NoError(err)
	req.Equal(exp, &act)
}

func TestMetadataRequestWithTopics(t *testing.T) {
	req := require.New(t)
	exp := &MetadataRequest{
		APIVersion: 1,
		Topics:     []string{"my-topic", "test-topic"},
	}
	b, err := Encode(exp)
	req.NoError(err)
	var act MetadataRequest
	err = Decode(b, &act, exp.Version())
	req.NoError(err)
	req.Equal(exp, &act)
}
