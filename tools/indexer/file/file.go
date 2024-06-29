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

// GetFolderFiles reads file names that are sittings inside the specified
// folder. File names are complemented with paths. A special flag enables
// recursive search into sub-folders.
func GetFolderFiles(folder string, readSubFolders bool, outFiles *[]string) (err error) {
	var entries []os.DirEntry
	entries, err = os.ReadDir(folder)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if readSubFolders {
				subPath := filepath.Join(folder, entry.Name())
				err = GetFolderFiles(subPath, readSubFolders, outFiles)
				if err != nil {
					return err
				}
			}

			continue
		}

		fileExt := strings.ToLower(filepath.Ext(entry.Name()))
		if fileExt != FileExtension_Torrent {
			continue
		}

		fileNameFull := filepath.Join(folder, entry.Name())
		*outFiles = append(*outFiles, fileNameFull)
	}

	return nil
}

// CleanFilePath is a "fixed" version of the standard 'filepath.Clean' method.
func CleanFilePath(path string) (cp string) {
	cp = filepath.Clean(path)

	// The 'filepath.Clean' method adds a dot symbol to an empty string.
	// We do not need it !
	if (len(cp) == 1) && (cp == `.`) {
		cp = ""
	}

	return cp
}
