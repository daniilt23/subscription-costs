package model

import "database/sql"

type Subscription struct {
	UserId      string
	ServiceName string
	Price       int
	DateStart   string
	DateEnd     sql.NullString
}

type SubscriptionFind struct {
	UserId      string
	ServiceName string
	DateStart   string
	DateEnd     string
}
