package protocol

type Record struct {
	Length         int64
	Attributes     int8
	TimestampDelta int64
	OffsetDelta    int64
	Key            []byte
	Value          []byte
	Headers        []*Header
}

func (r *Record) Encode(e PacketEncoder) error {
	var err error
	e.PutVarint(r.Length)
	e.PutInt8(r.Attributes)
	e.PutVarint(r.TimestampDelta)
	e.PutVarint(r.OffsetDelta)
	e.PutVarintBytes(r.Key)
	e.PutVarintBytes(r.Value)
	for _, h := range r.Headers {
		if err = h.Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (r *Record) Decode(d PacketDecoder) error {
	var err error
	if r.Length, err = d.Varint(); err != nil {
		return err
	}
	if r.Attributes, err = d.Int8(); err != nil {
		return err
	}
	if r.TimestampDelta, err = d.Varint(); err != nil {
		return err
	}
	if r.OffsetDelta, err = d.Varint(); err != nil {
		return err
	}
	if r.Key, err = d.VarintBytes(); err != nil {
		return err
	}
	if r.Value, err = d.VarintBytes(); err != nil {
		return err
	}
	var hl int64
	if hl, err = d.Varint(); err != nil {
		return err
	}
	for hl > 0 {
		m := new(Header)
		err = m.Decode(d)
		switch err {
		case nil:
			r.Headers = append(r.Headers, m)
			hl -= 1
		default:
			return err
		}
	}

	return nil
}
