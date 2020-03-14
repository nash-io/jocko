package protocol

import "fmt"

type RecordBatch struct {
	BaseOffset              int64
	BatchLength             int32
	PartitionLeaderEpoch    int32
	Magic                   int8
	Attributes              int16
	LastOffsetDelta         int32
	FirstTimestamp          int64
	MaxTimestamp            int64
	ProducerId              int64
	ProducerEpoch           int16
	BaseSequence            int32
	Records                 []*Record
	PartialTrailingMessages bool
}

func (rb *RecordBatch) String() string {
	return fmt.Sprintf(`RecordBatch{
BatchOffset=%d
BatchLength=%d
PartitionLeaderEpoch=%d
Magic=%d
Attributes=%d,
LastOffsetDelta=%d,
FirstTimestamp=%d,
MaxTimestamp=%d,
ProducerId=%d,
ProduceEpoch=%d,
BaseSequence=%d
Records=%d}`,
		rb.BaseOffset, rb.BatchLength, rb.PartitionLeaderEpoch,
		rb.Magic, rb.Attributes, rb.LastOffsetDelta,
		rb.FirstTimestamp, rb.MaxTimestamp, rb.ProducerId,
		rb.ProducerEpoch, rb.BaseSequence, len(rb.Records),
	)
}
func (rb *RecordBatch) Encode(e PacketEncoder) error {
	var err error
	e.PutInt64(rb.BaseOffset)
	e.PutInt32(rb.BatchLength)
	e.PutInt32(rb.PartitionLeaderEpoch)
	e.PutInt8(rb.Magic)
	e.Push(&CRCField{})
	e.PutInt16(rb.Attributes)
	e.PutInt32(rb.LastOffsetDelta)
	e.PutInt64(rb.FirstTimestamp)
	e.PutInt64(rb.MaxTimestamp)
	e.PutInt64(rb.ProducerId)
	e.PutInt16(rb.ProducerEpoch)
	e.PutInt32(rb.BaseSequence)
	e.PutInt32(int32(len(rb.Records)))
	for _, r := range rb.Records {
		if err = r.Encode(e); err != nil {
			return err
		}
	}
	e.Pop()
	return nil
}

func (rb *RecordBatch) Decode(d PacketDecoder) error {
	var err error
	if rb.BaseOffset, err = d.Int64(); err != nil {
		return err
	}
	if rb.BatchLength, err = d.Int32(); err != nil {
		return err
	}
	if rb.PartitionLeaderEpoch, err = d.Int32(); err != nil {
		return err
	}
	if rb.Magic, err = d.Int8(); err != nil {
		return err
	}
	if err = d.Push(&CRCField{}); err != nil {
		return err
	}
	if rb.Attributes, err = d.Int16(); err != nil {
		return err
	}
	if rb.LastOffsetDelta, err = d.Int32(); err != nil {
		return err
	}
	if rb.FirstTimestamp, err = d.Int64(); err != nil {
		return err
	}
	if rb.MaxTimestamp, err = d.Int64(); err != nil {
		return err
	}
	if rb.ProducerId, err = d.Int64(); err != nil {
		return err
	}
	if rb.ProducerEpoch, err = d.Int16(); err != nil {
		return err
	}
	if rb.BaseSequence, err = d.Int32(); err != nil {
		return err
	}
	var rl int32
	if rl, err = d.Int32(); err != nil {
		return err
	}
	for rl > 0 {
		r := new(Record)
		err = r.Decode(d)
		switch err {
		case nil:
			rb.Records = append(rb.Records, r)
			rl -= 1
		case ErrInsufficientData:
			rb.PartialTrailingMessages = true
			return nil
		default:
			return err
		}
	}
	return d.Pop()
}
