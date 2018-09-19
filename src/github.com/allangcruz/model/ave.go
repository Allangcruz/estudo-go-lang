package model

// Galinha representa uma ave do tipo galinha
type Galinha interface {
	Cacareja() string
}

// Pato repreesenta o pato
type Pato interface {
	Quack() string
}

// Ave represemta um animal
type Ave struct {
	Nome string
}

// CAcareja retorna um som emitindo por uma galinha
func (ave Ave) Cacareja() string {
	return "Cócóricó ..."
}

// Quack retorna o som emitido por um pato
func (pato Ave) Quack() string {
	return "Quack, quack ..."
}
