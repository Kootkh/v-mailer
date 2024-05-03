package util

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func NormalizeRecipient(items []string) (Address, Name, Mode string, err error) {
	var (
		address string // email address
		name    string // recipient name
		mode    string // send mode
		// modeAlt	string // send mode alternative (if specified in items[1])
		// inputItems          = items
		// inputItemsLen       = len(items)
		// itemsLen = len(items)

	)

	// Если в массиве строк items элементов больше чем три - что-то не так. Возвращаем ошибку
	if len(items) > 3 {
		return "", "", "", fmt.Errorf("entry: [%v] - len exceeded. Expected 3 or less, have: [%v]", items, len(items))
	}

	// Первый элемент массива строк items может быть ТОЛЬКО email 'address'.
	// Второй элемент массива строк items может быть как строка(возможно разделённая пробелами) 'name', так и 'mode'.
	// Третий элемент массива строк items может быть ТОЛЬКО send 'mode'.

	// Проходим циклом по всем элементам массива items
	for i, item := range items {

		// Удаляем лишние пробелы у каждого элемента item
		item = strings.TrimSpace(item)

		// Если нулевой или второй элемент массива items содержат значения разделённые пробелами - возвращаем ошибку
		if i != 1 && len(RemoveEmptyElementsFromArray(strings.Split(item, " "))) > 1 {
			return "", "", "", fmt.Errorf("entry: [%v], item(%v)[%v] - multiple values not allowed", items, i, item)
		}

		// Если нулевой элемент массива items не проходит проверку на email - возвращаем ошибку
		if i == 0 && !CheckIfEmail(item) {
			return "", "", "", fmt.Errorf("entry: [%v], item(%v)[%v] - email address not specified", items, i, item)
		} else if i == 0 && CheckIfEmail(item) {
			// Иначе, если нулевой элемент массива строк items проходит проверку на email - Вносим в address найденное значение.
			address = item
			// Итерируемся далее по циклу
			continue
		}

		// Если второй элемент массива строк items не проходит проверку на mode - возвращаем ошибку
		if i == 2 && item != "" && !CheckIfMode(item) {
			return "", "", "", fmt.Errorf("entry: [%v], item(%v)[%v] - bad value in mode", items, i, item)
		}

		// Если это первый элемент массива строк items содержит строку разделённую пробелами
		if i == 1 {

			if len(items) == 2 {
				// Если в массиве строк items всего два элемента и первый элемент массива строк items проходит проверку на "mode"
				if CheckIfMode(item) {
					//Вносим в "mode" найденное значение.
					mode = item
					// Итерируемся далее по циклу
					continue
					// Иначе, если в массиве строк items всего два элемента и первый элемент массива строк items проходит проверку на "email" - возвращаем ошибку (его тут быть не должно)
				} else if CheckIfEmail(item) {
					return "", "", "", fmt.Errorf("entry: [%v], item(%v)[%v] - additional email address not allowed", items, i, item)
				}
			}

			// В других случаях - первый элемент массива строк items должен быть "name"

			// пробуем разложить первый элемент массива строк items на суб-элементы по пробелам (удаляя пустые элементы) и проходим по ним циклом добавляя значения в "name" убирая лишние пробелы у каждого суб-элемента
			for _, itemValue := range RemoveEmptyElementsFromArray(strings.Split(item, " ")) {

				if CheckIfEmail(itemValue) {
					return "", "", "", fmt.Errorf("entry: [%v], item(%v)[%v] - additional email address not allowed", items, i, itemValue)
				}

				if CheckIfMode(itemValue) {
					// Если в массиве строк items три элемента и первый элемент массива строк items проходит проверку на "mode" - возвращаем ошибку (его тут быть не должно)
					return "", "", "", fmt.Errorf("entry: [%v], item(%v)[%v] - mode in name field", items, i, itemValue)
				}

				SetName(&name, strings.TrimSpace(itemValue))
			}

			// Итерируемся далее по циклу
			continue
		}
	}

	// Если после перебора всех элементов массива items "name" или "mode" не определено - заполняем "name" значением "address" и "mode" дефолтным значением "to".
	if name == "" {
		name = address
	}

	if mode == "" {
		mode = "to"
	}

	return address, name, mode, nil
}

