package abac

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/torabian/fireback/modules/fireback"
)

func (x *UserEntity) FullName() string {

	full := ""

	if x.FirstName != "" {
		full += x.FirstName
	}

	if x.LastName != "" {
		full += " " + x.LastName
	}

	return full

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

func init() {

	UserActions.SeederInit = func() *UserEntity {
		return &UserEntity{
			FirstName:      getRandomName(firstNames),
			LastName:       getRandomName(lastNames),
			BirthDate:      fireback.XDate((getRandomBirthDate())),
			Photo:          getRandomAvatarURL(),
			Gender:         fireback.NewInt(randomZeroOrOne()),
			LastIpAddress:  randomPublicIP(),
			PrimaryAddress: RandomUserPrimaryAddress(),
		}
	}

	// Tokens are related to users, so let's move them there.
	UserCliCommands = append(
		UserCliCommands,
		TokenCliFn(),
		CreateRootUser,
		AcceptInviteActionCmd,
		UserInvitationsActionCmd,
	)
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

func RandomUserPrimaryAddress() *UserPrimaryAddress {
	rand.Seed(time.Now().UnixNano())
	s := addresses[rand.Intn(len(addresses))]
	return &UserPrimaryAddress{
		AddressLine1:    (s.Address1),
		AddressLine2:    fireback.NewString(s.Address2),
		City:            fireback.NewString(s.City),
		StateOrProvince: fireback.NewString(s.State),
		PostalCode:      fireback.NewString(s.Postcode),
		CountryCode:     fireback.NewString(s.CountryCode),
	}
}
