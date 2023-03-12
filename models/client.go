package models

import "github.com/uadmin/uadmin"

type Client struct {
	uadmin.Model
	User          uadmin.User `uadmin:"read_only:edit"`
	UserID        uint
	ContactNumber string `uadmin:"read_only:edit"`
}

func (c *Client) String() string {
	return c.User.FirstName + " " + c.User.LastName
}
