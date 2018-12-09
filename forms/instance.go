package forms

//InstanceForm ...
type InstanceForm struct {
	Type   string `form:"type" json:"type" binding:"required,max=32"`
	Name   string `form:"name" json:"name" binding:"max=32"`
	ID	string	 `form:"id" json:"id" binding:"max=22"`
	Readme string `form:"readme" json:"readme" binding:"max=1000"`
}

type InstanceRsp struct {
	Page	Page	 `form:"page" json:"page" binding:"max=22"`
	List []map[string]interface{} `form:"list" json:"list" binding:"required,max=1000"`
}

type Page struct {
	Size int `form:"page" json:"page" binding:"max=22"`
	Current int `form:"current" json:"current" binding:"max=22"`
	Total int64 `form:"total" json:"total" binding:"max=22"`
}
