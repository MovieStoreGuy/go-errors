package errors

type Builder []error

func (b *Builder) WithError(err error) {
	if err == nil {
		return
	}
	*b = append(*b, err)
}

func (b Builder) Build() error {
	switch len(b) {
	case 0:
		return error(nil)
	case 1:
		return b[0]
	}
	return combined(b)
}
