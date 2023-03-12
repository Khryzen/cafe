package models

import "github.com/uadmin/uadmin"

type Cafe struct {
	uadmin.Model
	Name          string `uadmin:"required"`
	Logo          string `uadmin:"image"`
	BusinessName  string `uadmin:"list_exclude;required"`
	Slogan        string `uadmin:"list_exclude"`
	EmailAddress  string `uadmin:"email"`
	ContactNumber string `uadmin:"list_exclude;required"`
	Address       string `uadmin:"list_exclude"`
	Municipality  string `uadmin:"list_exclude"`
	Province      string `uadmin:"list_exclude"`
	Country       string `uadmin:"list_exclude"`
	ZipCode       string `uadmin:"list_exclude"`
	Active        bool   `uadmin:"required"`
}

func (c *Cafe) String() string {
	return c.Name
}

func (c *Cafe) Save() {
	cafe := Cafe{}
	if c.Active {
		if uadmin.Count(&cafe, "active = ?", true) > 0 {
			uadmin.Trail(uadmin.INFO, "Only one cafe should be active.")
		} else {
			uadmin.Save(c)
		}
	}
	uadmin.Save(c)
}