func NormalizeFlags(ctx context.Context, args []string) (Args []string, err error) {

	var (
		flagCount = map[string]int{}
		argo      string

		singleInstanceFlags = []string{
			"--from",         // FlagSenderAddressName
			"--config", "-c", // FlagConfigPathName, FlagConfigPathShorthand
			"--help", "-h", // FlagShowHelpName, FlagShowHelpShorthand
			"--info", "-i", // FlagShowSMTPInfoName, FlagShowSMTPInfoShorthand
			"--examples", "-e", // FlagShowExamplesName, FlagShowExamplesShorthand
			"--version", "-V", // FlagShowVersionName, FlagShowVersionShorthand
			"--debug",        // FlagDebugName
			"--smtp_server",  // FlagSMTPServerName
			"--smtp_port",    // FlagSMTPPortName
			"--domain", "-d", // FlagDomainName, FlagDomainShorthand
			"--log_level",      // FlagDebugLogLevelName
			"--log_format",     // FlagDebugLogFormatName
			"--log_file",       // FlagDebugLogFileName
			"--log_marker",     // FlagAppLogMarkerName
			"--verify_cert",    // FlagVerifyCertName
			"--ssl",            // FlagSSLName
			"--username", "-u", // FlagAuthUsernameName, FlagAuthUsernameShorthand
			"--password", "-p", // FlagAuthPasswordName, FlagAuthPasswordShorthand
			"--charset",       // FlagCharsetName
			"--subject", "-s", // FlagSubjectName, FlagSubjectShorthand
			"--message", "-m", // FlagBodyMessageName, FlagBodyMessageShorthand
			"--file", "-f", // FlagBodyFileName, FlagBodyFileShorthand
			"--mime-type", // FlagBodyMimeTypeName
		}
	)

	// Создаем массив нормализованных аргументов
	normalizedArgs := make([]string, 0)

	// Добавляем первый элемент массива (путь и имя исполняемого файла) в массив нормализованных аргументов
	normalizedArgs = append(normalizedArgs, args[0])

	// Итерируемся по всем элементам массива os.Args исключая первый элемент массива (путь и имя исполняемого файла)
	for _, arg := range args[1:] {

		// Если аргумент начинается с "-" - считаем его флагом
		if strings.HasPrefix(arg, "-") {

			// fmt.Printf("arg: %s\n", arg)

			// Считаем количество лидирующих "-" в аргументе
			countDashes := CountLeadingDashes(arg)

			// Заменяем все "-" на "_" в названии флага (исключая лидирующие "-/--")
			re := regexp.MustCompile("^-+")
			argo = re.ReplaceAllString(arg, "")
			argo = strings.ReplaceAll(argo, "-", "_")

			// Если флаг НЕ является "краткой записью" (одним символом) - трансформируем строку аргумента в нижний регистр
			//fmt.Printf("len arg: %v\n", len(argo))
			if len(argo) > 1 {
				argo = strings.ToLower(argo)
			}

			// fmt.Printf("argo: %s\n", argo)
			// fmt.Printf("dashes count: %d\n", countDashes)

			// Возвращаем нормализованным названиям флагов соответствующее количество лидирующих "-"
			switch countDashes {
			case 1:
				arg = "-" + argo
			case 2:
				arg = "--" + argo
			}

			// Инкрементируем счетчик использования данного флага
			flagCount[arg]++

			// fmt.Printf("normalized arg [%s] count: %d\n", arg, flagCount[arg])
		}

		// Если счетчик использования данного флага больше одного и этот флаг является одним из списка уникальности - возвращаем ошибку
		if flagCount[arg] > 1 && CheckValueInArray(arg, singleInstanceFlags) {
			return nil, errors.New("flag " + arg + " provided more than once")
			//Printf("Flag %s is in unique mode and is set more than once [%v]\n", arg, flagCount[arg])

		}

		// Добавляем аргумент в нормализованный массив
		normalizedArgs = append(normalizedArgs, arg)
		// fmt.Printf("normalized args: %v\n\n", normalizedArgs)

	}

	// Возвращаем нормализованный массив и нулевую ошибку
	return normalizedArgs, nil
}

