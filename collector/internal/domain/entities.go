package domain

type Service struct {
    ID          int32
    Name        string
    Description string
}

type User struct {
    ID   int32 `db:"id"`
    Name string `db:"name"`
}

type Call struct {
    UserID    int32 `db:"user_id"`
    ServiceID int32 `db:"service_id"`
    Count     int32 `db:"count"`
}