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

    file, stat, err := file.OpenFile(fPath)
    if err != nil {
        return nil, "", err
    }

    if stat.Size() > cfg.File.MaxStreamSize {
        return nil, "", except.NewHandledError(except.BadRequestErr, "File too large")
    }

    return file, mime, nil
}

