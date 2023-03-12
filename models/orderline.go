package models

import (
	"fmt"

	"github.com/uadmin/uadmin"
)

type OrderLine struct {
	uadmin.Model
	Orders      Orders `uadmin:"read_only:edit"`
	OrdersID    uint
	MenuItems   MenuItems `uadmin:"read_only:edit"`
	MenuItemsID uint
}

func (o *OrderLine) String() string {
	return fmt.Sprint(o.ID)
}
