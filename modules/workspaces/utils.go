package workspaces

import (
	"context"
	"math/rand"
	"reflect"
	"regexp"
	"strings"

	"encoding/hex"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc/metadata"
	"gorm.io/gorm"
)

func UUID() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 4) //equals 8 characters
	rand.Read(b)
	s := hex.EncodeToString(b)

	return s

	// This is the long version
	// u, _ := uuid.NewV4()
	// return u.String()
}

func UUID_Long() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 24) //equals 8 characters
	rand.Read(b)
	s := hex.EncodeToString(b)

	return s
}

func BuildGrpcQuery(in *QueryFilter) QueryDSL {

	if nil == in {
		f := QueryDSL{}
		return f
	}

	f := QueryDSL{
		Query:        in.Query,
		StartIndex:   int(in.StartIndex),
		ItemsPerPage: int(in.ItemsPerPage),
		Deep:         false,
		UniqueId:     in.UniqueId,
	}

	query := ParseAcceptLanguage(in.AcceptLanguage)
	f.AcceptLanguage = query

	if len(query) > 0 {
		f.Language = query[0].Lang
	}

	if f.ItemsPerPage > 50 || f.ItemsPerPage == 0 {
		f.ItemsPerPage = 50
	}

	return f
}

func BuildGrpcQuery2(ctx *context.Context, in *QueryFilter) QueryDSL {
	data, _ := metadata.FromIncomingContext(*ctx)
	fmt.Println(data)

	if nil == in {
		f := QueryDSL{}
		return f
	}

	f := QueryDSL{
		Query:        in.Query,
		StartIndex:   int(in.StartIndex),
		ItemsPerPage: int(in.ItemsPerPage),
		Deep:         false,
		UniqueId:     in.UniqueId,
	}

	query := ParseAcceptLanguage(in.AcceptLanguage)
	f.AcceptLanguage = query

	if len(query) > 0 {
		f.Language = query[0].Lang
	}

	if f.ItemsPerPage > 50 || f.ItemsPerPage == 0 {
		f.ItemsPerPage = 50
	}

	return f
}

func TouchFile(name string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		return err
	}

	return file.Close()
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func CamelCaseToWords(input string) string {
	// Use regular expression to find uppercase letters preceded by lowercase letters
	re := regexp.MustCompile("([a-z])([A-Z])")
	// Replace uppercase letters with space followed by lowercase letter
	output := re.ReplaceAllString(input, "${1} ${2}")
	// Convert the output to lowercase
	output = strings.ToLower(output)
	return output
}

func CamelCaseToWordsDashed(input string) string {
	// Use regular expression to find uppercase letters preceded by lowercase letters
	re := regexp.MustCompile("([a-z])([A-Z])")
	// Replace uppercase letters with space followed by lowercase letter
	output := re.ReplaceAllString(input, "${1}-${2}")
	// Convert the output to lowercase
	output = strings.ToLower(output)
	return output
}

func CamelCaseToWordsUnderlined(input string) string {
	// Use regular expression to find uppercase letters preceded by lowercase letters
	re := regexp.MustCompile("([a-z])([A-Z])")
	// Replace uppercase letters with space followed by lowercase letter
	output := re.ReplaceAllString(input, "${1}_${2}")
	// Convert the output to lowercase
	output = strings.ToLower(output)
	return output
}

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func OsGetDefaultDatabase() string {
	databasePath := "fireback-project-database.db"
	return databasePath
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func rangeIn(low, hi int) int {
	rand.Seed(time.Now().UnixNano())
	return low + rand.Intn(hi-low)
}

// Creates random numeric keygen, with given length
func GenerateRandomKey(length int) string {
	key := ""
	for i := 1; i <= length; i++ {
		key += fmt.Sprint(rangeIn(1, 9))
	}

	return key
}

func PolyglotCreateHandler[T any, P any](dto *T, dtoPolyGlot *P, query QueryDSL) {

	var dbref *gorm.DB = nil
	if query.Tx == nil {
		dbref = GetDbRef()
	} else {
		dbref = query.Tx
	}

	if dto == nil {
		return
	}

	// Detect if it's going to be editing or creation, our action would be different
	linkerId := GetFieldString(dto, "UniqueId")
	if linkerId == "" {
		linkerId = query.UniqueId
	}

	t := map[string]interface{}{}

	v := reflect.ValueOf(dto).Elem()
	for j := 0; j < v.NumField(); j++ {
		n := v.Type().Field(j).Name
		tag := v.Type().Field(j).Tag.Get("translate")
		fieldType := v.Field(j).Type().String()

		if tag == "true" && fieldType == "string" {
			t[ToSnakeCase(n)] = GetFieldString(dto, n)
		} else if tag == "true" && fieldType == "*string" {
			t[ToSnakeCase(n)] = GetFieldStringP(dto, n)
		}
	}

	update := dbref.Model(dtoPolyGlot).Where("linker_id = ? and language_id = ?", linkerId, query.Language).Updates(t)

	if update.Error != nil {
		fmt.Println("Update error", update.Error)
	}

	if update.RowsAffected == 0 {
		t["linker_id"] = linkerId
		t["language_id"] = query.Language
		if linkerId == "" || query.Language == "" {
			return
		}
		dbref.Model(dtoPolyGlot).Create(t)
	}
}
