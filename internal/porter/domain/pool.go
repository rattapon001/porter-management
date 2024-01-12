package domain

type Pool struct {
	Porters []*Porter
}

func CreateNewPool() *Pool {
	return &Pool{
		Porters: []*Porter{},
	}
}

func (p *Pool) FindAvailablePorter() *Porter {
	for _, porter := range p.Porters {
		if porter.Status == PorterStatusAvailable {
			return porter
		}
	}
	return nil
}

func (p *Pool) PorterRegister(porter *Porter, token string) {
	porter.Available()
	porter.InvokedToken(token)
	p.Porters = append(p.Porters, porter)
}

func (p *Pool) PorterAcceptJob(porter *Porter) {
	porter.AcceptJob()
	for i, _p := range p.Porters {
		if _p.ID == porter.ID {
			p.Porters = append(p.Porters[:i], p.Porters[i+1:]...)
		}
	}
}

func (p *Pool) PorterUnavailable(porter *Porter) {
	porter.Unavailable()
	porter.RevokedToken()
	for i, _p := range p.Porters {
		if _p.ID == porter.ID {
			p.Porters = append(p.Porters[:i], p.Porters[i+1:]...)
		}
	}
}
