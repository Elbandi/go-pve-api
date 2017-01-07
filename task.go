package pve

import (
	"encoding/json"
	//"fmt"
	"time"
)

type TaskResponse struct {
	Data []Task `json:"data"`
}

type Task struct {
	Id        string `json:"id"`
	Node      string `json:"node"`
	Saved     bool `json:"saved"`
	StartTime time.Time `json:"starttime"`
	EndTime   time.Time `json:"endtime"`
	Status    string `json:"status"`
	Type      string `json:"type"`
	UPid      string `json:"upid"`
	User      string `json:"user"`
}

func (task *Task) UnmarshalJSON(data []byte) error {
	type Alias Task
	aux := &struct {
		Saved     int8 `json:"saved,string"`
		StartTime int64 `json:"starttime,string"`
		Endtime   int64 `json:"endtime"`
		*Alias
	}{
		Alias: (*Alias)(task),
	}
	//fmt.Printf("%+v\n", string(data))
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	task.StartTime = time.Unix(aux.StartTime, 0)
	task.EndTime = time.Unix(aux.Endtime, 0)
	task.Saved = aux.Saved == 1
	return nil
}

func (client *PveClient) GetTasks() ([]Task, error) {
	if err := client.CheckLogin(); err != nil {
		return nil, err
	}
	var tasks TaskResponse
	_, err := client.sling.New().Get("cluster/tasks").ReceiveSuccess(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks.Data, nil
}

