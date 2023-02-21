package crm_structures

import "time"

type GetPing struct {
	DateTime time.Time `json:"dateTime"`
	Endpoint string    `json:"endpoint"`
	Valid    bool      `json:"valid"`
}
