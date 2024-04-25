package registry

import (
	"sync"
)

type R interface {
	Add(string)
}

type Registry struct {
	m sync.Map
	l []string
}

func New() *Registry {
	return new(Registry)
}

func (r *Registry) Add(host string) {
	_, loaded := r.m.LoadOrStore(host, true)
	if loaded {
		return
	}

	r.l = append(r.l, host)
}

func (r *Registry) GetAny() string {
	return r.l[0]
}
