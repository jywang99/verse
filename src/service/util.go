package service

import "jy.org/verse/src/config"

var cfg = config.Config

func SetToSlice[T comparable](set map[T]bool) []T {
    res := make([]T, 0, len(set))
    for k := range set {
        res = append(res, k)
    }
    return res
}

