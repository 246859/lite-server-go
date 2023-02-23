package response

// UserSimpleInfo
// @Date 2023-02-23 23:00:14
// @Description: 用户简单信息
type UserSimpleInfo struct {
	Id       uint   `json:"id" label:"用户id"`
	Avatar   string `json:"avatar" label:"头像"`
	Username string `json:"username" label:"用户名"`
	Nickname string `json:"nickname" label:"昵称"`
}
