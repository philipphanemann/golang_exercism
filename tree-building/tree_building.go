package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	var (
		root      *Node
		nodeTable = map[int]*Node{}
	)

	sort.SliceStable(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	for _, r := range records {
		nodeTable[r.ID] = &Node{ID: r.ID}

		if r.ID == 0 && r.Parent > 0 {
			return &Node{}, errors.New("root node shall not have parent")
		}

		if r.ID == 0 {
			root = nodeTable[0]
		} else {
			parentNode := nodeTable[r.Parent]
			parentNode.Children = append(parentNode.Children, nodeTable[r.ID])
		}

	}
	return root, nil

}
