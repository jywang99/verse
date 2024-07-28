package entity

import (
    cs "jy.org/verse/src/constant"
    "jy.org/verse/src/except"
)

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

func DefaultPaging() Paging {
    return Paging{
        By: cs.Id,
        Desc: false,
        PageSize: 20,
        Offset: 0,
    }
}

func (pg Paging) Validate() error {
    pgSize := pg.PageSize
    if pgSize < 1 || pgSize > 100 {
        return except.NewHandledError(except.BadRequestErr, "Invalid page size")
    }
    if pg.Offset < 0 {
        return except.NewHandledError(except.BadRequestErr, "Invalid offset")
    }
    return nil
}

type GetByIds struct {
    Ids []int `json:"ids"`
}

