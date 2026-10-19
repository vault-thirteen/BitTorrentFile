package main

import (
	"os"
	"path/filepath"

	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/errors"

	btf "github.com/vault-thirteen/BitTorrentFile"
)

// This is a simple example how to use the library.
// You can study this file in debugger and try changing the file name to other
// values (such as FileName2, FileName3 or FileName4) and see the results.

// Settings.
const (
	ExampleFolder = "example"
	DataFolder    = "data"
	FileName1     = "5942384.torrent"
	FileName2     = "DX12.torrent"
	FileName3     = "OneFile_V2.torrent"
	FileName4     = "MultipleFiles_V2.torrent"
)

const (
	ExitCode_Failure = 1
)

func main() {
	var err = openFile()
	mustBeNoError(err)
}

func mustBeNoError(err error) {
	if err != nil {
		errors.LogError(err)
		os.Exit(ExitCode_Failure)
	}
}

func openFile() (err error) {
	var tf = btf.NewBitTorrentFile(
		filepath.Join(ExampleFolder, DataFolder, FileName1),
	)

	err = tf.Open()
	if err != nil {
		return err
	}

	filesCount := len(tf.Files)

	cio.PrintMessage(cio.FormatMessage("Source: %s.", tf.Source.GetPath()))
	cio.PrintMessage(cio.FormatMessage("Version: %s.", tf.Version))
	cio.PrintMessage(cio.FormatMessage("Name: %s.", tf.Name))
	cio.PrintMessage(cio.FormatMessage("BTIH: %s.", tf.BTIH.Text))
	cio.PrintMessage(cio.FormatMessage("BTIH V2: %s.", tf.BTIH2.Text))
	cio.PrintMessage(cio.FormatMessage("%d file(s):", filesCount))
	for _, file := range tf.Files {
		cio.PrintMessage(cio.FormatMessage("\t - %v (%d bytes)", file.Path, file.Size))
	}
	cio.PrintMessage(cio.FormatMessage("Is broken: %t.", tf.IsBroken))

	return nil
}
