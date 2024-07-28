package service

import (
	"jy.org/verse/src/entity"
	"jy.org/verse/src/except"
)

func GetCollection(gc entity.GetCollections) (entity.GotCollections, error) {
    colls, err := conn.GetCollections(gc)
    if err != nil {
        logger.ERROR.Println("Error while getting collections: ", err)
        return colls, except.NewHandledError(except.DbErr, "Error while getting collections")
    }
    if !gc.Pg.GetTotal {
        return colls, nil
    }

    count, err := conn.CountCollections(gc)
    if err != nil {
        logger.ERROR.Println("Error while counting collections: ", err)
        return colls, except.NewHandledError(except.DbErr, "Error while counting collections")
    }
    colls.Pg = &entity.GotPaging{
        Total: &count,
    }

    return colls, nil
}

func GetCollectionsByIds(ids []int) ([]entity.GotCollection, error) {
    got, err := conn.GetCollectionsByIds(ids)
    if err != nil {
        return nil, err
    }

    return got, nil
}

