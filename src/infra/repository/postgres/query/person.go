package query

const (
	PersonID        = "person_id"
	PersonName      = "person_name"
	PersonBirthDate = "person_birth_date"
	PersonEmail     = "person_email"
	PersonCPF       = "person_cpf"
	PersonPhone     = "person_phone"
)

type PersonQueryBuilder interface {
	Insert() string
}

type personQueryBuilder struct{}

func Person() PersonQueryBuilder {
	return &personQueryBuilder{}
}

func (*personQueryBuilder) Insert() string {
	return `
		INSERT INTO person (name, birth_date, email, cpf, phone)
		VALUES ($1, $2, $3, $4, $5) RETURNING id
	`
}
