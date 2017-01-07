package pve

type qemustatusResponse struct {
	Data QemuStatus `json:"data"`
}

type QemuStatus struct {
	Cpu       float64 `json:"cpu"`
	Disk      uint8 `json:"disk"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Template  string `json:"template"`
	Pid       uint16 `json:"pid,string"`
	DiskRead  uint64 `json:"diskread"`
	DiskWrite uint64 `json:"diskwrite"`
	Maxmem    uint64 `json:"maxmem"`
	Maxdisk   uint64 `json:"maxdisk"`
	Uptime    uint32 `json:"uptime"`
	Mem       uint64 `json:"mem"`
	NetIn     uint64 `json:"netin"`
	NetOut    uint64 `json:"netout"`
}

func (client *PveClient) GetQemuStatus(name string, vmid string) (*QemuStatus, error) {
	if err := client.CheckLogin(); err != nil {
		return nil, err
	}
	var status qemustatusResponse
	_, err := client.sling.New().Get("nodes/" + name + "/qemu/" + vmid + "/status/current").ReceiveSuccess(&status)
	if err != nil {
		return nil, err
	}
	return &status.Data, nil
}
