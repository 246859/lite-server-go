package request

type IdWrap struct {
	Id uint `json:"id" form:"id" uri:"id" binding:"required,gte=1" label:"id"`
}
