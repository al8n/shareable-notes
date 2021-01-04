package flag

import (
	"fmt"
	"go.mongodb.org/mongo-driver/tag"
	"strings"
)

const errMongoTagSetSliceFormat = "%s must be formatted as [n1=v1 n2=v2 n3=v3]"

// -- mongoTagSetSlice Value
type mongoTagSetSliceValue struct {
	value   *[]tag.Set
	changed bool
}

func newMongoTagSetSliceValue(val []tag.Set, p *[]tag.Set) *mongoTagSetSliceValue {
	isv := new(mongoTagSetSliceValue)
	isv.value = p
	*isv.value = val
	return isv
}

// Format:
//
// 1. [n1=v1 n2=v2], [n3=v3 n4=v4 n5=v5]
func (s *mongoTagSetSliceValue) Set(val string) error {
	ss := strings.Split(val, ",")
	var out []tag.Set
	for _, set := range ss {
		var tags tag.Set
		rawTags := strings.Split(strings.Trim(strings.TrimSpace(set), "[]"), " ")
		for _, rawTag := range rawTags {
			if strings.TrimSpace(rawTag) == "" {
				continue
			}

			kv := strings.Split(rawTag, "=")
			if len(kv) != 2 {
				return fmt.Errorf(errMongoTagSetSliceFormat, rawTag)
			}

			t := tag.Tag{
				Name:  strings.TrimSpace(kv[0]),
				Value: strings.TrimSpace(kv[1]),
			}
			tags = append(tags, t)
		}

		if len(tags) != 0 {
			out = append(out, tags)
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

func (s *mongoTagSetSliceValue) Type() string {
	return "mongoTagSetSlice"
}

func (s *mongoTagSetSliceValue) String() string {

	out := make([]string, len(*s.value))
	for i, d := range *s.value {
		tmpd := strings.ReplaceAll(d.String(), ",", " ")
		out[i] = fmt.Sprintf("[%s]", tmpd)
	}
	return strings.Join(out, ",")
}

func (s *mongoTagSetSliceValue) Append(val string) error {
	ss := strings.Split(val, ",")
	for _, set := range ss {
		var tags tag.Set
		rawTags := strings.Split(strings.Trim(strings.TrimSpace(set), "[]"), " ")
		for _, rawTag := range rawTags {
			if strings.TrimSpace(rawTag) == "" {
				continue
			}

			kv := strings.Split(rawTag, "=")
			if len(kv) != 2 {
				return fmt.Errorf(errMongoTagSetSliceFormat, rawTag)
			}

			t := tag.Tag{
				Name:  strings.TrimSpace(kv[0]),
				Value: strings.TrimSpace(kv[1]),
			}
			tags = append(tags, t)
		}

		if len(tags) != 0 {
			*s.value = append(*s.value, tags)
		}
	}
	return nil
}

func (s *mongoTagSetSliceValue) Replace(val []string) error {
	var out []tag.Set
	for _, set := range val {
		rawTags := strings.Split(strings.Trim(set, "[]"), " ")

		var inner tag.Set
		for _, rawTag := range rawTags {
			kv := strings.Split(rawTag, "=")
			if len(kv) != 2 {
				return fmt.Errorf(errMongoTagSetSliceFormat, rawTag)
			}
			inner = append(inner, tag.Tag{
				Name:  strings.TrimSpace(kv[0]),
				Value: strings.TrimSpace(kv[1]),
			})
		}
		if len(inner) != 0 {
			out = append(out, inner)
		}
	}
	*s.value = out
	return nil
}

func (s *mongoTagSetSliceValue) GetSlice() []string {
	out := make([]string, len(*s.value))
	for i, d := range *s.value {
		innerOut := make([]string, len(d))
		for j, v := range d {
			innerOut[j] = fmt.Sprintf("%s=%s", v.Name, v.Value)
		}
		out[i] = "[" + strings.Join(innerOut, " ") + "]"
	}
	return out
}

func mongoTagSetSliceConv(val string) (interface{}, error) {
	sets := strings.Split(val, ",")
	var out []tag.Set

	for _, set := range sets {
		var tags tag.Set
		rawTags := strings.Split(strings.Trim(set, "[]"), " ")
		for _, rawTag := range rawTags {
			if strings.TrimSpace(rawTag) == "" {
				continue
			}
			
			kv := strings.Split(rawTag, "=")
			if len(kv) != 2 {
				return nil, fmt.Errorf(errMongoTagSetSliceFormat, rawTag)
			}

			t := tag.Tag{
				Name:  strings.TrimSpace(kv[0]),
				Value: strings.TrimSpace(kv[1]),
			}
			tags = append(tags, t)
		}

		if len(tags) != 0 {
			out = append(out, tags)
		}
	}

	return out, nil
}

// GetMongoTagSetSlice return the tag.Set value of a flag with the given name
func (f *FlagSet) GetMongoTagSetSlice(name string) ([]tag.Set, error) {
	val, err := f.getFlagType(name, "mongoTagSetSlice", mongoTagSetSliceConv)
	if err != nil {
		return []tag.Set{}, err
	}
	return val.([]tag.Set), nil
}

// MongoTagSetSliceVar defines a tag.Set Slice flag with specified name, default value, and usage string.
// The argument p points to a []tag.Set variable in which to store the value of the flag.
func (f *FlagSet) MongoTagSetSliceVar(p *[]tag.Set, name string, value []tag.Set, usage string) {
	f.VarP(newMongoTagSetSliceValue(value, p), name, "", usage)
}

// MongoTagSetSliceVarP is like MongoTagSetSliceVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) MongoTagSetSliceVarP(p *[]tag.Set, name, shorthand string, value []tag.Set, usage string) {
	f.VarP(newMongoTagSetSliceValue(value, p), name, shorthand, usage)
}

// MongoTagSetSlice defines a []tag.Set flag with specified name, default value, and usage string.
// The return value is the address of a tag.Set variable that stores the value of the flag.
func (f *FlagSet) MongoTagSetSlice(name string, value []tag.Set, usage string) *[]tag.Set {
	p := []tag.Set{}
	f.MongoTagSetSliceVarP(&p, name, "", value, usage)
	return &p
}

// MongoTagSetSliceP is like MongoTagSetSlice, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) MongoTagSetSliceP(name, shorthand string, value []tag.Set, usage string) *[]tag.Set {
	p := []tag.Set{}
	f.MongoTagSetSliceVarP(&p, name, shorthand, value, usage)
	return &p
}