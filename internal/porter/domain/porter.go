package domain

import "github.com/google/uuid"

type PorterId string
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
	Token      string
}

func CreatedNewPorter(name string, porterCode string, token string) (*Porter, error) {

	ID, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}
	return &Porter{
		ID:         PorterId(ID.String()),
		Name:       name,
		PorterCode: porterCode,
		Status:     PorterStatusUnavailable,
	}, nil
}

func (p *Porter) Working() {
	p.Status = PorterStatusWorking
}

func (p *Porter) Available() {
	p.Status = PorterStatusAvailable
}

func (p *Porter) Unavailable() {
	p.Status = PorterStatusUnavailable
}

func (p *Porter) InvokedToken(token string) {
	p.Token = token
}

func (p *Porter) RevokedToken() {
	p.Token = ""
}
