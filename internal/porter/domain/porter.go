package domain

import "github.com/google/uuid"

type PorterId int64
type PorterStatus string

const (
	PorterStatusAvailable   PorterStatus = "available"
	PorterStatusWorking     PorterStatus = "working"
	PorterStatusUnavailable PorterStatus = "unavailable"
)

type Porter struct {
	ID         PorterId
	Name       string
	PorterCode string
	Status     PorterStatus
}

func CreateNewPorter(name string, porterCode string) (*Porter, error) {

	ID, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}
	return &Porter{
		ID:         PorterId(ID.ID()),
		Name:       name,
		PorterCode: porterCode,
		Status:     PorterStatusUnavailable,
	}, nil
}

func (p *Porter) AcceptJob() {
	p.Status = PorterStatusWorking
}

func (p *Porter) Available() {
	p.Status = PorterStatusAvailable
}

func (p *Porter) Unavailable() {
	p.Status = PorterStatusUnavailable
}
