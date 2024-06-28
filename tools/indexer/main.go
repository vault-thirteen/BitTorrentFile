package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/vault-thirteen/BitTorrentFile/tools/indexer/cla"
	"github.com/vault-thirteen/BitTorrentFile/tools/indexer/file"
	ver "github.com/vault-thirteen/auxie/Versioneer"
)

const UsageHint = `Usage:
	[ObjectType] [ObjectPath] [Output]

Examples:
	indexer.exe File "123.torrent" "index.csv" 
	indexer.exe Folder "torrents" "index.csv"

Notes:
	Possible object types: File, Folder or Directory.
	Letter case is not important.
	Change directory (CD) to a working directory before usage.`

const MsgFProcessingFiles = "Processing %d files."

func main() {
	args, err := cla.NewCommandLineArguments()
	if err != nil {
		log.Println(err)
		showIntro()
		showUsage()
		os.Exit(1)
		return
	}

	switch args.ObjectType.ID() {
	case cla.ObjectTypeId_File:
		err = processFile(args)
		mustBeNoError(err)

	case cla.ObjectTypeId_Folder:
		err = processFolder(args)
		mustBeNoError(err)

	default:
		err = fmt.Errorf(cla.ErrFObjectTypeUnknown, args.ObjectType.ID())
		mustBeNoError(err)
	}
}

func mustBeNoError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func showIntro() {
	versioneer, err := ver.New()
	mustBeNoError(err)
	versioneer.ShowIntroText("indexer")
	versioneer.ShowComponentsInfoText()
	fmt.Println()
}

func showUsage() {
	fmt.Println(UsageHint)
}

func processFile(args *cla.CommandLineArguments) (err error) {
	fileExt := strings.ToLower(filepath.Ext(args.ObjectPath))
	if fileExt != file.FileExtension_Torrent {
		return fmt.Errorf(file.ErrFFileExtensionUnsupported, fileExt)
	}

	var files = []string{args.ObjectPath}
	return processFiles(files, args.Output)
}

func processFolder(args *cla.CommandLineArguments) (err error) {
	var files []string
	files, err = file.GetFolderFiles(args.ObjectPath)
	if err != nil {
		return err
	}

	return processFiles(files, args.Output)
}

func processFiles(files []string, output string) (err error) {
	fmt.Println(fmt.Sprintf(MsgFProcessingFiles, len(files)))

	return nil //TODO
}
