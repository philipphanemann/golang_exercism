package tree

import (
	"errors"
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

	var (
		root      *Node
		nodeTable = map[int]*Node{}
	)

	sort.SliceStable(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	for i, r := range records {

		// checking for errors
		if _, ok := nodeTable[r.ID]; ok {
			return nil, fmt.Errorf("ID %d should occur only once", r.ID)
		}

		if r.ID == 0 {
			if r.Parent > 0 {
				return nil, fmt.Errorf("Parent ID %d greater than root node", r.Parent)
			}
		} else {

			if i == 0 {
				return nil, errors.New("no root node provided")
			}
			if _, ok := nodeTable[r.ID-1]; !ok {
				return nil, fmt.Errorf("ID %d missing", r.ID-1)
			}
			if r.ID <= r.Parent {
				return nil, fmt.Errorf("ID %d can not be placed before or equal ID %d", r.ID, r.Parent)
			}
		}
		//end errors

		nodeTable[r.ID] = &Node{ID: r.ID}
		if r.ID == 0 {
			root = nodeTable[0]
		} else {
			parentNode := nodeTable[r.Parent]
			parentNode.Children = append(parentNode.Children, nodeTable[r.ID])
		}

	}
	return root, nil

}
