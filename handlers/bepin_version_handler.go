package handlers

import (
	"archive/zip"
	"errors"
	"os"
	"path"
	"strings"
)

type BepinVersionHandler struct {
	LtsBepinVersion    string
	BepinDownloadLinks map[string]BepinVersionInfo
}

type BepinVersionInfo struct {
	Id              string
	Name            string
	Version         string
	Link            string
	OperatingSystem string
	FileName        string
}

func InitBepinHandler() BepinVersionHandler {
	return BepinVersionHandler{
		LtsBepinVersion: "5.4.22",
		BepinDownloadLinks: map[string]BepinVersionInfo{
			"5.4.22-win": {
				Id:       "5.4.22-win",
				Name:     "BepInEx x64 - 5.4.22",
				Version:  "5.4.22",
				Link:     "https://github.com/BepInEx/BepInEx/releases/download/v5.4.22/BepInEx_x64_5.4.22.0.zip",
				FileName: "BepInEx_x64_5.4.22.0.zip",
			},
			"5.4.22-unix": {
				Id:       "5.4.22-unix",
				Name:     "BepInEx Unix - 5.4.22",
				Version:  "5.4.22",
				Link:     "https://github.com/BepInEx/BepInEx/releases/download/v5.4.22/BepInEx_unix_5.4.22.0.zip",
				FileName: "BepInEx_unix_5.4.22.0.zip",
			},
		},
	}
}

func (vh BepinVersionHandler) InstallBepinEx(version string, config ConfigHandler, gameDir string) error {

	err := MustDownloadFile(vh.BepinDownloadLinks[version].FileName, "./tmp", vh.BepinDownloadLinks[version].Link)

	if err != nil {
		return err
	}

	bepinZipFile, err := os.Open(path.Join("./tmp", vh.BepinDownloadLinks[version].FileName))

	if err != nil {
		return err
	}

	defer bepinZipFile.Close()

	zipFileSize, err := bepinZipFile.Stat()

	if err != nil {
		return err
	}

	zipReader, err := zip.NewReader(bepinZipFile, int64(zipFileSize.Size()))

	if err != nil {
		return err
	}

	installDir, ok := config.GameDirectories[gameDir]

	if !ok {
		return errors.New("invalid-install-dir")
	}

	for _, f := range zipReader.File {
		splitName := strings.Split(f.Name, "/")

		fileName := splitName[len(splitName)-1:][0]

		if len(splitName) > 1 {
			fileDir := path.Join(splitName[0 : len(splitName)-1]...)
			err := os.MkdirAll(path.Join(installDir, fileDir), os.ModePerm)
			if err != nil {
				return err
			}

			MustCopyZipFile(path.Join(installDir, fileDir, fileName), *f)
		} else {
			MustCopyZipFile(path.Join(installDir, fileName), *f)
		}
	}

	return nil
}
