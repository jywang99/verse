package db

import (
	"fmt"
	"strings"

	e "jy.org/verse/src/entity"
)

func arrayToString(a []int, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func getOrderBy(pg e.Paging) string {
    return fmt.Sprintf(" ORDER BY %s %s LIMIT %d OFFSET %d", pg.By, pg.GetDesc(), pg.PageSize, pg.Offset)
}

