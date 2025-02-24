package observer

type Listener[T any] func(T)

type Manager[T any] interface {
	Add(n string, l Listener[T])
	//Remove(n string, l Listener[T])
	Invoke(n string, args any)
}

type BaseManager[T any] struct {
	Lst map[string][]Listener[T]
}

func (m *BaseManager[T]) Add(n string, l Listener[T]) {
	m.Lst[n] = append(m.Lst[n], l)
}

/*func (m *BaseManager[T]) Remove(n string, l Listener[T]) {
	for index, ls := range m.Lst[n] {
		if ls == l {
			m.Lst[n] = append(m.Lst[n][:index], m.Lst[n][index+1:]...)
		}

	}
}*/

func (m *BaseManager[T]) Invoke(n string, args T) {
	for _, ls := range m.Lst[n] {
		ls(args)
	}
}
