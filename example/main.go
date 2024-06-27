package main

import (
	"fmt"
	"log"
	"path/filepath"

	btf "github.com/vault-thirteen/BitTorrentFile"
)

// Settings.
const (
	ExampleFolder = "example"
	DataFolder    = "data"
	FileName1     = "5942384.torrent"
	FileName2     = "DX12.torrent"
)

func main() {
	var err = openFile()
	mustBeNoError(err)
}

func mustBeNoError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func openFile() (err error) {
	var tf = btf.NewBitTorrentFile(
		filepath.Join(ExampleFolder, DataFolder, FileName2),
	)

	err = tf.Open()
	if err != nil {
		return err
	}

	fmt.Println(tf.Source.GetPath())
	fmt.Println(tf.BTIH.Text)

	return nil
}
