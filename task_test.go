package pve

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestGetTasks(t *testing.T) {
	httpClient, mux, server := testServer(t, true)
	defer server.Close()

	sampleItem := `{
		"data":[
		   {
		      "endtime": 1483342697,
		      "id": "",
		      "node": "node1",
		      "saved": "1",
		      "starttime": "1483334102",
		      "status": "OK",
		      "type": "vzdump",
		      "upid": "UPID:node1:00003A02:5F1C52E9:5869E1D6:vzdump::root@pam:",
		      "user": "root@pam"
		   },
		   {
		      "endtime": 1483331711,
		      "id": "",
		      "node": "node1",
		      "saved": "1",
		      "starttime": "1483331703",
		      "status": "OK",
		      "type": "aptupdate",
		      "upid": "UPID:node1:00000D77:5F18AA10:5869D877:aptupdate::root@pam:",
		      "user": "root@pam"
		   }
		]
	}`

	expectedItem := []Task{
		Task{
			EndTime: time.Unix(1483342697, 0),
			Id: "",
			Node: "node1",
			Saved: true,
			StartTime: time.Unix(1483334102, 0),
			Status: "OK",
			Type: "vzdump",
			UPid: "UPID:node1:00003A02:5F1C52E9:5869E1D6:vzdump::root@pam:",
			User: "root@pam",
		},
		Task{
			EndTime: time.Unix(1483331711, 0),
			Id: "",
			Node: "node1",
			Saved: true,
			StartTime: time.Unix(1483331703, 0),
			Status: "OK",
			Type: "aptupdate",
			UPid: "UPID:node1:00000D77:5F18AA10:5869D877:aptupdate::root@pam:",
			User: "root@pam",
		},
	}

	mux.HandleFunc("/api2/json/cluster/tasks", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	pveClient := NewPveClient(httpClient, "http://dummy.com/", "", "user", "pass", "pam")
	tasks, err := pveClient.GetTasks()

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, tasks)
}

