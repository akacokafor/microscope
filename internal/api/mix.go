package api

import (
	"encoding/json"
	"io/fs"
)

//Laravel mix helper
var mixCache map[string]string

func mix(fs fs.FS, path string, isProd bool) string {
	if len(path) == 0 {
		return path
	}

	if !isProd {
		return path[1:]
	}

	if len(mixCache) == 0 {
		mixFile, err := fs.Open("mix-manifest.json")
		if err != nil {
			return path[1:]
		}
		err = json.NewDecoder(mixFile).Decode(&mixCache)
		if err != nil {
			return path[1:]
		}
	}

	if hashedValue, ok := mixCache[path]; ok {
		return hashedValue[1:]
	}

	return path[1:]
}
