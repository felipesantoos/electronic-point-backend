package query

type ProfessionalQueryBuilder interface {
	Insert() string
}

type professionalQueryBuilder struct{}

func Professional() ProfessionalQueryBuilder {
	return &professionalQueryBuilder{}
}

func (*professionalQueryBuilder) Insert() string {
	return `
		INSERT INTO professional (person_id)
		VALUES ($1);
	`
}
