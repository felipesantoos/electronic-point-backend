package request

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internshipLocation"
)

type InternshipLocation struct {
	Name    string   `json:"name" example:"Estágio Exemplo"`
	Address string   `json:"address" example:"Rua Exemplo, 123"`
	City    string   `json:"city" example:"Cidade Exemplo"`
	Lat     *float64 `json:"lat" example:"-23.5505"`
	Long    *float64 `json:"long" example:"-46.6333"`
}

func (this *InternshipLocation) ToDomain() (internshipLocation.InternshipLocation, errors.Error) {
	_internshipLocation := internshipLocation.NewBuilder().
		WithName(this.Name).
		WithAddress(this.Address).
		WithCity(this.City).
		WithLat(this.Lat).
		WithLong(this.Long)
	return _internshipLocation.Build()
}
