package utils

import (
	"errors"
	"io"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

const (
	TempDir = "temp"
	Files   = TempDir + string(os.PathSeparator) + "stiller"
)

func WriteFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, os.ModePerm)
}

func SaveFile(data []byte, filename string) error {
	if err := WriteFile(filename, data); err != nil {
		return errors.New(color.RedString("error write screenshot to file"))
	}
	return nil
}

func CopyFileToDirectory(pathSourceFile string, pathDestFile string) error {
	sourceFile, err := os.Open(pathSourceFile)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(pathDestFile)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	err = destFile.Sync()
	if err != nil {
		return err
	}

	sourceFileInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}

	destFileInfo, err := destFile.Stat()
	if err != nil {
		return err
	}

	if sourceFileInfo.Size() == destFileInfo.Size() {
	} else {
		return err
	}
	return nil
}
