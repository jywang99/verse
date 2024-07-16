package entity

type GotCastLite struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Pic *string `json:"pic"`
}

type GetCasts struct {
    Pg Paging `json:"paging"`
    Keyword *string `json:"keyword"`
}

func NewGetCasts() *GetCasts {
    return &GetCasts{
        Pg: DefaultPaging(),
    }
}

func (gc GetCasts) Validate() error {
    return gc.Pg.Validate()
}

type GotCasts struct {
    Casts []GotCastLite `json:"casts"`
    Pg GotPaging `json:"paging"`
}

type GotCast struct {
    Meta GotCastLite `json:"meta"`
    Birthday *string `json:"birthday"`
    Desc *string `json:"desc"`
}

