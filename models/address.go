package models

import "github.com/uadmin/uadmin"

type Address struct {
	uadmin.Model
	Client       Client `uadmin:"read_only:edit"`
	ClientID     uint
	Name         string `uadmin:"required"`
	AddressOne   string
	AddressTwo   string
	Brgy         string
	Municipality string
	Province     string
	ZipCode      string
}

func (a *Address) String() string {
	return a.Name
}