func IPv4Normalizer(address, port string) (normAddress string, normPort int, err error) {

	const (
		errMsgAddr = "IPv4Normalizer. invalid address: %s"
		errMsgPort = "IPv4Normalizer. invalid port: %v"
	)

	segments := strings.Split(address, ".")

	if len(segments) != 4 {
		return "", 0, fmt.Errorf(errMsgAddr, address)
	}

	for _, segment := range segments {

		segmentInt, err := strconv.Atoi(segment)

		if err != nil {
			return "", 0, err
		}

		if _, err = ValidateIPv4Segment(segmentInt); err != nil {
			return "", 0, fmt.Errorf(errMsgAddr, address)
		}

	}

	normAddress = address

	if port != "" {
		if normPort, err = ValidatePort(port); err != nil {
			return "", 0, fmt.Errorf(errMsgPort, address)
		}
	}

	return normAddress, normPort, nil
}

func IPv6Normalizer(ipv6, ipv4, zone, port string) (normAddress, normZone string, normPort int, err error) {

	var (
		maxSegmentsLen = 8
		errMsg         = "IPv6Normalizer. [FAILED] - invalid address: %s"
	)

	// Если нашли более двух двоеточий подряд или более одного сочетания двух двоеточий в строке адреса  - ошибка
	if len(regexp.MustCompile(":{3,}").FindStringSubmatch(ipv6)) > 0 || strings.Count(ipv6, "::") > 1 {
		return "", "", 0, fmt.Errorf(errMsg, ipv6)
	}

	// segments, segmentsLen, notEmptyCount := IPv6ParseInput(ipv6)
	segments, segmentsLen, _ := IPv6ParseInput(ipv6)

	// Если есть "embedded IPv4" - валидируем её
	if ipv4 != "" {
		// fmt.Printf("IPv6Normalizer. Try to validate embedded IPv4 segment: %q\n", ipv4)
		if ipv4, _, err = IPv4Normalizer(ipv4, ""); err != nil {
			return "", "", 0, err
		}
		// fmt.Printf("IPv6Normalizer. embedded IPv4 segment (%q) is valid IPv4 address\n", ipv4)

		// Последний сегмент - это IPv4. Из количества сегментов вычитаем 1
		segmentsLen = segmentsLen - 1

		// В случае embedded IPv4 - количество сегментов ipv6 не может быть более 6
		maxSegmentsLen = 6
	}

	// Проверяем количество сегментов. Если менее 3 или больше maxSegmentsLen - ошибка
	if segmentsLen < 3 || segmentsLen > maxSegmentsLen {
		// Точно не ipv6 адрес.
		return "", "", 0, fmt.Errorf(errMsg, ipv6)
	}

	if IPv6AddressFiller(&normAddress, segments, segmentsLen); err != nil {
		return "", "", 0, err
	}

	// fmt.Printf("IPv6Normalizer. Normalized address: %s\n", normAddress)

	// Заменяем более двух двоеточий на два двоеточия
	normAddress = strings.TrimSuffix(regexp.MustCompile(":{2,}").ReplaceAllString(fmt.Sprintf("%s%s", normAddress, ipv4), "::"), ":")

	if !ValidateIPv6SegmentsLen(normAddress, ipv4) {
		return "", "", 0, fmt.Errorf(errMsg, normAddress)
	}

	// fmt.Printf("IPv6Normalizer. RESULT: Normalized address: %s\n", normAddress)

	if zone != "" && !CheckInvalidChars(zone, "ipv6zone") {
		// fmt.Printf("IPv6Normalizer. Zone: %s\n", zone)
		normZone = zone
	}

	if port != "" {
		if normPort, err = ValidatePort(port); err != nil {
			return "", "", 0, err
		}
	}

	return normAddress, normZone, normPort, nil

}

func DomainNormalizer(domain, ext, port string) (normAddress string, normPort int, err error) {

	if CheckInvalidChars(domain, "domain") {
		return "", 0, fmt.Errorf("DomainNormalizer. Invalid characters in domain name: %s", domain)
	}

	normAddress, err = ValidateDomain(domain)

	if err != nil {
		// fmt.Printf("DomainNormalizer. Validation domain error: %s\n", domain)
		return "", 0, err
	}

	if port != "" {
		if normPort, err = ValidatePort(port); err != nil {
			// fmt.Printf("DomainNormalizer. Validation port error: %s\n", domain)
			return "", 0, err
		}
	}

	// fmt.Printf("DomainNormalizer. Validation [SUCCESS]: Domain: %s ; Port: %d\n", normAddress, normPort)

	return normAddress, normPort, nil
}
