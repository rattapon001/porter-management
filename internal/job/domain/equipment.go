package domain

import "github.com/google/uuid"

type EquipmentId string

type Equipment struct {
	ID          EquipmentId `bson:"_id" gorm:"primaryKey;type:uuid"`
	EquipmentId EquipmentId `bson:"equipment_id" json:"equipmentId"`
	JobId       JobId       `bson:"job_id" json:"jobId"`
	Amount      int         `bson:"amount" json:"amount"`
}

func NewEquipment(sourceId EquipmentId, jobId JobId, amount int) (*Equipment, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &Equipment{
		ID:          EquipmentId(id.String()),
		EquipmentId: sourceId,
		JobId:       jobId,
		Amount:      amount,
	}, nil
}

func (e *Equipment) UpdateAmont(amont int) {
	e.Amount = amont
}
