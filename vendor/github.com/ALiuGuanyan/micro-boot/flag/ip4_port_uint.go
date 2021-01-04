package flag

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const errPortUint = "not a valid IPv4 port number"

type ip4PortUintValue struct {
	value *uint
}

func newIP4PortUintValue(val uint, p *uint) *ip4PortUintValue {
	ssv := new(ip4PortUintValue)
	ssv.value = p
	*ssv.value = val
	return ssv
}

func (s *ip4PortUintValue) Set(val string) (err error) {
	var(
		p64 uint64
	)

	value := strings.TrimSpace(val)
	p64, err = strconv.ParseUint(value, 10, 64)
	if err != nil {
		return errors.New(errPortUint)
	}
	if p64 > 65535 {
		return errors.New(errPortUint)
	}

	*s.value = uint(p64)
	return nil
}

func (s *ip4PortUintValue) Type() string {
	return "ip4PortUint"
}

func (s *ip4PortUintValue) String() string {
	return fmt.Sprintf("%d", *s.value)
}

func ip4PortUintValueConv(val string) (interface{}, error) {
	value := strings.TrimSpace(val)
	p, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return nil, errors.New(errPortUint)
	}
	if p > 65535 {
		return nil, errors.New(errPortUint)
	}
	return uint(p), nil
}

// GetIPv4PortUint return the uint value of a flag with the given name
func (f *FlagSet) GetIPv4PortUint(name string) (uint, error) {
	val, err := f.getFlagType(name, "ip4PortUint", ip4PortUintValueConv)
	if err != nil {
		return 0, err
	}
	return val.(uint), nil
}

// IPv4PortUintVar defines a string flag with specified name, default value, and usage string.
func (f *FlagSet) IPv4PortUintVar(p *uint, name string, value uint, usage string) {
	f.fs.VarP(newIP4PortUintValue(value, p), name, "", usage)
}

// IPv4PortUintVarP is like IPv4PortUintVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPv4PortUintVarP(p *uint, name, shorthand string, value uint, usage string) {
	f.fs.VarP(newIP4PortUintValue(value, p), name, shorthand, usage)
}

// IPv4PortUint defines a string flag with specified name, default value, and usage string.
func (f *FlagSet) IPv4PortUint(name string, value uint, usage string) *uint {
	p := new(uint)
	f.IPv4PortUintVarP(p, name, "", value, usage)
	return p
}

// IPv4PortUintP is like IPv4PortUint, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPv4PortUintP(name, shorthand string, value uint, usage string) *uint {
	p := new(uint)
	f.IPv4PortUintVarP(p, name, shorthand, value, usage)
	return p
}


