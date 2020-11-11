package types

type Champ struct {
	Id         string `gorm:"column:id;type:uuid;primary_key" json:"id"`
	Name       string `gorm:"column:name;type:varchar(255)" json:"name"`
	Rarity     int    `gorm:"column:rarity;type:int" json:"rarity"`
	AffinityId int    `gorm:"column:affinity_id;type:int" json:"affinity_id"`
	FactionId  int    `gorm:"column:faction_id;type:int" json:"faction_id"`
	//Affinity   Affinity `gorm:"foreignkey:id;references:affinity_id"`
	//Faction    Faction  `gorm:"foreignkey:id;references:faction_id"`
}

func (Champ) TableName() string {
	return "champ"
}
