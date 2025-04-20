package fireback

import (
	"encoding/json"
	"fmt"
)

type TreeNode[T any] struct {
	Parent   *TreeNode[T]
	Content  *T
	Children []*TreeNode[T]
	IsLeaf   bool
}

func FindLeafNodes[T any](roots []*TreeNode[T]) []*TreeNode[T] {
	var leafNodes []*TreeNode[T]

	var findLeaves func(node *TreeNode[T])
	findLeaves = func(node *TreeNode[T]) {
		if len(node.Children) == 0 && node.Parent != nil {
			leafNodes = append(leafNodes, node)
			node.IsLeaf = true
			return
		}
		for _, child := range node.Children {
			findLeaves(child)
		}
	}

	for _, root := range roots {
		findLeaves(root)
	}

	return leafNodes
}

func buildTreeStructure[T any](nodes []*TreeNode[T], parentKey func(T) string, childKey func(T) string) []*TreeNode[T] {
	nodeMap := make(map[string]*TreeNode[T])
	for _, node := range nodes {
		nodeMap[childKey(*node.Content)] = node
	}

	j, _ := json.MarshalIndent(nodeMap, "", "  ")
	fmt.Println("nodemap:", nodeMap, string(j))

	var buildSubTree func(parentID string, parentNode *TreeNode[T]) []*TreeNode[T]
	buildSubTree = func(parentID string, parentNode *TreeNode[T]) []*TreeNode[T] {
		var children []*TreeNode[T]
		for _, node := range nodes {
			fmt.Println("Comparing:", parentKey(*node.Content), "With", parentID)
			if parentKey(*node.Content) == parentID {
				node.Parent = parentNode
				node.Children = buildSubTree(childKey(*node.Content), node)
				children = append(children, node)
			}
		}
		return children
	}

	return buildSubTree("", nil)
}

func TraverseNodes[T any](tree []*TreeNode[T], callback func(*TreeNode[T]) error) error {
	for _, node := range tree {
		if err := callback(node); err != nil {
			return err
		}
		if len(node.Children) > 0 {
			if err := TraverseNodes(node.Children, callback); err != nil {
				return err
			}
		}
	}
	return nil
}

func CreateTreeWithLeaves[T any](items []*T, parentKey func(*T) string, childKey func(*T) string) ([]*TreeNode[*T], []*TreeNode[*T]) {
	tree := CreateTree(items, parentKey, childKey)
	leafNodes := FindLeafNodes(tree)
	return tree, leafNodes
}

type TreeOperation[T any] struct {
	tree   []*TreeNode[*T]
	leaves []*TreeNode[*T]
}

func NewTreeOperation[T any](items []*T, parentKey func(*T) string, childKey func(*T) string) *TreeOperation[T] {

	tree, leaves := CreateTreeWithLeaves(items, parentKey, childKey)
	return &TreeOperation[T]{
		tree:   tree,
		leaves: leaves,
	}
}

func (x *TreeOperation[T]) ToArray() []*T {
	return FlattenTreeToArray(x.tree)
}

func CreateTree[T any](items []*T, parentKey func(*T) string, childKey func(*T) string) []*TreeNode[*T] {
	var nodes []*TreeNode[*T]
	for _, item := range items {
		item := item
		nodes = append(nodes, &TreeNode[*T]{Content: &item})
	}

	nodes = buildTreeStructure(nodes, parentKey, childKey)
	return nodes
}

func FlattenTree[T any](roots []*TreeNode[T]) []*TreeNode[T] {
	var flatList []*TreeNode[T]

	var flatten func(nodes []*TreeNode[T])
	flatten = func(nodes []*TreeNode[T]) {
		for _, node := range nodes {
			flatList = append(flatList, node)
			flatten(node.Children)
		}
	}

	flatten(roots)
	return flatList
}

func FlattenTreeToArray[T any](roots []*TreeNode[T]) []T {
	var result []T
	for _, node := range FlattenTree(roots) {
		result = append(result, *node.Content)
	}
	return result
}
