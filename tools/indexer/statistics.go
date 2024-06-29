package main

import (
	"fmt"
	"os"
	"time"
)

// Statistics stores statistical data about program execution.
type Statistics struct {
	TimeStart    time.Time
	TimeDuration time.Duration

	OutputFileName string

	InspectedBTFilesCount uint
	ProcessedBTFiles      []string
	IndexedFilesCount     uint

	BrokenFilesCount uint
	BrokenFiles      []string

	SelfCheckErrorsCount  uint
	SelfCheckErroredFiles []string
}

// showResume shows the results of the program execution.
func showResume(stat *Statistics) (err error) {
	fmt.Println(MsgProcessingHasFinished)
	fmt.Println(fmt.Sprintf(MsgFTimeTakenSeconds, stat.TimeDuration.Seconds()))

	// Count bytes.
	var outputFileSize int
	outputFileSize, err = getFileSize(stat.OutputFileName)
	if err != nil {
		return err
	}

	var btFilesSize int = 0
	var fileSize int
	for _, btFile := range stat.ProcessedBTFiles {
		fileSize, err = getFileSize(btFile)
		if err != nil {
			return err
		}

		btFilesSize += fileSize
	}

	fmt.Println(fmt.Sprintf(MsgFBytesRead, btFilesSize))
	fmt.Println(fmt.Sprintf(MsgFBytesWritten, outputFileSize))

	// Count files.
	fmt.Println(fmt.Sprintf(MsgFInspectedBTFilesCount, stat.InspectedBTFilesCount))
	fmt.Println(fmt.Sprintf(MsgFFilesIndexed, stat.IndexedFilesCount))
	fmt.Println()

	fmt.Println(fmt.Sprintf(MsgFFilesBroken, stat.BrokenFilesCount))
	printList(stat.BrokenFiles)
	fmt.Println()

	fmt.Println(fmt.Sprintf(MsgFSelfCheckErrors, stat.SelfCheckErrorsCount))
	printList(stat.SelfCheckErroredFiles)
	fmt.Println()

	return nil
}

// getFileSize reads the size of a file.
func getFileSize(file string) (size int, err error) {
	var fi os.FileInfo
	fi, err = os.Stat(file)
	if err != nil {
		return 0, err
	}

	return int(fi.Size()), nil
}

// printList prints a list of strings adding a space delimiter and a '-' mark
// before each line.
func printList(list []string) {
	for _, s := range list {
		fmt.Println(fmt.Sprintf("  - %s", s))
	}
}
