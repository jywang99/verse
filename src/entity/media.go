package entity

type GetPartContent struct {
    RangeStart int64
    RangeEnd int64
}

type GotPartContent struct {
    ContentType string
    CRangeStart int64
    CRangeEnd int64
    ContentLength int64
    TotalLength int64
}

