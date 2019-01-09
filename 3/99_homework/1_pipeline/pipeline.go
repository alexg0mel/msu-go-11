package pipeline

import (
	"time"
)

type job func(in, out chan interface{})

type pipeline struct {
	job
	in  chan interface{}
	out chan interface{}
}

func (p *pipeline) run() {
	p.job(p.in, p.out)
}

func NewPipline(fun job) *pipeline {
	return &pipeline{
		fun,
		make(chan interface{}),
		make(chan interface{}),
	}

}

func Pipe(funcs ...job) {
	mapPipelines := make(map[int]*pipeline)
	var countPipes int

	for _, fun := range funcs {
		mapPipelines[countPipes] = NewPipline(fun)
		if countPipes > 0 {
			go func(ind int) {
				for o := range mapPipelines[ind-1].out {
					mapPipelines[ind].in <- o
				}
			}(countPipes)
		}
		go mapPipelines[countPipes].run()
		countPipes++
	}
	time.Sleep(time.Nanosecond)
	//использование sync.WaitGroup не получилось - в тестах sync.WaitGroup.Wait() зависает... и выдает deadlock
	// дошел до такого варианта - (see struct pipline) Первый вариант был гораздо проще и в лоб, но все тот же deadlock даже без использования мню горутин.
}
