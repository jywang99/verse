package service

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"jy.org/verse/src/constant"
	e "jy.org/verse/src/entity"
	"jy.org/verse/src/except"
)

func SeekVideo(pPath, subPath string, get e.GetPartContent) (*e.GotPartContent, *[]byte, error) {
    fPath := filepath.Join(cfg.File.MediaRoot, pPath, subPath)
    mime, err := getMime(fPath)
    if err != nil {
        return nil, nil, err
    }

    stat, err := os.Stat(fPath)
    if err != nil {
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
	defer file.Close()

    fileSize := stat.Size()
    // default rangeEnd
    if get.RangeEnd == 0 || get.RangeEnd > fileSize {
        get.RangeEnd = fileSize
    }
    // respect max length
    maxLen := cfg.File.MaxStreamSize
    l := get.RangeEnd - get.RangeStart + 1
    if l > maxLen || l < 0 {
        get.RangeEnd = get.RangeStart + maxLen - 1
    }

    // seek
    _, err = file.Seek(get.RangeStart, 0)
	if err != nil {
        logger.ERROR.Println("Failed to seek file: ", err)
        return nil, nil, errors.New("Unable to seek file")
	}

    buf := make([]byte, get.RangeEnd-get.RangeStart+1)
    _, err = file.Read(buf)
    if err != nil {
        logger.ERROR.Println("Failed to read file: ", err)
        return nil, nil, errors.New("Unable to read file")
    }
    
    // headers
    got := &e.GotPartContent{
        ContentType: mime,
        ContentLength: get.RangeEnd - get.RangeStart + 1,
        CRangeStart: get.RangeStart,
        CRangeEnd: get.RangeEnd,
        TotalLength: fileSize,
    }

    return got, &buf, nil
}

func GetStaticContent(pPath, subPath string) (*os.File, string, error) {
    fPath := filepath.Join(cfg.File.MediaRoot, pPath, subPath)
    mime, err := getMime(fPath)
    if err != nil {
        return nil, "", err
    }

    if mime == constant.Video {
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

    return file, mime, nil
}

func getMime(fpath string) (string, error) {
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

