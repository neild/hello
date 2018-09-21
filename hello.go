package hello

type Language int

const (
	English = Language(iota)
	Japanese
)

func Hello(lang Language) string {
	switch lang {
	case Japanese:
		return "Konnichiwa, sekai!"
	default:
		return "Hello, world!"
	}
}
