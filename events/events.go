package events

type Listener[T any] func(T)

type Manager[T any] interface {
	Add(n string, l Listener[T])
	Run()
}

type BaseManager[T any] struct {
	lst map[string][]Listener[T]
}

func (m *BaseManager[T]) Invoke(n string, args T) {
	for _, ls := range m.lst[n] {
		ls(args)
	}
}

func (m *BaseManager[T]) Add(n string, l Listener[T]) {
	m.lst[n] = append(m.lst[n], l)
}
