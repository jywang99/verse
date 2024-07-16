package entity

type GotTagLite struct {
    Id int `json:"id"`
    Name string `json:"name"`
}

type GotTag struct {
    Meta GotTagLite `json:"meta"`
    Parent *GotTagLite `json:"parent"`
    Children []GotTagLite `json:"children"`
}

func NewGotTag() *GotTag {
    return &GotTag{
        Meta: GotTagLite{},
    }
}

type GetTags struct {
    Pg Paging `json:"paging"`
    Ids *[]int `json:"ids"`
    Keyword *string `json:"keyword"`
    Parent *int `json:"parent"`
}

func NewGetTags() *GetTags {
    return &GetTags{
        Pg: DefaultPaging(),
        Keyword: nil,
    }
}

func (gt GetTags) Validate() error {
    return gt.Pg.Validate()
}

type GotTags struct {
    Tags []GotTagLite `json:"tags"`
    Pg GotPaging `json:"paging"`
}

func NewGotTags() *GotTags {
    return &GotTags{
        Tags: make([]GotTagLite, 0),
    }
}

