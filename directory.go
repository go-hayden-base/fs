package fs

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ListDirectory return subfile by callback function
func ListDirectory(dirname string, recursive bool, callback ListDirectCallbackFunc) {
	if callback == nil {
		return
	}
	if !DirectoryExists(dirname) {
		msg := "Directory path '" + dirname + "' is not exist!"
		callback(nil, errors.New(msg))
		return
	}
	dir, err := ioutil.ReadDir(dirname)
	if err != nil {
		callback(nil, err)
	}

	for _, fileinfo := range dir {
		fi := new(FileInfoBase)
		fi.FileInfo = fileinfo
		fi.Path = filepath.Join(dirname, fileinfo.Name())
		var newFileInfo FileInfo = fi
		callback(newFileInfo, nil)
		if recursive && newFileInfo.IsDir() {
			ListDirectory(newFileInfo.FilePath(), recursive, callback)
		}
	}
}

// DirectoryExists return a bool value that the directory is exist or not
func DirectoryExists(path string) bool {
	if len(path) == 0 {
		return false
	}
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}

	return stat.IsDir()
}
