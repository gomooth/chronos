package calculator

type AddDateOption struct {
	notOverflow bool
}

func WithOverflow(enabled bool) func(*AddDateOption) {
	return func(opt *AddDateOption) {
		opt.notOverflow = !enabled
	}
}
