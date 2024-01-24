package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Porter struct {
	Code string `bson:"code" json:"code"`
	Name string `bson:"name" json:"name"`
}

func (p Porter) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *Porter) Scan(value interface{}) error {
	if data, ok := value.([]uint8); ok {
		err := json.Unmarshal(data, &p)
		return err
	}
	return fmt.Errorf("failed to unmarshal PorterDB value: %v", value)
}
