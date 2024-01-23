package chan_unique

import (
	array_unique "github.com/FosenTS/unique-types/array-unique"
)

func ConvertToUniqueClosedChan(in chan any) chan any {
	var a array_unique.UniqueArray
	var el any
	var ok bool
	for {
		select {
		case el, ok = <-in:
		}
		if !ok {
			break
		}
		a.Add(el)
	}
	out := make(chan any)
	for _, el := range a.Elements() {
		out <- el
	}
	close(out)

	return out
}
