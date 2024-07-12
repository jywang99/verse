package db

import (
	"context"
	"fmt"

	cs "jy.org/verse/src/constant"
	e "jy.org/verse/src/entity"
)

func (conn *dbConn) GetUserByCreds(email, password string) (user e.AuthUser, err error) {
    query := fmt.Sprintf("SELECT %s, %s, %s FROM %s WHERE %s = $1 and %s = $2", cs.Id, cs.Name, cs.Email, cs.UserTable, cs.Email, cs.Password)
    err = conn.pool.QueryRow(context.Background(), query, email, password).Scan(&user.Id, &user.Name, &user.Email)
    return
}

func (conn *dbConn) CheckUniqueName(name string) (err error, unique bool) {
    query := fmt.Sprintf("SELECT count(*) FROM %s WHERE %s = $1", cs.UserTable, cs.Name)
    var count int
    err = conn.pool.QueryRow(context.Background(), query, name).Scan(&count)
    unique = count == 0
    return
}

func (conn *dbConn) CheckUniqueEmail(email string) (err error, unique bool) {
    query := fmt.Sprintf("SELECT count(*) FROM %s WHERE %s = $1", cs.UserTable, cs.Email)
    var count int
    err = conn.pool.QueryRow(context.Background(), query, email).Scan(&count)
    unique = count == 0
    return
}

func (conn *dbConn) Register(user e.RegistUser) (err error) {
    query := fmt.Sprintf("INSERT INTO %s (%s, %s, %s) VALUES ($1, $2, $3)", cs.UserTable, cs.Name, cs.Email, cs.Password)
    _, err = conn.pool.Exec(context.Background(), query, user.Name, user.Email, user.Password)
    return
}

