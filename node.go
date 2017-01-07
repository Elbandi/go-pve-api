package pve

import (
	"encoding/json"
	"strconv"
	"github.com/Elbandi/go-pve-api/resource"
)

type nodeResponse struct {
	Data Node `json:"data"`
}

type Size struct {
	Free  uint64 `json:"free"`
	Used  uint64 `json:"used"`
	Total uint64 `json:"total"`
}

type CpuInfo struct {
	Cpus    uint8 `json:"cpus"`
	Hvm     uint8 `json:"hvm"`
	Mhz     float64 `json:"mhz,string"`
	Model   string `json:"model"`
	Sockets uint8 `json:"sockets"`
	UserHz  uint8 `json:"user_hz"`
}

type Node struct {
	Cpu        float64 `json:"cpu"`
	Wait       float64 `json:"wait"`
	CpuInfo    CpuInfo `json:"cpuinfo"`
	KVersion   string `json:"kversion"`
	PveVersion string `json:"pveversion"`
	Uptime     int64 `json:"uptime"`
	LoadAvg    []float64 `json:"loadavg"`
	Memory     Size `json:"memory"`
	Swap       Size `json:"swap"`
	RootFS     Size `json:"rootfs"`
}

func (node *Node) UnmarshalJSON(data []byte) error {
	type Alias Node
	aux := &struct {
		LoadAvg []string `json:"loadavg"`
		*Alias
	}{
		Alias: (*Alias)(node),
	}
	//fmt.Printf("%+v\n", string(data))
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	for _, l := range aux.LoadAvg {
		f, err := strconv.ParseFloat(l, 64)
		if err != nil {
			return err
		}
		node.LoadAvg = append(node.LoadAvg, f)
	}
	return nil
}

func (client *PveClient) GetNodeList() ([]string, error) {
	if err := client.CheckLogin(); err != nil {
		return nil, err
	}
	var res Resource
	_, err := client.sling.New().Get("nodes").ReceiveSuccess(&res)
	if err != nil {
		return nil, err
	}
	var ret []string
	for _, r := range res.Data {
		ret = append(ret, r.(resource.Node).Node)
	}
	return ret, nil
}

func (client *PveClient) GetNode(name string) (*Node, error) {
	if err := client.CheckLogin(); err != nil {
		return nil, err
	}
	var node nodeResponse
	_, err := client.sling.New().Get("nodes/" + name + "/status").ReceiveSuccess(&node)
	if err != nil {
		return nil, err
	}
	return &node.Data, nil
}

