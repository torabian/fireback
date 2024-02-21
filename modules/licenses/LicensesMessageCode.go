package licenses

var LicensesMessageCode = newLicensesMessageCode()

func newLicensesMessageCode() *LicensesMessageCodeType {
	return &LicensesMessageCodeType{
		PrivateKeyIsMissing: "PrivateKeyIsMissing",
	}
}

type LicensesMessageCodeType struct {
	PrivateKeyIsMissing string
}
