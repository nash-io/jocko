package protocol

type Header struct {
	HeaderKeyLength   int64
	HeaderKey         int64
	HeaderValueLength int64
	Value             []byte
}

func (m *Header) Encode(e PacketEncoder) error {
	return nil
}

func (m *Header) Decode(d PacketDecoder) error {
	return nil
}
