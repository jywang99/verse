package file

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"jy.org/verse/src/constant"
	"jy.org/verse/src/except"
	"jy.org/verse/src/logging"
)

var logger = logging.Logger

func OpenFile(fPath string) (*os.File, fs.FileInfo, error) {
    stat, err := os.Stat(fPath)
    if err != nil {
        if err == os.ErrNotExist {
            return nil, nil, except.NewHandledError(except.NotFoundErr, "File not found")
        }
        logger.ERROR.Println("Failed to stat file: ", err)
        return nil, nil, errors.New("Unable to open file")
    }

    if stat.IsDir() {
        return nil, nil, except.NewHandledError(except.ForbiddenErr, "Invalid file path")
    }

    file, err := os.Open(fPath)
    if err != nil {
        logger.ERROR.Println("Failed to open file: ", err)
        return nil, nil, errors.New("Unable to open file")
    }

    return file, stat, nil
}

func GetStaticFile(fpath string) (*os.File, string, error) {
    mime, err := GetMime(fpath)
    if err != nil {
        return nil, "", err
    }

    if mime == constant.Video {
        // video files should be streamed partially
        return nil, "", except.NewHandledError(except.BadRequestErr, "Invalid file type")
    }

    file, _, err := OpenFile(fpath)
    if err != nil {
        return nil, "", err
    }

    return file, mime, nil
}

func GetMime(fpath string) (string, error) {
    ext := filepath.Ext(fpath)
    if len(ext) < 2 {
        return "", except.NewHandledError(except.BadRequestErr, "Invalid file type")
    }
    ext = ext[1:]
    ftype, got := constant.FileTypes.GetType(ext)
    if !got {
        return "", except.NewHandledError(except.BadRequestErr, "Invalid file type")
    }
    return fmt.Sprintf("%s/%s", ftype, ext), nil
}

