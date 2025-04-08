package abac

var PassportTypes = newPassportFactory()

func newPassportFactory() *passportType {
	return &passportType{
		EmailPassword: "EmailPassword",
		Google:        "Google",
		PhoneNumber:   "PhoneNumber",
	}
}

type passportType struct {
	EmailPassword string
	Google        string
	PhoneNumber   string
}
