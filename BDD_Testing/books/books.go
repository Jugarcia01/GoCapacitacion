package main

type Books struct {
	Titulo  string
	Autor   string
	Paginas int
}

func (b *Books) CategoryByLength() string {

	if b.Paginas >= 300 {
		return "NOVELA"
	}

	return "SINTESIS"
}
