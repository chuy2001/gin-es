package forms

//TableForm ...
type TableForm struct {
	Name   string `form:"name" json:"name" binding:"required,max=32"`
	Alias string `form:"alias" json:"alias" binding:"required,max=32"`
	Readme string `form:"readme" json:"readme" binding:"required,max=1000"`
	Fields []*FieldInfo `form:"fields" json:"fields" binding:"required,max=1000"`
}

//TableForm ...
type FieldInfo struct {
	Name   string `form:"name" json:"name" binding:"required,max=32"`
	Alias string `form:"alias" json:"alias" binding:"required,max=32"`
	Readme string `form:"readme" json:"readme" binding:"required,max=1000"`
	Is_multi bool `form:"is_multi" json:"is_multi" binding:"required"`
	Required bool `form:"required" json:"required" binding:"required"`
}