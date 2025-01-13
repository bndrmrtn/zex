package zex

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/google/uuid"
)

// RouterParamValidator is an interface that allows you to register custom route parameter validators.
type RouterParamValidator interface {
	// RouterParamValidator is an interface that allows you to register custom route parameter validators.
	// default validators: int, bool, uuid, alpha, alphanumeric.
	RegisterRouteParamValidator(name string, fn RouteParamValidatorFunc)
}

// RouteParamValidatorFunc is a function that validates a route parameter.
type RouteParamValidatorFunc func(value string) (string, error)

// registerDefaultRouteValidators registers the default route parameter validators.
func registerDefaultRouteValidators(router RouterParamValidator) {
	router.RegisterRouteParamValidator("int", validateInt)
	router.RegisterRouteParamValidator("bool", validateBool)
	router.RegisterRouteParamValidator("uuid", validateUUIDv4)
	router.RegisterRouteParamValidator("alpha", validateAlpha)
	router.RegisterRouteParamValidator("alphanumeric", validateAlphaNumeric)
}

// validateInt validates an integer.
func validateInt(value string) (string, error) {
	_, err := strconv.Atoi(value)
	if err != nil {
		return "", err
	}
	return value, nil
}

// validateBool validates a boolean.
func validateBool(value string) (string, error) {
	_, err := strconv.ParseBool(value)
	if err != nil {
		return "", err
	}
	return value, nil
}

// validateUUIDv4 validates a UUIDv4.
func validateUUIDv4(value string) (string, error) {
	_, err := uuid.Parse(value)
	if err != nil {
		return "", err
	}
	return value, nil
}

// validateAlpha validates an alpha string.
func validateAlpha(value string) (string, error) {
	ok := regexp.MustCompile("^[a-zA-Z]+$").MatchString(value)
	if !ok {
		return "", errors.New("param is not alpha")
	}
	return value, nil
}

// validateAlphaNumeric validates an alphanumeric string.
func validateAlphaNumeric(value string) (string, error) {
	ok := regexp.MustCompile("^[a-zA-Z0-9]+$").MatchString(value)
	if !ok {
		return "", errors.New("param is not alphanumeric")
	}
	return value, nil
}
