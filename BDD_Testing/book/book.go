package book

type Book struct {
	Titulo  string
	Autor   string
	Paginas int
}

func (b *Book) CategoryByLength() string {

	if b.Paginas >= 300 {
		return "NOVELA"
	}

	return "SINTESIS"
}
