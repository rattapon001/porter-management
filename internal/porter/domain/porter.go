package domain

import "github.com/google/uuid"

type PorterId string
type PorterStatus string
type PorterCode string

const (
	PorterStatusAvailable   PorterStatus = "available"
	PorterStatusWorking     PorterStatus = "working"
	PorterStatusUnavailable PorterStatus = "unavailable"
)

type Porter struct {
	ID     PorterId
	Name   string
	Code   PorterCode
	Status PorterStatus
	Token  string
}

func NewPorter(name string, code string, token string) (*Porter, error) {

	ID, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}
	porter := Porter{
		ID:     PorterId(ID.String()),
		Name:   name,
		Code:   PorterCode(code),
		Status: PorterStatusUnavailable,
	}
	porter.InvokedToken(token)
	return &porter, nil
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
