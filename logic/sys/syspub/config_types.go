package syspub

import "github.com/fengjx/go-halo/json"

type ConfigDTO struct {
	Scope string `json:"scope"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (dto *ConfigDTO) ToData(data any) error {
	if dto == nil || len(dto.Value) == 0 {
		return nil
	}
	return json.FromJson(dto.Value, data)
}
