package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	btf "github.com/vault-thirteen/BitTorrentFile"
	"github.com/vault-thirteen/BitTorrentFile/models"
	"github.com/vault-thirteen/BitTorrentFile/shitty/csv"
	"github.com/vault-thirteen/BitTorrentFile/tools/indexer/cla"
	"github.com/vault-thirteen/BitTorrentFile/tools/indexer/file"
	"github.com/vault-thirteen/auxie/errors"
)

func processFile(args *cla.CommandLineArguments) (stat *Statistics, err error) {
	fileExt := strings.ToLower(filepath.Ext(args.ObjectPath))
	if fileExt != file.FileExtension_Torrent {
		return nil, fmt.Errorf(file.ErrFFileExtensionUnsupported, fileExt)
	}

	var files = []string{args.ObjectPath}
	return processFiles(files, args.Output)
}

func processFolder(args *cla.CommandLineArguments) (stat *Statistics, err error) {
	var files []string
	err = file.GetFolderFiles(args.ObjectPath, true, &files)
	if err != nil {
		return nil, err
	}

	return processFiles(files, args.Output)
}

func processFiles(files []string, output string) (stat *Statistics, err error) {
	fmt.Println(fmt.Sprintf(MsgFProcessingFiles, len(files)))

	stat = &Statistics{
		OutputFileName:    output,
		ProcessedBTFiles:  make([]string, 0, len(files)),
		IndexedFilesCount: 0,
	}

	var csvFile *os.File
	csvFile, err = os.Create(output)
	if err != nil {
		return nil, err
	}

	defer func() {
		derr := csvFile.Close()
		if derr != nil {
			err = errors.Combine(err, derr)
		}
	}()

	// Unfortunately standard CSV library in Golang is absolutely useless.
	// It does not add quotes to strings when writes them into a CSV file !!!

	csvWriter := csv.NewWriter(csvFile)

	err = csvWriter.WriteRow(getCsvFileHeader())
	if err != nil {
		return nil, err
	}

	// For each BitTorrent file ...
	var storedFilesInfo []models.File
	var csvLine []any
	var n = 0
	for _, btFile := range files {
		n++
		fmt.Println(fmt.Sprintf("%6d. %s", n, btFile))

		storedFilesInfo, err = getStoredFilesInfo(btFile)
		if err != nil {
			return nil, err
		}

		// For each stored file ...
		for _, stFile := range storedFilesInfo {
			csvLine, err = prepareCsvLine(btFile, stFile)
			if err != nil {
				return nil, err
			}

			err = csvWriter.WriteRow(csvLine)
			if err != nil {
				return nil, err
			}

			stat.IndexedFilesCount++
		}

		stat.ProcessedBTFiles = append(stat.ProcessedBTFiles, btFile)
	}

	return stat, nil
}

func getStoredFilesInfo(btFile string) (fsi []models.File, err error) {
	var tf = btf.NewBitTorrentFile(btFile)

	err = tf.Open()
	if err != nil {
		return nil, err
	}

	return tf.Files, nil
}
