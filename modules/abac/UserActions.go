package abac

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/torabian/fireback/modules/fireback/complexes"
)

// user had permRewrite root.modules -> root.manage, and security: { writeOnRoot: true }
// in the old yaml: the old generated code set AllowOnRoot: true on Create/Update/Delete
// only, leaving Query/Get plain - preserved here exactly. The old yaml's rpc.query.qs
// (withImages) and events (Googoli2) blocks had zero real readers anywhere in the
// codebase, so neither needed an extension-file implementation.
var userPerms = NewCrudPermissionSet("root.manage", "user", "user")
var PERM_ROOT_USER = userPerms.Wildcard
var PERM_ROOT_USER_QUERY = userPerms.Query
var PERM_ROOT_USER_CREATE = userPerms.Create
var PERM_ROOT_USER_UPDATE = userPerms.Update
var PERM_ROOT_USER_DELETE = userPerms.Delete
var ALL_USER_PERMISSIONS = userPerms.All

var UserActions = NewEntityActionsBundle[abacdefs.UserEntity]()

func FullName(x *abacdefs.UserEntity) string {

	full := ""

	if x.FirstName != "" {
		full += x.FirstName
	}

	if x.LastName != "" {
		full += " " + x.LastName
	}

	return full

}

// Unlike almost every other entity in this package, abacdefs.UserEntity is NOT
// workspace-scoped (users belong to the system as a whole, root-only management, no
// workspace_id column at all) - so, same as CapabilityActions.go's
// GetCapabilitiesAction/CapabilityGetAction, this calls the entity's own generated,
// unscoped abacdefs.UserEntityActions.Browse directly instead of going through
// EntityActionsBundle/fireback.QueryEntitiesPointer, which unconditionally filters by
// workspace_id and fails outright with "column workspace_id does not exist".
func UserBrowseAction(c abacdefs.UserBrowseActionRequest) (*abacdefs.UserBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_USER_QUERY}})
	if err != nil {
		return nil, err
	}
	qs := abacdefs.UserBrowseActionQueryFromString(c.QueryParams.Encode())
	items, qrm, err2 := abacdefs.UserEntityActions.Browse(fireback.GetDbRef(), qs, "")
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	for _, item := range items {
		HydrateUserPrimaryAddress(item)
	}
	return &abacdefs.UserBrowseActionResponse{Payload: fireback.GResponseQuery(items, &fireback.QueryResultMeta{
		TotalItems: qrm.TotalItems,
		Cursor:     qrm.Cursor,
	}, query)}, nil
}

