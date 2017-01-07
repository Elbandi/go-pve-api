package resource

type Node struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	Node    string `json:"node"`
	Level   string `json:"level"`
	Uptime  uint32 `json:"uptime"`
	Cpu     float64 `json:"cpu"`
	MaxCpu  uint32 `json:"maxcpu"`
	Mem     uint64 `json:"mem"`
	MaxMem  uint64 `json:"maxmem"`
	Disk    uint64 `json:"disk"`
	MaxDisk uint64 `json:"maxdisk"`
}

