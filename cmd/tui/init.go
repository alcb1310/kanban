package tui

import "github.com/alcb1310/kanban/cmd/types"

func initList() map[int][]types.List {
	lists := make(map[int][]types.List)

	lists[types.TODO] = []types.List{
		{ID: 1, Title: "List 1"},
		{ID: 2, Title: "List 2"},
		{ID: 3, Title: "List 3"},
	}

	lists[types.INPROGRESS] = []types.List{
		{ID: 4, Title: "List 4"},
		{ID: 5, Title: "List 5"},
		{ID: 6, Title: "List 6"},
	}

	lists[types.DONE] = []types.List{
		{ID: 7, Title: "List 7"},
		{ID: 8, Title: "List 8"},
		{ID: 9, Title: "List 9"},
	}

	return lists
}
