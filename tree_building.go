package tree

import (
	"fmt"
	"sort"
)

// Record is a struct containing int fields ID and Parent
type Record struct {
	ID     int
	Parent int
}

// Node is a struct containing int field ID and []*Node field Children
type Node struct {
	ID       int
	Children []*Node
}

// Build is a function for building a tree layout
func Build(records []Record) (*Node, error) {
	node := make(map[int]*Node)
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID || r.ID <= 0 && r.Parent > 0 {
			return nil, fmt.Errorf("not in sequence or has bad parent: %v", r)
		}
		node[r.ID] = &Node{ID: r.ID}
		if i > 0 {
			parent := node[r.Parent]
			parent.Children = append(parent.Children, node[r.ID])
		}
	}
	return node[0], nil
}
