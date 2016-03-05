package slob

import (
	"fmt"
	"os"
)

var genParams struct {
	Dir      string
	Prefix   string
	Suffix   string
	FileType string
}

func SetGenParams(dir, prefix, suffix, fileType string) {
	genParams.Dir = dir
	genParams.Prefix = prefix
	genParams.Suffix = suffix
	genParams.FileType = fileType
}

func SetGenDir(dirPath string) {
	genParams.Dir = dirPath
}

func SetGenPrefix(fix string) {
	genParams.Prefix = fix
}

func SetGenSuffix(fix string) {
	genParams.Suffix = fix
}

func SetFileType(fileType string) {
	genParams.FileType = fileType
}

func GetGenDir() string {
	return genParams.Dir
}

func getFilePath(structName string) string {
	fileName := structName
	if genParams.Prefix != "" {
		fileName = fmt.Sprintf("%s_%s", genParams.Prefix, fileName)
	}
	if genParams.Suffix != "" {
		fileName = fmt.Sprintf("%s_%s", fileName, genParams.Suffix)
	}
	if genParams.FileType != "" {
		fileName = fmt.Sprintf("%s.%s", fileName, genParams.FileType)
	}
	if genParams.Dir != "" {
		fileName = fmt.Sprintf("%s/%s", genParams.Dir, fileName)
	}

	return fileName
}

func GenDir() error {
	if genParams.Dir == "" {
		return nil
	}
	return os.MkdirAll(genParams.Dir, 0777)
}

func GenFileHandle(structName string) (*os.File, error) {
	fileName := getFilePath(structName)
	err := GenDir()
	if err != nil {
		return nil, err
	}
	return os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0777)
}
