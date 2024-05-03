package util

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/go-playground/validator/v10"
)

// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------

func Init() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

// ----------------------------------------------------------------------------

// Validator - объявляем интерфейс валидатора
/* type Validator interface {
	Validate(interface{}) (bool, error) // Сигнатура метода Validate, принимает на вход значение валидируемого интерфейса и возвращает булево значение и ошибку (если она есть)
} */

func ValidateStruct(ctx context.Context, st interface{}) {
	//logging.L(ctx).Debug("v_validator: validate", logging.AnyAttr("struct", st))
	if err := L(ctx).Struct(st); err != nil {
		ProcessValidationErrors(ctx, err)
		os.Exit(1)
	}
}

// ----------------------------------------------------------------------------

func ProcessValidationErrors(ctx context.Context, err error) {

	// // this check is only needed when your code could produce
	// // an invalid value for validation such as interface with nil
	// // value most including myself do not usually have code like this.
	// if _, ok := err.(*validator.InvalidValidationError); ok {
	// 	fmt.Println(err)
	// 	return
	// }

	fmt.Println("\nvalidation errors:")

	for _, err := range err.(validator.ValidationErrors) {
		//fmt.Println()
		//logging.L(ctx).Error("main: failed to validate configuration", logging.ErrAttr(err))

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		//fmt.Fprintln(w, "StructNamespace:\t", err.StructNamespace(), "\t| StructField:\t", err.StructField(), "\t| ActualTag:\t", err.ActualTag(), "\t| Type:\t", err.Type(), "\t| Accepted values:\t", err.Param())
		//fmt.Fprintln(w, "Namespace:\t", err.Namespace(), "\t| Field:\t", err.Field(), "\t| Tag:\t", err.Tag(), "\t| Kind:\t", err.Kind(), "\t| Value:\t", err.Value())
		fmt.Fprintln(w, "Field:\t", err.Field(), "\t| Tag:\t", err.Tag(), "\t| Kind:\t", err.Kind(), "\t| Value:\t", err.Value())
		w.Flush()
	}
	fmt.Println()

	// errors := err.(validator.ValidationErrors)
	// logging.L(ctx).Error("main: failed to validate configuration", logging.ErrAttr(errors))

	os.Exit(1)
}

type ctxValidator struct{}

// ContextWithLogger adds logger to context.
func ContextWithValidator(ctx context.Context, v *validator.Validate) context.Context {
	return context.WithValue(ctx, ctxValidator{}, v)
}

/* func LToContext(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, l)
} */

// loggerFromContext returns logger from context.
func validatorFromContext(ctx context.Context) *validator.Validate {
	if v, ok := ctx.Value(ctxValidator{}).(*validator.Validate); ok {
		return v
	}
	return validator.New(validator.WithRequiredStructEnabled())
}

func L(ctx context.Context) *validator.Validate {
	return validatorFromContext(ctx)
}

func ValidateName(name string) bool {
	// Собираем регулярку для фильтрации имени.
	patternName := regexp.MustCompile(`^[A-ZА-Яa-zа-я0-9._-]+$`)
	//patternName := regexp.MustCompile(`^[A-ZА-Яa-zа-я0-9._-]*$`)
	return patternName.MatchString(name)
}

func ValidateDomain(domain string) (string, error) {

	const errMsg = "invalid domain name: %s"

	var (
		// Регулярки для фильтрации.
		validPart = regexp.MustCompile(`^[a-zA-Zа-яА-Я0-9-_]+$`)
		validLast = regexp.MustCompile(`^[a-zA-Zа-яА-Я]+$`)
		// validPart = regexp.MustCompile(`^[a-zA-Zа-яА-Я0-9-_]*$`)
		// validLast = regexp.MustCompile(`^[a-zA-Zа-яА-Я]*$`)
	)

	if domain == "localhost" {
		return domain, nil
	}

	parts := strings.Split(domain, ".")

	for i, part := range parts {

		if len(parts) < 2 ||
			part == "" ||
			CheckNonPrintableASCII(part) ||
			(i != len(parts)-1 && !validPart.MatchString(part)) ||
			(i == len(parts)-1 && !validLast.MatchString(part)) {
			return "", fmt.Errorf(errMsg, domain)
		}

	}
	return domain, nil
}

func ValidateIPv4Segment(segmentInt int) (int, error) {

	const errMsg = "invalid IPv4 segment: %v"

	if segmentInt < 0 || segmentInt > 255 {
		return 0, fmt.Errorf(errMsg, segmentInt)
	}

	return segmentInt, nil
}

// func ValidateIPv6Segment(segment string) (int64, error) {
func ValidateIPv6Segment(segment string) (string, error) {

	const errMsg = "invalid IPv6 segment: %s"

	//fmt.Printf("ValidateIPv6Segment: %s\n", segment)
	// Если сегмент пустой - возвращаем
	if segment == "" {
		return "", nil
	}
	// Собираем регулярку для фильтрации hexadecimal IPv6 segment.
	pattern := regexp.MustCompile(`^[0-9A-Fa-f]{1,4}$`)

	// Если сегмент не соответствует регулярке - возвращаем ошибку
	if !pattern.MatchString(segment) {
		return "", fmt.Errorf(errMsg, segment)
	}

	// Пробуем конвертировать string segment в int
	// hexValue, err := strconv.ParseInt(segment, 16, 64)

	// if err != nil {
	// 	return 0, err
	// }

	// return hexValue, nil
	return segment, nil
}

func ValidateIPv6SegmentsLen(address, ipv4 string) bool {

	_, segmentsLen, _ := IPv6ParseInput(address)

	if ipv4 != "" && segmentsLen < 7 && strings.Count(address, "::") != 1 {
		return false
	}

	if ipv4 == "" && segmentsLen < 8 && strings.Count(address, "::") != 1 {
		return false
	}

	return true
}

func ValidatePort(portStr string) (int, error) {

	const errMsg = "invalid port: %s"

	// Пробуем конвертировать string portStr в int
	port, err := strconv.Atoi(portStr)
	// Если произошла ошибка или порт меньше 1 или больше 65535 - возвращаем ошибку
	if err != nil || port < 1 || port > 65535 {
		return 0, fmt.Errorf(errMsg, portStr)
	}

	return port, nil
}
