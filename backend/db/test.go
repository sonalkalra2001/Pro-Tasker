package db

type Test struct {
	TestID int    `json:"test_id" gorm:"primaryKey;autoIncrement"`
	Name   string `json:"name" gorm:"column:name;default:null"`
}
