package workspaces

import (
	"fmt"
	"math/rand"
	"time"
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

			FirstName: getRandomName(firstNames),
			LastName:  getRandomName(lastNames),
			BirthDate: XDate((getRandomBirthDate())),
			Photo:     getRandomAvatarURL(),
			Gender:    NewInt(randomZeroOrOne()),
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
