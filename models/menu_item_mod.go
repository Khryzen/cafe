package models

import "github.com/uadmin/uadmin"

type MenuItemMod struct {
	uadmin.Model
	Name        string  `uadmin:"read_only:edit"`
	Description string  `uadmin:"read_only:edit"`
	Price       float32 `uadmin:"required"`
	MenuItems   MenuItems
	MenuItemsID uint
	Multiple    bool `uadmin:"required"`
}

func (m *MenuItemMod) String() string {
	return m.Name
}
