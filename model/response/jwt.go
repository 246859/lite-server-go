package response

// Jwt
// @Date 2023-02-23 23:04:03
// @Description: Jwt展示信息
type Jwt struct {
	Access  string `json:"access" form:"access" binding:"required" label:"Access-Token"`
	Refresh string `json:"refresh" form:"refresh" binding:"required" label:"Refresh-Token"`
}
