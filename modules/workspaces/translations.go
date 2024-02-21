package workspaces

var EN = "en"
var FA = "fa"
var PL = "pl"

var BasicTranslations = map[string]map[string]string{
	"NOT_FOUND": {
		EN: "The data you have requeted does not exist. It might be deleted by someone else",
		FA: "اطلاعات مورد نظر شما وجود ندارد. ممکن است توسط شخص دیگری حذف شده باشد",
		PL: "Nie ma takie informacji.",
	},
	"required": {
		EN: "This field is required",
		FA: "این فیلد الزامی است",
		PL: "To pole jest wymagane",
	},
	"VALIDATION_FAILED_ON_SOME_FIELDS": {
		EN: "Validation failed on some fields, check them individually",
		FA: "اعتبارسنجی در برخی از فیلدها انجام نشد، آنها را به صورت جداگانه بررسی کنید",
		PL: "Walidacja niektórych pól nie powiodła się, sprawdź je pojedynczo",
	},
	"NOT_ENOUGH_PERMISSION": {
		EN: "User has not enough permission for this area",
		FA: "کاربر فعلی دسترسی کافی برای این بخش ندارد",
	},
}

type TranslatedString struct {
	En string `yaml:"en"`
	Fa string `yaml:"fa"`
}

func GetTranslationKey(data map[string]interface{}, key string, language string) string {
	fieldTranslations := data["fieldTranslations"].(map[interface{}]interface{})
	return fieldTranslations[key].(map[interface{}]interface{})[language].(string)
}
