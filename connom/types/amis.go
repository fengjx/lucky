package types

// amis 接口协议适配

type SelectResp struct {
	Options []*Option `json:"options"`
}

type Option struct {
	Label    string    `json:"label"`
	Value    any       `json:"value"`
	Children []*Option `json:"children,omitempty"`
}

type AmisPageResp[T any] struct {
	Items []T   `json:"items"`
	Total int64 `json:"total"`
}
