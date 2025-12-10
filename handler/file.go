// Package handler does all the file handling operations after decoder finishes
package handler

import (
	"fmt"
	"os"

	"github.com/Prabesh-Sharma/suckless-client/decode"
)

func FileHandler() {
	data, err := os.ReadFile("/home/albert/ssd/suckless-client/test-file/test.torrent")
	if err != nil {
		fmt.Println("error while reading file: ", err)
	}
	p := decode.NewBencodeParser(data)
	decoded, err := p.Decoder()
	if err != nil {
		panic(err)
	}
	dict, ok := decoded.(map[string]any)
	if !ok {
		panic("typecasting failed")
	}
	infoDict := dict["info"].(map[string]any)
	fmt.Println(infoDict["piece length"])
}
