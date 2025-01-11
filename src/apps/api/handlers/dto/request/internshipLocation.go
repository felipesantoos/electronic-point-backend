package request

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internshipLocation"
)

type InternshipLocation struct {
	Name         string  `json:"name" example:"Estágio Exemplo"`
	Number       string  `json:"number" example:"123"`
	Street       string  `json:"street" example:"Rua Exemplo"`
	Neighborhood string  `json:"neighborhood" example:"Bairro Exemplo"`
	City         string  `json:"city" example:"Cidade Exemplo"`
	ZipCode      string  `json:"zip_code" example:"57260-000"`
	Lat          float64 `json:"lat" example:"-23.5505"`
	Long         float64 `json:"long" example:"-46.6333"`
}

func (this *InternshipLocation) ToDomain() (internshipLocation.InternshipLocation, errors.Error) {
	_internshipLocation := internshipLocation.NewBuilder().
		WithName(this.Name).
		WithNumber(this.Number).
		WithStreet(this.Street).
		WithNeighborhood(this.Neighborhood).
		WithCity(this.City).
		WithZipCode(this.ZipCode).
		WithLat(this.Lat).
		WithLong(this.Long)
	return _internshipLocation.Build()
}
