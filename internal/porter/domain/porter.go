package domain

type PorterId int64
type PorterStatus string

const (
	PorterStatusAvailable   PorterStatus = "available"
	PorterStatusWorking     PorterStatus = "working"
	PorterStatusUnavailable PorterStatus = "unavailable"
)

type Porter struct {
	Id         PorterId
	Name       string
	PorterCode string
	Status     PorterStatus
}

func CreateNewPorter(name string, porterCode string) (*Porter, error) {
	return &Porter{
		Name:       name,
		PorterCode: porterCode,
		Status:     PorterStatusAvailable,
	}, nil
}

func (p *Porter) AcceptJob() {
	p.Status = PorterStatusWorking
}

func (p *Porter) CompleteJob() {
	p.Status = PorterStatusAvailable
}

func (p *Porter) Unavailable() {
	p.Status = PorterStatusUnavailable
}
