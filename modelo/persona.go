package modelo

type Comunidad struct {
	Nombre string `json:"nombre"`
}

//tipo de dato comunidades slice
type Comunidades []Comunidad

type Persona struct {
	Nombre      string      `json:"nombre"`
	Edad        uint8       `json:"edad"`
	Comunidades Comunidades `json:"comunidades"`
}

//slice de perona tipo de dato
type Personas []Persona
