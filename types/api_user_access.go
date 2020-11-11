package types

type ApiUserAccess struct {
	Username string `gorm:"column:username;type:varchar(255);primary_key" json:"username"`
	Route    string `gorm:"column:route;type:varchar(128);primary_key" json:"route"`
	Method   string `gorm:"column:method;type:varchar(7);primary_key" json:"method"`
}

func (ApiUserAccess) TableName() string {
	return "api_user_access"
}

func (a *ApiUserAccess) IsValidMethod(method string) bool {
	switch method {
	case
		"GET",
		"PUT",
		"POST",
		"DELETE",
		"PATCH":
		return true
	}

	return false
}
