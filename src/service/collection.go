package service

import (
	"jy.org/verse/src/entity"
	"jy.org/verse/src/except"
)

func GetCollections(gc entity.GetCollections) (entity.GotCollections, error) {
    colls, err := conn.GetCollections(gc.Pg)
    if err != nil {
        logger.ERROR.Println("Error while getting collections: ", err)
        return colls, except.NewHandledError(except.DbErr, "Error while getting collections")
    }
    return colls, nil
}

