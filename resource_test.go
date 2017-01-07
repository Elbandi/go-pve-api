package pve

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/Elbandi/go-pve-api/resource"
)

func TestGetResources(t *testing.T) {
	httpClient, mux, server := testServer(t, true)
	defer server.Close()

	sampleItem := `{
		"data":[
		   {
		      "cpu": 0.120845194852149,
		      "disk": 3923591168,
		      "id": "node/node1",
		      "level": "c",
		      "maxcpu": 24,
		      "maxdisk": 7739711488,
		      "maxmem": 135167451136,
		      "mem": 69552300032,
		      "node": "node1",
		      "type": "node",
		      "uptime": 16797164
		   },
		   {
		      "disk": 3137686503424,
		      "id": "storage/node1/lvmvolume",
		      "maxdisk": 4055438983168,
		      "node": "node1",
		      "storage": "lvmvolume",
		      "type": "storage"
		   },
		   {
		      "cpu": 0.0286182416314956,
		      "disk": 0,
		      "diskread": 420098048,
		      "diskwrite": 57117724672,
		      "id": "qemu/111",
		      "maxcpu": 2,
		      "maxdisk": 8589934592,
		      "maxmem": 2147483648,
		      "mem": 1465028608,
		      "name": "test",
		      "netin": 445749761684,
		      "netout": 21036491302,
		      "node": "node1",
		      "status": "running",
		      "template": 0,
		      "type": "qemu",
		      "uptime": 14555394,
		      "vmid": 111
		   }
		]
	}`

	expectedItem := []interface{}{
		resource.Node{
			Cpu: 0.120845194852149,
			Disk: 3923591168,
			Id: "node/node1",
			Level: "c",
			MaxCpu: 24,
			MaxDisk: 7739711488,
			MaxMem: 135167451136,
			Mem: 69552300032,
			Node: "node1",
			Type: "node",
			Uptime: 16797164,
		},
		resource.Storage{
			Disk: 3137686503424,
			Id: "storage/node1/lvmvolume",
			MaxDisk: 4055438983168,
			Node: "node1",
			Storage: "lvmvolume",
			Type: "storage",
		},
		resource.Vm{
			Cpu: 0.0286182416314956,
			Disk: 0,
			DiskRead: 420098048,
			DiskWrite: 57117724672,
			Id: "qemu/111",
			MaxCpu: 2,
			MaxDisk: 8589934592,
			MaxMem: 2147483648,
			Mem: 1465028608,
			Name: "test",
			NetIn: 445749761684,
			NetOut: 21036491302,
			Node: "node1",
			Status: "running",
			Type: "qemu",
			Uptime: 14555394,
			Vmid: 111,
		},
	}

	mux.HandleFunc("/api2/json/cluster/resources", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	pveClient := NewPveClient(httpClient, "http://dummy.com/", "", "user", "pass", "pam")
	status, err := pveClient.GetResources()

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if !reflect.DeepEqual(expectedItem, status) {
		t.Errorf("expected %v, got %v", expectedItem, status)
	}
}

