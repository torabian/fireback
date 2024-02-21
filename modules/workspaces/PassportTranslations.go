package workspaces

type TranslationMap = map[string]map[string]string

var PassportTranslations TranslationMap = TranslationMap{
	PassportMessageCode.UserDoesNotExist: {
		"en": "User you are looking for does not exists, check the details you have entered and, for example if password is correct",
		"fa": "کاربر مورد نظر شما وجود ندارد، مشخصاتی که وارد کرده اید را بررسی کنید و مثلاً اگر رمز عبور صحیح است",
	},
	PassportMessageCode.PassportNotAvailable: {
		"en": "This account name is not available. It might be used by other people or blocked already. If you are the owner, try login instead.",
		"fa": "این نام حساب در دسترس نیست. ممکن است توسط افراد دیگر استفاده شود یا قبلا مسدود شده باشد. اگر شما مالک هستید، به جای آن وارد سیستم شوید.",
	},
}
