package model

import "time"

type Freelancer struct {
	ID                int       `json:"id" valid:"int, optional"`
	AccountId         int64       `json:"accountId" valid:"int, optional"`
	RegistrationDate  time.Time `json:"registrationDate" valid:"-"`
	Country           string    `json:"country" valid:"utfletter"`
	City              string    `json:"city" valid:"utfletter"`
	Address           string    `json:"address" valid:"-"`
	Phone             string    `json:"phone" valid:"-"`
	TagLine           string    `json:"tagline" valid:"-"`
	Overview          string    `json:"overview" valid:"-"`
	ExperienceLevelId int       `json:"experienceLevelId" valid:"in(1|2|3)"`
	SpecialityId      int       `json:"specialityId,string" valid:"int"`
}
