package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	btf "github.com/vault-thirteen/BitTorrentFile"
	"github.com/vault-thirteen/BitTorrentFile/tools/indexer/cla"
	"github.com/vault-thirteen/BitTorrentFile/tools/indexer/file"
	"github.com/vault-thirteen/auxie/csv"
	"github.com/vault-thirteen/auxie/errors"
	"github.com/vault-thirteen/bencode"
)

// processFile processes a single file.
func processFile(args *cla.CommandLineArguments) (stat *Statistics, err error) {
	fileExt := strings.ToLower(filepath.Ext(args.ObjectPath))
	if fileExt != file.FileExtension_Torrent {
		return nil, fmt.Errorf(file.ErrFFileExtensionUnsupported, fileExt)
	}

	var files = []string{args.ObjectPath}
	return processFiles(files, args.Output)
}

// processFolder processes files of a folder.
func processFolder(args *cla.CommandLineArguments) (stat *Statistics, err error) {
	var files []string
	err = file.GetFolderFiles(args.ObjectPath, true, &files)
	if err != nil {
		return nil, err
	}

	return processFiles(files, args.Output)
}

// processFiles processes specified files.
func processFiles(files []string, output string) (stat *Statistics, err error) {
	fmt.Println(fmt.Sprintf(MsgFProcessingFiles, len(files)))

	stat = &Statistics{
		OutputFileName:        output,
		ProcessedBTFiles:      make([]string, 0, len(files)),
		IndexedFilesCount:     0,
		BrokenFiles:           make([]string, 0, len(files)),
		SelfCheckErroredFiles: make([]string, 0, len(files)),
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
	var torrentFileInfo *btf.BitTorrentFile
	var csvLine []any
	for _, btFile := range files {
		stat.InspectedBTFilesCount++
		fmt.Println(fmt.Sprintf("%6d. %s", stat.InspectedBTFilesCount, btFile))

		torrentFileInfo, err = getTorrentFileInfo(btFile)
		if err != nil {
			// Self-check error means that file is damaged. E.g. it may have
			// additional data below the official format space. Such behaviour
			// is popular among stupid hackers or people who try to edit
			// BitTorrent files manually without proper knowledge how to do it.
			// It may also be an exploit of some bug or something else. We
			// advise to skip processing of such files.
			if err.Error() == bencode.ErrSelfCheck {
				stat.SelfCheckErrorsCount++
				stat.SelfCheckErroredFiles = append(stat.SelfCheckErroredFiles, btFile)
				log.Println(err.Error())

				if SkipHackedFiles {
					continue
				}
			}

			// On other errors we stop.
			return nil, err
		}

		// Some idiots like to violate specifications. Fools will always be
		// fools. You can try processing such files while it is not harmful.
		if torrentFileInfo.IsBroken {
			stat.BrokenFilesCount++
			stat.BrokenFiles = append(stat.BrokenFiles, btFile)
			log.Println(ErrBrokenFile)

			if SkipBrokenFiles {
				continue
			}
		}

		// For each stored file ...
		for _, stFile := range torrentFileInfo.Files {
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

// getTorrentFileInfo parses the BitTorrent file and returns information about
// it.
func getTorrentFileInfo(btFile string) (tf *btf.BitTorrentFile, err error) {
	tf = btf.NewBitTorrentFile(btFile)

	err = tf.Open()
	if err != nil {
		return nil, err
	}

	return tf, nil
}
