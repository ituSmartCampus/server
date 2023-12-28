package models

type Data struct {
	Base
	MeasuredValue string `gorm:"" json:"measuredValue"`
	SensorType    string `gorm:"" json:"sensorType"`
	Time          string `gorm:"" json:"time"`
	Location      string `gorm:"" json:"location"`
	Email         string `gorm:"" json:"email"`
}
