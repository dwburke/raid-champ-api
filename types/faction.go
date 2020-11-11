package types

type Faction struct {
	Id   int    `gorm:"column:id;type:int;primary_key" json:"id"`
	Name string `gorm:"column:name;type:varchar(255)" json:"name"`
}

func (Faction) TableName() string {
	return "faction"
}
