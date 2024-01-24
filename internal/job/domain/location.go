package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Location struct {
	From string `bson:"from" json:"from"`
	To   string `bson:"to" json:"to"`
}

func (l Location) Value() (driver.Value, error) {
	return json.Marshal(l)
}

func (l *Location) Scan(value interface{}) error {
	if data, ok := value.([]uint8); ok {
		err := json.Unmarshal(data, &l)

		return err
	}
	return fmt.Errorf("failed to unmarshal Location value: %v", value)
}
