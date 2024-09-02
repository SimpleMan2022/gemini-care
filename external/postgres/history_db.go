package postgres

type History struct {
	Id         int    `json:"id" gorm:"primaryKey,autoIncrement"`
	symptoms   string `json:"symptoms" gorm:"type:text[];notnull"`
	Diagnosis  string `json:"diagnosis" gorm:"type:varchar(255);notnull"`
	Confidence uint8  `json:"confidence" gorm:"type:smallint;null"`
}
