package service

import (
	"os"
	"path/filepath"

	"jy.org/verse/src/service/file"
)

func GetThumb(thumbPath string) (*os.File, string, error) {
    fPath := filepath.Join(cfg.File.ThumbRoot, thumbPath)
    return file.GetStaticFile(fPath)
}

func GetCastPic(thumbPath string) (*os.File, string, error) {
    fPath := filepath.Join(cfg.File.CastRoot, thumbPath)
    return file.GetStaticFile(fPath)
}

