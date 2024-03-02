package tui

import (
	"github.com/alcb1310/kanban/cmd/database"
	"github.com/alcb1310/kanban/cmd/types"
)

func initList(db database.Database) map[int][]types.List {
	lists := make(map[int][]types.List)

	lists[types.TODO] = db.GetByStatus(types.TODO)
	lists[types.INPROGRESS] = db.GetByStatus(types.INPROGRESS)
	lists[types.DONE] = db.GetByStatus(types.DONE)

	return lists
}
