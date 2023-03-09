package dto

type CreateTeam struct {
	Name string `json:"name" validate:"required,alpha"`
}
