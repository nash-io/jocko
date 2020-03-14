package jocko

import (
	"testing"

	"github.com/nash-io/jocko/jocko/structs"
	"github.com/stretchr/testify/require"
)

func TestNewReplicaLookup(t *testing.T) {
	lookup := NewReplicaLookup()
	rep := &Replica{BrokerID: 1, Partition: structs.Partition{Topic: "test-topic", ID: 1}}

	lookup.AddReplica(rep)
	got, err := lookup.Replica("test-topic", 1)
	require.NoError(t, err)
	require.Equal(t, got, rep)

	lookup.RemoveReplica(rep)

	got, err = lookup.Replica("test-topic", 1)
	require.Error(t, err)
	require.Nil(t, got)
}
