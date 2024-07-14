package db

import (
	"context"
	"fmt"
	"strings"

	cs "jy.org/verse/src/constant"
	e "jy.org/verse/src/entity"
)

func (conn *dbConn) GetEntries(ge e.GetEntries) (e.GotEntries, error) {
    // select from
    query := fmt.Sprintf(`SELECT 
        e.%s, e.%s, e.%s, e.%s, e.%s, e.%s, e.%s, e.%s, e.%s,
        (SELECT array_agg(%s) FROM %s WHERE %s=e.%s) as tag_ids,
        (SELECT array_agg(%s) FROM %s WHERE %s=e.%s) as cast_ids
        FROM %s e`,
        cs.Id, cs.DispName, cs.Desc, cs.ThumbStatic, cs.ThumbDynamic, cs.Created, cs.Updated, cs.Aired, cs.Parent,
        cs.TagId, cs.EntryTagTable, cs.EntryId, cs.Id,
        cs.CastId, cs.EntryCastTable, cs.EntryId, cs.Id, 
        cs.EntryTable,
    )
    // where
    where, args := getEntryWhere(ge)
    query += where
    // order by
    query += getOrderBy(ge.Pg)

    rows, err := conn.pool.Query(context.Background(), query, args...)
    if err != nil {
        return e.GotEntries{}, err
    }

    colls := e.GotEntries{
        Entries: []e.GotEntryLite{},
    }
    for rows.Next() {
        var gel e.GotEntryLite
        err := rows.Scan(&gel.Id, &gel.Name, &gel.Desc, &gel.ThumbStatic, &gel.ThumbDynamic, &gel.Created, &gel.Updated, &gel.Aired, &gel.ParentId, &gel.TagIds, &gel.CastIds)
        if err != nil {
            return e.GotEntries{}, err
        }
        colls.Entries = append(colls.Entries, gel)
    }

    return colls, nil
}

func (conn *dbConn) CountEntries(ge e.GetEntries) (int, error) {
    query := fmt.Sprintf("SELECT COUNT(*) FROM %s e", cs.EntryTable)
    where, args := getEntryWhere(ge)
    query += where

    row := conn.pool.QueryRow(context.Background(), query, args...)
    var count int
    err := row.Scan(&count)
    if err != nil {
        return 0, err
    }
    return count, nil
}

func getEntryWhere(ge e.GetEntries) (string, []any) {
    args := []any{}
    stmts := []string{}

    if ge.Keyword != nil {
        stmts = append(stmts, fmt.Sprintf("(%s ILIKE $1 or %s ILIKE $1)", cs.DispName, cs.Desc))
        args = append(args, fmt.Sprintf("%%%s%%", *ge.Keyword))
    }

    if ge.ParentIds != nil {
        stmts = append(stmts, fmt.Sprintf("e.%s IN (%s)", cs.Parent, arrayToString(*ge.ParentIds, ",")))
    }

    if ge.CastIds != nil {
        stmts = append(stmts, fmt.Sprintf(`e.%s IN (
                SELECT %s
                FROM %s
                WHERE %s IN (%s)
                GROUP BY %s
                HAVING COUNT(DISTINCT %s) = %d
            )`, 
            cs.Id, cs.EntryId, cs.EntryCastTable, cs.CastId,
            arrayToString(*ge.CastIds, ","),
            cs.EntryId, cs.CastId,
            len(*ge.CastIds)),
        )
    }

    if ge.TagIds != nil {
        stmts = append(stmts, fmt.Sprintf(`e.%s IN (
                SELECT %s
                FROM %s
                WHERE %s IN (%s)
                GROUP BY %s
                HAVING COUNT(DISTINCT %s) = %d
            )`, 
            cs.Id, cs.EntryId, cs.EntryTagTable, cs.TagId,
            arrayToString(*ge.TagIds, ","),
            cs.EntryId, cs.TagId,
            len(*ge.TagIds)),
        )
    }

    where := ""
    if len(stmts) > 0 {
        where = " WHERE " + strings.Join(stmts, " AND ")
    }

    return where, args
}

func (conn *dbConn) GetEntry(id int) (e.GotEntry, error) {
    query := fmt.Sprintf(`SELECT 
        e.%s, e.%s, e.%s, e.%s, e.%s, e.%s, e.%s, e.%s, e.%s, e.%s,
        (SELECT array_agg(%s) FROM %s WHERE %s=e.%s) as tag_ids,
        (SELECT array_agg(%s) FROM %s WHERE %s=e.%s) as cast_ids
        FROM %s e WHERE e.%s = %d`,
        cs.Id, cs.DispName, cs.Desc, cs.ThumbStatic, cs.ThumbDynamic, cs.Created, cs.Updated, cs.Aired, cs.Parent, cs.Path,
        cs.TagId, cs.EntryTagTable, cs.EntryId, cs.Id,
        cs.CastId, cs.EntryCastTable, cs.EntryId, cs.Id, 
        cs.EntryTable, cs.Id, id,
    )

    row := conn.pool.QueryRow(context.Background(), query)
    gel := e.NewGotEntry()
    err := row.Scan(&gel.Meta.Id, &gel.Meta.Name, &gel.Meta.Desc, &gel.Meta.ThumbStatic, &gel.Meta.ThumbDynamic, &gel.Meta.Created, &gel.Meta.Updated, &gel.Meta.Aired, &gel.Meta.ParentId, &gel.Path, &gel.Meta.TagIds, &gel.Meta.CastIds)
    if err != nil {
        return e.GotEntry{}, err
    }

    return gel, nil
}

