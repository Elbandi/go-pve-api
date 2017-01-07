package pve

import (
	"encoding/json"
)

type NetworkResponse struct {
	Data []Network `json:"data"`
}

type Network struct {
	Active             bool `json:"active"`
	AutoStart          bool `json:"autostart"`
	Exists             bool `json:"exists"`
	BridgeForwardDelay int16 `json:"bridge_fd,string"`
	BridgePorts        string `json:"bridge_ports"`
	BridgeStp          bool `json:"bridge_stp"`
	Iface              string `json:"iface"`
	Method             string `json:"method"`
	Address            string `json:"address"`
	NetMask            string `json:"netmask"`
	Gateway            string `json:"gateway"`
	Type               string `json:"type"`
	Priority           uint8 `json:"priority"`
}

func (task *Network) UnmarshalJSON(data []byte) error {
	type Alias Network
	aux := &struct {
		Active    int8 `json:"active"`
		AutoStart int8 `json:"autostart"`
		Exists    int8 `json:"exists"`
		BridgeStp string `json:"bridge_stp"`
		*Alias
	}{
		Alias: (*Alias)(task),
	}
	//fmt.Printf("%+v\n", string(data))
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	task.Active = aux.Active == 1
	task.AutoStart = aux.AutoStart == 1
	task.Exists = aux.Exists == 1
	task.BridgeStp = aux.BridgeStp == "on"
	return nil
}

func (client *PveClient) GetNodeNetworks(name string) ([]Network, error) {
	if err := client.CheckLogin(); err != nil {
		return nil, err
	}
	var network NetworkResponse
	_, err := client.sling.New().Get("nodes/" + name + "/network").ReceiveSuccess(&network)
	if err != nil {
		return nil, err
	}
	return network.Data, nil
}

