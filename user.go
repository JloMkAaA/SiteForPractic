package SiteForPractic

import "errors"

type User struct {
	Id              int    `json:"-" db:"id"`
	Phone_number    uint64 `json:"phone_number" binding:"required"`
	Password        string `json:"password" binding:"required"`
	Position        string `json:"position" binding:"required"`
	Expirience      uint16 `json:"expirience" binding:"required"`
	Education_level string `json:"education_level" binding:"required"`
	Description     string `json:"description" binding:"required"`
}

type UpdateUser struct {
	Phone_number    *uint64 `json:"phone_number"`
	Password        *string `json:"password"`
	Position        *string `json:"position"`
	Expirience      *uint16 `json:"expirience"`
	Education_level *string `json:"education_level"`
	Description     *string `json:"description"`
}

func (i UpdateUser) Validate() error {
	if i.Phone_number == nil && i.Password == nil && i.Position == nil && i.Expirience == nil && i.Education_level == nil && i.Description == nil {
		return errors.New("update structure has not values")
	}

	return nil
}
