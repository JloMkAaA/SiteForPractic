package SiteForPractic

type User struct {
	Id              int    `gorm:"primaryKey;autoIncrement;unique" json:"id"`
	Phone_number    int    `json:"phone_number"`
	Password        string `json:"password"`
	Position        string `json:"position"`
	Expirience      uint16 `json:"expirience"`
	Education_level int    `json:"education_level"`
	Description     string `json:"description"`
}
