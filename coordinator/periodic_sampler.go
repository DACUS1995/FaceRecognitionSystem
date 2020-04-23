package periodicsampler

type Sampler interface {
	Sample() []byte
}

type Sampler struct {
	interval int
}
