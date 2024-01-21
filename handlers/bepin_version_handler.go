package handlers

type BepinVersionHandler struct {
	ltsBepinVersion    string
	bepinDownloadLinks map[string]BepinVersionInfo
}

type BepinVersionInfo struct {
	version         string
	link            string
	operatingSystem string
}

func InitBepinHandler() BepinVersionHandler {
	return BepinVersionHandler{
		ltsBepinVersion: "5.4.22",
		bepinDownloadLinks: map[string]BepinVersionInfo{
			"5.4.22-win": {
				version:         "5.4.22",
				link:            "https://github.com/BepInEx/BepInEx/releases/download/v5.4.22/BepInEx_x64_5.4.22.0.zip",
				operatingSystem: "win",
			},
			"5.4.22-unix": {
				version:         "5.4.22",
				link:            "https://github.com/BepInEx/BepInEx/releases/download/v5.4.22/BepInEx_unix_5.4.22.0.zip",
				operatingSystem: "unix",
			},
		},
	}
}

func (vh BepinVersionHandler) installBepinEx(bepinVersion BepinVersionInfo) error {
	return nil
}
