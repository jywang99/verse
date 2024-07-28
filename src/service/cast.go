package service

import "jy.org/verse/src/entity"

func GetCasts(gc entity.GetCasts) (*entity.GotCasts, error) {
    casts, err := conn.GetCasts(gc)
    if err != nil {
        return nil, err
    }

    got := &entity.GotCasts{
        Casts: casts,
    }

    if !gc.Pg.GetTotal {
        return got, nil
    }

    count, err := conn.CountCasts(gc)
    if err != nil {
        return nil, err
    }
    got.Pg = entity.GotPaging{
        Total: &count,
    }

    return got, nil
}

func GetCastById(id int) (*entity.GotCast, error) {
    got, err := conn.GetCastById(id)
    if err != nil {
        return nil, err
    }

    return got, nil
}

func GetCastByIds(ids []int) ([]entity.GotCastLite, error) {
    got, err := conn.GetCastsByIds(ids)
    if err != nil {
        return nil, err
    }

    return got, nil
}

