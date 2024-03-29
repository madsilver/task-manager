package db

type DB interface {
	Query(query string, args any, fn func(scan func(dest ...any) error) error) error
	QueryRow(query string, args any, fn func(scan func(dest ...any) error) error) error
	Save(query string, args ...any) (any, error)
	Update(query string, args ...any) error
}
