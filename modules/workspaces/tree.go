package workspaces

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Tree can be any map with:
// 1. Key that has method 'String() string'
// 2. Value is Tree itself
// You can replace this with your own tree
type Tree map[string]Tree

type NestedNode struct {
	UniqueId string       `json:"uniqueId"`
	Children []NestedNode `json:"children"`
}

func (tree Tree) Add(path string, spliter string) {
	frags := strings.Split(path, spliter)
	tree.add(frags)
}

func (tree Tree) add(frags []string) {
	if len(frags) == 0 {
		return
	}

	nextTree, ok := tree[frags[0]]
	if !ok {
		nextTree = Tree{}
		tree[frags[0]] = nextTree
	}

	nextTree.add(frags[1:])
}

func (tree Tree) Json() string {
	data, _ := json.MarshalIndent(tree, "", "  ")
	return string(data)
}

func (tree Tree) ToObject(root bool) []NestedNode {
	if tree == nil {
		return []NestedNode{}
	}

	items := []NestedNode{}

	index := 0

	for k, v := range tree {

		children := []NestedNode{}
		if len(v) > 0 {
			children = v.ToObject(false)
		}
		items = append(items, NestedNode{UniqueId: k, Children: children})

		index++
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].UniqueId < items[j].UniqueId
	})

	return items
}

func (tree Tree) Fprint(w io.Writer, root bool, padding string) {
	if tree == nil {
		return
	}

	index := 0
	for k, v := range tree {
		fmt.Fprintf(w, "%s%s\n", padding+getPadding(root, getBoxType(index, len(tree))), k)
		v.Fprint(w, false, padding+getPadding(root, getBoxTypeExternal(index, len(tree))))
		index++
	}
}

type BoxType int

const (
	Regular BoxType = iota
	Last
	AfterLast
	Between
)

func (boxType BoxType) String() string {
	switch boxType {
	case Regular:
		return "\u251c" // ├
	case Last:
		return "\u2514" // └
	case AfterLast:
		return " "
	case Between:
		return "\u2502" // │
	default:
		panic("invalid box type")
	}
}

func getBoxType(index int, len int) BoxType {
	if index+1 == len {
		return Last
	} else if index+1 > len {
		return AfterLast
	}
	return Regular
}

func getBoxTypeExternal(index int, len int) BoxType {
	if index+1 == len {
		return AfterLast
	}
	return Between
}

func getPadding(root bool, boxType BoxType) string {
	if root {
		return ""
	}

	return boxType.String() + " "
}
