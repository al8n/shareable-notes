package flag

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const errPortString = "not a valid IPv4 port number"

type ip4PortStringValue struct {
	value *string
}

func newIP4PortStringValue(val string, p *string) *ip4PortStringValue {
	ssv := new(ip4PortStringValue)
	ssv.value = p
	*ssv.value = val
	return ssv
}

func (s *ip4PortStringValue) Set(val string) (err error) {
	value := strings.TrimSpace(val)
	p64, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return errors.New(errPortString)
	}
	if p64 > 65535 {
		return errors.New(errPortString)
	}
	*s.value = value
	return nil
}

func (s *ip4PortStringValue) Type() string {
	return "ip4PortString"
}

func (s *ip4PortStringValue) String() string {
	return fmt.Sprintf("%s", *s.value)
}

func ip4PortStringValueConv(val string) (interface{}, error) {
	value := strings.TrimSpace(val)
	p, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return nil, errors.New(errPortString)
	}
	if p > 65535 {
		return nil, errors.New(errPortString)
	}
	return fmt.Sprintf("%d", p), nil
}

// GetIPv4PortString return the uint value of a flag with the given name
func (f *FlagSet) GetIPv4PortString(name string) (string, error) {
	val, err := f.getFlagType(name, "ip4PortString", ip4PortStringValueConv)
	if err != nil {
		return "", err
	}
	return val.(string), nil
}

// IPv4PortStringVar defines a string flag with specified name, default value, and usage string.
func (f *FlagSet) IPv4PortStringVar(p *string, name, value, usage string) {
	f.fs.VarP(newIP4PortStringValue(value, p), name, "", usage)
}

// IPv4PortStringVarP is like IPv4PortStringVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPv4PortStringVarP(p *string, name, shorthand, value, usage string) {
	f.fs.VarP(newIP4PortStringValue(value, p), name, shorthand, usage)
}

// IPv4PortString defines a string flag with specified name, default value, and usage string.
func (f *FlagSet) IPv4PortString(name, value, usage string) *string {
	p := new(string)
	f.IPv4PortStringVarP(p, name, "", value, usage)
	return p
}

// IPv4PortStringP is like IPv4PortString, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPv4PortStringP(name, shorthand, value, usage string) *string {
	p := new(string)
	f.IPv4PortStringVarP(p, name, shorthand, value, usage)
	return p
}
