package utils

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Util struct{}

func ConcatenateStrings(values ...string) string {
	var sb strings.Builder

	for _, str := range values {
		sb.WriteString(str)
	}

	return sb.String()
}

func StringIsNotEmpty(value string) bool {
	return len(strings.TrimSpace(value)) > 0
}

func CapitalizeFirstLetter(text string) string {
	return cases.Title(language.Und, cases.Compact).String(text)
}

func ValidateEmail(email string) bool {
	email = strings.TrimSpace(email)
	result, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&’*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	return result
}

func ValidateUpperCharacters(value string, expectedQuantity int) bool {
	countUppers := 0
	for _, char := range value {
		if unicode.IsUpper(char) {
			countUppers++
			if countUppers >= expectedQuantity {
				break
			}
		}
	}

	return countUppers == expectedQuantity
}

func ValidateSpecialCharacters(value string, expectedQuantity int) bool {
	pattern := `[!@#$%^&*()_+\-=\[\]{}|;':",./<>?]`

	re := regexp.MustCompile(pattern)

	specialChars := re.FindAllString(value, -1)

	return len(specialChars) >= expectedQuantity
}

func RandomStringAndNumber(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}
