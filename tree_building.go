package tree

import (
	"errors"
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

	root := make(map[int]*Node)
	processed := make(map[int]Record)

	if len(records) <= 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	for _, record := range records {
		if len(processed) > 0 {
			if _, ok := processed[record.ID-1]; !ok {
				return nil, errors.New("non-continuous")
			}
		}
		if record.ID < record.Parent {
			return nil, errors.New("higher id parent of lower id")
		}
		if record.ID > 0 && record.Parent >= 0 {
			if record.ID == record.Parent {
				return nil, errors.New("cycle directly")
			}
			if _, ok := processed[record.ID]; ok {
				return nil, errors.New("duplicate node")
			}
			if parent, ok := root[record.Parent]; ok {
				if len(parent.Children) <= 0 {
					parent.Children = make([]*Node, 0)
				}
				child := &Node{ID: record.ID}
				root[child.ID] = child
				parent.Children = append(parent.Children, child)
			}
		} else {
			if record.Parent > 0 {
				return nil, errors.New("Root node has parent")
			}
			if _, ok := root[0]; ok {
				return nil, errors.New("duplicate root")
			} else {
				root[0] = &Node{ID: 0}
			}
		}
		processed[record.ID] = record
	}
	if node, ok := root[0]; !ok || node.ID != 0 {
		return nil, errors.New("no root node")
	}
	return root[0], nil
}
