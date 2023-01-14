package data

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"syreclabs.com/go/faker"
	"syreclabs.com/go/faker/locales"
)

type stringGeneratorFunction func(minLength int, maxLength int) string

func defaultFormattedStringGenerators(generator *rangedTextGenerator) map[string]stringGeneratorFunction {

	// Default Locale is English
	faker.Locale = locales.Ko

	base64 := &base64Generator{generator: generator}
	html := &htmlGenerator{random: generator.random}

	return map[string]stringGeneratorFunction{
		// ADDRESS
		"address": func(_ int, _ int) string {
			return faker.Address().String()
		},
		"city": func(_ int, _ int) string {
			return faker.Address().City()
		},
		"street": func(_ int, _ int) string {
			return faker.Address().StreetName()
		},
		"streetAddress": func(_ int, _ int) string {
			return faker.Address().StreetAddress()
		},
		"secondaryAddress": func(_ int, _ int) string {
			return faker.Address().SecondaryAddress()
		},
		"buildingNumber": func(_ int, _ int) string {
			return faker.Address().BuildingNumber()
		},
		"postcode": func(_ int, _ int) string {
			return faker.Address().ZipCode()
		},
		"postcodeByState": func(_ int, _ int) string {
			return faker.Address().ZipCodeByState("IN")
		},
		"zipCode": func(_ int, _ int) string {
			return faker.Address().ZipCode()
		},
		"zipCodeByState": func(_ int, _ int) string {
			return faker.Address().ZipCodeByState("IN")
		},
		"timeZone": func(_ int, _ int) string {
			return faker.Address().TimeZone()
		},
		"cityPrefix": func(_ int, _ int) string {
			return faker.Address().CityPrefix()
		},
		"citySuffix": func(_ int, _ int) string {
			return faker.Address().CitySuffix()
		},
		"streetSuffix": func(_ int, _ int) string {
			return faker.Address().StreetSuffix()
		},
		"state": func(_ int, _ int) string {
			return faker.Address().State()
		},
		"stateAbbr": func(_ int, _ int) string {
			return faker.Address().StateAbbr()
		},
		"country": func(_ int, _ int) string {
			return faker.Address().Country()
		},
		"countryCode": func(_ int, _ int) string {
			return faker.Address().CountryCode()
		},
		"latitude": func(_ int, _ int) string {
			return fmt.Sprintf("%f", faker.Address().Latitude())
		},
		"longitude": func(_ int, _ int) string {
			return fmt.Sprintf("%f", faker.Address().Longitude())
		},
		// APP
		"appName": func(_ int, _ int) string {
			return faker.App().Name()
		},
		"appVersion": func(_ int, _ int) string {
			return faker.App().Version()
		},
		"appAuthor": func(_ int, _ int) string {
			return faker.App().Author()
		},
		// AVATAR
		"avatar": func(_ int, _ int) string {
			return faker.Avatar().String()
		},
		// BITCOIN
		"bitcoinAddress": func(_ int, _ int) string {
			return faker.Bitcoin().Address()
		},
		// BUSINESS
		"creditCardNumber": func(_ int, _ int) string {
			return faker.Finance().CreditCard()
		},
		"creditCardType": func(_ int, _ int) string {
			return faker.Business().CreditCardType()
		},
		"creditCardExpiry": func(_ int, _ int) string {
			return faker.Business().CreditCardExpiryDate()
		},
		// CODE
		"isbn10": func(_ int, _ int) string {
			return faker.Code().Isbn10()
		},
		"isbn13": func(_ int, _ int) string {
			return faker.Code().Isbn13()
		},
		"ean8": func(_ int, _ int) string {
			return faker.Code().Ean8()
		},
		"ean13": func(_ int, _ int) string {
			return faker.Code().Ean13()
		},
		"abn": func(_ int, _ int) string {
			return faker.Code().Abn()
		},
		"rut": func(_ int, _ int) string {
			return faker.Code().Rut()
		},
		// COMMERCE
		"color": func(_ int, _ int) string {
			return faker.Commerce().Color()
		},
		"department": func(_ int, _ int) string {
			return faker.Commerce().Department()
		},
		"productName": func(_ int, _ int) string {
			return faker.Commerce().ProductName()
		},
		"price": func(_ int, _ int) string {
			return fmt.Sprintf("%f", faker.Commerce().Price())
		},
		// COMPANY
		"companyName": func(_ int, _ int) string {
			return faker.Company().Name()
		},
		"companySuffix": func(_ int, _ int) string {
			return faker.Company().Suffix()
		},
		"companyCatchPhrase": func(_ int, _ int) string {
			return faker.Company().CatchPhrase()
		},
		"companyBs": func(_ int, _ int) string {
			return faker.Company().Bs()
		},
		"Ein": func(_ int, _ int) string {
			return faker.Company().Ein()
		},
		"DunsNumber": func(_ int, _ int) string {
			return faker.Company().DunsNumber()
		},
		"Logo": func(_ int, _ int) string {
			return faker.Company().Logo()
		},
		// Date
		"date": func(_ int, _ int) string {
			date := generateRandomTime()
			return fmt.Sprintf("%d-%02d-%02d", date.Year(), int(date.Month()), date.Day())
		},

		"date-time": func(_ int, _ int) string {
			date := generateRandomTime()

			return date.Format(time.RFC3339)
		},

		// FINANCE
		"creditCard": func(_ int, _ int) string {
			return faker.Finance().CreditCard()
		},

		// HACKER
		"saySomethingSmart": func(_ int, _ int) string {
			return faker.Hacker().SaySomethingSmart()
		},
		"abbreviation": func(_ int, _ int) string {
			return faker.Hacker().Abbreviation()
		},
		"adjective": func(_ int, _ int) string {
			return faker.Hacker().Adjective()
		},
		"noun": func(_ int, _ int) string {
			return faker.Hacker().Noun()
		},
		"verb": func(_ int, _ int) string {
			return faker.Hacker().Verb()
		},
		"ingVerb": func(_ int, _ int) string {
			return faker.Hacker().IngVerb()
		},
		"phrases": func(_ int, _ int) string {
			return faker.Hacker().Phrases()[0] // FIXME []string 이나 0번째만 리턴한다.
		},
		// IMAGE
		"imageUrl": func(_ int, _ int) string {
			return "https://loremflickr.com/640/480"
		},
		"imageUrls": func(_ int, _ int) string {
			return "https://picsum.photos/v2/list"
		},
		// INTERNET
		"email": func(_ int, _ int) string {
			return faker.Internet().Email()
		},
		"freeEmail": func(_ int, _ int) string {
			return faker.Internet().FreeEmail()
		},
		"safeEmail": func(_ int, _ int) string {
			return faker.Internet().SafeEmail()
		},
		"userName": func(_ int, _ int) string {
			return faker.Internet().UserName()
		},
		"password": func(_ int, _ int) string {
			return faker.Internet().Password(8, 16)
		},
		"domainName": func(_ int, _ int) string {
			return faker.Internet().DomainName()
		},
		"domainWord": func(_ int, _ int) string {
			return faker.Internet().DomainWord()
		},
		"domainSuffix": func(_ int, _ int) string {
			return faker.Internet().DomainSuffix()
		},
		"macAddress": func(_ int, _ int) string {
			return faker.Internet().MacAddress()
		},
		"ipv4Address": func(_ int, _ int) string {
			return faker.Internet().IpV4Address()
		},
		"ipv6Address": func(_ int, _ int) string {
			return faker.Internet().IpV6Address()
		},
		"url": func(_ int, _ int) string {
			return faker.Internet().Url()
		},
		"slug": func(_ int, _ int) string {
			return faker.Internet().Slug()
		},
		//--------------------------------
		"hostname": func(_ int, _ int) string {
			return faker.Internet().DomainName()
		},
		"ipv4": func(_ int, _ int) string {
			return faker.Internet().IpV4Address()
		},
		"ipv6": func(_ int, _ int) string {
			return faker.Internet().IpV6Address()
		},

		// LOREM
		"character": func(_ int, _ int) string {
			return faker.Lorem().Character()
		},
		"characters": func(length int, _ int) string {
			return faker.Lorem().Characters(length)
		},
		"word": func(_ int, _ int) string {
			return faker.Lorem().Word()
		},
		"words": func(length int, _ int) string {
			return strings.Join(faker.Lorem().Words(length), " ")
		},
		"sentence": func(length int, _ int) string {
			return faker.Lorem().Sentence(length)
		},
		"sentences": func(length int, _ int) string {
			return strings.Join(faker.Lorem().Sentences(length), " ")
		},
		"paragraph": func(length int, _ int) string {
			return faker.Lorem().Paragraph(length)
		},
		"paragraphs": func(length int, _ int) string {
			return strings.Join(faker.Lorem().Paragraphs(length), " ")
		},

		// NAME
		"name": func(_ int, _ int) string {
			return faker.Name().Name()
		},
		"firstName": func(_ int, _ int) string {
			return faker.Name().FirstName()
		},
		"lastName": func(_ int, _ int) string {
			return faker.Name().LastName()
		},
		"prefix": func(_ int, _ int) string {
			return faker.Name().Prefix()
		},
		"suffix": func(_ int, _ int) string {
			return faker.Name().Suffix()
		},
		"title": func(_ int, _ int) string {
			return faker.Name().Title()
		},
		// PHONE NUMBER
		"phoneNumber": func(_ int, _ int) string {
			return faker.PhoneNumber().PhoneNumber()
		},
		"cellPhone": func(_ int, _ int) string {
			return faker.PhoneNumber().CellPhone()
		},
		"areaCode": func(_ int, _ int) string {
			return faker.PhoneNumber().AreaCode()
		},
		"exchangeCode": func(_ int, _ int) string {
			return faker.PhoneNumber().ExchangeCode()
		},
		"subscriberNumber": func(length int, _ int) string {
			return faker.PhoneNumber().SubscriberNumber(length)
		},
		// TEAM
		"team": func(_ int, _ int) string {
			return faker.Team().Name()
		},
		"teamCreature": func(_ int, _ int) string {
			return faker.Team().Creature()
		},
		"teamState": func(_ int, _ int) string {
			return faker.Team().State()
		},

		// TIME
		"time": func(_ int, _ int) string {
			return generateRandomTime().Format(time.RFC3339)
		},

		// UUID
		"uuid": func(_ int, _ int) string {
			return uuid.Must(uuid.NewV4()).String()
		},

		"byte": base64.GenerateBase64Text,
		"html": html.GenerateHTML,
	}
}

func generateRandomTime() time.Time {
	return faker.Date().Between(
		time.Date(1800, 1, 1, 1, 1, 1, 1, time.UTC),
		time.Date(2100, 1, 1, 1, 1, 1, 1, time.UTC),
	)
}
