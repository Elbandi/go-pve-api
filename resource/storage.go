package resource

type Storage struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	Storage string `json:"storage"`
	Node    string `json:"node"`
	Disk    uint64 `json:"disk"`
	MaxDisk uint64 `json:"maxdisk"`
}
