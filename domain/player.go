package domain

type Player struct {
	ID     string `json:"id"`
	Secret string `json:"-"`
	Name   string `json:"name"`
}

func (p *Player) GetAuth() Auth {
	return Auth{
		ID:     p.ID,
		Secret: p.Secret,
		Name:   p.Name,
	}
}

type Auth struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
	Name   string `json:"name"`
}

func (a *Auth) GetPlayer() Player {
	return Player{
		ID:     a.ID,
		Secret: a.Secret,
		Name:   a.Name,
	}
}
