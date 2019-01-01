package pipeline

type job func(in, out chan interface{})

func Pipe(funcs ...job) {
	return
}
