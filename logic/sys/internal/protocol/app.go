package protocol

import "github.com/fengjx/lucky/logic/sys/syspub"

type AppDataResp struct {
	Config map[string][]*syspub.ConfigDTO `json:"config"`
	Dict   map[string][]*syspub.DictDTO   `json:"dict"`
}
