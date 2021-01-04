package flag

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
)

// -- stringToFloat64 Value
type stringToFloat64Value struct {
	value   *map[string]float64
	changed bool
}

func newStringToFloat64Value(val map[string]float64, p *map[string]float64) *stringToFloat64Value {
	ssv := new(stringToFloat64Value)
	ssv.value = p
	*ssv.value = val
	return ssv
}

// Format: a=1.9,b=3.6
func (s *stringToFloat64Value) Set(val string) error {
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

	out := make(map[string]float64, len(ss))
	for _, pair := range ss {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", pair)
		}



		val, err := strconv.ParseFloat(strings.TrimSpace(kv[1]), 64)
		if err != nil {
			return fmt.Errorf("%s cannot be parsed as float64", kv[1])
		}

		out[strings.TrimSpace(kv[0])] = val
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

func (s *stringToFloat64Value) Type() string {
	return "stringToFloat64"
}

func (s *stringToFloat64Value) String() string {
	records := make([]string, 0, len(*s.value)>>1)
	for k, v := range *s.value {
		records = append(records, fmt.Sprintf("%s=%f", k, v))
	}

	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	if err := w.Write(records); err != nil {
		panic(err)
	}
	w.Flush()
	return "[" + strings.TrimSpace(buf.String()) + "]"
}

func stringToFloat64ValueConv(val string) (interface{}, error) {
	val = strings.Trim(val, "[]")
	// An empty string would cause an empty map
	if len(val) == 0 {
		return map[string]float64{}, nil
	}
	r := csv.NewReader(strings.NewReader(val))
	ss, err := r.Read()
	if err != nil {
		return nil, err
	}
	out := make(map[string]float64, len(ss))
	for _, pair := range ss {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("%s must be formatted as key=value", pair)
		}


		val, err := strconv.ParseFloat(strings.TrimSpace(kv[1]), 64)
		if err != nil {
			return nil, fmt.Errorf("%s cannot be parsed as float64", kv[1])
		}
		out[strings.TrimSpace(kv[0])] = val
	}
	return out, nil
}

// GetStringToFloat64 return the map[string]float64 value of a flag with the given name
func (f *FlagSet) GetStringToFloat64(name string) (map[string]float64, error) {
	val, err := f.getFlagType(name, "stringToFloat64", stringToFloat64ValueConv)
	if err != nil {
		return map[string]float64{}, err
	}
	return val.(map[string]float64), nil
}

// StringToFloat64Var defines a string flag with specified name, default value, and usage string.
// The argument p points to a map[string]float64 variable in which to store the values of the multiple flags.
// The value of each argument will not try to be separated by comma
func (f *FlagSet) StringToFloat64Var(p *map[string]float64, name string, value map[string]float64, usage string) {
	f.fs.VarP(newStringToFloat64Value(value, p), name, "", usage)
}

// StringToFloat64VarP is like StringToFloat64Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringToFloat64VarP(p *map[string]float64, name, shorthand string, value map[string]float64, usage string) {
	f.fs.VarP(newStringToFloat64Value(value, p), name, shorthand, usage)
}

// StringToFloat64 defines a string flag with specified name, default value, and usage string.
// The return value is the address of a map[string]float64 variable that stores the value of the flag.
// The value of each argument will not try to be separated by comma
func (f *FlagSet) StringToFloat64(name string, value map[string]float64, usage string) *map[string]float64 {
	p := map[string]float64{}
	f.StringToFloat64VarP(&p, name, "", value, usage)
	return &p
}

// StringToFloat64P is like StringToFloat64, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringToFloat64P(name, shorthand string, value map[string]float64, usage string) *map[string]float64 {
	p := map[string]float64{}
	f.StringToFloat64VarP(&p, name, shorthand, value, usage)
	return &p
}

