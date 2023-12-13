package models

import "time"

type Cashier struct {
	Id        uint      `JSON:"id"`
	Name      string    `JSON:"name"`
	Passcode  string    `JSON:"passcode"`
	CreatedAt time.Time `JSON:"created_at"`
	UpdatedAt time.Time `JSON:"updated_at"`
}
