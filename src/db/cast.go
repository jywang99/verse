package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	cs "jy.org/verse/src/constant"
	e "jy.org/verse/src/entity"
	"jy.org/verse/src/except"
)

func (conn *dbConn) GetCastById(id int) (*e.GotCast, error) {
    query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s WHERE %s=$1", cs.Id, cs.Name, cs.BirthDay, cs.Desc, cs.PicPath, cs.CastTable, cs.Id)

    row := conn.pool.QueryRow(context.Background(), query, id)
    got := e.GotCast{}
    err := row.Scan(&got.Meta.Id, &got.Meta.Name, &got.Birthday, &got.Desc, &got.Meta.Pic)
    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, except.NewHandledError(except.NotFoundErr, "Cast not found")
        }
        logger.ERROR.Println("Error while getting cast: ", err)
        return nil, except.NewHandledError(except.DbErr, "Error while getting cast")
    }

    return &got, nil
}

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
            return nil, except.NewHandledError(except.DbErr, "Error while scanning casts")
        }
        casts = append(casts, cast)
    }

    return casts, nil
}

func (conn *dbConn) GetCasts(gc e.GetCasts) ([]e.GotCastLite, error) {
    query := fmt.Sprintf("SELECT %s, %s, %s FROM %s", cs.Id, cs.Name, cs.PicPath, cs.CastTable)
    where, args := getCastsWhere(gc)
    query += where
    query += getOrderBy(gc.Pg)

    rows, err := conn.pool.Query(context.Background(), query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var casts []e.GotCastLite
    for rows.Next() {
        var cast e.GotCastLite
        err = rows.Scan(&cast.Id, &cast.Name, &cast.Pic)
        if err != nil {
            logger.ERROR.Println("Error while scanning casts: ", err)
            return nil, except.NewHandledError(except.DbErr, "Error while scanning casts")
        }
        casts = append(casts, cast)
    }

    return casts, nil
}

func (conn *dbConn) CountCasts(gc e.GetCasts) (int, error) {
    query := fmt.Sprintf("SELECT COUNT(*) FROM %s", cs.CastTable)
    where, args := getCastsWhere(gc)
    query += where

    row := conn.pool.QueryRow(context.Background(), query, args...)
    var count int
    err := row.Scan(&count)
    if err != nil {
        logger.ERROR.Println("Error while counting casts: ", err)
        return 0, except.NewHandledError(except.DbErr, "Error while counting casts")
    }
    return count, nil
}

func getCastsWhere(gc e.GetCasts) (string, []any) {
    var where string
    var args []interface{}

    if gc.Keyword != nil {
        where += fmt.Sprintf(" WHERE %s ILIKE $1", cs.Name)
        args = append(args, fmt.Sprintf("%%%s%%", *gc.Keyword))
    }

    return where, args
}

