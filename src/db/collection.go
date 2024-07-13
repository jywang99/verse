package db

import (
	"context"
	"fmt"

	cs "jy.org/verse/src/constant"
	e "jy.org/verse/src/entity"
)

func (conn *dbConn) GetCollections(gc e.GetCollections) (e.GotCollections, error) {
    // select from
    query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s",
        cs.Id, cs.DispName, cs.Desc, cs.Created, cs.Parent, cs.CollectionTable,
    )

    where, args := GetCollectionArgs(gc)
    query += where

    // order by
    query += fmt.Sprintf(" ORDER BY %s %s LIMIT $%d OFFSET $%d", gc.Pg.By, gc.Pg.GetDesc(), len(args)+1, len(args)+2)
    args = append(args, gc.Pg.PageSize, gc.Pg.Offset)

    rows, err := conn.pool.Query(context.Background(), query, args...)
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

func (conn *dbConn) CountCollections(gc e.GetCollections) (int, error) {
    query := fmt.Sprintf("SELECT COUNT(*) FROM %s", cs.CollectionTable)
    where, args := GetCollectionArgs(gc)
    query += where

    row := conn.pool.QueryRow(context.Background(), query, args...)
    var count int
    err := row.Scan(&count)
    if err != nil {
        return 0, err
    }
    return count, nil
}

func GetCollectionArgs(gc e.GetCollections) (string, []any) {
    query := ""
    args := []any{}
    if gc.Keyword != nil {
        query += fmt.Sprintf(" WHERE %s ILIKE $1", cs.DispName)
        args = append(args, fmt.Sprintf("%%%s%%", *gc.Keyword))
    }

    return query, args
}

