package util

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func CheckIfEmail(addr string) bool {
	addr = strings.TrimSpace(addr)

	// Если email содержит внутри пробел - возвращаем false
	if strings.Contains(addr, " ") {
		return false
	}

	// Если "@" в email не делит его на две части - возвращаем false
	parts := strings.Split(addr, "@")
	if len(parts) != 2 {
		return false
	}

	// Если имя не проходит валидацию - возвращаем false
	if !CheckInvalidChars(parts[0], "email-name") {
		return false
	}

	// Возвращаем результат валидации домена
	if _, err := ValidateDomain(parts[1]); err != nil {
		return false
	}
	return true
}

func CheckIfMode(s string) bool {
	return s == "to" || s == "cc" || s == "bcc"
}

func CheckValueInArray(val string, array []string) bool {
	for _, item := range array {
		if item == val {
			return true
		}
	}
	return false
}

func CheckStringForDuplicates(input string) []string {

	entries := strings.Split(input, ";")

	duplicates := make(map[string]bool)
	result := []string{}

	for _, entry := range entries {
		if duplicates[strings.TrimSpace(entry)] {
			result = append(result, strings.TrimSpace(entry))
		} else {
			duplicates[strings.TrimSpace(entry)] = true
		}
	}

	return result
}

func CheckIsStringCommented(s string) bool {
	if strings.HasPrefix(s, "#") {
		return true
	}

	for i, char := range s {
		if !unicode.IsSpace(char) {
			return strings.HasPrefix(s[i:], "#") || strings.HasPrefix(s[i:], "//")
		}
	}
	return false
}

func CheckStringNotEmpty(s string) (bool, error) {
	if s == "" {
		return false, errors.New("value is empty")
	}
	return true, nil
}

func CheckIPv6Segments(segment string) (bool, error) {
	// Регулярка для фильтрации IPv6 сегментов
	/*
		patternIPv6HEX := regexp.MustCompile(`^[0-9A-Fa-f]{1,4}$`)

		if segment != "" && !patternIPv6HEX.MatchString(segment) {
			// Convert hex string to integer
			num, err := strconv.ParseInt(segment, 16, 64)
		}
	*/

	return true, nil
}

func CheckInvalidChars(input, classification string) bool {

	var (
		invalidAddressChars    = " \t\n\r~!@#$^&*()+={}|\";’`'<,>? "
		invalidIPv6ZoneIDChars = " \t\n\r%[]|\";<>?/"
		validEmailNameChars    = `^[a-zA-Z0-9._]+$`
		invalidDomainChars     = ` &‘@*(),!?_#$€£₽^;:\\/%+=<>`
		invalidSubjectChars    = `\t\n\r\u\x`
		// invalidAddressChars = " \t\n\r~!@#$^&*()+={}|\";’`'<,>/? "
		// invalidIPv6Chars    = " \t\n\r%&*+={}|;“’<,>?/`"
		// invalidChars        = " \t\n\r~!@#$%^&*()+={}|;\\“’<,>?/`"
		// validNameChars      = `^[a-zа-яA-ZА-Я0-9.,_!@№#$%^&*()+={}|\\“’<> ]+$`

		/*
			\x20	// 32 Space
			\x21	// 33 !	Exclamation mark
			\x22	// 34 "	Double quotes (or speech marks)
			\x23	// 35 #	Number sign
			\x24	// 36 $	Dollar
			\x25	// 37 %	Per cent sign
			\x26	// 38 &	Ampersand
			\x27	// 39 '	Single quote
			\x28	// 40 (	Open parenthesis (or open bracket)
			\x29	// 41 )	Close parenthesis (or close bracket)
			\x2A	// 42 *	Asterisk
			\x2B	// 43 +	Plus
			\x2C	// 44 ,	Comma
			\x2D	// 45 -	Hyphen-minus
			\x2E	// 46 .	Period, dot or full stop
			\x2F	// 47 /	Slash or divide
			\x30	// 48 0	Zero
			\x31	// 49 1	One
			\x32	// 50 2	Two
			\x33	// 51 3	Three
			\x34	// 52 4	Four
			\x35	// 53 5	Five
			\x36	// 54 6	Six
			\x37	// 55 7	Seven
			\x38	// 56 8	Eight
			\x39	// 57 9	Nine
			\x3A	// 58 :	Colon
			\x3B	// 59 ;	Semicolon
			\x3C	// 60 <	Less than (or open angled bracket)
			\x3D	// 61 =	Equals
			\x3E	// 62 >	Greater than (or close angled bracket)
			\x3F	// 63 ?	Question mark
			\x40	// 64 @	At sign
		*/

	)

	switch classification {

	case "address":
		//return !strings.ContainsAny(input, invalidAddressChars) && !CheckNonASCII(input)
		return CheckNonPrintableASCII(input) || !strings.ContainsAny(input, invalidAddressChars)

	case "ipv6zone":
		return strings.ContainsAny(input, invalidIPv6ZoneIDChars)

	case "domain":
		return CheckNonPrintableASCII(input) || strings.ContainsAny(input, invalidDomainChars)

	case "subject":
		/*
			a := strings.ContainsAny(input, invalidSubjectChars)
			b := CheckNonPrintableASCII(input)
			fmt.Printf("invalidSubjectChars: %t\n", a)
			fmt.Printf("CheckNonPrintableASCII: %t\n", b)
		*/
		return CheckNonPrintableASCII(input) || strings.ContainsAny(input, invalidSubjectChars)

	case "email-name":
		// Валидация имени электронной почты .
		// re := regexp.MustCompile(`^[a-zA-Z0-9._]+$`)
		re := regexp.MustCompile(validEmailNameChars)
		// Если имя не проходит валидацию - возвращаем false
		return !CheckNonPrintableASCII(input) || re.MatchString(input)

		/*
			case "name":
				// Валидация имени электронной почты .
					// re := regexp.MustCompile(`^[a-zA-Z0-9._]+$`)
					re := regexp.MustCompile(validNameChars)
					// Если имя не проходит валидацию - возвращаем false
					return !CheckNonPrintableASCII(input) || re.MatchString(input)
					//return !strings.ContainsAny(input, invalidChars) && !CheckNonASCII(input)
		*/

	default:
		return false

	}
}

func CheckNonPrintableASCII(input string) bool {
	// ASCII printable characters (character code 32-127)
	for i, char := range input {
		if char < 32 {
			fmt.Printf("non-printable char: (%d)%v %v\n", i, strconv.QuoteRune(char), char)
			return true
		}
	}
	return false
}

func CheckASCIILatCyr(input string) bool {

	// 65-90 - A-Z
	// 97-122 - a-z
	// 126-159 - Extended ASCII characters А-Я
	// 160-175 - Extended ASCII characters а-п
	// 224-239 - Extended ASCII characters р-я
	// 240-241 - Extended ASCII characters Ёё

	for _, char := range input {

		if (char > 64 && char < 91) || (char > 96 && char < 123) ||
			(char > 125 && char < 160) || (char > 159 && char < 176) ||
			(char > 223 && char < 240) || (char == 240 || char == 241) {
			return true
		}

	}
	return false
}
