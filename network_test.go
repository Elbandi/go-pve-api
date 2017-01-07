package pve

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetNodeNetworks(t *testing.T) {
	httpClient, mux, server := testServer(t, true)
	defer server.Close()

	sampleItem := `{
		"data":[
		   {
		      "active": 1,
		      "autostart": 1,
		      "bridge_fd": "0",
		      "bridge_ports": "eth0.3",
		      "bridge_stp": "off",
		      "families": [
		         "inet"
		      ],
		      "iface": "vmbr3",
		      "method": "manual",
		      "method6": "manual",
		      "priority": 7,
		      "type": "bridge"
		   },
		   {
		      "active": 1,
		      "address": "10.0.0.1",
		      "autostart": 1,
		      "bridge_fd": "0",
		      "bridge_ports": "eth1",
		      "bridge_stp": "off",
		      "families": [
		         "inet"
		      ],
		      "iface": "vmbr0",
		      "method": "static",
		      "method6": "manual",
		      "netmask": "255.255.255.0",
		      "priority": 6,
		      "type": "bridge"
		   },
		   {
		      "active": 1,
		      "exists": 1,
		      "families": [
		         "inet"
		      ],
		      "iface": "eth1",
		      "method": "manual",
		      "method6": "manual",
		      "priority": 5,
		      "type": "eth"
		   },
		   {
		      "active": 1,
		      "exists": 1,
		      "families": [
		         "inet"
		      ],
		      "iface": "eth0",
		      "method": "manual",
		      "method6": "manual",
		      "priority": 4,
		      "type": "eth"
		   }
		]
	}`

	expectedItem := []Network{
		Network{
			Active: true,
			AutoStart: true,
			BridgeForwardDelay: 0,
			BridgePorts: "eth0.3",
			BridgeStp: false,
			Iface: "vmbr3",
			Method: "manual",
			Priority: 7,
			Type: "bridge",
		},
		Network{
			Active: true,
			Address: "10.0.0.1",
			AutoStart: true,
			BridgeForwardDelay: 0,
			BridgePorts: "eth1",
			BridgeStp: false,
			Iface: "vmbr0",
			Method: "static",
			NetMask: "255.255.255.0",
			Priority: 6,
			Type: "bridge",
		},
		Network{
			Active: true,
			Exists: true,
			Iface: "eth1",
			Method: "manual",
			Priority: 5,
			Type: "eth",
		},
		Network{
			Active: true,
			Exists: true,
			Iface: "eth0",
			Method: "manual",
			Priority: 4,
			Type: "eth",
		},
	}

	mux.HandleFunc("/api2/json/nodes/node1/network", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	pveClient := NewPveClient(httpClient, "http://dummy.com/", "", "user", "pass", "pam")
	networks, err := pveClient.GetNodeNetworks("node1")

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, networks)
}
