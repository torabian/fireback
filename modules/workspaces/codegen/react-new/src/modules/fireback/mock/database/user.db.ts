import { random, sample, times } from "lodash";
import {
  UserEntity,
  UserPrimaryAddress,
} from "../../sdk/modules/abac/UserEntity";
import { MemoryEntity } from "./memory-db";

const generateRandomIp = (): string => {
  const isPrivate = Math.random() > 0.5; // Randomly decide whether to generate a private or public IP

  if (isPrivate) {
    // Generate a random private IP address
    const privateIpRange = Math.floor(Math.random() * 3); // Pick one of the private ranges (0 = Class A, 1 = Class B, 2 = Class C)
    switch (privateIpRange) {
      case 0: // Class A: 10.x.x.x
        return `10.${Math.floor(Math.random() * 256)}.${Math.floor(
          Math.random() * 256
        )}.${Math.floor(Math.random() * 256)}`;
      case 1: // Class B: 172.16.x.x to 172.31.x.x
        return `172.${Math.floor(16 + Math.random() * 16)}.${Math.floor(
          Math.random() * 256
        )}.${Math.floor(Math.random() * 256)}`;
      case 2: // Class C: 192.168.x.x
        return `192.168.${Math.floor(Math.random() * 256)}.${Math.floor(
          Math.random() * 256
        )}`;
      default:
        return `192.168.${Math.floor(Math.random() * 256)}.${Math.floor(
          Math.random() * 256
        )}`;
    }
  } else {
    // Generate a random public IP address (avoid private ranges)
    return `${Math.floor(Math.random() * 223) + 1}.${Math.floor(
      Math.random() * 256
    )}.${Math.floor(Math.random() * 256)}.${Math.floor(Math.random() * 256)}`;
  }
};

// Arrays for first names, last names, and addresses with more varied data

// First names from different countries/cultures
const firstNames = [
  "Ali",
  "Behnaz",
  "Carlos",
  "Daniela",
  "Ethan",
  "Fatima",
  "Gustavo",
  "Helena",
  "Isla",
  "Javad",
  "Kamila",
  "Leila",
  "Mateo",
  "Nasim",
  "Omid",
  "Parisa",
  "Rania",
  "Saeed",
  "Tomas",
  "Ursula",
  "Vali",
  "Wojtek",
  "Zara",
  "Alice",
  "Bob",
  "Charlie",
  "Diana",
  "George",
  "Mohammed",
  "Julia",
  "Khalid",
  "Lena",
  "Mohammad",
  "Nina",
  "Oscar",
  "Quentin",
  "Rosa",
  "Sam",
  "Tina",
  "Umar",
  "Vera",
  "Waleed",
  "Xenia",
  "Yara",
  "Ziad",
  "Maxim",
  "Johann",
  "Krzysztof",
  "Baris",
  "Mehmet",
];

// Last names from different countries/cultures
const lastNames = [
  "Smith",
  "Johnson",
  "Williams",
  "Brown",
  "Jones",
  "Garcia",
  "Miller",
  "Davis",
  "Rodriguez",
  "Martinez",
  "Hernandez",
  "Lopez",
  "Gonzalez",
  "Wilson",
  "Anderson",
  "Thomas",
  "Taylor",
  "Moore",
  "Jackson",
  "Martin",
  "Lee",
  "Perez",
  "Thompson",
  "White",
  "Harris",
  "Sanchez",
  "Clark",
  "Ramirez",
  "Lewis",
  "Robinson",
  "Walker",
  "Young",
  "Allen",
  "King",
  "Wright",
  "Scott",
  "Torres",
  "Nguyen",
  "Hill",
  "Flores",
  "Green",
  "Adams",
  "Nelson",
  "Baker",
  "Hall",
  "Rivera",
  "Campbell",
  "Mitchell",
  "Carter",
  "Roberts",
  "Kowalski",
  "Nowak",
  "Jankowski",
  "Zieliński",
  "Wiśniewski",
  "Lewandowski",
  "Kaczmarek",
  "Bąk",
  "Pereira",
  "Altıntaş",
];

