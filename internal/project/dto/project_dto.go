package dto

import "time"

type StatusEnum string

const (
	StatusPending StatusEnum = "PENDING"

	StatusApproved StatusEnum = "APPROVED"
	StatusDenied   StatusEnum = "DENIED"

	StatusHalted   StatusEnum = "HALTED"
	StatusFinished StatusEnum = "FINISHED"

	StatusDeleted StatusEnum = "DELETED"
)

type CategoryEnum string

const (
	CategoryFood         CategoryEnum = "FOOD"
	CategoryHealth       CategoryEnum = "HEALTH"
	CategoryEducation    CategoryEnum = "EDUCATION"
	CategoryEnvironment  CategoryEnum = "ENVIRONMENT"
	CategoryReligion     CategoryEnum = "RELIGION"
	CategoryHumanitarian CategoryEnum = "HUMANTARIAN"
	CategoryHousing      CategoryEnum = "HOUSING"
	CategoryOther        CategoryEnum = "OTHER"
)

// Values provides list valid values for Enum.
func (CategoryEnum) Values() (kinds []string) {
	for _, s := range []CategoryEnum{
		CategoryFood,
		CategoryHealth,
		CategoryEducation,
		CategoryEnvironment,
		CategoryReligion,
		CategoryHousing,
		CategoryOther,
	} {
		kinds = append(kinds, string(s))
	}
	return
}

type CreateProjectRequestDto struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Goal        float64 `json:"goal"`

	StartDate int64 `json:"startDate"`
	EndDate   int64 `json:"endDate"`

	Category    string `json:"category"`
	CountryCode string `json:"countryCode"`
}

type ProjectResponseDto struct {
	ReadableId string `json:"readableId"`

	Name        string  `json:"name"`
	Description string  `json:"description"`
	Goal        float64 `json:"goal"`

	StartDate int64 `json:"startDate"`
	EndDate   int64 `json:"endDate"`

	Category    string `json:"category"`
	CountryCode string `json:"countryCode"`
}

type SaveProjectEntDto struct {
	Name        string
	Description string
	Goal        float64
	StartDate   time.Time
	EndDate     time.Time
	Category    string
	CountryCode string
}

type MessageResponseDto struct {
	Message string `json:"message"`
}

type ErrorResponseDto struct {
	Message    string `json:"message"`
	StatusCode uint   `json:"statusCode"`
}

func (d *CreateProjectRequestDto) ToSaveProjectEntDto() *SaveProjectEntDto {
	return &SaveProjectEntDto{
		Name:        d.Name,
		Description: d.Description,
		Goal:        d.Goal,
		StartDate:   time.UnixMilli(d.StartDate),
		EndDate:     time.UnixMilli(d.EndDate),
		Category:    d.Category,
		CountryCode: d.CountryCode,
	}
}
