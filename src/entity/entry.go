package entity

import (
	"slices"
	"time"

	"jy.org/verse/src/constant"
	"jy.org/verse/src/except"
)

type GetEntries struct {
    Pg Paging `json:"paging"`
    Keyword *string `json:"keyword"`
    ParentIds *[]int `json:"parents"`
    TagIds *[]int `json:"tagIds"`
    CastIds *[]int `json:"castIds"`
}

func DefaultGetEntries() *GetEntries {
    return &GetEntries{
        Pg: DefaultPaging(),
    }
}

func (ge GetEntries) Validate() error {
    if !slices.Contains(constant.CollectionCols, ge.Pg.By) {
        return except.NewHandledError(except.BadRequestErr, "Invalid sort column")
    }
    return ge.Pg.Validate()
}

type GotEntryLite struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Desc *string `json:"desc"`
    CastIds []int `json:"casts"`
    TagIds []int `json:"tags"`
    ParentId *int `json:"parentId"`
    ThumbStatic *string `json:"thumbStatic"`
    ThumbDynamic *[]string `json:"thumbDynamic"`
    Created time.Time `json:"created"`
    Updated time.Time `json:"updated"`
    Aired *time.Time `json:"aired"`
}

type GotEntry struct {
    Meta GotEntryLite `json:"meta"`
    Path string `json:"path"`
    Files []string `json:"files"`
    Tags []GotTagLite `json:"tags"`
    Casts []GotCastLite `json:"casts"`
    Parent *GotCollection `json:"parent"`
}

func NewGotEntry() GotEntry {
    return GotEntry{
        Meta: GotEntryLite{},
        Tags: make([]GotTagLite, 0),
        Casts: make([]GotCastLite, 0),
    }
}

type GotEntries struct {
    Entries []GotEntryLite `json:"entries"`
    Pg *GotPaging `json:"paging"`
    Casts *[]GotCastLite `json:"casts"`
    Tags *[]GotTagLite `json:"tags"`
}

