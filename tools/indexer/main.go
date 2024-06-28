package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/vault-thirteen/BitTorrentFile/tools/indexer/cla"
	ver "github.com/vault-thirteen/auxie/Versioneer"
)

const ToolName = "Indexer"

func main() {
	showIntro()

	args, err := cla.NewCommandLineArguments()
	if err != nil {
		log.Println(err)
		showUsage()
		os.Exit(1)
		return
	}

	timeStart := time.Now()
	var stat *Statistics

	switch args.ObjectType.ID() {
	case cla.ObjectTypeId_File:
		stat, err = processFile(args)
		mustBeNoError(err)

	case cla.ObjectTypeId_Folder:
		stat, err = processFolder(args)
		mustBeNoError(err)

	default:
		err = fmt.Errorf(cla.ErrFObjectTypeUnknown, args.ObjectType.ID())
		mustBeNoError(err)
	}

	// Résumé.
	stat.TimeStart = timeStart
	stat.TimeDuration = time.Now().Sub(stat.TimeStart)
	err = showResume(stat)
	mustBeNoError(err)
}

func mustBeNoError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func showIntro() {
	versioneer, err := ver.New()
	mustBeNoError(err)
	versioneer.ShowIntroText(ToolName)
	versioneer.ShowComponentsInfoText()
	fmt.Println()
}

func showUsage() {
	fmt.Println(UsageHint)
}
