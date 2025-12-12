// Package decode implements decoding of the bencoded
// .torrent files
package decode

import (
	"fmt"
)

type BencodeParser struct {
	data      []byte
	offset    int
	InfoStart int
	InfoEnd   int
	depth     int
}

// NewBencodeParser creates a BencodeParser instance
func NewBencodeParser(data []byte) *BencodeParser {
	return &BencodeParser{data: data, offset: 0, InfoStart: 0, InfoEnd: 0, depth: 0}
}

// Decoder is recursive descent parser for bencoded files
func (p *BencodeParser) Decoder() (any, error) {
	if p.offset >= len(p.data) {
		return nil, fmt.Errorf("no data at current offset")
	}

	switch s := p.data[p.offset]; s {
	case L:
		return p.decodeList()
	case D:
		return p.decodeDict()
	case I:
		return p.decodeInt()
	default:
		if s >= '0' && s <= '9' {
			return p.decodeString()
		}
		return nil, fmt.Errorf("invalid starting char: %q", s)
	}
}
