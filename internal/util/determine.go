package util

import (
	"fmt"
	"regexp"
)

func DetermineAddress(address string) (Address, Zone string, Port int, err error) {

	var (
		resultBad    = "DetermineAddress. RESULT: [FAILED] - Bad address: [%s]\n\n"
		validAddress string
		validPort    int
		validZone    string
	)

	// Проверяем содержит ли строка адреса ASCII non printable characters
	//if !CheckInvalidChars(address, "address") {
	if CheckNonPrintableASCII(address) {
		// fmt.Println("DetermineAddress. NonPrintableASCII check is [FAILED]")
		//fmt.Printf(resultBad, address)
		return validAddress, validZone, validPort, fmt.Errorf(resultBad, address)
	}

	//fmt.Println("DetermineAddress. BadSymbols check is [PASSED]")

	addrType, addrIPv6, addrIPv4, domain, zone, ext, port, err := parseAddress(address)

	if err != nil {
		return "", "", 0, err
	}

	// fmt.Printf("DetermineAddress. addrType: %s\n", addrType)

	switch addrType {
	case "uriIPv6", "IPv6":
		validAddress, validZone, validPort, err = IPv6Normalizer(addrIPv6, addrIPv4, zone, port)
		if err != nil {
			return "", "", 0, err
		}

	case "IPv4":
		validAddress, validPort, err = IPv4Normalizer(addrIPv4, port)
		if err != nil {
			return "", "", 0, err
		}

	case "fqdn":
		validAddress, validPort, err = DomainNormalizer(domain, ext, port)
		if err != nil {
			return "", "", 0, err
		}

	default:
		// В остальных случаях - возвращаем ошибку
		return "", "", 0, fmt.Errorf("DetermineAddress. error: bad address: %s", address)
	}

	// fmt.Printf("DetermineAddress. Valid address: %s ; Valid zone: %s ; Valid port: %d\n", validAddress, validZone, validPort)
	return validAddress, validZone, validPort, nil
}

