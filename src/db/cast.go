package db

import (
	"context"
	"fmt"

	cs "jy.org/verse/src/constant"
	e "jy.org/verse/src/entity"
)

func (conn *dbConn) GetCastsByIds(ids []int) ([]e.GotCastLite, error) {
    query := fmt.Sprintf("SELECT %s, %s, %s FROM %s WHERE %s IN (%s)", cs.Id, cs.Name, cs.PicPath, cs.CastTable, cs.Id, arrayToString(ids, ","))
    rows, err := conn.pool.Query(context.Background(), query)
    if err != nil {
        logger.ERROR.Println("Error while getting casts: ", err)
        return nil, err
    }
    defer rows.Close()

    var casts []e.GotCastLite
    for rows.Next() {
        var cast e.GotCastLite
        err = rows.Scan(&cast.Id, &cast.Name, &cast.Pic)
        if err != nil {
            logger.ERROR.Println("Error while scanning casts: ", err)
            return nil, err
        }
        casts = append(casts, cast)
    }

    return casts, nil
}

