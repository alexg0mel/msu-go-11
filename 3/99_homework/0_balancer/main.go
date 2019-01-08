package main

type RoundRobinBalancer struct {
	ch     chan bool
	worker chan int
	nodes  []int
	count  int
}

func (r *RoundRobinBalancer) Init(size int) {
	r.ch = make(chan bool)
	r.worker = make(chan int)
	for i := 0; i < size; i++ {
		r.nodes = append(r.nodes, 0)
	}
	r.count = 0
	go r.do()
}

func (r *RoundRobinBalancer) GiveStat() []int {

	return r.nodes
}

func (r *RoundRobinBalancer) GiveNode() int {
	r.ch <- true
	return <-r.worker
}

func (r *RoundRobinBalancer) do() {
start:
	for {
		<-r.ch
		for index, item := range r.nodes {
			if item == r.count {
				r.nodes[index]++
				r.worker <- index
				continue start
			}
		}
		r.count++
		r.nodes[0]++
		r.worker <- 0
	}

}
