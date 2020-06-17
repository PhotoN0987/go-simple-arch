package models

import "time"

// CommonColumn すべてのDBに共通する項目
type CommonColumn struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
