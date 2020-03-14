package protocol

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecordBatchDecodeKafka011(t *testing.T) {
	b, err := ioutil.ReadFile("kafka_011.log")
	if err != nil {
		t.Errorf(err.Error())
	}
	var rb RecordBatch
	if err = Decode0(b, &rb); err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println(rb.Records[0])
	firstRecordValue := rb.Records[0].Value
	require.Equal(t, string(firstRecordValue), "magnus", "The two words should be the same.")

}
