package fs

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// FileExists return a bool value that file is exist or not
func FileExists(path string) bool {
	if len(path) == 0 {
		return false
	}
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !stat.IsDir()
}

// ReadLine is a function to read file which specify by filePath
func ReadLine(filePath string, callback ReadLineCallbackFunc) {
	if callback == nil {
		return
	}

	f, err := os.Open(filePath)
	defer f.Close()

	stop := false

	if err != nil {
		callback("", true, err, &stop)
		return
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				callback("", true, nil, &stop)
				break
			} else {
				callback("", false, err, &stop)
			}
		} else {
			callback(string(b), false, nil, &stop)
		}

		if stop {
			break
		}
	}
}

// FileMD5 return file md5
func FileMD5(file *os.File) (string, error) {
	md5Ctx := md5.New()
	_, err := io.Copy(md5Ctx, file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(md5Ctx.Sum(nil)), nil
}
