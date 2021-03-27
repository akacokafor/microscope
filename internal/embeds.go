package internal

import (
	"embed"
	"io/fs"
	"os"

	"github.com/sirupsen/logrus"
)

//go:embed templates/dist
var UITemplates embed.FS

func GetStaticFileSystem(isProduction bool) (fs.FS, error) {
	logrus.Printf("Is Production Current Env: %s\n", isProduction)
	if isProduction {
		uiFileSys, err := fs.Sub(UITemplates, "templates/dist")
		if err != nil {
			return nil, err
		}
		return uiFileSys, nil
	}
	return os.DirFS("internal/templates/dist"), nil
}
