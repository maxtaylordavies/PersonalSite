package design

type Theme = struct {
	Name       string
	Background string
	Color      string
}

func GetTheme(name string) Theme {
	themes := map[string]Theme{
		"teal": {
			Name:       "teal",
			Background: "rgba(152, 255, 212, 0.15)",
			Color:      "#50E0A4",
		},
		"purple": {
			Name:       "purple",
			Background: "rgba(145, 71, 255, 0.15)",
			Color:      "rgb(145, 71, 255)",
		},
	}

	if theme, ok := themes[name]; ok {
		return theme
	}

	return themes["teal"]
}
