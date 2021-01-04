package flag

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
)

// -- float64ToFloat64 Value
type float64ToFloat64Value struct {
	value   *map[float64]float64
	changed bool
}

func newFloat64ToFloat64Value(val map[float64]float64, p *map[float64]float64) *float64ToFloat64Value {
	ssv := new(float64ToFloat64Value)
	ssv.value = p
	*ssv.value = val
	return ssv
}

// Format: 1.7=1.9,2.4=3.6
func (s *float64ToFloat64Value) Set(val string) error {
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

	out := make(map[float64]float64, len(ss))
	for _, pair := range ss {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", pair)
		}

		key, err := strconv.ParseFloat(strings.TrimSpace(kv[0]), 64)
		if err != nil {
			return fmt.Errorf("%s cannot be parsed as float64", kv[0])
		}

		val, err := strconv.ParseFloat(strings.TrimSpace(kv[1]), 64)
		if err != nil {
			return fmt.Errorf("%s cannot be parsed as float64", kv[1])
		}

		out[key] = val
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

func (s *float64ToFloat64Value) Type() string {
	return "float64ToFloat64"
}

func (s *float64ToFloat64Value) String() string {
	records := make([]string, 0, len(*s.value)>>1)
	for k, v := range *s.value {
		records = append(records, fmt.Sprintf("%f=%f", k, v))
	}

	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	if err := w.Write(records); err != nil {
		panic(err)
	}
	w.Flush()
	return "[" + strings.TrimSpace(buf.String()) + "]"
}

func float64ToFloat64ValueConv(val string) (interface{}, error) {
	val = strings.Trim(val, "[]")
	// An empty string would cause an empty map
	if len(val) == 0 {
		return map[float64]float64{}, nil
	}
	r := csv.NewReader(strings.NewReader(val))
	ss, err := r.Read()
	if err != nil {
		return nil, err
	}
	out := make(map[float64]float64, len(ss))
	for _, pair := range ss {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("%s must be formatted as key=value", pair)
		}
		key, err := strconv.ParseFloat(strings.TrimSpace(kv[0]), 64)
		if err != nil {
			return nil, fmt.Errorf("%s cannot be parsed as float64", kv[0])
		}

		val, err := strconv.ParseFloat(strings.TrimSpace(kv[1]), 64)
		if err != nil {
			return nil, fmt.Errorf("%s cannot be parsed as float64", kv[1])
		}
		out[key] = val
	}
	return out, nil
}

// GetFloat64ToFloat64 return the map[float64]float64 value of a flag with the given name
func (f *FlagSet) GetFloat64ToFloat64(name string) (map[float64]float64, error) {
	val, err := f.getFlagType(name, "float64ToFloat64", float64ToFloat64ValueConv)
	if err != nil {
		return map[float64]float64{}, err
	}
	return val.(map[float64]float64), nil
}

// Float64ToFloat64Var defines a string flag with specified name, default value, and usage string.
// The argument p points to a map[float64]float64 variable in which to store the values of the multiple flags.
// The value of each argument will not try to be separated by comma
func (f *FlagSet) Float64ToFloat64Var(p *map[float64]float64, name string, value map[float64]float64, usage string) {
	f.fs.VarP(newFloat64ToFloat64Value(value, p), name, "", usage)
}

// Float64ToFloat64VarP is like Float64ToFloat64Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Float64ToFloat64VarP(p *map[float64]float64, name, shorthand string, value map[float64]float64, usage string) {
	f.fs.VarP(newFloat64ToFloat64Value(value, p), name, shorthand, usage)
}

// Float64ToFloat64 defines a string flag with specified name, default value, and usage string.
// The return value is the address of a map[float64]float64 variable that stores the value of the flag.
// The value of each argument will not try to be separated by comma
func (f *FlagSet) Float64ToFloat64(name string, value map[float64]float64, usage string) *map[float64]float64 {
	p := map[float64]float64{}
	f.Float64ToFloat64VarP(&p, name, "", value, usage)
	return &p
}

// Float64ToFloat64P is like Float64ToFloat64, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) Float64ToFloat64P(name, shorthand string, value map[float64]float64, usage string) *map[float64]float64 {
	p := map[float64]float64{}
	f.Float64ToFloat64VarP(&p, name, shorthand, value, usage)
	return &p
}



