package model


type ExtendFreelancer struct {
	F *Freelancer `json:"freelancer"`
	FirstName string `json:"firstName"`
	SecondName string `json:"secondName"`
}