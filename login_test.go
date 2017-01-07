package pve

import (
	"fmt"
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	httpClient, mux, server := testServer(t, false)
	defer server.Close()

	sampleItem := `{
	   "data":{
	      "CSRFPreventionToken":"5876C03B:token",
	      "username":"root@pam",
	      "ticket":"PVE:root@pam:5876C03B::ticket",
	      "cap":{
		 "dc":{
		    "Sys.Audit":1
		 },
		 "access":{
		    "Group.Allocate":1,
		    "User.Modify":1
		 },
		 "storage":{
		    "Datastore.AllocateSpace":1,
		    "Datastore.Allocate":1,
		    "Permissions.Modify":1,
		    "Datastore.AllocateTemplate":1,
		    "Datastore.Audit":1
		 },
		 "vms":{
		    "VM.Backup":1,
		    "Permissions.Modify":1,
		    "VM.Config.Network":1,
		    "VM.Config.CPU":1,
		    "VM.Audit":1,
		    "VM.Config.CDROM":1,
		    "VM.Clone":1,
		    "VM.Migrate":1,
		    "VM.Config.Memory":1,
		    "VM.Config.Disk":1,
		    "VM.Allocate":1,
		    "VM.Monitor":1,
		    "VM.Config.HWType":1,
		    "VM.Snapshot":1,
		    "VM.Console":1,
		    "VM.Config.Options":1,
		    "VM.PowerMgmt":1
		 },
		 "nodes":{
		    "Sys.Console":1,
		    "Sys.Audit":1,
		    "Sys.PowerMgmt":1,
		    "Sys.Modify":1,
		    "Sys.Syslog":1
		 }
	      }
	   }
	}`

	expectedItem := &Login{
		CSRFPreventionToken:"5876C03B:token",
		Username:"root@pam",
		Ticket:"PVE:root@pam:5876C03B::ticket",
	}

	mux.HandleFunc("/api2/json/access/ticket", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostForm(t, map[string]string{"username": "user", "password": "pass", "realm": "pam"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, sampleItem)
	})

	pveClient := NewPveClient(httpClient, "http://dummy.com/", "", "user", "pass", "pam")
	networks, err := pveClient.Login()

	assert.Nil(t, err)
	assert.Equal(t, expectedItem, networks)
}
