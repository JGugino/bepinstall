package handlers

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func MustDownloadFile(fileName string, downloadPath string, downloadURL string) error {

	newFile, err := os.Create(path.Join(downloadPath, fileName))

	if err != nil {
		panic(err)
	}

	defer newFile.Close()

	downloadResp, err := http.Get(downloadURL)

	if err != nil {
		panic(err)
	}

	defer downloadResp.Body.Close()

	if downloadResp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: %s", downloadResp.Status)
	}

	_, err = io.Copy(newFile, downloadResp.Body)

	if err != nil {
		panic(err)
	}

	return nil
}

func MustCopyZipFile(copyLocation string, file zip.File) {
	currentFile, err := file.Open()
	if err != nil {
		panic(err)
	}

	newFile, err := os.Create(copyLocation)

	if err != nil {
		panic(err)
	}

	_, err = io.Copy(newFile, currentFile)

	if err != nil {
		panic(err)
	}

	defer currentFile.Close()
	defer newFile.Close()
}
