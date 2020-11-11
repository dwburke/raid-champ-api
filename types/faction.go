package types

type Faction struct {
	Id   int    `gorm:"column:id;type:int;primary_key" mapstructure:"id"`
	Name string `gorm:"column:name;type:varchar(255)" mapstructure:"name"`
}

func (Faction) TableName() string {
	return "faction"
}
