package binparser

import (
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
)

type BinParser struct {
	Logger *zap.SugaredLogger
}

func (b *BinParser) castFormat(name string, file []byte) {
	switch strings.ToLower(name) {
	case "areaset":

	}
}

func (b *BinParser) Parse(binPath string) error {
	err := filepath.Walk(binPath, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			b.Logger.Errorw("Error read file", "err", err)

			return err
		}

		return nil
	})

	return err
}
