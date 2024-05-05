package common

import "github.com/fengjx/daox"

// GetConditionVal 获取指定字段条件值
func GetConditionVal(field string, conds []daox.Condition) []any {
	for _, cond := range conds {
		if !cond.Disable && field == cond.Field {
			return cond.Vals
		}
	}
	return nil
}
