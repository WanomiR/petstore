package entities

type User struct {
	Id         int    `json:"id,int"`
	Username   string `json:"username" example:"johndoe001"`
	FirstName  string `json:"firstName" example:"John"`
	LastName   string `json:"lastName" example:"Doe"`
	Email      string `json:"email" example:"johndoe@example.com"`
	Password   string `json:"password" example:"123456"`
	Phone      string `json:"phone" example:"7-999-999-99-99"`
	UserStatus int    `json:"userStatus,int" example:"0"`
}
