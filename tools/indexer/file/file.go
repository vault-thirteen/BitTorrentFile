package file

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	FileExtension_Torrent = ".torrent"
)

const (
	ErrFFileExtensionUnsupported = "file extension unsupported: %s"
)

func GetFolderFiles(folder string) (files []string, err error) {
	var entries []os.DirEntry
	entries, err = os.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		fileExt := strings.ToLower(filepath.Ext(entry.Name()))
		if fileExt != FileExtension_Torrent {
			continue
		}

		fileNameFull := filepath.Join(folder, entry.Name())
		files = append(files, fileNameFull)
	}

	return files, nil
}
