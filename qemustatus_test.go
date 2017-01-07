package pve

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
)

func TestGetQemuStatus(t *testing.T) {
	httpClient, mux, server := testServer(t, true)
	defer server.Close()

	sampleItem := `{
		"data": {
		   "balloon": 2147483648,
		   "ballooninfo": {
		      "actual": 2147483648,
		      "free_mem": 655257600,
		      "last_update": 1484173167,
		      "major_page_faults": 7748,
		      "max_mem": 2147483648,
		      "mem_swapped_in": 13955072,
		      "mem_swapped_out": 31518720,
		      "minor_page_faults": 20858086598,
		      "total_mem": 2105524224
		   },
		   "blockstat": {
		      "virtio0": {
		         "account_failed": true,
		         "account_invalid": true,
		         "failed_flush_operations": 0,
		         "failed_rd_operations": 0,
		         "failed_wr_operations": 0,
		         "flush_operations": 1431238,
		         "flush_total_time_ns": 170125661143,
		         "idle_time_ns": 11831295749,
		         "invalid_flush_operations": 0,
		         "invalid_rd_operations": 0,
		         "invalid_wr_operations": 0,
		         "rd_bytes": 420098048,
		         "rd_merged": 5848,
		         "rd_operations": 21511,
		         "rd_total_time_ns": 15531500850,
		         "timed_stats": [],
		         "wr_bytes": 57117212672,
		         "wr_highest_offset": 8522534912,
		         "wr_merged": 27353,
		         "wr_operations": 4638474,
		         "wr_total_time_ns": 10546623776405
		      }
		   },
		   "cpu": 0.0151572733196034,
		   "cpus": 2,
		   "disk": 0,
		   "diskread": 420098048,
		   "diskwrite": 57117212672,
		   "freemem": 655257600,
		   "ha": {
		      "managed": 0
		   },
		   "maxdisk": 8589934592,
		   "maxmem": 2147483648,
		   "mem": 1450266624,
		   "name": "test",
		   "netin": 445735332443,
		   "netout": 21036187680,
		   "nics": {
		      "tap111i0": {
		         "netin": "445735332443",
		         "netout": "21036187680"
		      }
		   },
		   "pid": "5741",
		   "qmpstatus": "running",
		   "status": "running",
		   "template": "",
		   "uptime": 14555060
		}
	}`

	expectedItem := &QemuStatus{
		Cpu: 0.0151572733196034,
		Disk: 0,
		DiskRead: 420098048,
		DiskWrite: 57117212672,
		Maxdisk: 8589934592,
		Maxmem: 2147483648,
		Mem: 1450266624,
		Name: "test",
		NetIn: 445735332443,
		NetOut: 21036187680,
		Pid: 5741,
		Status: "running",
		Template: "",
		Uptime: 14555060,
	}

	mux.HandleFunc("/api2/json/nodes/node2/qemu/123/status/current", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	pveClient := NewPveClient(httpClient, "http://dummy.com/", "", "user", "pass", "pam")
	status, err := pveClient.GetQemuStatus("node2", "123")

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if !reflect.DeepEqual(expectedItem, status) {
		t.Errorf("expected %v, got %v", expectedItem, status)
	}
}
