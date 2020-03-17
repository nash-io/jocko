package protocol

import (
	"testing"
)

//TODO! Add a log file with Record batch data to recover this tests.
func TestRecordBatchDecodeKafka011(t *testing.T) {
	/* 	b, err := ioutil.ReadFile("kafka_011.log")
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
	   	require.Equal(t, string(firstRecordValue), "magnus", "The two words should be the same.") */

}
