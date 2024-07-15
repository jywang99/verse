package db

import (
	"context"
	"fmt"
	"strings"

	cs "jy.org/verse/src/constant"
	"jy.org/verse/src/entity"
	e "jy.org/verse/src/entity"
	"jy.org/verse/src/except"
)

func (conn *dbConn) GetTagById(id int) (*e.GotTag, error) {
    query := fmt.Sprintf("SELECT %s, %s, %s FROM %s WHERE %s=$1", cs.Id, cs.Name, cs.Parent, cs.TagTable, cs.Id)

    row := conn.pool.QueryRow(context.Background(), query, id)
    got := e.NewGotTag()
    var parentId *int
    err := row.Scan(&got.Meta.Id, &got.Meta.Name, &parentId)
    if err != nil {
        logger.ERROR.Println("Error while getting tag: ", err)
        return nil, except.NewHandledError(except.DbErr, "Error while getting tag")
    }

    // set parent id
    if parentId != nil {
        got.Parent = &e.GotTagLite{
            Id: *parentId,
        }
    }

    return got, nil
}

func (conn *dbConn) GetTagChildren(id int) ([]e.GotTagLite, error) {
    query := fmt.Sprintf("SELECT %s, %s FROM %s WHERE %s=$1", cs.Id, cs.Name, cs.TagTable, cs.Parent)
    rows, err := conn.pool.Query(context.Background(), query, id)
    if err != nil {
        logger.ERROR.Println("Error while getting tag children: ", err)
        return nil, except.NewHandledError(except.DbErr, "Error while getting tag children")
    }
    defer rows.Close()

    var tags []e.GotTagLite
    for rows.Next() {
        var tag e.GotTagLite
        err = rows.Scan(&tag.Id, &tag.Name)
        if err != nil {
            logger.ERROR.Println("Error while scanning tag children: ", err)
            return nil, except.NewHandledError(except.DbErr, "Error while scanning tag children")
        }
        tags = append(tags, tag)
    }

    return tags, nil
}

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
            return nil, except.NewHandledError(except.DbErr, "Error while scanning tags")
        }
        tags = append(tags, tag)
    }

    return tags, nil
}

func (conn *dbConn) GetTags(gt e.GetTags) (*e.GotTags, error) {
    query := fmt.Sprintf("SELECT %s, %s FROM %s", cs.Id, cs.Name, cs.TagTable)

    where, args := getTagsWhere(gt)
    query += where

    query += getOrderBy(gt.Pg)

    rows, err := conn.pool.Query(context.Background(), query, args...)
    if err != nil {
        logger.ERROR.Println("Error while getting tags: ", err)
        return nil, except.NewHandledError(except.DbErr, "Error while getting tags")
    }

    got := entity.NewGotTags()
    for rows.Next() {
        var tag e.GotTagLite
        err = rows.Scan(&tag.Id, &tag.Name)
        if err != nil {
            logger.ERROR.Println("Error while scanning tags: ", err)
            return nil, except.NewHandledError(except.DbErr, "Error while scanning tags")
        }
        got.Tags = append(got.Tags, tag)
    }

    return got, nil
}

func getTagsWhere(gt e.GetTags) (string, []any) {
    args := []any{}
    stmts := []string{}
    i := 1

    if gt.Keyword != nil {
        stmts = append(stmts, fmt.Sprintf("%s ILIKE $%d", cs.Name, i))
        i++
        args = append(args, fmt.Sprintf("%%%s%%", *gt.Keyword))
    }

    if gt.Parent != nil {
        stmts = append(stmts, fmt.Sprintf("%s=$%d", cs.Parent, i))
        i++
        args = append(args, *gt.Parent)
    }

    where := ""
    if len(stmts) > 0 {
        where = " WHERE " + strings.Join(stmts, " AND ")
    }

    return where, args
}

