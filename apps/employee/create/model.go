package create

import "time"

type Auth struct {
	PublicId string
	Email    string
	Password string
	Role     string
	IsActive bool
}

type Employee struct {
	Id        int
	PublicId  string
	Name      string
	Profile   string
	AuthId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a Auth) IsExist() bool {
	return a == Auth{}
}
