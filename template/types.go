package template

import "fmt"

func bigIncrements(s string) string {
	return fmt.Sprintf("bigIncrements('%s')", s)
}

func bigInteger(s string) string {
	return fmt.Sprintf("bigInteger('%s')", s)
}

func date(s string) string {
	return fmt.Sprintf("date('%s')", s)
}

func uuid(s string) string {
	return fmt.Sprintf("uuid('%s')", s)
}

func str(s string) string {
	return fmt.Sprintf("string('%s')", s)
}

func id(s string) string {
	return fmt.Sprintf("id('%s')", s)
}

func text(s string) string {
	return fmt.Sprintf("text('%s)", s)
}

func integer(s string) string {
	return fmt.Sprintf("integer('%s')", s)
}

func boolean(s string) string {
	return fmt.Sprintf("boolean('%s')", s)
}

