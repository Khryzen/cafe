package models

import (
	"fmt"

	"github.com/uadmin/uadmin"
)

type Orders struct {
	uadmin.Model
	Client         Client `uadmin:"read_only:edit"`
	ClientID       uint
	SpecialRequest string `uadmin:"read_only:edit"`
	Status         int    `uadmin:"read_only:edit"`
	Paid           bool   `uadmin:"read_only:edit"`
}

type Status int

func (s *Status) Preparing() int {
	return 1
}

func (s *Status) Serving() int {
	return 2
}

func (s *Status) Pending() int {
	return 3
}
func (s *Status) Cancelled() int {
	return 4
}
func (s *Status) Served() int {
	return 5
}

func (o *Orders) String() string {
	uadmin.Preload(o.Client)
	return fmt.Sprint(o.ID) + " - " + o.Client.User.FirstName + o.Client.User.LastName
}
