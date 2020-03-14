package commitlog_test

import (
	"testing"

	"github.com/nash-io/jocko/commitlog"
	"github.com/stretchr/testify/require"
)

func TestSegmentScanner(t *testing.T) {
	var err error

	l := setupWithOptions(t, commitlog.Options{
		MaxSegmentBytes: 1000,
		MaxLogBytes:     1000,
	})
	defer cleanup(t, l)

	for _, msgSet := range msgSets {
		_, err = l.Append(msgSet)
		require.NoError(t, err)
	}

	segments := l.Segments()
	segment := segments[0]

	scanner := commitlog.NewSegmentScanner(segment)
	ms, err := scanner.Scan()
	require.NoError(t, err)
	require.Equal(t, msgSets[0], ms)
}
