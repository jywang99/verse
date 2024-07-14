package service

import (
	"github.com/jackc/pgx/v5"
	e "jy.org/verse/src/entity"
	"jy.org/verse/src/except"
)

func GetEntries(ge e.GetEntries) (e.GotEntries, error) {
    es, err := conn.GetEntries(ge)
    if err != nil {
        logger.ERROR.Println("Error while counting collections: ", err)
        return es, except.NewHandledError(except.DbErr, "Error while getting entries")
    }

    casts, tags, err := getCastsAndTags(es)
    if err != nil {
        return es, err
    }
    es.Casts = casts
    es.Tags = tags

    if !ge.Pg.GetTotal {
        return es, nil
    }

    cnt, err := conn.CountEntries(ge)
    if err != nil {
        logger.ERROR.Println("Error while counting collections: ", err)
        return es, except.NewHandledError(except.DbErr, "Error while counting entries")
    }

    es.Pg = &e.GotPaging{
        Total: &cnt,
    }
    return es, nil
}

func getCastsAndTags(es e.GotEntries) (*[]e.GotCastLite, *[]e.GotTagLite, error) {
    castSet := make(map[int]bool)
    tagSet := make(map[int]bool)

    for _, e := range es.Entries {
        for _, c := range e.CastIds {
            castSet[c] = true
        }
        for _, t := range e.TagIds {
            tagSet[t] = true
        }
    }

    if len(castSet) == 0 && len(tagSet) == 0 {
        return nil, nil, nil
    }

    tagIds := SetToSlice(tagSet)
    castIds := SetToSlice(castSet)

    var casts *[]e.GotCastLite
    if len(castSet) > 0 {
        castVs, err := conn.GetCastsByIds(castIds)
        if err != nil {
            logger.ERROR.Println("Error while getting casts: ", err)
            return nil, nil, except.NewHandledError(except.DbErr, "Error while getting casts")
        }
        casts = &castVs
    }

    var tags *[]e.GotTagLite
    if len(tagSet) > 0 {
        tagVs, err := conn.GetTagsByIds(tagIds)
        if err != nil {
            logger.ERROR.Println("Error while getting tags: ", err)
            return nil, nil, except.NewHandledError(except.DbErr, "Error while getting tags")
        }
        tags = &tagVs
    }

    return casts, tags, nil
}

func GetEntry(id int) (e.GotEntry, error) {
    entry, err := conn.GetEntry(id)
    if err != nil {
        if err == pgx.ErrNoRows {
            return entry, except.NewHandledError(except.NotFoundErr, "Entry not found")
        }
        logger.ERROR.Println("Error while getting entry: ", err)
        return entry, except.NewHandledError(except.DbErr, "Error while getting entry")
    }

    return entry, nil
}
