package main

import (
	"fmt"

	"github.com/Prabesh-Sharma/suckless-client/handler"
)

func main() {
	tf := handler.NewTorrentFile()
	err := tf.TorrentFileHandler("/home/albert/ssd/suckless-client/test-file/test.torrent")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", tf)
}
