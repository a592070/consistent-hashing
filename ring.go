package consistent_hashing

import (
	"github.com/pkg/errors"
	"slices"
	"sort"
	"sync"
)

var ErrNodeNotFound = errors.New("node not found")

type Ring struct {
	Nodes Nodes
	sync.RWMutex
}

func NewRing() *Ring {
	return &Ring{Nodes: Nodes{}}
}

func (r *Ring) AddNode(id string) {
	r.Lock()
	defer r.Unlock()
	node := NewNode(id)
	r.Nodes = append(r.Nodes, node)
	slices.SortStableFunc(r.Nodes, func(a, b *Node) int {
		if a.HashId == b.HashId {
			return 0
		}
		if a.HashId > b.HashId {
			return 1
		}
		return -1
	})
}

func (r *Ring) search(id string) int {
	return sort.Search(r.Nodes.Len(), func(i int) bool {
		return r.Nodes[i].HashId >= hashId(id)
	})
}

func (r *Ring) RemoveNode(id string) error {
	r.Lock()
	defer r.Unlock()

	i := r.search(id)
	if i >= r.Nodes.Len() || r.Nodes[i].Id != id {
		return ErrNodeNotFound
	}
	r.Nodes = append(r.Nodes[:i], r.Nodes[i+1:]...)

	return nil
}

func (r *Ring) Get(id string) string {
	i := r.search(id)
	if i >= r.Nodes.Len() {
		i = 0
	}
	return r.Nodes[i].Id
}
