package flag

import (
	"errors"
	"fmt"
	"github.com/openzipkin/zipkin-go"
	"strings"
)

const errExtractFailurePolicyStr = "invalid string being converted to zipkin.ExtractFailurePolicy"

type extractFailurePolicyValue struct {
	value *zipkin.ExtractFailurePolicy
}

func newZipkinExtractFailurePolicyValue(val zipkin.ExtractFailurePolicy, p *zipkin.ExtractFailurePolicy) *extractFailurePolicyValue {
	ssv := new(extractFailurePolicyValue)
	ssv.value = p
	*ssv.value = val
	return ssv
}

func (e extractFailurePolicyValue) Type() string {
	return "zipkinExtractFailurePolicy"
}

func (e extractFailurePolicyValue) String() (result string) {
	return fmt.Sprintf("%d", *e.value)
}

func (e *extractFailurePolicyValue) Set(val string) (err error)  {
	value := strings.ToLower(dashBlankReplacer.Replace(val))
	if strings.Contains(value, "tag") {
		*e.value = zipkin.ExtractFailurePolicyTagAndRestart
		return nil
	}

	if strings.Contains(value, "restart") {
		*e.value = zipkin.ExtractFailurePolicyRestart
		return nil
	}

	if strings.Contains(value, "error") {
		*e.value = zipkin.ExtractFailurePolicyError
		return nil
	}

	return errors.New("unsupported extract failure policy")
}

func extractFailurePolicyConv(val string) (interface{}, error) {
	switch val {
	case "0":
		return zipkin.ExtractFailurePolicyRestart, nil
	case "1":
		return zipkin.ExtractFailurePolicyError, nil
	case "2":
		return zipkin.ExtractFailurePolicyTagAndRestart, nil
	default:
		return -1, fmt.Errorf(errExtractFailurePolicyStr)
	}
}

// GetZipkinExtractFailurePolicy return the zipkin.ExtractFailurePolicy value of a flag with the given name
func (f *FlagSet) GetZipkinExtractFailurePolicy(name string) (zipkin.ExtractFailurePolicy, error) {
	val, err := f.getFlagType(name, "zipkinExtractFailurePolicy", extractFailurePolicyConv)
	if err != nil {
		return -1, err
	}
	return val.(zipkin.ExtractFailurePolicy), nil
}

// ZipkinExtractFailurePolicyVar defines a zipkin.ExtractFailurePolicy flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (f *FlagSet) ZipkinExtractFailurePolicyVar(p *zipkin.ExtractFailurePolicy, name string, value zipkin.ExtractFailurePolicy, usage string) {
	f.fs.VarP(newZipkinExtractFailurePolicyValue(value, p), name, "", usage)
}

// ZipkinExtractFailurePolicyVarP is like ZipkinExtractFailurePolicyVar, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) ZipkinExtractFailurePolicyVarP(p *zipkin.ExtractFailurePolicy, name, shorthand string, value zipkin.ExtractFailurePolicy, usage string) {
	f.fs.VarP(newZipkinExtractFailurePolicyValue(value, p), name, shorthand, usage)
}


// ZipkinExtractFailurePolicy defines a zipkin.ExtractFailurePolicy flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func (f *FlagSet) ZipkinExtractFailurePolicy(name string, value zipkin.ExtractFailurePolicy, usage string) *zipkin.ExtractFailurePolicy {
	p := new(zipkin.ExtractFailurePolicy)
	f.ZipkinExtractFailurePolicyVarP(p, name, "", value, usage)
	return p
}


// ZipkinExtractFailurePolicyP is like ZipkinExtractFailurePolicy, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) ZipkinExtractFailurePolicyP(name, shorthand string, value zipkin.ExtractFailurePolicy, usage string) *zipkin.ExtractFailurePolicy {
	p := new(zipkin.ExtractFailurePolicy)
	f.ZipkinExtractFailurePolicyVarP(p, name, shorthand, value, usage)
	return p
}