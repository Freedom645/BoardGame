package game

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) Point {
	return Point{x, y}
}

func Zero() Point {
	return Point{0, 0}
}

func Up() Point {
	return Point{0, -1}
}

func Down() Point {
	return Point{0, 1}
}

func Left() Point {
	return Point{-1, 0}
}

func Right() Point {
	return Point{1, 0}
}

func (p *Point) X() int {
	return p.x
}

func (p *Point) Y() int {
	return p.y
}

/* 和 */
func Add(args ...Point) Point {
	var res = Point{0, 0}
	for _, p := range args {
		res.x += p.x
		res.y += p.y
	}
	return res
}

/* 和 */
func (p *Point) Plus(v Point) Point {
	return Point{p.x + v.x, p.y + v.y}
}

/* 積 */
func (p *Point) Product(s int) Point {
	return Point{p.x * s, p.y * s}
}

/* 等価: = */
func (p *Point) Eq(v Point) bool {
	return p.x == v.x && p.y == v.y
}

/* 未満: < */
func (p *Point) LessThan(v Point) bool {
	return p.x < v.x && p.y < v.y
}

/* 以下: <= */
func (p *Point) Less(v Point) bool {
	return p.x <= v.x && p.y <= v.y
}

/* 8近傍 */
func ConstDir8() []Point {
	return []Point{Up(), Down(), Left(), Right(), {-1, -1}, {1, -1}, {-1, 1}, {1, 1}}
}
