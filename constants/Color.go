package constants

// Color enumeration
type Color int

const (
	White Color = iota
	Black
	Blue
	Green
	Red
	Yellow
	Pink
	Violet
	Orange
)

var s2c = map[string]Color{
	"White":  White,
	"Black":  Black,
	"Blue":   Blue,
	"Green":  Green,
	"Red":    Red,
	"Yellow": Yellow,
	"Pink":   Pink,
	"Violet": Violet,
	"Orange": Orange,
}

var c2s = map[Color]string{
	White:  "White",
	Black:  "Black",
	Blue:   "Blue",
	Green:  "Green",
	Red:    "Red",
	Yellow: "Yellow",
	Pink:   "Pink",
	Violet: "Violet",
	Orange: "Orange",
}

func StringToColor(color string) Color {
	return s2c[color]
}

func (c Color) String() string {
	return c2s[c]
}
