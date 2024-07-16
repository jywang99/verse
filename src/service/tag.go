package service

import "jy.org/verse/src/entity"

func GetTags(gt entity.GetTags) (*entity.GotTags, error) {
    got, err := conn.GetTags(gt)
    if err != nil {
        return nil, err
    }

    if !gt.Pg.GetTotal {
        return got, nil
    }

    count, err := conn.CountTags(gt)
    if err != nil {
        return nil, err
    }
    got.Pg.Total = &count

    return got, nil
}

func GetTagById(id int) (*entity.GotTag, error) {
    got, err := conn.GetTagById(id)
    if err != nil {
        return nil, err
    }

    if got.Parent != nil {
        parent, err := conn.GetTagById(got.Parent.Id)
        if err != nil {
            return nil, err
        }
        got.Parent = &parent.Meta
    }

    children, err := conn.GetTagChildren(id)
    if err != nil {
        return nil, err
    }
    got.Children = children

    return got, nil
}

