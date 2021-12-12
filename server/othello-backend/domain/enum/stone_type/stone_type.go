package stone_type

type StoneType string

const (
	None  StoneType = "none"
	Black StoneType = "black"
	White StoneType = "white"
)

func RevStone(s StoneType) StoneType {
	if s == Black {
		return White
	}
	if s == White {
		return Black
	}
	return None
}
