package flag

import (
	"errors"
	"fmt"
	"github.com/openzipkin/zipkin-go/idgenerator"
	"strings"
)

type ZipkinIDGenerator int

const (
	Random64 ZipkinIDGenerator = iota
	Random128
	RandomTimestamped
)

const errZipkinZipkinIDGeneratorStr = "invalid string being converted to idgenerator.IDGenerator"

type extractZipkinIDGeneratorValue struct {
	value *idgenerator.IDGenerator
	typ *ZipkinIDGenerator
}

func mapZipkinIDGenerator(val ZipkinIDGenerator)	(typ ZipkinIDGenerator, value idgenerator.IDGenerator) {
	switch val {
	case Random64:
		return val, idgenerator.NewRandom64()
	case Random128:
		return val, idgenerator.NewRandom128()
	case RandomTimestamped:
		return val, idgenerator.NewRandomTimestamped()
	default:
		return Random64, idgenerator.NewRandom64()
	}
}

func newZipkinZipkinIDGeneratorValue(val ZipkinIDGenerator, p *idgenerator.IDGenerator) *extractZipkinIDGeneratorValue {

	ssv := new(extractZipkinIDGeneratorValue)
	ssv.value = p
	ssv.typ = new(ZipkinIDGenerator)
	*ssv.typ, *ssv.value = mapZipkinIDGenerator(val)
	return ssv
}

func (e extractZipkinIDGeneratorValue) Type() string {
	return "zipkinIDGenerator"
}

func (e extractZipkinIDGeneratorValue) String() (result string) {
	return fmt.Sprintf("%d", *e.typ)
}

func (e *extractZipkinIDGeneratorValue) Set(val string) (err error)  {
	value := strings.ToLower(dashBlankReplacer.Replace(val))
	if strings.Contains(value, "64") {
		*e.value = idgenerator.NewRandom64()
		*e.typ = Random64
		return nil
	} else if strings.Contains(value, "128") {
		*e.value = idgenerator.NewRandom128()
		*e.typ = Random128
		return nil
	} else if strings.Contains(value, "timestamp") {
		*e.value = idgenerator.NewRandomTimestamped()
		*e.typ= RandomTimestamped
		return nil
	} else {
		return errors.New("unsupported zipkin id generator")
	}
}

func zipkinZipkinIDGeneratorConv(val string) (interface{}, error) {
	switch val {
	case "0":
		return idgenerator.NewRandom64(), nil
	case "1":
		return idgenerator.NewRandom128(), nil
	case "2":
		return idgenerator.NewRandomTimestamped(), nil
	default:
		return nil, fmt.Errorf("%s: %s", errZipkinZipkinIDGeneratorStr, val)
	}
}

// GetZipkinIDGenerator return the idgenerator.IDGenerator value of a flag with the given name
func (f *FlagSet) GetZipkinIDGenerator(name string) (idgenerator.IDGenerator, error) {
	val, err := f.getFlagType(name, "zipkinIDGenerator", zipkinZipkinIDGeneratorConv)
	if err != nil {
		return nil, err
	}
	return val.(idgenerator.IDGenerator), nil
}

// ZipkinIDGeneratorVar defines a idgenerator.IDGenerator flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (f *FlagSet) ZipkinIDGeneratorVar(p *idgenerator.IDGenerator, name string, value ZipkinIDGenerator, usage string) {
	f.fs.VarP(newZipkinZipkinIDGeneratorValue(value, p), name, "", usage)
}

// ZipkinIDGeneratorVarP is like ZipkinIDGeneratorVarVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) ZipkinIDGeneratorVarP(p *idgenerator.IDGenerator, name, shorthand string, value ZipkinIDGenerator, usage string) {
	f.fs.VarP(newZipkinZipkinIDGeneratorValue(value, p), name, shorthand, usage)
}


// ZipkinExtractFailurePolicy defines a ExtractFailurePolicy flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func (f *FlagSet) ZipkinIDGenerator(name string, value ZipkinIDGenerator, usage string) *idgenerator.IDGenerator {
	p := new(idgenerator.IDGenerator)
	f.ZipkinIDGeneratorVarP(p, name, "", value, usage)
	return p
}


// ZipkinExtractFailurePolicyP is like ZipkinExtractFailurePolicy, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) ZipkinIDGeneratorP(name, shorthand string, value ZipkinIDGenerator, usage string) *idgenerator.IDGenerator {
	p := new(idgenerator.IDGenerator)
	f.ZipkinIDGeneratorVarP(p, name, shorthand, value, usage)
	return p
}