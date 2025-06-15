package parse

import "time"

type FromStringOption struct {
	layouts []string
	loc     *time.Location
}

func WithFromStringLayout(layout string, others ...string) func(*FromStringOption) {
	return func(o *FromStringOption) {
		if o.layouts == nil {
			o.layouts = make([]string, 0)
		}
		layouts := append([]string{layout}, others...)
		o.layouts = append(o.layouts, layouts...)
	}
}

func WithFromStringLocation(loc *time.Location) func(*FromStringOption) {
	return func(o *FromStringOption) {
		o.loc = loc
	}
}
