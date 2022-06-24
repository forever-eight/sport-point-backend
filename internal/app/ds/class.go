package ds

import "fmt"

type ClassInput struct {
	StudioID    int32  `json:"studio_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TypeID      int32  `json:"type_id"`
	Weekday     int32  `json:"weekday"`
	StartedAt   int32  `json:"started_at"`
	Duration    int32  `json:"duration"`
	Amount      int32  `json:"amount"`
}

func (i *ClassInput) Validate() error {
	if i == nil {
		return fmt.Errorf("empty input")
	}
	if i.StudioID <= 0 {
		return fmt.Errorf("invalid studio ID")
	}
	if i.Title == "" {
		return fmt.Errorf("invalid title")
	}
	if i.Description == "" {
		return fmt.Errorf("invalid title")
	}
	if i.TypeID <= 0 {
		return fmt.Errorf("invalid type ID")
	}
	if i.Weekday <= 0 || i.Weekday > 7 {
		return fmt.Errorf("invalid weekday")
	}
	if i.StartedAt >= 24*60 || i.StartedAt <= 0 {
		return fmt.Errorf("invalid started at")
	}
	if i.Duration <= 0 {
		return fmt.Errorf("invalid duration")
	}
	if i.Amount <= 0 {
		return fmt.Errorf("invalid amount")
	}

	return nil
}

func (i *ClassInput) ToDB() *Class {
	return &Class{
		StudioID:    i.StudioID,
		Title:       i.Title,
		Description: i.Description,
		TypeID:      i.TypeID,
		Weekday:     i.Weekday,
		StartedAt:   i.StartedAt,
		Duration:    i.Duration,
		Amount:      i.Amount,
	}
}

type Class struct {
	StudioID    int32  `db:"studio_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	TypeID      int32  `db:"type_id"`
	Weekday     int32  `db:"weekday"`
	StartedAt   int32  `db:"started_at"`
	Duration    int32  `db:"duration"`
	Amount      int32  `db:"amount"`
}

func (i *Class) ToOutput(ID uint32, Studio string, Type string) *ClassOutput {
	return &ClassOutput{
		ID:          ID,
		Studio:      Studio,
		Title:       i.Title,
		Description: i.Description,
		Type:        Type,
		Weekday:     i.Weekday,
		StartedAt:   i.StartedAt,
		Duration:    i.Duration,
		Amount:      i.Amount,
	}
}

type ClassOutput struct {
	ID          uint32 `json:"id"`
	Studio      string `json:"studio"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Weekday     int32  `json:"weekday"`
	StartedAt   int32  `json:"started_at"`
	Duration    int32  `json:"duration"`
	Amount      int32  `json:"amount"`
}

type Studio struct {
	ID      uint32 `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Address string `json:"address" db:"address"`
}
