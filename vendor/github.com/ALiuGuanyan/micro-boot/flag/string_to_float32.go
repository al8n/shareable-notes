package flag

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
)

// -- stringToFloat32 Value
type stringToFloat32Value struct {
	value   *map[string]float32
	changed bool
}

func newStringToFloat32Value(val map[string]float32, p *map[string]float32) *stringToFloat32Value {
	ssv := new(stringToFloat32Value)
	ssv.value = p
	*ssv.value = val
	return ssv
}

// Format: a=1.9,b=3.6
func (s *stringToFloat32Value) Set(val string) error {
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

	out := make(map[string]float32, len(ss))
	for _, pair := range ss {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", pair)
		}



		val, err := strconv.ParseFloat(strings.TrimSpace(kv[1]), 32)
		if err != nil {
			return fmt.Errorf("%s cannot be parsed as float32", kv[1])
		}

		out[strings.TrimSpace(kv[0])] = float32(val)
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

func (s *stringToFloat32Value) Type() string {
	return "stringToFloat32"
}

func (s *stringToFloat32Value) String() string {
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

func stringToFloat32ValueConv(val string) (interface{}, error) {
	val = strings.Trim(val, "[]")
	// An empty string would cause an empty map
	if len(val) == 0 {
		return map[string]float32{}, nil
	}
	r := csv.NewReader(strings.NewReader(val))
	ss, err := r.Read()
	if err != nil {
		return nil, err
	}
	out := make(map[string]float32, len(ss))
	for _, pair := range ss {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("%s must be formatted as key=value", pair)
		}


		val, err := strconv.ParseFloat(strings.TrimSpace(kv[1]), 32)
		if err != nil {
			return nil, fmt.Errorf("%s cannot be parsed as float32", kv[1])
		}
		out[strings.TrimSpace(kv[0])] = float32(val)
	}
	return out, nil
}

// GetStringToFloat32 return the map[string]float32 value of a flag with the given name
func (f *FlagSet) GetStringToFloat32(name string) (map[string]float32, error) {
	val, err := f.getFlagType(name, "stringToFloat32", stringToFloat32ValueConv)
	if err != nil {
		return map[string]float32{}, err
	}
	return val.(map[string]float32), nil
}

// StringToFloat32Var defines a string flag with specified name, default value, and usage string.
// The argument p points to a map[string]float32 variable in which to store the values of the multiple flags.
// The value of each argument will not try to be separated by comma
func (f *FlagSet) StringToFloat32Var(p *map[string]float32, name string, value map[string]float32, usage string) {
	f.fs.VarP(newStringToFloat32Value(value, p), name, "", usage)
}

// StringToFloat32VarP is like StringToFloat32Var, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringToFloat32VarP(p *map[string]float32, name, shorthand string, value map[string]float32, usage string) {
	f.fs.VarP(newStringToFloat32Value(value, p), name, shorthand, usage)
}

// StringToFloat32 defines a string flag with specified name, default value, and usage string.
// The return value is the address of a map[string]float32 variable that stores the value of the flag.
// The value of each argument will not try to be separated by comma
func (f *FlagSet) StringToFloat32(name string, value map[string]float32, usage string) *map[string]float32 {
	p := map[string]float32{}
	f.StringToFloat32VarP(&p, name, "", value, usage)
	return &p
}

// StringToFloat32P is like StringToFloat32, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) StringToFloat32P(name, shorthand string, value map[string]float32, usage string) *map[string]float32 {
	p := map[string]float32{}
	f.StringToFloat32VarP(&p, name, shorthand, value, usage)
	return &p
}
