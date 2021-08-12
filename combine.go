package errors

import "strings"

type combined []error

var (
	_ error                              = (*combined)(nil)
	_ interface{ Is(target error) bool } = (*combined)(nil)
)

func (cd combined) Error() string {
	var sb strings.Builder
	for i, err := range cd {
		sb.WriteString(err.Error())
		if i != len(cd)-1 {
			sb.WriteString(": ")
		}
	}
	return sb.String()
}

func (cd combined) Is(target error) bool {
	if t, ok := target.(combined); ok {
		if len(t) != len(cd) {
			return false
		}
		for _, err := range t {
			if !cd.Is(err) {
				return false
			}
		}
		return true
	}
	for _, err := range cd {
		if Is(err, target) {
			return true
		}
	}
	return false
}

func Combine(errs ...error) error {
	var eb Builder
	for _, err := range errs {
		eb.WithError(err)
	}
	return eb.Build()
}
