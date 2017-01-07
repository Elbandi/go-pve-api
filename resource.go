package pve

import (
	"encoding/json"
	"errors"
	"github.com/Elbandi/go-pve-api/resource"
)

type Resource struct {
	Data []interface{} `json:"data"`
}

type ResourceResponse struct {
	Data []json.RawMessage `json:"data"`
}

func (r *Resource) UnmarshalJSON(data []byte) error {
	var alias ResourceResponse
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	for _, element := range alias.Data {
		//		fmt.Printf("%+v %+v\n", index, string(element))
		var res resource.ResourceBase
		if err := json.Unmarshal(element, &res); err != nil {
			return err
		}
		switch res.Type {
		case "pool":
			var pool resource.Pool
			if err := json.Unmarshal(element, &pool); err != nil {
				return err
			}
			r.Data = append(r.Data, pool)
		case "qemu":
			var vm resource.Vm
			if err := json.Unmarshal(element, &vm); err != nil {
				return err
			}
			r.Data = append(r.Data, vm)
		case "node":
			var node resource.Node
			if err := json.Unmarshal(element, &node); err != nil {
				return err
			}
			r.Data = append(r.Data, node)
		case "storage":
			var storage resource.Storage
			if err := json.Unmarshal(element, &storage); err != nil {
				return err
			}
			r.Data = append(r.Data, storage)
		default:
			return errors.New("invalid type")
		}
	}
	return nil
}

func (client *PveClient) GetResources() ([]interface{}, error) {
	if err := client.CheckLogin(); err != nil {
		return nil, err
	}
	var response Resource
	_, err := client.sling.New().Get("cluster/resources").ReceiveSuccess(&response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}