// Addresses with a variety of global samples
const addresses: UserPrimaryAddress[] = [
  {
    addressLine1: "123 Main St",
    addressLine2: "Apt 4",
    city: "Berlin",
    stateOrProvince: "Berlin",
    postalCode: "10115",
    countryCode: "DE",
  },
  {
    addressLine1: "456 Elm St",
    addressLine2: "Apt 23",
    city: "Paris",
    stateOrProvince: "Île-de-France",
    postalCode: "75001",
    countryCode: "FR",
  },
  {
    addressLine1: "789 Oak Dr",
    addressLine2: "Apt 9",
    city: "Warszawa",
    stateOrProvince: "Mazowieckie",
    postalCode: "01010",
    countryCode: "PL",
  },
  {
    addressLine1: "101 Maple Ave",
    addressLine2: "",
    city: "Tehran",
    stateOrProvince: "تهران",
    postalCode: "11365",
    countryCode: "IR",
  },
  {
    addressLine1: "202 Pine St",
    addressLine2: "Apt 7",
    city: "Madrid",
    stateOrProvince: "Community of Madrid",
    postalCode: "28001",
    countryCode: "ES",
  },
  {
    addressLine1: "456 Park Ave",
    addressLine2: "Suite 5",
    city: "New York",
    stateOrProvince: "NY",
    postalCode: "10001",
    countryCode: "US",
  },
  {
    addressLine1: "789 Sunset Blvd",
    addressLine2: "Unit 32",
    city: "Los Angeles",
    stateOrProvince: "CA",
    postalCode: "90001",
    countryCode: "US",
  },
  {
    addressLine1: "12 Hauptstrasse",
    addressLine2: "Apt 2",
    city: "Munich",
    stateOrProvince: "Bavaria",
    postalCode: "80331",
    countryCode: "DE",
  },
  {
    addressLine1: "75 Taksim Square",
    addressLine2: "Apt 12",
    city: "Istanbul",
    stateOrProvince: "Istanbul",
    postalCode: "34430",
    countryCode: "TR",
  },
  {
    addressLine1: "321 Wierzbowa",
    addressLine2: "",
    city: "Kraków",
    stateOrProvince: "Małopolskie",
    postalCode: "31000",
    countryCode: "PL",
  },
  {
    addressLine1: "55 Rue de Rivoli",
    addressLine2: "Apt 10",
    city: "Paris",
    stateOrProvince: "Île-de-France",
    postalCode: "75004",
    countryCode: "FR",
  },
  {
    addressLine1: "1001 Tehran Ave",
    addressLine2: "",
    city: "Tehran",
    stateOrProvince: "تهران",
    postalCode: "14155",
    countryCode: "IR",
  },
  {
    addressLine1: "9 Calle de Alcalá",
    addressLine2: "Apt 6",
    city: "Madrid",
    stateOrProvince: "Madrid",
    postalCode: "28009",
    countryCode: "ES",
  },
  {
    addressLine1: "222 King St",
    addressLine2: "Suite 1B",
    city: "London",
    stateOrProvince: "London",
    postalCode: "E1 6AN",
    countryCode: "GB",
  },
  {
    addressLine1: "15 St. Peters Rd",
    addressLine2: "",
    city: "Toronto",
    stateOrProvince: "Ontario",
    postalCode: "M5A 1A2",
    countryCode: "CA",
  },
  {
    addressLine1: "1340 Via Roma",
    addressLine2: "",
    city: "Rome",
    stateOrProvince: "Lazio",
    postalCode: "00100",
    countryCode: "IT",
  },
  {
    addressLine1: "42 Nevsky Prospekt",
    addressLine2: "Apt 1",
    city: "Saint Petersburg",
    stateOrProvince: "Leningradskaya",
    postalCode: "190000",
    countryCode: "RU",
  },
  {
    addressLine1: "3 Rüdesheimer Str.",
    addressLine2: "Apt 9",
    city: "Frankfurt",
    stateOrProvince: "Hessen",
    postalCode: "60326",
    countryCode: "DE",
  },
  {
    addressLine1: "271 Süleyman Demirel Bulvarı",
    addressLine2: "Apt 45",
    city: "Ankara",
    stateOrProvince: "Ankara",
    postalCode: "06100",
    countryCode: "TR",
  },
  {
    addressLine1: "7 Avenues des Champs-Élysées",
    addressLine2: "",
    city: "Paris",
    stateOrProvince: "Île-de-France",
    postalCode: "75008",
    countryCode: "FR",
  },
  {
    addressLine1: "125 E. 9th St.",
    addressLine2: "Apt 12",
    city: "Chicago",
    stateOrProvince: "IL",
    postalCode: "60606",
    countryCode: "US",
  },
  {
    addressLine1: "30 Rue de la Paix",
    addressLine2: "",
    city: "Paris",
    stateOrProvince: "Île-de-France",
    postalCode: "75002",
    countryCode: "FR",
  },
  {
    addressLine1: "16 Zlote Tarasy",
    addressLine2: "Apt 18",
    city: "Warszawa",
    stateOrProvince: "Mazowieckie",
    postalCode: "00-510",
    countryCode: "PL",
  },
  {
    addressLine1: "120 Váci utca",
    addressLine2: "",
    city: "Budapest",
    stateOrProvince: "Budapest",
    postalCode: "1056",
    countryCode: "HU",
  },
  {
    addressLine1: "22 Sukhbaatar Sq.",
    addressLine2: "",
    city: "Ulaanbaatar",
    stateOrProvince: "Central",
    postalCode: "14190",
    countryCode: "MN",
  },
  {
    addressLine1: "34 Princes Street",
    addressLine2: "Flat 1",
    city: "Edinburgh",
    stateOrProvince: "Scotland",
    postalCode: "EH2 4AY",
    countryCode: "GB",
  },
  {
    addressLine1: "310 Alzaibiyah",
    addressLine2: "",
    city: "Amman",
    stateOrProvince: "Amman",
    postalCode: "11183",
    countryCode: "JO",
  },
  {
    addressLine1: "401 Taksim Caddesi",
    addressLine2: "Apt 25",
    city: "Istanbul",
    stateOrProvince: "Istanbul",
    postalCode: "34430",
    countryCode: "TR",
  },
  {
    addressLine1: "203 High Street",
    addressLine2: "Unit 3",
    city: "London",
    stateOrProvince: "London",
    postalCode: "W1T 2LQ",
    countryCode: "GB",
  },
  {
    addressLine1: "58 Via Nazionale",
    addressLine2: "",
    city: "Rome",
    stateOrProvince: "Lazio",
    postalCode: "00184",
    countryCode: "IT",
  },
  {
    addressLine1: "47 Gloucester Road",
    addressLine2: "",
    city: "London",
    stateOrProvince: "London",
    postalCode: "SW7 4QA",
    countryCode: "GB",
  },
  {
    addressLine1: "98 Calle de Bravo Murillo",
    addressLine2: "",
    city: "Madrid",
    stateOrProvince: "Madrid",
    postalCode: "28039",
    countryCode: "ES",
  },
  {
    addressLine1: "57 Mirza Ghalib Street",
    addressLine2: "",
    city: "Tehran",
    stateOrProvince: "تهران",
    postalCode: "15996",
    countryCode: "IR",
  },
  {
    addressLine1: "35 Królewska St",
    addressLine2: "",
    city: "Warszawa",
    stateOrProvince: "Mazowieckie",
    postalCode: "00-065",
    countryCode: "PL",
  },
  {
    addressLine1: "12 5th Ave",
    addressLine2: "",
    city: "New York",
    stateOrProvince: "NY",
    postalCode: "10128",
    countryCode: "US",
  },
];

