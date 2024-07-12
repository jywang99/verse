package db

import (
	"context"
	"fmt"

	cs "jy.org/verse/src/constant"
	e "jy.org/verse/src/entity"
)

func (conn *dbConn) GetCollections(pg e.Paging) (e.GotCollections, error) {
    query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s ORDER BY $1 %s LIMIT $2 OFFSET $3",
        cs.Id, cs.DispName, cs.Desc, cs.Created, cs.Parent, cs.CollectionTable, pg.GetDesc(),
    )
    rows, err := conn.pool.Query(context.Background(), query, pg.By, pg.PageSize, pg.Offset)
    if err != nil {
        return e.GotCollections{}, err
    }

    colls := e.GotCollections{
        Collections: []e.GotCollection{},
    }
    for rows.Next() {
        var col e.GotCollection
        err := rows.Scan(&col.Id, &col.Name, &col.Desc, &col.Created, &col.Parent)
        if err != nil {
            return e.GotCollections{}, err
        }
        colls.Collections = append(colls.Collections, col)
    }

    return colls, nil
}

