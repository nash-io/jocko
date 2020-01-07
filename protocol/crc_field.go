package protocol

import (
	"errors"
	"hash/crc32"
)

var table = crc32.MakeTable(crc32.Castagnoli)

type CRCField struct {
	StartOffset int
}

func (f *CRCField) SaveOffset(in int) {
	f.StartOffset = in
}

func (f *CRCField) ReserveSize() int {
	return 4
}

func (f *CRCField) Fill(curOffset int, buf []byte) error {
	crc := crc32.Checksum(buf[f.StartOffset+4:curOffset], table)
	Encoding.PutUint32(buf[f.StartOffset:], crc)
	return nil
}

func (f *CRCField) Check(curOffset int, buf []byte) error {
	crc := crc32.Checksum(buf[f.StartOffset+4:curOffset], table)
	if crc != Encoding.Uint32(buf[f.StartOffset:]) {
		return errors.New("crc didn't match")
	}
	return nil
}
