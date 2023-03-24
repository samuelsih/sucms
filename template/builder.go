package template

type Builder struct {
	name string
	datatype string
	skeleton []string
	hasErr error
}

func (b *Builder) SetName(name string) *Builder {
	b.name = name
	return b
}

func (b *Builder) SetType(datatype string) *Builder {
	b.datatype = datatype
	return b
}

func (b *Builder) Build() error {
	return nil
}
