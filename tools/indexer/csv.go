package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/vault-thirteen/BitTorrentFile/models"
	"github.com/vault-thirteen/BitTorrentFile/tools/indexer/file"
)

// CSV file header.
const (
	// CsvColumn1 is a base name of the indexed BitTorrent file.
	// Base name does not contain file extension.
	CsvColumn1 = "Torrent file"

	// CsvColumn2 is a path to the indexed BitTorrent file.
	// Path does not contain the file name.
	CsvColumn2 = "Torrent path"

	// CsvColumn3 is a name of a file stored inside the indexed
	// BitTorrent file. Name contains file extension.
	CsvColumn3 = "Stored file"

	// CsvColumn4 is a relative path to a file stored inside the indexed
	// BitTorrent file. Path does not contain the file name.
	CsvColumn4 = "Stored path"

	// CsvColumn5 is a size of a file stored inside the indexed
	// BitTorrent file. Size is set in bytes.
	CsvColumn5 = "Stored size"
)

// getCsvFileHeader returns the header line for the CSV table.
func getCsvFileHeader() []any {
	return []any{CsvColumn1, CsvColumn2, CsvColumn3, CsvColumn4, CsvColumn5}
}

// prepareCsvLine returns the line (row) for the CSV table.
func prepareCsvLine(torrentFilePath string, storedFileInfo models.File) (line []any, err error) {
	torrentFileFolder, torrentFileName := filepath.Split(torrentFilePath)
	torrentFileFolder = escapeBackSlashes(file.CleanFilePath(torrentFileFolder))
	torrentFileExt := filepath.Ext(torrentFileName)
	torrentFileBaseName := strings.TrimSuffix(torrentFileName, torrentFileExt)

	if torrentFileExt != file.FileExtension_Torrent {
		return nil, fmt.Errorf(file.ErrFFileExtensionUnsupported, torrentFileExt)
	}

	storedFilePath := filepath.Join(storedFileInfo.Path...)
	storedFileFolder, storedFileName := filepath.Split(storedFilePath)
	storedFileFolder = escapeBackSlashes(file.CleanFilePath(storedFileFolder))

	return []any{torrentFileBaseName, torrentFileFolder, storedFileName, storedFileFolder, storedFileInfo.Size}, nil
}

// escapeBackSlashes escapes all the backward slashes. Some SQL tools, such as
// DBeaver, require backslashes to be escaped in CSV files.
func escapeBackSlashes(s string) string {
	return strings.ReplaceAll(s, `\`, `\\`)
}
