package stone_type

type StoneType int

const (
	None StoneType = iota
	Black
	White
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
