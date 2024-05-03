package util

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func FileExists(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

/*
	func readFile(path string) ([]byte, error) {
		b, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		return b, nil
	}
*/

func ReadFile(path string) (records [][]string, err error) {

	// Open the CSV file
	//fmt.Printf("file: [%v]\n", path)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)
	//reader.Comma = ';'   // Use semicolon as the delimiter
	reader.Comment = '#' // Use '#' as the comment character
	reader.FieldsPerRecord = -1

	// Read all records from the CSV file
	Records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return nil, err
	}

	return Records, nil
}

func RemoveItemFromList[T comparable](l []T, item T) []T {
	/* fmt.Println("\n\t\tRemover\t---------------")
	fmt.Printf("\t\t\tInput list: [%v]\n", strings.Join(strings.Fields(fmt.Sprint(l)), "|"))
	fmt.Printf("\t\t\tRemoved item: [%v]\n", item) */
	out := make([]T, 0)
	for _, element := range l {
		if element != item {
			out = append(out, element)
		}
	}
	/* fmt.Printf("\t\t\tReturned list: [%v]\n", strings.Join(strings.Fields(fmt.Sprint(out)), "|"))
	fmt.Println("\t\t\t---------------\n") */
	return out
}

func CountLeadingDashes(arg string) int {
	re := regexp.MustCompile("^-+")
	match := re.FindString(arg)
	return len(match)
}

func SetName(name *string, item string) {
	if *name == "" {
		*name = strings.TrimSpace(item)
	} else {
		*name = *name + " " + strings.TrimSpace(item)
	}
	/* fmt.Printf("name:\t\t\t[%v]\n", *name) */
}

func RemoveEmptyElementsFromArray(arr []string) []string {
	result := []string{}
	for _, s := range arr {
		if s != "" {
			result = append(result, strings.TrimSpace(s))
		}
	}
	return result
}

func RemoveWhitespacesFromArray(arr []string) []string {
	result := make([]string, len(arr))
	for i, str := range arr {
		result[i] = strings.TrimSpace(str)
	}
	return result
}

func NotEmptyElementsInArrayCount(arr []string) int {
	count := 0
	for _, s := range arr {
		if s != "" {
			count++
		}
	}
	return count
}

func IPv6AddressFiller(normalizedAddr *string, segments []string, segmentsLen int) error {

	for i, segment := range segments[:segmentsLen] {

		// fmt.Printf("IPv6Normalizer. Validate segment: (%v)[%s] from segments: %s\n", i, segment, segments)

		if segment, err := ValidateIPv6Segment(segment); err != nil {
			// fmt.Printf("IPv6AddressNormalizer. Validation error: %s\n", err)
			return err

		} else if segment == "" {
			if i == 0 {
				*normalizedAddr = ":"
			} else {
				*normalizedAddr += ":"
			}

		} else {
			if i == 0 {
				*normalizedAddr = fmt.Sprintf("%s:", segment)
			} else {
				*normalizedAddr += fmt.Sprintf("%s:", segment)
			}
		}

	}
	return nil
}

/* func IPv6AddressExpander(address *string, notEmptyCount, segments int) (string, error) {

	doubleColonsPattern := regexp.MustCompile("::")
	founded := doubleColonsPattern.FindAllStringIndex(*address, -1)

	fmt.Printf("IPv6AddressExpander. Found shorthands: %d\n", len(founded))

	if len(founded) != 1 {
		fmt.Printf("IPv6AddressExpander. bad ipv6 address: %s\n", *address)
		return "", fmt.Errorf("IPv6AddressExpander. bad ipv6 address: %s", *address)
	}
	//if segmentsLen <= 6 && len(founded) == 1 {

	fmt.Printf("IPv6AddressExpander. Expanding shorthand: %s\n", *address)
	// Заменяем "::" на необходимое количество нулей
	if doubleColonsPattern.MatchString(*address) {
		missingSegments := segments - notEmptyCount
		fmt.Printf("IPv6AddressExpander. Missing Segments: %d\n", missingSegments)
		*address = strings.TrimLeft(doubleColonsPattern.ReplaceAllString(*address, strings.Repeat(":0", missingSegments)+":"), ":")
		fmt.Printf("IPv6AddressExpander. Expanded address: %s\n", *address)
	}

	return *address, nil
} */

func IPv6ParseInput(address string) (segments []string, segmentsLen, notEmptyCount int) {
	// Пробуем разложить строку "адреса" на суб-элементы по ":"
	segments = strings.Split(address, ":")

	segmentsLen = len(segments)
	// fmt.Printf("IPv6ParseInput. Segments Len: %v\n", segmentsLen)

	// Считаем количество НЕ пустых элементов
	notEmptyCount = NotEmptyElementsInArrayCount(segments)

	return segments, segmentsLen, notEmptyCount
}
