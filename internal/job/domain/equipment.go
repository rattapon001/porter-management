package domain

type EquipmentId string

type Equipment struct {
	ID     EquipmentId `bson:"_id" gorm:"primaryKey;type:uuid"`
	JobId  JobId       `bson:"job_id" json:"job_id"`
	Name   string      `bson:"name" json:"name"`
	Amount int         `bson:"amount" json:"amount"`
}

func NewEquipment(id EquipmentId, jobId JobId, name string, amount int) (*Equipment, error) {
	return &Equipment{
		ID:     id,
		JobId:  jobId,
		Name:   name,
		Amount: amount,
	}, nil
}

func (e *Equipment) UpdateAmont(amont int) {
	e.Amount = amont
}
