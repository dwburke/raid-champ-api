package types

type Champ struct {
	Id         int    `gorm:"column:id;type:int;primary_key" mapstructure:"id"`
	Name       string `gorm:"column:name;type:varchar(255)" mapstructure:"name"`
	Rarity     int    `gorm:"column:rarity;type:int" mapstructure:"rarity"`
	AffinityId int    `gorm:"column:affinity_id;type:int" mapstructure:"affinity_id"`
	FactionId  int    `gorm:"column:faction_id;type:int" mapstructure:"faction_id"`
	//Affinity   Affinity `gorm:"foreignkey:id;references:affinity_id"`
	//Faction    Faction  `gorm:"foreignkey:id;references:faction_id"`
}

func (Champ) TableName() string {
	return "champ"
}
