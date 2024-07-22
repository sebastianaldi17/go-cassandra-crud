package entity

import "time"

type CartCount struct {
	UserID     string    `json:"userid"`
	ItemCount  int64     `json:"item_count"`
	LastUpdate time.Time `json:"last_update_timestamp"`
}
