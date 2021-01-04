package flag

import (
	"encoding/csv"
	"fmt"
	"go.mongodb.org/mongo-driver/tag"
	"strings"
)

// -- mongoTagSet Value
type mongoTagSetValue struct {
	value   *tag.Set
	changed bool
}

func newMongoTagSetValue(val tag.Set, p *tag.Set) *mongoTagSetValue {
	isv := new(mongoTagSetValue)
	isv.value = p
	*isv.value = val
	return isv
}

func (s *mongoTagSetValue) Set(val string) error {
	ss := strings.Split(val, ",")
	n := strings.Count(val, "=")
	switch n {
	case 0:
		return fmt.Errorf("%s must be formatted as name=value", val)
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
	
	out := make(tag.Set, len(ss))
	for i, pair := range ss {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", pair)
		}
		
		out[i] = tag.Tag{
			Name: kv[0],
			Value: kv[1],
		}
	}
	if !s.changed {
		*s.value = out
	} else {
		*s.value = append(*s.value, out...)
	}
	s.changed = true
	return nil
}

func (s *mongoTagSetValue) Type() string {
	return "mongoTagSet"
}

func (s *mongoTagSetValue) String() string {
	out := make([]string, len(*s.value))
	for i, d := range *s.value {
		out[i] = fmt.Sprintf("%s=%s", d.Name, d.Value)
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (s *mongoTagSetValue) Append(val string) error {
	ss := strings.Split(val, ",")
	n := strings.Count(val, "=")
	switch n {
	case 0:
		return fmt.Errorf("%s must be formatted as name=value", val)
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

	for _, pair := range ss {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", pair)
		}
		*s.value = append(*s.value, tag.Tag{
			Name:  kv[0],
			Value: kv[1],
		})
	}


	return nil
}

func (s *mongoTagSetValue) Replace(val []string) error {
	var out tag.Set
	for _, d := range val {
		ss := strings.Split(d, ",")
		n := strings.Count(d, "=")
		switch n {
		case 0:
			return fmt.Errorf("%s must be formatted as name=value", d)
		case 1:
			ss = append(ss, strings.Trim(d, `"`))
		default:
			r := csv.NewReader(strings.NewReader(d))
			var err error
			ss, err = r.Read()
			if err != nil {
				return err
			}
		}

		innerOut := make([]tag.Tag, len(ss))
		for i, pair := range ss {
			kv := strings.SplitN(pair, "=", 2)
			if len(kv) != 2 {
				return fmt.Errorf("%s must be formatted as key=value", pair)
			}

			innerOut[i] = tag.Tag{
				Name: kv[0],
				Value: kv[1],
			}
		}

		out = append(out, innerOut...)
	}
	*s.value = out
	return nil
}

func (s *mongoTagSetValue) GetSlice() []string {
	out := make([]string, len(*s.value))
	for i, d := range *s.value {
		out[i] = fmt.Sprintf("%s=%s", d.Name, d.Value)
	}
	return out
}

func mongoTagSetConv(val string) (interface{}, error) {
	val = strings.Trim(val, "[]")
	// Empty string would cause a slice with one (empty) entry
	if len(val) == 0 {
		return tag.Set{}, nil
	}
	ss := strings.Split(val, ",")
	n := strings.Count(val, "=")
	switch n {
	case 0:
		return nil, fmt.Errorf("%s must be formatted as name=value", val)
	case 1:
		ss = append(ss, strings.Trim(val, `"`))
	default:
		r := csv.NewReader(strings.NewReader(val))
		var err error
		ss, err = r.Read()
		if err != nil {
			return nil, err
		}
	}

	out := make(tag.Set, len(ss))
	for i, pair := range ss {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("%s must be formatted as key=value", pair)
		}
		out[i] = tag.Tag{
			Name: kv[0],
			Value: kv[1],
		}
	}
	return out, nil
}

// GetMongoTagSet return the tag.Set value of a flag with the given name
func (f *FlagSet) GetMongoTagSet(name string) (tag.Set, error) {
	val, err := f.getFlagType(name, "mongoTagSet", mongoTagSetConv)
	if err != nil {
		return []tag.Tag{}, err
	}
	return val.([]tag.Tag), nil
}

// MongoTagSetVar defines a intSlice flag with specified name, default value, and usage string.
// The argument p points to a tag.Set variable in which to store the value of the flag.
func (f *FlagSet) MongoTagSetVar(p *tag.Set, name string, value tag.Set, usage string) {
	f.VarP(newMongoTagSetValue(value, p), name, "", usage)
}

// MongoTagSetVarP is like MongoTagSetVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) MongoTagSetVarP(p *tag.Set, name, shorthand string, value tag.Set, usage string) {
	f.VarP(newMongoTagSetValue(value, p), name, shorthand, usage)
}

// MongoTagSet defines a tag.Set flag with specified name, default value, and usage string.
// The return value is the address of a tag.Set variable that stores the value of the flag.
func (f *FlagSet) MongoTagSet(name string, value tag.Set, usage string) *tag.Set {
	p := tag.Set{}
	f.MongoTagSetVarP(&p, name, "", value, usage)
	return &p
}

// MongoTagSetP is like MongoTagSet, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) MongoTagSetP(name, shorthand string, value tag.Set, usage string) *tag.Set {
	p := tag.Set{}
	f.MongoTagSetVarP(&p, name, shorthand, value, usage)
	return &p
}