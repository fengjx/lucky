// Code generated by "daox.gen"; DO NOT EDIT.
package entity



// CmsSubscription 用户订阅主题表
type CmsSubscription struct {
    ID int64 `json:"id"` // -
    UID int64 `json:"uid"` // 用户ID
    Topic string `json:"topic"` // 主题
}

func (m *CmsSubscription) GetID() interface{} {
	return m.ID
}