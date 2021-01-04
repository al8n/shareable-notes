package flag

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
)

// -- stringToInt32Value Value
type stringToInt32Value struct {
	value   *map[string]int32
	changed bool
}

func newStringToInt32Value(val map[string]int32, p *map[string]int32) *stringToInt32Value {
	ssv := new(stringToInt32Value)
	ssv.value = p
	*ssv.value = val
	return ssv
}

// Format: a=1.9,b=3.6
func (s *stringToInt32Value) Set(val string) error {
	var ss []string
	n := strings.Count(val, "=")
	switch n {
	case 0:
		return fmt.Errorf("%s must be formatted as key=value", val)
	case 1:
		ss = append(ss, strings.Trim(val, `"`))
	default:
		r := csv.NewReader(strings.NewReader(val))
		var err error
		ss, err = r.Read()
		if err != nil {
			return err
		}
	}

	out := make(map[string]int32, len(ss))
	for _, pair := range ss {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", pair)
		}



		val, err := strconv.ParseInt(strings.TrimSpace(kv[1]), 10, 32)
		if err != nil {
			return fmt.Errorf("%s cannot be parsed as int32", kv[1])
		}

		out[strings.TrimSpace(kv[0])] = int32(val)
	}
	if !s.changed {
		*s.value = out
	} else {
		for k, v := range out {
			(*s.value)[k] = v
		}
	}
	s.changed = true
	return nil
}

func (s *stringToInt32Value) Type() string {
	return "stringToInt32"
}

func (s *stringToInt32Value) String() string {
	records := make([]string, 0, len(*s.value)>>1)
	for k, v := range *s.value {
		records = append(records, fmt.Sprintf("%s=%d", k, v))
	}

	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	if err := w.Write(records); err != nil {
		panic(err)
	}
	w.Flush()
	return "[" + strings.TrimSpace(buf.String()) + "]"
}

func stringToInt32ValueConv(val string) (interface{}, error) {
	val = strings.Trim(val, "[]")
	// An empty string would cause an empty map
	if len(val) == 0 {
		return map[string]int32{}, nil
	}
	r := csv.NewReader(strings.NewReader(val))
	ss, err := r.Read()
	if err != nil {
		return nil, err
	}
	out := make(map[string]int32, len(ss))
	for _, pair := range ss {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("%s must be formatted as key=value", pair)
		}

		val, err := strconv.ParseInt(strings.TrimSpace(kv[1]), 10, 32)
		if err != nil {
			return nil, fmt.Errorf("%s cannot be parsed as int32", kv[1])
		}
		out[strings.TrimSpace(kv[0])] = int32(val)
	}
	return out, nil
}

// GetStringToInt32 return the map[string]int32 value of a flag with the given name
func (f *FlagSet) GetStringToInt32(name string) (map[string]int32, error) {
	val, err := f.getFlagType(name, "stringToInt32", stringToInt32ValueConv)
	if err != nil {
		return map[string]int32{}, err
	}
	return val.(map[string]int32), nil
}

// StringToInt32Var defines a string flag with specified name, default value, and usage string.
// The argument p points to a map[string]int32 variable in which to store the values of the multiple flags.
// The value of each argument will not try to be separated by comma
func (f *FlagSet) StringToInt32Var(p *map[string]int32, name string, value map[string]int32, usage string) {
	f.fs.VarP(newStringToInt32Value(value, p), name, "", usage)
}

// StringToInt32VarP is like StringToInt32Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringToInt32VarP(p *map[string]int32, name, shorthand string, value map[string]int32, usage string) {
	f.fs.VarP(newStringToInt32Value(value, p), name, shorthand, usage)
}

// StringToInt32 defines a string flag with specified name, default value, and usage string.
// The return value is the address of a map[string]int32 variable that stores the value of the flag.
// The value of each argument will not try to be separated by comma
func (f *FlagSet) StringToInt32(name string, value map[string]int32, usage string) *map[string]int32 {
	p := map[string]int32{}
	f.StringToInt32VarP(&p, name, "", value, usage)
	return &p
}

// StringToInt32P is like StringToInt32, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringToInt32P(name, shorthand string, value map[string]int32, usage string) *map[string]int32 {
	p := map[string]int32{}
	f.StringToInt32VarP(&p, name, shorthand, value, usage)
	return &p
}

