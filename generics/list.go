package generics

type List[T any] struct {
	Items []T
}

func (l *List[T]) Add(item T) {
	l.Items = append(l.Items, item)
}

func (l *List[T]) Get(index int) (T, bool) {
	if index < 0 || index >= len(l.Items) {
		var zero T
		return zero, false
	}
	return l.Items[index], true
}
