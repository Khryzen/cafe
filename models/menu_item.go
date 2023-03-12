package models

import "github.com/uadmin/uadmin"

type MenuItems struct {
	uadmin.Model
	Name        string `uadmin:"required"`
	Description string
	Picture     string   `uadmin:"required;image"`
	Category    Category `uadmin:"required"`
	CategoryID  uint
	Price       float32 `uadmin:"required"`
	Active      bool    `uadmin:"required"`
}

func (m *MenuItems) String() string {
	return m.Name
}
