package service

import (
	"github.com/jackc/pgx/v5"
	e "jy.org/verse/src/entity"
	"jy.org/verse/src/except"
)

func Authenticate(email, password string) (e.UserAuth, error) {
    user, err := conn.GetUserByCreds(email, password)
    if err != nil {
        if err != pgx.ErrNoRows {
            msg := "Error while authenticating user:"
            logger.ERROR.Println(msg, err)
            return e.UserAuth{}, except.NewHandledError(except.DbErr, msg)
        }
        return e.UserAuth{}, except.NewHandledError(except.AuthErr, "Invalid email or password")
    }
    return user, nil
}

func Register(user e.UserRegist) error {
    err, unique := conn.CheckUniqueName(user.Name)
    if err != nil {
        logger.ERROR.Println("Error while checking unique name: ", err)
        return except.NewHandledError(except.DbErr, "Error while checking unique name")
    }
    if !unique {
        return except.NewHandledError(except.AuthErr, "Name already taken")
    }

    err, unique = conn.CheckUniqueEmail(user.Email)
    if err != nil {
        logger.ERROR.Println("Error while checking unique email: ", err)
        return except.NewHandledError(except.DbErr, "Error while checking unique email")
    }
    if !unique {
        return except.NewHandledError(except.AuthErr, "Email already taken")
    }

    err = conn.Register(user)
    if err != nil {
        logger.ERROR.Println("Error while registering user: ", err)
        return  except.NewHandledError(except.DbErr, "Error while registering user")
    }
    return nil
}

