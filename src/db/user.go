package db

import (
	"context"

	e "jy.org/verse/src/entity"
)

func (conn *dbConn) GetUserByCreds(email, password string) (user e.UserAuth, err error) {
    err = conn.pool.QueryRow(context.Background(), 
        "SELECT id, name, email FROM \"user\" WHERE email = $1 and password = $2", email, password).Scan(
            &user.Id, &user.Name, &user.Email,
        )
    return
}

func (conn *dbConn) CheckUniqueName(name string) (err error, unique bool) {
    var count int
    err = conn.pool.QueryRow(context.Background(), 
        "SELECT count(*) FROM \"user\" WHERE name = $1", name).Scan(&count)
    unique = count == 0
    return
}

func (conn *dbConn) CheckUniqueEmail(email string) (err error, unique bool) {
    var count int
    err = conn.pool.QueryRow(context.Background(), 
        "SELECT count(*) FROM \"user\" WHERE email = $1", email).Scan(&count)
    unique = count == 0
    return
}

func (conn *dbConn) Register(user e.UserRegist) (err error) {
    _, err = conn.pool.Exec(context.Background(), 
        "INSERT INTO \"user\" (name, email, password) VALUES ($1, $2, $3)", 
        user.Name, user.Email, user.Password)
    return
}

