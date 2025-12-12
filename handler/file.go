// Package handler does all the file handling operations after decoder finishes
package handler

import (
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/Prabesh-Sharma/suckless-client/decode"
)

type TorrentFile struct {
	Announce    string
	InfoHash    [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

func NewTorrentFile() *TorrentFile {
	return &TorrentFile{}
}

// TorrentFileHandler creates a TorrentFile from given FilePath
func (tf *TorrentFile) TorrentFileHandler(TorFilePath string) error {
	data, err := os.ReadFile(TorFilePath)
	if err != nil {
		return err
	}

	// parser initialized
	p := decode.NewBencodeParser(data)
	decoded, err := p.Decoder()
	if err != nil {
		return err
	}

	// typecasting for decoded file
	fileDict, ok := decoded.(map[string]any)
	if !ok {
		return fmt.Errorf("typecasting failed fileDict")
	}

	announce, ok := fileDict["announce"].(string)
	if !ok {
		return fmt.Errorf("typecasting failed announce")
	}
	tf.Announce = announce

	infoDict, ok := fileDict["info"].(map[string]any)
	if !ok {
		return fmt.Errorf("typecasting failed infoDict")
	}

	infoLen, ok := infoDict["length"].(int)
	if !ok {
		return fmt.Errorf("typecasting failed infoLen")
	}
	tf.Length = infoLen

	infoName, ok := infoDict["name"].(string)
	if !ok {
		return fmt.Errorf("typecasting failed infoName")
	}
	tf.Name = infoName

	pieceLen, ok := infoDict["piece length"].(int)
	if !ok {
		return fmt.Errorf("typecasting failed pieceLen")
	}
	tf.PieceLength = pieceLen

	tf.InfoHash = calculateInfoHash(p, data)

	pieces := infoDict["pieces"]
	pieceHashes, err := arrangePieceHashes([]byte(pieces.(string)))
	if err != nil {
		return err
	}
	tf.PieceHashes = pieceHashes
	return nil
}

func calculateInfoHash(p *decode.BencodeParser, data []byte) [20]byte {
	infoBytes := data[p.InfoStart:p.InfoEnd]
	hash := sha1.Sum(infoBytes)
	return hash
}

func arrangePieceHashes(data any) ([][20]byte, error) {
	raw, ok := data.([]byte)
	if !ok {
		return [][20]byte{}, fmt.Errorf("typecasting error raw")
	}
	pieceHashes := make([][20]byte, 0, len(raw)/20)
	for i := 0; i < len(raw); i += 20 {
		pieceHashes = append(pieceHashes, [20]byte(raw[i:i+20]))
	}
	return pieceHashes, nil
}
