package entities

import "time"

type Order struct {
	Id         int       `json:"id,int"`
	PetId      int       `json:"petId,int"`
	Quantity   int       `json:"quantity,int"`
	ShipDate   time.Time `json:"shipDate"`
	Status     string    `json:"status"` // placed | approved | delivered
	IsComplete bool      `json:"complete"`
}

type Inventory map[string]int
