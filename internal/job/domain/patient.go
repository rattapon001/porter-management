package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Patient struct {
	Name string `bson:"name" json:"name"`
	HN   string `bson:"hn" json:"hn"`
}

func (p Patient) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *Patient) Scan(value interface{}) error {
	if data, ok := value.([]uint8); ok {
		err := json.Unmarshal(data, &p)
		return err
	}
	return fmt.Errorf("failed to unmarshal PatientDB value: %v", value)
}
