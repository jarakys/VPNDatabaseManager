package Models

type ConfigItemModel struct {
	Id     uint `gorm:"id"`
	UserId string
	Config string
	Type   string
	Name   string
	Date   int64
}

func (ConfigItemModel) TableName() string {
	return "configs"
}
