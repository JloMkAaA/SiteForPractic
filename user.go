package SiteForPractic

type User struct {
	Id              int    `gorm:"primaryKey;autoIncrement;unique" json:"id"`
	Phone_number    int    `json:"phone_number" binding:"required"`
	Password        string `json:"password" binding:"required"`
	Position        string `json:"position" binding:"required"`
	Expirience      uint16 `json:"expirience" binding:"required"`
	Education_level string `json:"education_level" binding:"required"`
	Description     string `json:"description" binding:"required"`
}