func parseAddress(address string) (Type, AddrIPv6, AddrIPv4, Domain, Zone, Ext, Port string, err error) {

	// Проверяем содержит ли строка адреса URI(значения в квадратных скобках), Порт(значения после двоеточия) и Zone ID Index[RFC6874] (значения после знака процента)
	// Собираем регулярку для проверки значения в квадратных скобках.

	addrRe := regexp.MustCompile(`(^((?P<proto1>(http|https))(:\/\/))?((?P<subdomain1>www)\.)?\[(?P<addr1ipv6>(([0-9a-fA-F]{0,4}:){2,8}[0-9a-fA-F]{0,4}?)((?P<addr1ipv4>(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3,3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]){1,3}))?)?(%(?P<zone1>.*))?\]:?(?P<port1>\d{1,5})?$)|(^((?P<proto2>(http|https))(:\/\/))?((?P<subdomain2>www)\.)?(?P<addr2ipv6>(([0-9a-fA-F]{0,4}:){2,8}[0-9a-fA-F]{0,4}?)((?P<addr2ipv4>(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3,3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]){1,3}))?)?(%(?P<zone2>.*))?$)|(^((?P<proto3>(http|https))(:\/\/))?((?P<subdomain3>www)\.)?(?P<addr3ipv4>(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3,3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]){1,3}):?(?P<port3>\d{1,5})?$)|(^((?P<proto4>(http|https))(:\/\/))?((?P<subdomain4>www)\.)?(?P<addr4>(([a-zA-Zа-яА-Я0-9_-]{1,61}\.)*(?P<ext4>[A-Za-zА-Яа-я]{2,}))):?(?P<port4>\d{1,5})?$)`)

	// Проверяем содержит ли строка адреса значения по регулярке
	// foundedValues := addrURIRe.FindStringSubmatch(address)
	foundedValues := addrRe.FindStringSubmatch(address)

	if len(foundedValues) > 0 {
		// fmt.Println("Determine address. values founded.")

		Proto1 := foundedValues[addrRe.SubexpIndex("proto1")]
		Subdomain1 := foundedValues[addrRe.SubexpIndex("subdomain1")]
		Addr1IPv6 := foundedValues[addrRe.SubexpIndex("addr1ipv6")]
		Addr1IPv4 := foundedValues[addrRe.SubexpIndex("addr1ipv4")]
		Zone1 := foundedValues[addrRe.SubexpIndex("zone1")]
		Port1 := foundedValues[addrRe.SubexpIndex("port1")]

		Proto2 := foundedValues[addrRe.SubexpIndex("proto2")]
		Subdomain2 := foundedValues[addrRe.SubexpIndex("subdomain2")]
		Addr2IPv6 := foundedValues[addrRe.SubexpIndex("addr2ipv6")]
		Addr2IPv4 := foundedValues[addrRe.SubexpIndex("addr2ipv4")]
		Zone2 := foundedValues[addrRe.SubexpIndex("zone2")]

		Proto3 := foundedValues[addrRe.SubexpIndex("proto3")]
		Subdomain3 := foundedValues[addrRe.SubexpIndex("subdomain3")]
		Addr3IPv4 := foundedValues[addrRe.SubexpIndex("addr3ipv4")]
		Port3 := foundedValues[addrRe.SubexpIndex("port3")]

		Proto4 := foundedValues[addrRe.SubexpIndex("proto4")]
		Subdomain4 := foundedValues[addrRe.SubexpIndex("subdomain4")]
		Addr4 := foundedValues[addrRe.SubexpIndex("addr4")]
		Ext4 := foundedValues[addrRe.SubexpIndex("ext4")]
		Port4 := foundedValues[addrRe.SubexpIndex("port4")]

		//uriAddr := foundedValues[addrURIRe.SubexpIndex("address")]
		//uriZone := foundedValues[addrURIRe.SubexpIndex("zone")]
		//uriPort := foundedValues[addrURIRe.SubexpIndex("port")]

		// fmt.Printf("Determine address. Input: %s\n", address)

		/*
			fmt.Printf("Determine address. Founded URI address: %s\n", uriAddr)
			fmt.Printf("Determine address. Founded URI Zone: %s\n", uriZone)
			fmt.Printf("Determine address. Founded URI Port: %s\n", uriPort)
		*/

		/*
			fmt.Println("================================================================================")
			fmt.Printf("Determine address. URI IPv6 protocol:\t\t%s\n", Proto1)
			fmt.Printf("Determine address. URI IPv6 subdomain:\t\t%s\n", Subdomain1)
			fmt.Printf("Determine address. URI IPv6 address:\t\t%s\n", Addr1IPv6)
			fmt.Printf("Determine address. URI IPv6 embedded IPv4:\t%s\n", Addr1IPv4)
			fmt.Printf("Determine address. URI IPv6 ZoneID:\t\t%s\n", Zone1)
			fmt.Printf("Determine address. URI IPV6 Port:\t\t%s\n", Port1)
			fmt.Println("================================================================================")
			fmt.Printf("Determine address. IPv6 protocol:\t\t%s\n", Proto2)
			fmt.Printf("Determine address. IPv6 subdomain:\t\t%s\n", Subdomain2)
			fmt.Printf("Determine address. IPv6 address:\t\t%s\n", Addr2IPv6)
			fmt.Printf("Determine address. IPv6 embedded IPv4:\t%s\n", Addr2IPv4)
			fmt.Printf("Determine address. IPv6 ZoneID:\t\t\t%s\n", Zone2)
			fmt.Println("================================================================================")
			fmt.Printf("Determine address. IPv4 protocol:\t\t%s\n", Proto3)
			fmt.Printf("Determine address. IPv4 subdomain:\t\t%s\n", Subdomain3)
			fmt.Printf("Determine address. IPv4 address:\t\t%s\n", Addr3IPv4)
			fmt.Printf("Determine address. IPv4 Port:\t\t\t%s\n", Port3)
			fmt.Println("================================================================================")
			fmt.Printf("Determine address. FQDN protocol:\t\t%s\n", Proto4)
			fmt.Printf("Determine address. FQDN subdomain:\t\t%s\n", Subdomain4)
			fmt.Printf("Determine address. FQDN:\t\t\t%s\n", Addr4)
			fmt.Printf("Determine address. FQDN Ext:\t\t\t%s\n", Ext4)
			fmt.Printf("Determine address. FQDN Port:\t\t\t%s\n", Port4)
			fmt.Println("================================================================================")
		*/

		if Proto1 != "" || Proto2 != "" || Proto3 != "" || Proto4 != "" || Subdomain1 != "" || Subdomain2 != "" || Subdomain3 != "" || Subdomain4 != "" {
			// fmt.Printf("DetermineAddress. parseAddress ERROR: address is web resource: %s\n", address)
			//fmt.Printf(resultBad, address)
			return "", "", "", "", "", "", "", fmt.Errorf("DetermineAddress. parseAddress ERROR: address is web resource: %s", address)
		}

		// Type
		// AddrIPv6
		// AddrIPv4
		// Domain
		// Zone
		// Ext
		// Port
		// err

		if Addr1IPv6 != "" {
			return "uriIPv6", Addr1IPv6, Addr1IPv4, "", Zone1, "", Port1, nil
		} else if Addr2IPv6 != "" {
			return "IPv6", Addr2IPv6, Addr2IPv4, "", Zone2, "", "", nil
		} else if Addr3IPv4 != "" {
			return "IPv4", "", Addr3IPv4, "", "", "", Port3, nil
		} else if Addr4 != "" {
			return "fqdn", "", "", Addr4, "", Ext4, Port4, nil
		}
	}

	return "", "", "", "", "", "", "", fmt.Errorf("DetermineAddress. parseAddress ERROR: invalid address: %s", address)
}
