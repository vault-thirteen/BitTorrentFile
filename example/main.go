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
	FileName3     = "OneFile_V2.torrent"
	FileName4     = "MultipleFiles_V2.torrent"
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
		filepath.Join(ExampleFolder, DataFolder, FileName4),
	)

	err = tf.Open()
	if err != nil {
		return err
	}

	filesCount := len(tf.Files)

	fmt.Println(fmt.Sprintf("Source: %s.", tf.Source.GetPath()))
	fmt.Println(fmt.Sprintf("Version: %s.", tf.Version))
	fmt.Println(fmt.Sprintf("Name: %s.", tf.Name))
	fmt.Println(fmt.Sprintf("BTIH: %s.", tf.BTIH.Text))
	fmt.Println(fmt.Sprintf("BTIH V2: %s.", tf.BTIH2.Text))
	fmt.Println(fmt.Sprintf("%d file(s):", filesCount))
	for _, file := range tf.Files {
		fmt.Println(fmt.Sprintf("\t - %v (%d bytes)", file.Path, file.Size))
	}

	return nil
}
