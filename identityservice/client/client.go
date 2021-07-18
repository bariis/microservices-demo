package client

import (
	"github.com/bariis/microservices-demo/db"
	"github.com/google/uuid"
)

type Client struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Password  string    `gorm:"size:255;not null;"`
	FullName  string    `gorm:"size:255;not null;"`
	Email     string    `gorm:"size:255;not null;"`
}

func (c *Client) DoesExist(email string) bool {
	client := &Client{}
	db.Database.Find(&client, "email = ?", email)
	return client.Email != ""
}

