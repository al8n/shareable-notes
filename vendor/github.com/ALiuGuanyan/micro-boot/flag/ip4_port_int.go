package flag

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const errPortInt = "not a valid IPv4 port number"

type ip4PortIntValue struct {
	value *int
}

func newIP4PortIntValue(val int, p *int) *ip4PortIntValue {
	ssv := new(ip4PortIntValue)
	ssv.value = p
	*ssv.value = val
	return ssv
}

func (s *ip4PortIntValue) Set(val string) (err error) {
	var(
		p64 int64
	)

	value := strings.TrimSpace(val)
	p64, err = strconv.ParseInt(value, 10, 64)
	if err != nil {
		return errors.New(errPortInt)
	}
	if p64 > 65535 || p64 < 0 {
		return errors.New(errPortInt)
	}

	*s.value = int(p64)
	return nil
}

func (s *ip4PortIntValue) Type() string {
	return "ip4PortInt"
}

func (s *ip4PortIntValue) String() string {
	return fmt.Sprintf("%d", *s.value)
}

func ip4PortIntValueConv(val string) (interface{}, error) {
	value := strings.TrimSpace(val)
	p, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return nil, errors.New(errPortInt)
	}
	if p > 65535 || p < 0 {
		return nil, errors.New(errPortInt)
	}
	return int(p), nil
}

// GetIPv4PortInt return the uint value of a flag with the given name
func (f *FlagSet) GetIPv4PortInt(name string) (int, error) {
	val, err := f.getFlagType(name, "ip4PortInt", ip4PortIntValueConv)
	if err != nil {
		return 0, err
	}
	return val.(int), nil
}

// IPv4PortIntVar defines a string flag with specified name, default value, and usage string.
func (f *FlagSet) IPv4PortIntVar(p *int, name string, value int, usage string) {
	f.fs.VarP(newIP4PortIntValue(value, p), name, "", usage)
}

// IPv4PortIntVarP is like IPv4PortIntVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPv4PortIntVarP(p *int, name, shorthand string, value int, usage string) {
	f.fs.VarP(newIP4PortIntValue(value, p), name, shorthand, usage)
}

// IPv4PortInt defines a string flag with specified name, default value, and usage string.
func (f *FlagSet) IPv4PortInt(name string, value int, usage string) *int {
	p := new(int)
	f.IPv4PortIntVarP(p, name, "", value, usage)
	return p
}

// IPv4PortIntP is like IPv4PortInt, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) IPv4PortIntP(name, shorthand string, value int, usage string) *int {
	p := new(int)
	f.IPv4PortIntVarP(p, name, shorthand, value, usage)
	return p
}