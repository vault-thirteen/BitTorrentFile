package main

import (
	"os"
	"time"

	ver "github.com/vault-thirteen/auxie/Versioneer/classes/Versioneer"
	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/errors"

	"github.com/vault-thirteen/BitTorrentFile/tools/indexer/cla"
)

// Program's entry point.
func main() {
	showIntro()

	args, err := cla.NewCommandLineArgumentsFromOs()
	if err != nil {
		errors.LogError(err)
		showUsage()
		os.Exit(ExitCode_CLAError)
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
		err = errors.ErrorFromFormattedText(cla.Errf_ObjectTypeUnknown, args.ObjectType.ID())
		mustBeNoError(err)
	}

	// Résumé.
	stat.TimeStart = timeStart
	stat.TimeDuration = time.Now().Sub(stat.TimeStart)
	err = showResume(stat)
	mustBeNoError(err)
}

// mustBeNoError exits program on error.
func mustBeNoError(err error) {
	if err != nil {
		errors.LogError(err)
		os.Exit(ExitCode_Other)
	}
}

// showIntro shows introductory information about the program.
func showIntro() {
	versioneer, err := ver.New(false)
	mustBeNoError(err)
	versioneer.ShowIntroText(ToolName)
	versioneer.ShowComponentsInfoText()
	cio.PrintMessage()
}

// showUsage prints the information about how to use the program.
func showUsage() {
	cio.PrintMessage(UsageHint)
}
