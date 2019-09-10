package main

import "time"

type User struct {
	ID                    uint      `json:"id" gorm:"primary_key"`
	Mobile                string    `json:"mobile" gorm:"not null;size:11;unique"`
	CreatedAt             time.Time `json:"created_at" gorm:"not null"`
	BlockedAt             *time.Time
	LatestAuthenticatedAt *time.Time

	Roles []Role `gorm:"many2many:user_roles"`
}

type Role struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Slug  string `json:"slug"`
	Title string `json:"title"`

	Users []User `gorm:"many2many:user_roles"`
}

type CreateRoleReq struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
}
