package workspaces

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var parent = "parent1"

var mockWorkspaces = []*WorkspaceEntity{
	{
		UniqueId: parent,
	},
	{
		UniqueId: "parent2",
		ParentId: &parent,
	},
}

func TestCreateTree(t *testing.T) {
	tree, _ := CreateTreeWithLeaves(
		mockWorkspaces,
		func(w *WorkspaceEntity) string {
			if w.ParentId == nil {
				return ""
			}
			return *w.ParentId
		},
		func(w *WorkspaceEntity) string {
			return w.UniqueId
		},
	)

	assert.Equal(t, 1, len(tree), "one parent item")
	assert.Equal(t, 1, len(tree[0].Children), "one child")
}
