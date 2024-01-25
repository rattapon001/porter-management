package domain

type EquipmentId string

type Equipment struct {
	ID    EquipmentId `bson:"_id" gorm:"primaryKey"`
	JobId JobId       `bson:"job_id" json:"job_id"`
	Name  string      `bson:"name" json:"name"`
	Amont int         `bson:"amont" json:"amont"`
}

func NewEquipment(id EquipmentId, jobId JobId, name string, amont int) (*Equipment, error) {
	return &Equipment{
		ID:    id,
		JobId: jobId,
		Name:  name,
		Amont: amont,
	}, nil
}

func (e *Equipment) UpdateAmont(amont int) {
	e.Amont = amont
}
