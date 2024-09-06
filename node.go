package consistent_hashing

import "hash/crc32"

type Node struct {
	Id     string
	HashId uint32
}

func NewNode(id string) *Node {
	return &Node{
		Id:     id,
		HashId: hashId(id),
	}
}

type Nodes []*Node

func (n Nodes) Len() int           { return len(n) }
func (n Nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n Nodes) Less(i, j int) bool { return n[i].HashId < n[j].HashId }

func hashId(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}
