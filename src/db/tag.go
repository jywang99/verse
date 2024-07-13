package db

import (
	"context"
	"fmt"

	cs "jy.org/verse/src/constant"
	e "jy.org/verse/src/entity"
)

func (conn *dbConn) GetTagsByIds(ids []int) ([]e.GotTagLite, error) {
    query := fmt.Sprintf("SELECT %s, %s FROM %s WHERE %s IN (%s)", cs.Id, cs.Name, cs.TagTable, cs.Id, arrayToString(ids, ","))
    rows, err := conn.pool.Query(context.Background(), query)
    if err != nil {
        logger.ERROR.Println("Error while getting tags: ", err)
        return nil, err
    }
    defer rows.Close()

    var tags []e.GotTagLite
    for rows.Next() {
        var tag e.GotTagLite
        err = rows.Scan(&tag.Id, &tag.Name)
        if err != nil {
            logger.ERROR.Println("Error while scanning tags: ", err)
            return nil, err
        }
        tags = append(tags, tag)
    }

    return tags, nil
}
