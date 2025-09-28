package models

type Employees struct {
	ID    int    `db:"id" json:"id"`
	Age   int    `db:"age" json:"age"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}
