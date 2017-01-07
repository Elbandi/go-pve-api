package resource

type Pool struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Pool   string `json:"pool"`
	Uptime uint32 `json:"uptime"`
	Cpu    float64 `json:"cpu"`
	MaxCpu uint32 `json:"maxcpu"`
	Mem    uint64 `json:"mem"`
	MaxMem uint64 `json:"maxmem"`
}
