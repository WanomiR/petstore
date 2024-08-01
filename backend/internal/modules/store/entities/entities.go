package entities

import "time"

type Order struct {
	Id         int       `json:"id,int"`
	PetId      int       `json:"petId,int" example:"1" binding:"required"`
	Quantity   int       `json:"quantity,int" example:"1" binding:"required"`
	ShipDate   time.Time `json:"shipDate" example:"2024-08-01T07:25:40.698Z"`
	Status     string    `json:"status" example:"placed"` // placed | approved | delivered
	IsComplete bool      `json:"complete" example:"true"`
}

type Inventory map[string]int
