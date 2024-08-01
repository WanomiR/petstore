package entities

type Category struct {
	Id   int    `json:"id,int"`
	Name string `json:"name" example:"cat"`
}

type Tag struct {
	Id   int    `json:"id,int"`
	Name string `json:"name" example:"fluffy"`
}

type PhotoUrl struct {
	Id    int    `json:"id,int"`
	PetId int    `json:"petId,int"`
	Url   string `json:"url"`
}

type Pet struct {
	Id        int      `json:"id,int"`
	Category  Category `json:"category" binding:"required"`
	Name      string   `json:"name" binding:"required" example:"doggy"`
	PhotoUrls []string `json:"photoUrls"`
	Tags      []Tag    `json:"tags"`
	Status    string   `json:"status" binding:"required" example:"available"` // available | pending | sold
}

type Pets []Pet

type PetTag struct {
	Id    int `db:"id,int"`
	PetId int `db:"pet_id,int"`
	TagId int `db:"tag_id,int"`
}
