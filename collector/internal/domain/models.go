package domain

type CallsFilter struct {
    UserID    *int `validate:"omitempty,min=1"`
    ServiceID *int `validate:"omitempty,min=1"`
    Page      *int `validate:"omitempty,min=1"`
    PerPage   *int `validate:"omitempty,min=1"`
    Sort      string `validate:"required,oneof=ASC DESC"`
}
