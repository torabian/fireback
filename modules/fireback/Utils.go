package fireback

import (
	"math/rand"
	"reflect"
	"regexp"
	"strings"

	"encoding/hex"
	"fmt"
	"os"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func UUID() string {

	return UUID_NANO()
}

func UUID_NANO() string {
	id, _ := gonanoid.New()

	// I don't want dashes and underlines. Hurts the URL
	id = strings.ReplaceAll(id, "_", "")
	id = strings.ReplaceAll(id, "-", "")

	return id
}

func UUID_SHORT() string {
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

func ToSnakeUpper(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToUpper(snake)
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

// Checks if the underlying value is a nil error.
// Using IError sometimes causes the err to not be nil, but underlying value is nil
func IsErr(err error) bool {
	return err != nil && !reflect.ValueOf(err).IsNil()
}
