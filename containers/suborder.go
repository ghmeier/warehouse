package containers

import (
	"database/sql"
	"github.com/pborman/uuid"
)

type SubOrder struct {
	ID            uuid.UUID `json: "id"`
	OrderID				uuid.UUID `json: "order_id"`
	UserID        uuid.UUID `json: "user_id"`
	//corresponding to an item in inventory
	ItemId 				uuid.UUID `json: "item_id"`
}

func SubOrderFromSql(rows *sql.Rows) ([]*SubOrder, error) {
	suborder := make([]*SubOrder, 0)

	for rows.Next() {
		s := &SubOrder{}
		rows.Scan(&s.ID, &s.OrderID, &s.UserID, &s.ItemID)
		suborder = append(order, s)
	}

	return suborder, nil
}
