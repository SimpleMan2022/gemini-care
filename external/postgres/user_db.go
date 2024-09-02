package postgres

type User struct {
	Id       int    `json:"id" gorm:"primaryKey,autoIncrement"`
	Username string `json:"username" gorm:"unique;type:varchar(255);notnull"`
	Email    string `json:"email" gorm:"unique;type:varchar(255);notnull"`
	Password string `json:"password" gorm:"type:varchar(255);null"`
	Provider string `json:"provider" gorm:"type:varchar(255);notnull"`
}
