package zipkin

// CountingSampler is the settings of counting sampler for creating a standard zipkin
// counting sampler
type CountingSampler struct {
	Rate float64 `json:"rate" yaml:"rate"`
}

// BoundarySampler is the settings of boundary sampler for
// creating a standard zipkin boundary sampler
type BoundarySampler struct {
	Rate float64 `json:"rate" yaml:"rate"`
	Salt int64 `json:"salt" yaml:"salt"`
}

// ModuloSampler is the settings of modulo sampler for creating a standard zipkin
// module sampler
type ModuloSampler struct {
	Mod uint64 `json:"mod" yaml:"mod"`
}
