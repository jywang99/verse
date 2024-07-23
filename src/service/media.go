package service

import (
	"os"
	"path/filepath"

	"jy.org/verse/src/constant"
	"jy.org/verse/src/except"
	"jy.org/verse/src/service/file"
)

func GetStaticContent(pPath, subPath string) (*os.File, string, error) {
    fPath := filepath.Join(cfg.File.MediaRoot, pPath, subPath)
    mime, err := file.GetMime(fPath)
    if err != nil {
        return nil, "", err
    }

    if mime == constant.Video {
        // video files should be streamed partially
        return nil, "", except.NewHandledError(except.BadRequestErr, "Invalid file type")
    }

    file, _, err := file.OpenFile(fPath)
    if err != nil {
        return nil, "", err
    }

    return file, mime, nil
}

