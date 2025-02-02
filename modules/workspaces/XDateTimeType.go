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

type XDateTime string

type XDateTimeMetaData struct {
	Formatted *string `json:"formatted,omitempty"`
	Locale    *string `json:"startLocale,omitempty"`
	DaysLeft  *int64  `json:"daysLeft,omitempty"`
}

func (date *XDateTime) MarshalCSV() (string, error) {
	return date.String(), nil
}

type XDateTimeComputed struct {
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

func GetXDateTimeLocaleInfo(startTime time.Time, query QueryDSL) (int64, string, string) {
	DaysLeftToStart := int64(startTime.Sub(time.Now()).Hours() / 24)
	locale := startTime.Format(time.RFC3339)
	StartFormatted := startTime.Format("Monday, 02 Jan 2006")
	switch query.Region {
	case "IR":
		pt := ptime.New(startTime)
		locale = (pt.Format("yyyy-MM-dd"))
		StartFormatted = persian.ToPersianDigits(pt.Format("E dd MMM yy"))
	}

	return DaysLeftToStart, locale, StartFormatted
}

func ComputeXDateTimeMetaData(date *XDate, query QueryDSL) XDateTimeMetaData {
	data := XDateTimeMetaData{}
	endTime, err1 := date.GetTime()

	if err1 == nil {
		count, locale, formatter := GetXDateTimeLocaleInfo(endTime, query)
		data.DaysLeft = &count
		data.Formatted = &formatter
		data.Locale = &locale
	}

	return data

}

func XDateTimeComputeDateRange(start XDateTime, end XDateTime, query QueryDSL) XDateTimeComputed {

	data := XDateTimeComputed{}

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
		count, locale, formatter := GetXDateTimeLocaleInfo(startTime, query)
		data.DaysLeftToStart = &count
		data.StartFormatted = &formatter
		data.StartLocale = &locale
	}

	if err1 == nil {
		count, locale, formatter := GetXDateTimeLocaleInfo(endTime, query)
		data.DaysLeftToEnd = &count
		data.EndFormatted = &formatter
		data.EndLocale = &locale
	}

	return data
}

func (date *XDateTime) GetTime() (time.Time, error) {
	return dateparse.ParseAny(string(*date))
}

func XDateTimeFromString(input string, date *XDateTime) {
	fmt.Println("XDateTimeFromString:", input, date)
	if p, err3 := dateparse.ParseAny(input); err3 == nil {
		*date = XDateTime(p.Format(time.RFC3339))

		// If the year is less than 1500, it means it's an iranian date.
		// note please, if you are working with ancient project this might become
		// a problem. Remove this part of code, if that's an issue
		if p.Year() < 1500 {
			pt := ptime.Date(p.Year(), ptime.Month(p.Month()), p.Day(), 0, 0, 0, 0, ptime.Iran())
			*date = XDateTime(pt.Time().Format(time.RFC3339))
		}
	}
}

func XDateTimeFromTime(p time.Time) *XDateTime {
	utcTime := p.UTC()                              // Ensure time is in UTC
	date := XDateTime(utcTime.Format(time.RFC3339)) // Store as RFC3339 with timezone

	// Handle Iranian calendar conversion if needed
	if p.Year() < 1500 {
		pt := ptime.Date(p.Year(), ptime.Month(p.Month()), p.Day(), p.Hour(), p.Minute(), p.Second(), 0, ptime.Iran())
		date = XDateTime(pt.Time().UTC().Format(time.RFC3339))
	}

	return &date
}

func (date *XDateTime) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	if err == nil {
		*date = XDateTime(nullTime.Time.Format(time.RFC3339))
	} else {
		if v, ok := value.(string); ok {
			XDateTimeFromString(v, date)
		}
	}
	return
}

func (date XDateTime) Value() (driver.Value, error) {
	if date == "" {
		return nil, nil
	}
	return string(date), nil

}

// GormDataType gorm common data type
func (date XDateTime) GormDataType() string {
	return "date"
}

func (date XDateTime) GobEncode() ([]byte, error) {
	return []byte(date), nil
}

func (date *XDateTime) GobDecode(b []byte) error {
	*date = XDateTime(string(b))
	return nil
}

func (j XDateTime) String() string {
	return string(j)
}

func (j XDateTime) MarshalJSON() ([]byte, error) {

	return []byte(`"` + string(j) + `"`), nil
}

func (j *XDateTime) UnmarshalJSON(b []byte) error {

	XDateTimeFromString(strings.ReplaceAll(string(b), "\"", ""), j)

	return nil
}
