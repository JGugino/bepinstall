package handlers

import (
	"archive/zip"
	"os"
	"path"
	"strings"
)

type BepinVersionHandler struct {
	ltsBepinVersion    string
	bepinDownloadLinks map[string]BepinVersionInfo
}

type BepinVersionInfo struct {
	version         string
	link            string
	operatingSystem string
	fileName        string
}

func InitBepinHandler() BepinVersionHandler {
	return BepinVersionHandler{
		ltsBepinVersion: "5.4.22",
		bepinDownloadLinks: map[string]BepinVersionInfo{
			"5.4.22-win": {
				version:         "5.4.22",
				link:            "https://github.com/BepInEx/BepInEx/releases/download/v5.4.22/BepInEx_x64_5.4.22.0.zip",
				operatingSystem: "win",
				fileName:        "BepInEx_x64_5.4.22.0.zip",
			},
			"5.4.22-unix": {
				version:         "5.4.22",
				link:            "https://github.com/BepInEx/BepInEx/releases/download/v5.4.22/BepInEx_unix_5.4.22.0.zip",
				operatingSystem: "unix",
				fileName:        "BepInEx_unix_5.4.22.0.zip",
			},
		},
	}
}

func (vh BepinVersionHandler) InstallBepinEx(version string, config ConfigHandler, gameDir string) error {

	err := MustDownloadFile(vh.bepinDownloadLinks[version].fileName, "./tmp", vh.bepinDownloadLinks[version].link)

	if err != nil {
		return err
	}

	bepinZipFile, err := os.Open(path.Join("./tmp", vh.bepinDownloadLinks[version].fileName))

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

	installDir := config.GameDirectories[gameDir]

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
