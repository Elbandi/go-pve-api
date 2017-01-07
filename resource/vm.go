package resource

type Vm struct {
	Id        string `json:"id"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	Vmid      uint64 `json:"vmid"`
	Node      string `json:"node"`
	Status    string `json:"status"`
	Uptime    uint32 `json:"uptime"`
	//Template  string `json:"template"`
	Cpu       float64 `json:"cpu"`
	MaxCpu    uint16 `json:"maxcpu"`
	Mem       uint64 `json:"mem"`
	MaxMem    uint64 `json:"maxmem"`
	Disk      uint64 `json:"disk"`
	MaxDisk   uint64 `json:"maxdisk"`
	DiskRead  uint64 `json:"diskread"`
	DiskWrite uint64 `json:"diskwrite"`
	NetIn     uint64 `json:"netin"`
	NetOut    uint64 `json:"netout"`
}
