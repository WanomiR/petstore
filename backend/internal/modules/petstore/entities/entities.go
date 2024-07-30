package entities

import "time"

type Category struct {
	Id   int    `json:"id,int"`
	Name string `json:"name"`
}

type Tag struct {
	Id   int    `json:"id,int"`
	Name string `json:"name"`
}

type PhotoUrl struct {
	Id    int    `json:"id,int"`
	PetId int    `json:"pet_id,int"`
	Url   string `json:"url"`
}

type Pet struct {
	Id        int      `json:"id,int"`
	Category  Category `json:"category"`
	Name      string   `json:"name" binding:"required" example:"doggy"`
	PhotoUrls []string `json:"photoUrls" binding:"required"`
	Tags      []Tag    `json:"tags"`
	Status    string   `json:"status" example:"active"` // available | pending | sold
}

type Order struct {
	Id       int       `json:"id,int"`
	PetId    int       `json:"petId,int"`
	Quantity int       `json:"quantity,int"`
	ShipDate time.Time `json:"shipDate"`
	Status   string    `json:"status"` // placed | approved | delivered
	Complete bool      `json:"complete"`
}

type User struct {
	Id         int    `json:"id,int"`
	Username   string `json:"username"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	UserStatus int    `json:"userStatus,int"`
}
