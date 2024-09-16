package dto

import "github.com/google/uuid"

type BidResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	AuthorType string `json:"authorType"`
	AuthorID   uuid.UUID   `json:"authorId"`
	Version    uint   `json:"version"`
	CreatedAt  string `json:"createdAt"`
}
