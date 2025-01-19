package http

const (
	// AdminAPI 管理后台接口前缀
	AdminAPI = "/admin"
	// API 接口前缀，需要鉴权
	API = "/api"
	// OpenAPI 不需要鉴权的接口前缀
	OpenAPI = API + "/open"

	RequestHeaderAdminUID   = "X-Admin-UID"
	RequestHeaderAdminToken = "X-Admin-Token"

	ResponseHeaderServer       = "Server"
	ResponseHeaderRefreshToken = "X-Refresh-Token"
)
