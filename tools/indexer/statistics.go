package main

import (
	"fmt"
	"os"
	"time"
)

type Statistics struct {
	TimeStart         time.Time
	TimeDuration      time.Duration
	OutputFileName    string
	ProcessedBTFiles  []string
	IndexedFilesCount uint
}

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
	fmt.Println(fmt.Sprintf(MsgFFilesIndexed, stat.IndexedFilesCount))

	return nil
}

func getFileSize(file string) (size int, err error) {
	var fi os.FileInfo
	fi, err = os.Stat(file)
	if err != nil {
		return 0, err
	}

	return int(fi.Size()), nil
}
