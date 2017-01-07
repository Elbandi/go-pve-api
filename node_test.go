package pve

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetNodeList(t *testing.T) {
	httpClient, mux, server := testServer(t, true)
	defer server.Close()

	sampleItem := `{
		"data":[
		   {
		      "cpu" : 0.108209685911794,
		      "disk" : 2408931328,
		      "id" : "node/node1",
		      "level" : "c",
		      "maxcpu" : 24,
		      "maxdisk" : 7739711488,
		      "maxmem" : 135168663552,
		      "mem" : 56629956608,
		      "node" : "node1",
		      "type" : "node",
		      "uptime" : 16793637
		   },
		   {
		      "cpu" : 0.135433498620548,
		      "disk" : 3921641472,
		      "id" : "node/node2",
		      "level" : "c",
		      "maxcpu" : 24,
		      "maxdisk" : 7739711488,
		      "maxmem" : 135167451136,
		      "mem" : 69540794368,
		      "node" : "node2",
		      "type" : "node",
		      "uptime" : 16794484
		   }
		]
	}`

	expectedItem := []string{"node1", "node2"}

	mux.HandleFunc("/api2/json/nodes", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	pveClient := NewPveClient(httpClient, "http://dummy.com/", "", "user", "pass", "pam")
	nodes, err := pveClient.GetNodeList()

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, nodes)
}

func TestGetNode(t *testing.T) {
	httpClient, mux, server := testServer(t, true)
	defer server.Close()

	sampleItem := `{
		"data": {
		   "cpu" : 0,
		   "cpuinfo" : {
		      "cpus" : 24,
		      "hvm" : 1,
		      "mhz" : "2401.710",
		      "model" : "Intel(R) Xeon(R) CPU E5-2620 v2 @ 2.10GHz",
		      "sockets" : 2,
		      "user_hz" : 100
		   },
		   "idle" : 0,
		   "ksm" : {
		      "shared" : 0
		   },
		   "kversion" : "Linux 4.4.13-1-pve #1 SMP Tue Jun 28 10:16:33 CEST 2016",
		   "loadavg" : [
		      "3.55",
		      "2.97",
		      "2.85"
		   ],
		   "memory" : {
		      "free" : 78531592192,
		      "total" : 135168663552,
		      "used" : 56637071360
		   },
		   "pveversion" : "pve-manager/4.2-15/6669ad2c",
		   "rootfs" : {
		      "avail" : 5153824768,
		      "free" : 2744705024,
		      "total" : 7739711488,
		      "used" : 2409119744
		   },
		   "swap" : {
		      "free" : 7276183552,
		      "total" : 7998533632,
		      "used" : 722350080
		   },
		   "uptime" : 16793894,
		   "wait" : 0
		}
	}`

	expectedItem := &Node{
		Cpu : 0,
		CpuInfo : CpuInfo{
			Cpus : 24,
			Hvm : 1,
			Mhz : 2401.710,
			Model : "Intel(R) Xeon(R) CPU E5-2620 v2 @ 2.10GHz",
			Sockets : 2,
			UserHz : 100,
		},
		//		idle : 0,
		//		"ksm" : {
		//			"shared" : 0
		//		},
		KVersion : "Linux 4.4.13-1-pve #1 SMP Tue Jun 28 10:16:33 CEST 2016",
		LoadAvg : []float64{
			3.55,
			2.97,
			2.85,
		},
		Memory : Size{
			Free : 78531592192,
			Total : 135168663552,
			Used : 56637071360,
		},
		PveVersion : "pve-manager/4.2-15/6669ad2c",
		RootFS : Size{
			//			avail : 5153824768,
			Free : 2744705024,
			Total : 7739711488,
			Used : 2409119744,
		},
		Swap : Size{
			Free : 7276183552,
			Total : 7998533632,
			Used : 722350080,
		},
		Uptime : 16793894,
		Wait : 0,
	}

	mux.HandleFunc("/api2/json/nodes/node1/status", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	pveClient := NewPveClient(httpClient, "http://dummy.com/", "", "user", "pass", "pam")
	node, err := pveClient.GetNode("node1")

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, node)
}
