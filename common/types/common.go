package types

// AddRsp 添加数据返回结果
type AddRsp struct {
	ID int64 `json:"id"`
}

// DelReq 批量删除
type DelReq struct {
	IDs string `json:"ids"`
}

// OKRsp 返回成功或失败
type OKRsp struct {
	Success bool `json:"success"`
}

// BatchUpdate 批量更新参数，rows 需包含 ID 字段
type BatchUpdate struct {
	Rows []map[string]any `json:"rows"`
}

type Empty struct {
}
