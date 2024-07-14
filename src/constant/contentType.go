package constant

import (
	"strings"
)

const (
    Video = "video"
    Image = "image"
)

var videoSet = map[string]bool{
    "mp4": true,
    "mkv": true,
    "avi": true,
    "mov": true,
    "flv": true,
    "wmv": true,
    "webm": true,
}

var imageSet = map[string]bool{
    "jpg": true,
    "jpeg": true,
    "png": true,
    "gif": true,
    "bmp": true,
    "webp": true,
}

type filetype struct{
    types map[string]map[string]bool
}

var FileTypes = filetype{
    types: map[string]map[string]bool{
        Video: videoSet,
        Image: imageSet,
    },
}

func (ft filetype) GetType(ext string) (string, bool) {
    ext = strings.ToLower(ext)
    for k, v := range ft.types {
        if v[ext] {
            return k, true
        }
    }
    return "", false
}

