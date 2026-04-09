package fireback

import (
	"fmt"
	"log"
	"os"
	reflect "reflect"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func structToEnvMap(config interface{}) (map[string]string, error) {
	envMap := make(map[string]string)
	val := reflect.ValueOf(config).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		envKey := field.Tag.Get("envconfig")

		if envKey != "" {
			value := val.Field(i).Interface()

			switch v := value.(type) {
			case string:
				envMap[envKey] = v
			case bool:
				envMap[envKey] = strconv.FormatBool(v)
			case int, int64:
				envMap[envKey] = strconv.FormatInt(reflect.ValueOf(v).Int(), 10)
			case float64:
				envMap[envKey] = strconv.FormatFloat(reflect.ValueOf(v).Float(), 'f', -1, 64)
			default:
				return nil, fmt.Errorf("unsupported type: %s", reflect.TypeOf(v))
			}
		}
	}

	return envMap, nil
}

func SaveEnvFile(config interface{}, filename string) error {

	envMap, err := structToEnvMap(config)

	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for key, value := range envMap {
		if value == "" {
			continue
		}

		_, err := file.WriteString(fmt.Sprintf("%s=%s\n", key, value))
		if err != nil {
			return err
		}
	}

	return nil
}

func HandleEnvVars(spec interface{}) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	filename := ".env." + env
	err := godotenv.Load(filename)
	if err != nil {
		log.Printf("environment variable file expected: %s was not loaded. Error: %v", filename, err)
	}

	envconfig.MustProcess("", spec)
}
