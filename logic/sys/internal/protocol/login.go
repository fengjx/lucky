package protocol

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfoResp struct {
	UserInfo *UserInfo `json:"user_info"`
}

type LoginResp struct {
	UserInfoResp
	Token string `json:"token"`
}
