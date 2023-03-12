package models

import "github.com/uadmin/uadmin"

type Category struct {
	uadmin.Model
	Name   string `uadmin:"required"`
	Active bool   `uadmin:"required"`
}

func (c *Category) String() string {
	return c.Name
}
