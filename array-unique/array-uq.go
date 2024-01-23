package array_unique

type UniqueArray[T comparable] struct {
	elements map[T]bool
}

func NewUniqueArray[T any]() {

}

func (ua *UniqueArray[T]) Add(elem T) {
	if !ua.elements[elem] {
		ua.elements[elem] = true
	}
}

func (ua *UniqueArray[T]) Remove(elem T) {
	delete(ua.elements, elem)
}

func (ua *UniqueArray[T]) Elements() []T {
	keys := make([]T, len(ua.elements))
	i := 0
	for k := range ua.elements {
		keys[i] = k
		i++
	}
	return keys
}

func CopySliceToUnique[T comparable](in []T) UniqueArray[T] {
	out := UniqueArray[T]{}
	for _, elem := range in {
		out.Add(elem)
	}

	return out
}
