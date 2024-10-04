package workspaces

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/mavihq/persian"
	ptime "github.com/yaa110/go-persian-calendar"
)

type XDate string

type XDateMetaData struct {
	Formatted *string `json:"formatted,omitempty"`
	Locale    *string `json:"startLocale,omitempty"`
	DaysLeft  *int64  `json:"daysLeft,omitempty"`
}

func (date *XDate) MarshalCSV() (string, error) {
	return date.String(), nil
}

type XDateComputed struct {
	IsRangePassed   *bool   `json:"isRangePassed,omitempty"`
	IsRangeStarted  *bool   `json:"isRangeStarted,omitempty"`
	InRange         *bool   `json:"inRange,omitempty"`
	RangeLength     *int64  `json:"rangeLength,omitempty"`
	DaysLeftToStart *int64  `json:"daysLeftToStart,omitempty"`
	DaysLeftToEnd   *int64  `json:"daysLeftToEnd,omitempty"`
	StartFormatted  *string `json:"startFormatted,omitempty"`
	StartLocale     *string `json:"startLocale,omitempty"`
	EndLocale       *string `json:"endLocale,omitempty"`
	EndFormatted    *string `json:"endFormatted,omitempty"`
}

func GetDateLocaleInfo(startTime time.Time, query QueryDSL) (int64, string, string) {
	DaysLeftToStart := int64(startTime.Sub(time.Now()).Hours() / 24)
	locale := startTime.Format("2006-01-02")
	StartFormatted := startTime.Format("Monday, 02 Jan 2006")
	switch query.Region {
	case "IR":
		pt := ptime.New(startTime)
		locale = (pt.Format("yyyy-MM-dd"))
		StartFormatted = persian.ToPersianDigits(pt.Format("E dd MMM yy"))
	}

	return DaysLeftToStart, locale, StartFormatted
}

func ComputeXDateMetaData(date *XDate, query QueryDSL) XDateMetaData {
	data := XDateMetaData{}
	endTime, err1 := date.GetTime()

	if err1 == nil {
		count, locale, formatter := GetDateLocaleInfo(endTime, query)
		data.DaysLeft = &count
		data.Formatted = &formatter
		data.Locale = &locale
	}

	return data

}

func ComputeDateRange(start XDate, end XDate, query QueryDSL) XDateComputed {

	data := XDateComputed{}

	endTime, err1 := end.GetTime()
	startTime, err2 := start.GetTime()

	if err1 == nil && err2 == nil {
		m := int64(endTime.Sub(startTime).Hours() / 24)
		data.RangeLength = &m

		if time.Now().After(startTime) && time.Now().Before(endTime) {
			data.InRange = &TRUE
		}
	}

	if err2 == nil {
		count, locale, formatter := GetDateLocaleInfo(startTime, query)
		data.DaysLeftToStart = &count
		data.StartFormatted = &formatter
		data.StartLocale = &locale
	}

	if err1 == nil {
		count, locale, formatter := GetDateLocaleInfo(endTime, query)
		data.DaysLeftToEnd = &count
		data.EndFormatted = &formatter
		data.EndLocale = &locale
	}

	return data
}

func (date *XDate) GetTime() (time.Time, error) {
	return dateparse.ParseAny(string(*date))
}

func FromString(input string, date *XDate) {
	if p, err3 := dateparse.ParseAny(input); err3 == nil {
		*date = XDate(p.Format("2006-01-02"))

		// If the year is less than 1500, it means it's an iranian date.
		// note please, if you are working with ancient project this might become
		// a problem. Remove this part of code, if that's an issue
		if p.Year() < 1500 {
			pt := ptime.Date(p.Year(), ptime.Month(p.Month()), p.Day(), 0, 0, 0, 0, ptime.Iran())
			*date = XDate(pt.Time().Format("2006-01-02"))
		}
	}
}

func (date *XDate) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	if err == nil {
		*date = XDate(nullTime.Time.Format("2006-01-02"))
	} else {
		if v, ok := value.(string); ok {
			FromString(v, date)
		}
	}
	return
}

func (date XDate) Value() (driver.Value, error) {
	if date == "" {
		return nil, nil
	}
	return string(date), nil

}

// GormDataType gorm common data type
func (date XDate) GormDataType() string {
	return "date"
}

func (date XDate) GobEncode() ([]byte, error) {
	return []byte(date), nil
}

func (date *XDate) GobDecode(b []byte) error {
	*date = XDate(string(b))
	return nil
}

func (j XDate) String() string {
	return string(j)
}

func (j XDate) MarshalJSON() ([]byte, error) {

	return []byte(`"` + string(j) + `"`), nil
}

func (j *XDate) UnmarshalJSON(b []byte) error {
	fmt.Println(string(b))
	// fmt.Println(strings.ReplaceAll(string(b), "\"", ""))
	// return j.Scan(strings.ReplaceAll(string(b), "\"", ""))

	FromString(strings.ReplaceAll(string(b), "\"", ""), j)
	// return err
	// return err

	return nil
}
