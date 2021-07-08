package tree

import (
	"fmt"
	"sort"
)

//Record tracked
type Record struct {
	ID, Parent int
}

//Node of record tree
type Node struct {
	ID       int
	Children []*Node
}

//Build creates tree from records
func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	var tree = map[int]*Node{}

	sort.SliceStable(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	for i, r := range records {

		// checking for errors
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, fmt.Errorf("not in sequence or has bad parent: %v", r)
		}

		tree[r.ID] = &Node{ID: r.ID}
		if r.ID > 0 {
			parentNode := tree[r.Parent]
			parentNode.Children = append(parentNode.Children, tree[r.ID])
		}
	}

	return tree[0], nil

}
