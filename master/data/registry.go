package data

import (
	"sync"
	"sync/atomic"
)

type Registry struct {
	count atomic.Uint32

	m     sync.RWMutex
	nodes map[string]*Node
}

func NewRegistry() *Registry {
	return &Registry{
		nodes: make(map[string]*Node),
	}
}

func (r *Registry) Add(host string, port int32) {
	r.m.Lock()
	defer r.m.Unlock()

	r.nodes[host] = &Node{
		Host: host,
		Port: port,
	}
}

func (r *Registry) Get() *Node {
	r.m.RLock()
	defer r.m.RUnlock()

	nodes := r.getNodes()
	index := int(r.count.Load()) % len(nodes)
	r.count.Add(1)

	return nodes[index]
}

func (r *Registry) GetAll() Nodes {
	r.m.RLock()
	defer r.m.RUnlock()

	return r.getNodes()
}

func (r *Registry) getNodes() Nodes {
	var nodes Nodes
	for _, node := range r.nodes {
		nodes = append(nodes, node)
	}

	return nodes
}
