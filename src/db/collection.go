package db

import (
	"context"
	"fmt"

	cs "jy.org/verse/src/constant"
	e "jy.org/verse/src/entity"
	"jy.org/verse/src/except"
)

func (conn *dbConn) GetCollections(gc e.GetCollections) (e.GotCollections, error) {
    // select from
    query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s",
        cs.Id, cs.DispName, cs.Desc, cs.Created, cs.Parent, cs.CollectionTable,
    )
    // where
    where, args := getCollectiionWhere(gc)
    query += where
    // order by
    query += getOrderBy(gc.Pg)

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
    where, args := getCollectiionWhere(gc)
    query += where

    row := conn.pool.QueryRow(context.Background(), query, args...)
    var count int
    err := row.Scan(&count)
    if err != nil {
        return 0, err
    }
    return count, nil
}

func getCollectiionWhere(gc e.GetCollections) (string, []any) {
    query := ""
    args := []any{}
    if gc.Keyword != nil {
        query += fmt.Sprintf(" WHERE %s ILIKE $1", cs.DispName)
        args = append(args, fmt.Sprintf("%%%s%%", *gc.Keyword))
    }

    return query, args
}

func (conn *dbConn) GetCollectionsByIds(ids []int) ([]e.GotCollection, error) {
    query := fmt.Sprintf("SELECT %s, %s FROM %s WHERE %s IN (%s)", cs.Id, cs.DispName, cs.CollectionTable, cs.Id, arrayToString(ids, ","))
    rows, err := conn.pool.Query(context.Background(), query)
    if err != nil {
        logger.ERROR.Println("Error while getting collections: ", err)
        return nil, err
    }
    defer rows.Close()

    colls := make([]e.GotCollection, len(ids))
    i := 0
    for rows.Next() {
        var coll e.GotCollection
        err = rows.Scan(&coll.Id, &coll.Name)
        if err != nil {
            logger.ERROR.Println("Error while scanning collections: ", err)
            return nil, except.NewHandledError(except.DbErr, "Error while scanning collections")
        }
        colls[i] = coll
        i ++
    }

    return colls, nil
}

