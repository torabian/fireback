package workspaces

import (
	"fmt"
	"time"

	"github.com/mavihq/persian"
	ptime "github.com/yaa110/go-persian-calendar"
)

func FormatDateBasedOnQuery(nsec int64, query QueryDSL) string {

	t := time.Unix(0, nsec)

	if query.Language == "fa" {
		pt := ptime.New(t)
		return persian.ToPersianDigits(fmt.Sprint(pt.Day(), pt.Month(), pt.Year()))

	}
	return t.Local().Format("2006/01/02 15:04:05")
}
