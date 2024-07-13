package entity

import (
	"slices"
	"time"

	"jy.org/verse/src/constant"
	"jy.org/verse/src/except"
)

type GetCollections struct {
    Pg Paging `json:"paging"`
    Keyword *string `json:"keyword"`
}

func DefaultGetCollections() *GetCollections {
    return &GetCollections{
        Pg: DefaultPaging(),
    }
}

func (gc GetCollections) Validate() error {
    if !slices.Contains(constant.CollectionCols, gc.Pg.By) {
        return except.NewHandledError(except.BadRequestErr, "Invalid sort column")
    }
    return gc.Pg.Validate()
}

type GotCollection struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Desc *string `json:"desc"`
    Created time.Time `json:"created"`
    Parent *int `json:"parent"`
}

type GotPaging struct {
    Total *int `json:"total"`
}

type GotCollections struct {
    Collections []GotCollection `json:"collections"`
    Pg *GotPaging `json:"paging"`
}

