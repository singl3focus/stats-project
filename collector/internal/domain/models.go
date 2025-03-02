package domain

type CallsFilter struct {
    UserID    *int
    ServiceID *int
    Page      *int
    PerPage   *int
    Sort      string
}
