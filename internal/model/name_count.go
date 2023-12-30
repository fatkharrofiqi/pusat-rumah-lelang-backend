package model

type NameCount struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	PropertyCount int    `json:"property_count"`
}
