package service

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"jy.org/verse/src/constant"
	"jy.org/verse/src/except"
)

func SeekVideo(pPath, subPath string) error {
    fPath := filepath.Join(cfg.File.MediaRoot, pPath, subPath)

    stat, err := os.Stat(fPath)
    if err != nil {
        logger.ERROR.Println("Failed to stat file: ", err)
        return errors.New("Unable to open file")
    }
    if stat.IsDir() {
        return except.NewHandledError(except.ForbiddenErr, "Invalid file path")
    }

    // TODO partial content
    
    return nil
}

func GetStaticContent(pPath, subPath string) (*os.File, string, error) {
    fPath := filepath.Join(cfg.File.MediaRoot, pPath, subPath)
    ext := filepath.Ext(subPath)[1:]
    logger.INFO.Println("GetStaticContent: ", fPath, ext)

    // content type
    ftype, got := constant.FileTypes.GetType(ext)
    logger.INFO.Println("GetStaticContent: ", ftype, got)
    if !got {
        return nil, "", except.NewHandledError(except.BadRequestErr, "Invalid file type")
    }
    if ftype == constant.Video {
        // video files should be streamed partially
        return nil, "", except.NewHandledError(except.BadRequestErr, "Invalid file type")
    }

    // open file
    stat, err := os.Stat(fPath)
    if err != nil {
        logger.ERROR.Println("Failed to stat file: ", err)
        return nil, "", errors.New("Unable to open file")
    }
    if stat.IsDir() {
        return nil, "", except.NewHandledError(except.ForbiddenErr, "Invalid file path")
    }
    file, err := os.Open(fPath)
    if err != nil {
        logger.ERROR.Println("Failed to open file: ", err)
        return nil, "", errors.New("Unable to open file")
    }

    return file, fmt.Sprintf("%s/%s", ftype, ext), nil
}