// See UserBrowseAction - not workspace-scoped, so this calls the entity's own generated,
// unscoped abacdefs.UserEntityActions.Get directly.
func UserGetAction(c abacdefs.UserGetActionRequest) (*abacdefs.UserGetActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_USER_QUERY}}); err != nil {
		return nil, err
	}
	item, err2 := abacdefs.UserEntityActions.Get(fireback.GetDbRef(), c.Params.UniqueId)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	HydrateUserPrimaryAddress(item)
	return &abacdefs.UserGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func UserCreateAction(c abacdefs.UserCreateActionRequest) (*abacdefs.UserCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_USER_CREATE}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	// firstName/lastName are validate:"required" (see UserDto.go) but nothing was
	// actually enforcing it - same fix as
	// EmailProviderCreateAction/TableViewSizingCreateAction/PassportCreateAction.
	if err2 := fireback.CommonStructValidatorPointer(&c.Body, false); err2 != nil {
		return nil, err2
	}
	entity := &abacdefs.UserEntity{
		FirstName:     c.Body.FirstName,
		LastName:      c.Body.LastName,
		Photo:         c.Body.Photo,
		Gender:        c.Body.Gender,
		Title:         c.Body.Title,
		BirthDate:     c.Body.BirthDate,
		Avatar:        c.Body.Avatar,
		LastIpAddress: c.Body.LastIpAddress,
		PhoneNumber:   c.Body.PhoneNumber,
		JobTitle:      c.Body.JobTitle,
		Company:       c.Body.Company,
		Bio:           c.Body.Bio,
	}
	if v, ok := c.Body.PrimaryAddress.Get(); ok && v != nil {
		entity.PrimaryAddress = emigo.NullableOf(abacdefs.UserEntityPrimaryAddress{
			AddressLine1:    v.AddressLine1,
			AddressLine2:    v.AddressLine2,
			City:            v.City,
			StateOrProvince: v.StateOrProvince,
			PostalCode:      v.PostalCode,
			CountryCode:     v.CountryCode,
		})
	}
	created, err2 := UserActionCreate(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.UserCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

// See UserBrowseAction - not workspace-scoped, so this calls the entity's own generated,
// unscoped abacdefs.UserEntityActions.Update directly. UserEntityUpdateFn takes the
// UserOptionalDto request body straight (unlike UserEntityCreateFn's *UserEntity), so
// there's no manual field-by-field mapping to a *UserEntity needed here at all - see
// UserEntityUpdateFn's own body (UserEntity.go) for how it applies primaryAddress's
// sub-fields directly onto the embedded address_line1/city/... columns.
func UserUpdateAction(c abacdefs.UserUpdateActionRequest) (*abacdefs.UserUpdateActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_USER_UPDATE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	updated, err2 := abacdefs.UserEntityActions.Update(fireback.GetDbRef(), c.Params.UniqueId, c.Body)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	HydrateUserPrimaryAddress(updated)
	return &abacdefs.UserUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func UserAwareDeletePreviewAction(c abacdefs.UserAwareDeletePreviewActionRequest) (*abacdefs.UserAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_USER_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	uniqueIds := abacdefs.UserAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := abacdefs.UserEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.UserAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func UserAwareDeleteAction(c abacdefs.UserAwareDeleteActionRequest) (*abacdefs.UserAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_USER_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	if err2 := abacdefs.UserEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.UserAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

// --- Hand business logic recovered from the pre-migration UserActions.go/abacdefs.UserEntity.go ---

/*
This file holds in memory some temporary solution to tokens that need to be used upon redirection
They can be stored, and accessed only once, so third party can get the token in the other domain
*/

type exchangeItem struct {
	token string
}

var exchangePool map[string]exchangeItem = map[string]exchangeItem{}

func PutTokenInExchangePool(token string) string {
	uniqueId := fireback.UUID()

	exchangePool[uniqueId] = exchangeItem{
		token: token,
	}

	return uniqueId
}

func GetTokenFromExchangePoolAction(query fireback.QueryDSL) (*abacdefs.ExchangeKeyInformationDto, *fireback.IError) {
	item := exchangePool[query.UniqueId]
	token := item.token
	delete(exchangePool, query.UniqueId)

	if token == "" {
		return nil, fireback.Create401Error(&AbacMessages.InvalidExchangeKey, []string{})
	}

	return &abacdefs.ExchangeKeyInformationDto{Key: token}, nil
}

func GetUserFromToken(tokenString string) (*abacdefs.UserEntity, error) {

	var item abacdefs.TokenEntity

	if err := fireback.GetDbRef().Where(fireback.RealEscape("token = ?", tokenString)).First(&item).Error; err != nil {
		return &abacdefs.UserEntity{}, err
	}

	// Not workspace-scoped (see UserBrowseAction's own comment) - abacdefs.UserEntityActions.Get
	// is the entity's own generated, unscoped lookup.
	user, _ := abacdefs.UserEntityActions.Get(fireback.GetDbRef(), item.UserId.OrDefault(""))
	HydrateUserPrimaryAddress(user)
	return user, nil
}

// UserActionCreate is the shared create path for every way a UserEntity row comes into
// existence - not just UserCreateAction's own admin endpoint, but the signup/OS-login/
// OAuth path too (see UnsafeGenerateUser) and UserMockActionCliHandler (UserCli.go). Not
// workspace-scoped (see UserBrowseAction's own comment) - abacdefs.UserEntityActions.Create
// is the entity's own generated, unscoped insert, and (unlike the old
// EntityActionsBundle.Create this used to go through) actually copies dto.PrimaryAddress
// onto the embedded PrimaryAddressRow columns before inserting - see UserEntityCreateFn,
// UserEntity.go.
func UserActionCreate(
	dto *abacdefs.UserEntity, query fireback.QueryDSL,
) (*abacdefs.UserEntity, *fireback.IError) {
	created, err := abacdefs.UserEntityActions.Create(fireback.GetDbRef(), dto)
	if err != nil {
		return nil, fireback.GormErrorToIError(err)
	}
	HydrateUserPrimaryAddress(created)
	return created, nil
}

func getRandomAvatarURL() string {
	rand.Seed(time.Now().UnixNano()) // Seed to ensure randomness
	randomNum := rand.Intn(20) + 1   // Generate number between 1 and 20
	return fmt.Sprintf("https://cdn.jsdelivr.net/gh/alohe/avatars/png/vibrent_%d.png", randomNum)
}

func randomZeroOrOne() int {
	rand.Seed(time.Now().UnixNano()) // Ensure randomness
	return rand.Intn(2)              // Generates 0 or 1
}

var firstNames = []string{
	"Ali", "Mohammad", "John", "David", "Maria", "Sarah", "Hassan", "Omar", "James", "Robert",
	"Emily", "Sophia", "Daniel", "Michael", "Jessica", "Olivia", "Amir", "Reza", "Alex", "Emma",
	"Chris", "Elena", "Noah", "Liam", "Ethan", "Mason", "Lucas", "Henry", "Nathan", "Jack",
	"Isabella", "Charlotte", "Mia", "Layla", "Ava", "Ella", "Benjamin", "Jacob", "Matthew", "Sofia",
	"Zahra", "Fatima", "Amin", "Mehdi", "Tomas", "Victor", "Leon", "Julian", "Max", "Leo",
	"Arthur", "Elias", "Hugo", "Theo", "Oscar", "Gabriel", "William", "Daniela", "Samuel", "Adam",
	"Alexander", "Freddie", "Edward", "Joseph", "Harry", "Charlie", "Sebastian", "Ryan", "Evelyn", "Anna",
	"Adrian", "Diego", "Mateo", "Dylan", "Jason", "Carter", "Ezra", "Milo", "Jasper", "Axel",
	"Leonardo", "Caleb", "Hunter", "Isaiah", "Andrew", "Cooper", "Nathaniel", "Elliot", "Brody", "Parker",
	"Sadie", "Ruby", "Violet", "Luna", "Clara", "Madeline", "Stella", "Nora", "Lily", "Hazel",
}

var lastNames = []string{
	"Torabi", "Johnson", "Smith", "Williams", "Brown", "Taylor", "Anderson", "Thomas", "Jackson", "White",
	"Harris", "Martin", "Thompson", "Garcia", "Martinez", "Robinson", "Clark", "Rodriguez", "Lewis", "Lee",
	"Walker", "Hall", "Allen", "Young", "King", "Wright", "Lopez", "Hill", "Scott", "Green",
	"Adams", "Baker", "Gonzalez", "Nelson", "Carter", "Mitchell", "Perez", "Roberts", "Turner", "Phillips",
	"Campbell", "Parker", "Evans", "Edwards", "Collins", "Stewart", "Sanchez", "Morris", "Rogers", "Reed",
	"Cook", "Morgan", "Bell", "Murphy", "Bailey", "Rivera", "Cooper", "Richardson", "Cox", "Howard",
	"Ward", "Torres", "Peterson", "Gray", "Ramirez", "James", "Watson", "Brooks", "Kelly", "Sanders",
	"Price", "Bennett", "Wood", "Barnes", "Ross", "Henderson", "Coleman", "Jenkins", "Perry", "Powell",
	"Long", "Patterson", "Hughes", "Flores", "Washington", "Butler", "Simmons", "Foster", "Gonzales", "Bryant",
	"Alexander", "Russell", "Griffin", "Diaz", "Hayes", "Myers", "Ford", "Hamilton", "Graham", "Sullivan",
}

func getRandomName(names []string) string {
	return names[rand.Intn(len(names))]
}

func getRandomBirthDate() string {
	// Get the current year
	currentYear := time.Now().Year()

	// Generate a random number of years between 10 and 20
	yearsAgo := rand.Intn(11) + 10 // Random number between 10 and 20

	// Calculate the year for the birthdate
	birthYear := currentYear - yearsAgo

	// Randomly select a month (1 to 12)
	birthMonth := rand.Intn(12) + 1

	// Randomly select a day (1 to 28, for simplicity)
	birthDay := rand.Intn(28) + 1

	// Format the birthdate as YYYY-MM-DD
	return time.Date(birthYear, time.Month(birthMonth), birthDay, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
}

func randomPublicIP() string {
	rand.Seed(time.Now().UnixNano())

	for {
		a := rand.Intn(256)
		b := rand.Intn(256)
		c := rand.Intn(256)
		d := rand.Intn(256)

		// Skip reserved/private ranges
		if a == 10 || // 10.0.0.0/8
			(a == 172 && b >= 16 && b <= 31) || // 172.16.0.0 – 172.31.255.255
			(a == 192 && b == 168) || // 192.168.0.0/16
			a == 127 || // loopback
			a >= 224 { // multicast/reserved
			continue
		}
		return fmt.Sprintf("%d.%d.%d.%d", a, b, c, d)
	}
}

type sampleAddress struct {
	CountryCode, City, State, Address1, Address2, Postcode string
}

var addresses = []sampleAddress{
	{"US", "Springfield", "IL", "742 Evergreen Terrace", "Apt 4B", "62704"},
	{"DE", "Berlin", "Berlin", "Musterstraße 12", "EG", "10115"},
	{"IR", "تهران", "تهران", "خیابان ولیعصر، پلاک ۲۳", "طبقه سوم", "1599616313"},
	{"PL", "Warszawa", "Mazowieckie", "ul. Długa 45", "mieszkanie 12", "00-238"},
	{"FR", "Paris", "Île-de-France", "10 Rue de Rivoli", "5ème étage", "75001"},
	{"UK", "London", "Greater London", "221B Baker Street", "Flat 2", "NW1 6XE"},
	{"IT", "Rome", "Lazio", "Via Nazionale 75", "Scala B", "00184"},
	{"ES", "Madrid", "Community of Madrid", "Calle Mayor 3", "Piso 1", "28013"},
	{"CA", "Toronto", "Ontario", "123 Queen St W", "Unit 1502", "M5H 2M9"},
	{"AU", "Sydney", "NSW", "88 George Street", "Suite 7", "2000"},
	{"IN", "Mumbai", "Maharashtra", "12 Linking Road", "Flat 501", "400050"},
	{"RU", "Moscow", "Moscow", "ул. Тверская, д. 7", "кв. 23", "125009"},
	{"CN", "Beijing", "Beijing", "东直门南大街 5号", "三层", "100007"},
	{"JP", "Tokyo", "Tokyo", "1-2-3 Shibuya", "Apt 301", "150-0002"},
	{"BR", "São Paulo", "SP", "Av. Paulista, 1000", "Ap 102", "01310-100"},
	{"MX", "Mexico City", "CDMX", "Av. Reforma 222", "Depto 33", "06600"},
	{"AR", "Buenos Aires", "CABA", "Calle Florida 100", "Piso 2", "1005"},
	{"TR", "Istanbul", "Istanbul", "İstiklal Caddesi 56", "Kat 4", "34433"},
	{"NL", "Amsterdam", "North Holland", "Damrak 89", "2nd Floor", "1012 LP"},
	{"SE", "Stockholm", "Stockholm", "Drottninggatan 50", "Lgh 1101", "11121"},
	{"NO", "Oslo", "Oslo", "Karl Johans gate 15", "Etasje 3", "0159"},
	{"FI", "Helsinki", "Uusimaa", "Mannerheimintie 10", "Asunto 2A", "00100"},
	{"DK", "Copenhagen", "Capital Region", "Strøget 20", "2. sal", "1154"},
	{"CH", "Zurich", "Zurich", "Bahnhofstrasse 10", "Stock 3", "8001"},
	{"BE", "Brussels", "Brussels-Capital", "Rue Neuve 15", "Etage 2", "1000"},
	{"AT", "Vienna", "Vienna", "Mariahilfer Str. 99", "Top 6", "1060"},
	{"GR", "Athens", "Attica", "Ermou 20", "2ος Όροφος", "10563"},
	{"PT", "Lisbon", "Lisbon", "Avenida da Liberdade 144", "Apartamento 5D", "1250-146"},
	{"RO", "Bucharest", "Bucharest", "Strada Lipscani 35", "Etaj 1", "030036"},
	{"BG", "Sofia", "Sofia", "Vitosha Blvd 18", "Ap. 12", "1000"},
	{"HU", "Budapest", "Budapest", "Andrássy út 45", "2nd floor", "1061"},
	{"CZ", "Prague", "Prague", "Wenceslas Square 1", "Suite 4", "110 00"},
	{"SK", "Bratislava", "Bratislava", "Obchodná 12", "Byt 6", "811 06"},
	{"HR", "Zagreb", "Zagreb", "Ilica 50", "Kat 1", "10000"},
	{"SI", "Ljubljana", "Ljubljana", "Slovenska cesta 25", "Nadstropje 3", "1000"},
	{"EE", "Tallinn", "Harju", "Pikk 23", "Korter 5", "10133"},
	{"LV", "Riga", "Riga", "Brīvības iela 100", "Dzīvoklis 8", "LV-1011"},
	{"LT", "Vilnius", "Vilnius", "Gedimino pr. 9", "Butas 2", "01103"},
	{"UA", "Kyiv", "Kyiv", "Khreshchatyk St, 22", "kv 10", "01001"},
	{"RS", "Belgrade", "Belgrade", "Knez Mihailova 14", "Sprat 3", "11000"},
	{"BA", "Sarajevo", "Sarajevo", "Ferhadija 12", "Stan 7", "71000"},
	{"MK", "Skopje", "Skopje", "Makedonija Str. 5", "Apartment 11", "1000"},
	{"AL", "Tirana", "Tirana", "Rruga Myslym Shyri 77", "Kati 2", "1001"},
	{"GE", "Tbilisi", "Tbilisi", "Rustaveli Ave 40", "Apt 9", "0108"},
	{"AM", "Yerevan", "Yerevan", "Abovyan St 22", "Flat 4", "0001"},
	{"KZ", "Almaty", "Almaty", "Dostyk Ave 34", "Kv 16", "050010"},
	{"AZ", "Baku", "Baku", "Nizami St 78", "Mənzil 5", "AZ1000"},
	{"SA", "Riyadh", "Riyadh", "Olaya St 234", "Floor 6", "12211"},
}

func RandomUserPrimaryAddress() *abacdefs.UserEntityPrimaryAddress {
	rand.Seed(time.Now().UnixNano())
	s := addresses[rand.Intn(len(addresses))]
	return &abacdefs.UserEntityPrimaryAddress{
		AddressLine1:    (s.Address1),
		AddressLine2:    emigo.NullableOf(s.Address2),
		City:            emigo.NullableOf(s.City),
		StateOrProvince: emigo.NullableOf(s.State),
		PostalCode:      emigo.NullableOf(s.Postcode),
		CountryCode:     emigo.NullableOf(s.CountryCode),
	}
}

// jobTitles/companies/bios back both the "user mock" cli command (UserMockActionCliHandler
// in UserCli.go) and the seeders framework (UserActions.SeederInit below) with realistic
// values for the profile detail fields (phoneNumber/jobTitle/company/bio) added alongside
// primaryAddress.
var jobTitles = []string{
	"Support Engineer", "Product Manager", "Sales Representative", "Software Engineer",
	"Data Analyst", "Marketing Specialist", "Customer Success Manager", "HR Coordinator",
	"Operations Manager", "QA Engineer", "UX Designer", "Financial Analyst",
	"Account Executive", "DevOps Engineer", "Business Analyst", "Office Administrator",
	"Project Manager", "Solutions Architect", "Content Writer", "Recruiter",
}

var companies = []string{
	"Acme Corp", "Globex Industries", "Initech", "Umbrella Logistics", "Hooli",
	"Stark Enterprises", "Wayne Technologies", "Wonka Foods", "Cyberdyne Systems",
	"Soylent Corp", "Vandelay Industries", "Pied Piper", "Massive Dynamic",
	"Aperture Science", "Gringotts Financial", "Northwind Traders", "Contoso Ltd",
	"Fabrikam Inc", "Tyrell Corporation", "Oceanic Airlines",
}

var bios = []string{
	"Enjoys solving problems and helping teammates ship better software.",
	"Coffee enthusiast, weekend hiker, and long-time open source contributor.",
	"Focused on making customers happy, one ticket at a time.",
	"Loves clean data, clear dashboards, and even clearer deadlines.",
	"Believes good documentation is a love letter to your future self.",
	"Spends most weekends cycling and most weekdays debugging.",
	"Passionate about design systems that actually get used.",
	"Ex-teacher turned engineer, still can't stop explaining things.",
	"Runs marathons for fun, runs migrations for a living.",
	"Firm believer that every bug is a feature request in disguise.",
}

func getRandomPhoneNumber() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("+1-%03d-%03d-%04d", rand.Intn(800)+200, rand.Intn(800)+200, rand.Intn(10000))
}

func init() {

	UserActions.SeederInit = func() *abacdefs.UserEntity {
		return &abacdefs.UserEntity{
			FirstName:      getRandomName(firstNames),
			LastName:       getRandomName(lastNames),
			BirthDate:      complexes.XDate((getRandomBirthDate())),
			Photo:          getRandomAvatarURL(),
			Gender:         emigo.NullableOf(randomZeroOrOne()),
			LastIpAddress:  randomPublicIP(),
			PrimaryAddress: emigo.NullableOf(*RandomUserPrimaryAddress()),
			PhoneNumber:    emigo.NullableOf(getRandomPhoneNumber()),
			JobTitle:       emigo.NullableOf(getRandomName(jobTitles)),
			Company:        emigo.NullableOf(getRandomName(companies)),
			Bio:            emigo.NullableOf(getRandomName(bios)),
		}
	}

	// Note: UserActions (the EntityActionsBundle above) is only actually used for its
	// SeederInit now - Browse/Get/Create/Update/AwareDelete all go through the entity's
	// own generated abacdefs.UserEntityActions instead (see UserBrowseAction's own
	// comment for why: UserEntity has no workspace_id column, and
	// EntityActionsBundle/fireback.QueryEntitiesPointer unconditionally filters by it).
}

// HydrateUserPrimaryAddress copies UserEntity.PrimaryAddressRow (populated by GORM from
// the embedded address_line1/city/... columns whenever a row is read from the DB) onto
// PrimaryAddress, the field API responses actually serialize (`json:"primaryAddress"` vs
// PrimaryAddressRow's `json:"-"`) - nothing does this on its own, since PrimaryAddress is
// `gorm:"-"` (GORM never touches it) and UserEntity has no MarshalJSON override either.
// Called after every abacdefs.UserEntityActions.Get/Browse/Update in this file. A no-op
// if there's no row to hydrate from, or the row has no address line set (an empty
// primaryAddress was never given in the first place).
func HydrateUserPrimaryAddress(user *abacdefs.UserEntity) {
	if user == nil || user.PrimaryAddressRow == nil {
		return
	}
	if user.PrimaryAddressRow.AddressLine1 == "" {
		return
	}
	user.PrimaryAddress = emigo.NullableOf(*user.PrimaryAddressRow)
}
