package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Records []Record

func (records Records) Len() int {
	return len(records)
}

func (records Records) Less(i, j int) bool {
	return records[i].ID < records[j].ID
}

func (records Records) Swap(i, j int) {
	records[i], records[j] = records[j], records[i]
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	sort.Sort(Records(records))
	nodes := make([]*Node, len(records))
	for i, record := range records {
		nodes[i] = &Node{ID: record.ID}
		if i == 0 && (record.ID != 0 || record.Parent != 0) {
			return nil, fmt.Errorf("invalid root element: ID = %d; ParentID = %d", record.ID, record.Parent)
		}
		if record.ID == 0 {
			continue
		}
		if i != record.ID || record.ID <= record.Parent {
			return nil, fmt.Errorf("invalid element: ID = %d; ParentID = %d", record.ID, record.Parent)
		}
		if parent := nodes[record.Parent]; parent != nil {
			parent.Children = append(parent.Children, nodes[i])
		}
	}
	return nodes[0], nil
}
