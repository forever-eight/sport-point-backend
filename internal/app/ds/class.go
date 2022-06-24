package ds

type ClassInput struct {
	StudioID    int32  `json:"studio_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TypeID      int32  `json:"type_id"`
	Weekday     int32  `json:"weekday"`
	StartedAt   int32  `json:"started_at"`
	Duration    int32  `json:"duration"`
	Amount      int32  `json:"amount"`
	Currency    int32  `json:"currency"`
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
		Currency:    i.Currency,
	}
}

type Class struct {
	StudioID    int32  `db:"studio_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	TypeID      int32  `json:"type_id"`
	Weekday     int32  `db:"weekday"`
	StartedAt   int32  `db:"started_at"`
	Duration    int32  `db:"duration"`
	Amount      int32  `db:"amount"`
	Currency    int32  `db:"currency"`
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
		Currency:    i.Currency,
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
	Currency    int32  `json:"currency"`
}
