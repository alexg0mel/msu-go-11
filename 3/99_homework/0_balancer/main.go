package main

import "sync"

type RoundRobinBalancer struct {
	sync.Mutex
	nodes []int
	count int
}

func (r *RoundRobinBalancer) Init(size int) {
	for i := 0; i < size; i++ {
		r.nodes = append(r.nodes, 0)
	}
	r.count = 0
}

func (r *RoundRobinBalancer) GiveStat() []int {

	return r.nodes
}

func (r *RoundRobinBalancer) GiveNode() int {
	r.Lock()
	defer r.Unlock()
	for index, item := range r.nodes {
		if item == r.count {
			r.nodes[index]++
			return index
		}
	}
	r.count++
	r.nodes[0]++
	return 0
}
