package entity

import (
	"time"

	cs "jy.org/verse/src/constant"
	"jy.org/verse/src/except"
)

type AuthUser struct {
    Id int
    Name string
    Email string
}

type RegistUser struct {
    Name string
    Email string
    Password string
}

type Paging struct {
    By string `json:"by"`
    Desc bool `json:"desc"`
    PageSize int `json:"pageSize"`
    Offset int `json:"offset"`
    GetTotal bool `json:"getTotal"`
}

func (pg Paging) GetDesc() string {
    if pg.Desc {
        return cs.Descend
    }
    return cs.Ascend
}

type GetCollections struct {
    Pg Paging `json:"paging"`
    Keyword *string `json:"keyword"`
}

func DefaultCollectionGet() *GetCollections {
    return &GetCollections{
        Pg: Paging{
            By: cs.Id,
            Desc: false,
            PageSize: 20,
            Offset: 0,
        },
    }
}

func (gc GetCollections) Validate() error {
    pgSize := gc.Pg.PageSize
    if pgSize < 1 || pgSize > 100 {
        return except.NewHandledError(except.BadRequestErr, "Invalid page size")
    }
    if gc.Pg.Offset < 0 {
        return except.NewHandledError(except.BadRequestErr, "Invalid offset")
    }
    return nil
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

type GetEntities struct {
    Pg Paging `json:"paging"`
    GetTotal bool `json:"getTotal"`
}

