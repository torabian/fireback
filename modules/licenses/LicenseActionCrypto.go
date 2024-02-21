package licenses

import (
	"encoding/json"
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/hyperboloide/lk"
	"github.com/torabian/fireback/modules/workspaces"
)

type LicenseContentPermission struct {
	CapabilityId string `json:"capabilityId"`
}

type LicenseContent struct {
	Email             string                     `json:"email"`
	MachineId         string                     `json:"machineId"`
	WorkspaceId       string                     `json:"workspaceId"`
	Owner             string                     `json:"owner"`
	UserId            string                     `json:"userId"`
	ValidityStartDate time.Time                  `json:"validityStartDate"`
	ValidityEndDate   time.Time                  `json:"validityEndDate"`
	Permissions       []LicenseContentPermission `json:"permissions"`
}

type LicenseConfigurationFile struct {
	PrivateKey string `yaml:"privateKey"`
	PublicKey  string `yaml:"publicKey"`
	License    string `yaml:"license"`
}

// This is for fireback admins only
func GenertePrivatePublicKeySet() (*LicenseConfigurationFile, error) {
	key, err := lk.NewPrivateKey()

	if err != nil {
		return nil, err
	}

	actualKey, err := key.ToB32String()

	if err != nil {
		return nil, err
	}

	pk, err := lk.PrivateKeyFromB32String(actualKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := pk.GetPublicKey()

	config := &LicenseConfigurationFile{
		PrivateKey: actualKey,
		PublicKey:  publicKey.ToB32String(),
	}

	return config, nil

}

func GenerateLicense(doc LicenseContent, privateKeyBase32 string) (string, error) {

	privateKey, err := lk.PrivateKeyFromB32String(privateKeyBase32)
	if err != nil {
		return "", err
	}

	// Define the data you need in your license,
	// here we use a struct that is marshalled to json, but ultimately all you need is a []byte.

	// marshall the document to []bytes (this is the data that our license will contain).
	docBytes, err := json.Marshal(doc)
	if err != nil {
		return "", err
	}

	// generate your license with the private key and the document
	license, err := lk.NewLicense(privateKey, docBytes)
	if err != nil {
		return "", err
	}

	// the b32 representation of our license, this is what you give to your customer.
	licenseB32, err := license.ToB32String()
	if err != nil {
		log.Fatal(err)
	}

	return licenseB32, nil
}

func ValidateLicense() bool {

	invalidLicenseMessage := "Your app is not licensed, you need to activate your fireback by visiting https://pixelplux.com/en/fireback/license\n\nLimitation applies to your instance."
	data := &LicenseConfigurationFile{}
	workspaces.ReadYamlFile("fireback-license.yml", data)

	// A previously generated license b32 encoded. In real life you should read it from a file...
	// const licenseB32 = "FT7YOAYBAEDUY2LDMVXHGZIB76EAAAIDAECEIYLUMEAQUAABAFJAD74EAAAQCUYB76CAAAAABL7YGBIBAL7YMAAAAD73H74IAFEHWITFNVQWS3BCHIRHIZLTORAGK6DBNVYGYZJOMNXW2IRMEJSW4ZBCHIRDEMBRHAWTCMBNGI3FIMJSHIYTSORTGMXDOMBZG43TIMJYHAVTAMR2GAYCE7IBGEBAPXB37ROJCUOYBVG4LAL3MSNKJKPGIKNT564PYK5X542NH62V7TAUEYHGLEOPZHRBAPH7M4SC55OHAEYQEXMKGG3JPO6BSHTDF3T5H6T42VUD7YAJ3TY5AP5MDE5QW4ZYWMSAPEK24HZOUXQ3LJ5YY34XYPVXBUAA===="
	licenseB32 := data.License

	// the public key b32 encoded from the private key using: lkgen pub my_private_key_file`.
	// It should be hardcoded somewhere in your app.
	// const publicKeyBase32 = "ARIVIK3FHZ72ERWX6FQ6Z3SIGHPSMCDBRCONFKQRWSDIUMEEESQULEKQ7J7MZVFZMJDFO6B46237GOZETQ4M2NE32C3UUNOV5EUVE3OIV72F5LQRZ6DFMM6UJPELARG7RLJWKQRATUWD5YT46Q2TKQMPPGIA===="
	publicKeyBase32 := workspaces.UNIQUE_APP_PUBLIC_KEY
	c := color.New(color.FgRed).Add(color.Underline)
	// Unmarshal the public key.
	publicKey, err := lk.PublicKeyFromB32String(publicKeyBase32)
	if err != nil {
		c.Println("The license is wrong, or not has been added. Use `fireback set-license --value '.....' to set the license of the app")
		return false
	}

	// Unmarshal the customer license.
	license, err := lk.LicenseFromB32String(licenseB32)

	if err != nil {
		c.Println("Your license key is invalid")
		c.Println(invalidLicenseMessage)
		return false
	}

	// validate the license signature.
	if ok, err := license.Verify(publicKey); err != nil {
		c.Println(err)
		return false
	} else if !ok {
		c.Println("Your license signature is invalid, cannot be verfied")
		return false
	}

	result := LicenseContent{}

	if err := json.Unmarshal(license.Data, &result); err != nil {
		c.Println("Your license key is invalid")
		c.Println(invalidLicenseMessage)
		return false
	}

	// Now you just have to check that the end date is after time.Now() then you can continue!

	if workspaces.UNIQUE_MACHINE_ID != result.MachineId {
		c.Println("License issued for a different machine with id", result.MachineId)
		return false
	}
	if result.ValidityEndDate.Before(time.Now()) {

		c.Println("License expired on:", result.ValidityEndDate.Format("2006-01-02"), "refresh your license on https://pixelplux.com/en/fireback/license")
		return false
	} else {

		return true
	}
}
