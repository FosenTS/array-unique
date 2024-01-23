package array_unique

type UniqueArray struct {
	elements []any
}

func (ua *UniqueArray) Add(elem any) {
	for _, e := range ua.elements {
		if e == elem {
			return
		}
	}
	ua.elements = append(ua.elements, elem)
}

func (ua *UniqueArray) Remove(elem any) {
	index := -1
	for i, existingElement := range ua.elements {
		if existingElement == elem {
			index = i
			break
		}
	}
	if index != -1 {
		ua.elements = append(ua.elements[:index], ua.elements[index+1:]...)
	}
}

func (ua *UniqueArray) Elements() []any {
	return ua.elements
}

func CopySliceToUnique(in []any) UniqueArray {
	out := UniqueArray{}
	for _, elem := range in {
		out.Add(elem)
	}

	return out
}