const generateUniqueId = (): string => {
  // Generate a cryptographically secure random string (36 characters)
  const array = new Uint8Array(18); // 18 * 2 = 36 characters
  window.crypto.getRandomValues(array);

  // Convert to base36 (alphanumeric) format
  const randomStr = Array.from(array)
    .map((byte) => byte.toString(36).padStart(2, "0")) // convert each byte to a base36 string, pad with 0 if needed
    .join("");

  // Combine random string with timestamp for additional uniqueness
  const timestamp = Date.now().toString(36); // Base36 timestamp
  return timestamp + randomStr.slice(0, 30 - timestamp.length); // Ensure the final length is 30
};

const createMockUser = (): UserEntity => {
  return {
    uniqueId: generateUniqueId(),
    firstName: sample(firstNames),
    lastName: sample(lastNames),
    photo: `https://randomuser.me/api/portraits/men/${Math.floor(
      Math.random() * 100
    )}.jpg`, // Random photo URL
    birthDate: ((((new Date().getDate() as any) +
      "/" +
      new Date().getMonth()) as any) +
      "/" +
      new Date().getFullYear()) as any,
    gender: Math.random() > 0.5 ? 1 : 0, // Randomly assign gender (1 = Male, 0 = Female)
    title: Math.random() > 0.5 ? "Mr." : "Ms.", // Random title
    avatar: `https://randomuser.me/api/portraits/men/${Math.floor(
      Math.random() * 100
    )}.jpg`, // Random avatar URL
    lastIpAddress: generateRandomIp(),
    primaryAddress: sample(addresses),
  };
};

export const MockUsers = new MemoryEntity<UserEntity>(
  times(10000, () => createMockUser())
);
