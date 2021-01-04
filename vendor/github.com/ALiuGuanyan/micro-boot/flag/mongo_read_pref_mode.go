package flag

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"strings"
)

var errMongoReadPreferenceMode = errors.New("unsupported mongo read preference mode")

type mongoReadPreferenceModeValue struct {
	value *readpref.Mode
}

func newMongoReadPreferenceModeValue(val readpref.Mode, p *readpref.Mode) *mongoReadPreferenceModeValue {
	ssv := new(mongoReadPreferenceModeValue)
	ssv.value = p
	*ssv.value = val
	return ssv
}

func (m mongoReadPreferenceModeValue) Type() string {
	return "mongoReadPreferenceMode"
}

func (m mongoReadPreferenceModeValue) String() (result string) {
	switch *m.value {
	case 1:
		return "primary"
	case 2:
		return "primarypreferred"
	case 3:
		return "secondary"
	case 4:
		return "secondarypreferred"
	case 5:
		return "near"
	default:
		return ""
	}
}

func normalizeValue(val string) (uint8, error)  {
	value := strings.ToLower(dashBlankReplacer.Replace(val))

	if strings.Contains(value, "primary") {
		if strings.Contains(value, "preferred") {
			return 2, nil
		}
		return 1, nil
	}

	if strings.Contains(value, "secondary") {
		if strings.Contains(value, "preferred") {
			return 4, nil
		}
		return 3, nil
	}

	if strings.Contains(value, "near") {
		return 5, nil
	}

	return 0, errMongoReadPreferenceMode
}

func (m *mongoReadPreferenceModeValue) Set(val string) (err error)  {
	value, err := normalizeValue(val)
	if err != nil {
		return err
	}
	*m.value = (readpref.Mode)(value)
	return nil
}

func mongoReadPreferenceModeConv(val string) (interface{}, error) {
	return normalizeValue(val)
}

// GetMongoReadPreferenceMode return the readpref.Mode value of a flag with the given name
func (f *FlagSet) GetMongoReadPreferenceMode(name string) (readpref.Mode, error) {
	val, err := f.getFlagType(name, "mongoReadPreferenceMode", mongoReadPreferenceModeConv)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	
	return (readpref.Mode)(val.(uint8)), nil
}

// MongoReadPreferenceModeVar defines a readpref.Mode flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (f *FlagSet) MongoReadPreferenceModeVar(p *readpref.Mode, name string, value readpref.Mode, usage string) {
	f.fs.VarP(newMongoReadPreferenceModeValue(value, p), name, "", usage)
}

// MongoReadPreferenceModeVarP is like MongoReadPreferenceModeVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) MongoReadPreferenceModeVarP(p *readpref.Mode, name, shorthand string, value readpref.Mode, usage string) {
	f.fs.VarP(newMongoReadPreferenceModeValue(value, p), name, shorthand, usage)
}


// MongoReadPreferenceMode defines a readpref.Mode flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func (f *FlagSet) MongoReadPreferenceMode(name string, value readpref.Mode, usage string) *readpref.Mode {
	p := new(readpref.Mode)
	f.MongoReadPreferenceModeVarP(p, name, "", value, usage)
	return p
}


// MongoReadPreferenceModeP is like MongoReadPreferenceMode, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) MongoReadPreferenceModeP(name, shorthand string, value readpref.Mode, usage string) *readpref.Mode {
	p := new(readpref.Mode)
	f.MongoReadPreferenceModeVarP(p, name, shorthand, value, usage)
	return p
}

