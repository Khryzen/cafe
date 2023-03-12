package models

type MenuItemSize struct {
	Name        string
	Description string
	MenuItems   MenuItems
	MenuItemsID uint
	Price       float32
}
