package utils

import (
	"errors"
	"fmt"
)


type Methods = int8

type Type = int8

const (
	Request Methods = iota
	Response
)

const (
	GRPC Type =  iota
	Thrift
	HTTP
)

func errorCodecCasting(name string, method Methods) string {
	switch method {
	case Request:
		return fmt.Sprintf("error when casting request in %s", name)
	case Response:
		return fmt.Sprintf("error when casting response in %s", name)
	default:
		return fmt.Sprintf("error when casting in %s", name)
	}
}

func ErrorCodecCasting(name string, method Methods, typ Type) error  {
	switch typ {
	case GRPC:
		return errors.New(fmt.Sprintf("GRPC: %s", errorCodecCasting(name, method)))
	case Thrift:
		return errors.New(fmt.Sprintf("Thrift: %s", errorCodecCasting(name, method)))
	case HTTP:
		return errors.New(fmt.Sprintf("HTTP: %s", errorCodecCasting(name, method)))
	}

	return errors.New(fmt.Sprintf("Unkown: %s", errorCodecCasting(name, method)))
}

// These annoying helper functions are required to translate Go error types to
// and from strings, which is the type we use in our IDLs to represent errors.
// There is special casing to treat empty strings as nil errors.
func Str2Err(s string) error {
	if s == "" {
		return nil
	}
	return errors.New(s)
}

func Err2Str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}