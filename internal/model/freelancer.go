package model

import (
	"github.com/microcosm-cc/bluemonday"
)

type Freelancer struct {
	ID                int64  `json:"id" valid:"int, optional"`
	AccountId         int64  `json:"accountId" valid:"int, optional"`
	Country           int64  `json:"country" valid:"utfletter"`
	City              int64  `json:"city" valid:"utfletter"`
	Address           string `json:"address" valid:"-"`
	Phone             string `json:"phone" valid:"-"`
	TagLine           string `json:"tagline" valid:"-"`
	Overview          string `json:"overview" valid:"-"`
	ExperienceLevelId int64  `json:"experienceLevelId" valid:"in(1|2|3)"`
	SpecialityId      int64  `json:"specialityId,string" valid:"int"`
	Avatar            string `json:"avatar"`
}

type FreelancerContracts struct {
	ActiveContracts		int64	`json:"active"`
	SuccessContracts	int64	`json:"success"`
}

type FreelancerOutput struct {
	ID                int64  `json:"id" valid:"int, optional"`
	AccountId         int64  `json:"accountId" valid:"int, optional"`
	Country           string `json:"country" valid:"utfletter"`
	City              string `json:"city" valid:"utfletter"`
	Address           string `json:"address" valid:"-"`
	Phone             string `json:"phone" valid:"-"`
	TagLine           string `json:"tagline" valid:"-"`
	Overview          string `json:"overview" valid:"-"`
	ExperienceLevelId int64  `json:"experienceLevelId" valid:"in(1|2|3)"`
	SpecialityId      int64  `json:"specialityId,string" valid:"int"`
	Avatar            string `json:"avatar"`
}

func (freel *Freelancer) Sanitize(sanitizer *bluemonday.Policy) {
	freel.Address = sanitizer.Sanitize(freel.Address)
	freel.Phone = sanitizer.Sanitize(freel.Phone)
	freel.TagLine = sanitizer.Sanitize(freel.TagLine)
	freel.Overview = sanitizer.Sanitize(freel.Overview)
}

func (freel *FreelancerOutput) Sanitize(sanitizer *bluemonday.Policy) {
	freel.Country = sanitizer.Sanitize(freel.Country)
	freel.City = sanitizer.Sanitize(freel.City)
	freel.Address = sanitizer.Sanitize(freel.Address)
	freel.Phone = sanitizer.Sanitize(freel.Phone)
	freel.TagLine = sanitizer.Sanitize(freel.TagLine)
	freel.Overview = sanitizer.Sanitize(freel.Overview)
}
