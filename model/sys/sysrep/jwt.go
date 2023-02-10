package sysrep

type Jwt struct {
	Access  string `json:"access" form:"access" binding:"required" label:"Access-Token"`
	Refresh string `json:"refresh" form:"refresh" binding:"required" label:"Refresh-Token"`
}
