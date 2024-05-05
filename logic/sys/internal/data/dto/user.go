package dto

// UserInfoDTO dto for entity.SysUser
type UserInfoDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
	Status   string `json:"status"`
}
