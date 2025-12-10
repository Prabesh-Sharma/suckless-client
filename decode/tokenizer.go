package decode

import (
	"bytes"
	"fmt"
	"strconv"
)

const (
	S = byte('s')
	L = byte('l')
	D = byte('d')
	I = byte('i')
	E = byte('e')
)

func (p *BencodeParser) decodeDict() (any, error) {
	if p.data[p.offset] != D {
		return nil, fmt.Errorf("malformed file doesnot start with dict")
	}
	p.offset += 1 // consume 'd'
	dict := make(map[string]any)
	for p.data[p.offset] != E {
		key, err := p.decodeString()
		if err != nil {
			return nil, err
		}
		val, err := p.Decoder()
		if err != nil {
			return nil, err
		}
		dict[key] = val
	}
	p.offset += 1 // consume 'e'
	return dict, nil
}

func (p *BencodeParser) decodeList() ([]any, error) {
	p.offset += 1 // consume 'l'
	var list []any
	for p.data[p.offset] != E {
		item, err := p.Decoder()
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	p.offset += 1 // consume 'e'
	return list, nil
}

func (p *BencodeParser) decodeString() (string, error) {
	s := bytes.SplitN(p.data[p.offset:], []byte(":"), 2)[0]
	x := string(s)
	i, err := strconv.Atoi(x)
	if err != nil {
		return "", err
	}
	p.offset += len(x) + 1 // consume no. of letters + colon
	key := string(p.data[p.offset : p.offset+i])
	p.offset += i
	return key, nil
}

func (p *BencodeParser) decodeInt() (int, error) {
	p.offset += 1 // consume 'i'
	start := p.offset
	for p.data[p.offset] != E {
		p.offset += 1
	}
	i, err := strconv.Atoi(string(p.data[start:p.offset]))
	if err != nil {
		return 0, err
	}
	p.offset += 1 // consume 'e'
	return i, nil
}
