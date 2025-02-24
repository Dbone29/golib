package events

type Listener[T any] func(T)

type Manager[T any] interface {
	Add(n string, l Listener[T])
	Run()
}

type BaseManager[T any] struct {
	Lst map[string][]Listener[T]
}

func (m *BaseManager[T]) Invoke(n string, args T) {
	for _, ls := range m.Lst[n] {
		ls(args)
	}
}

func (m *BaseManager[T]) Add(n string, l Listener[T]) {
	m.Lst[n] = append(m.Lst[n], l)
}
