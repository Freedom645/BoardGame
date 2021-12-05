package board

import (
	"errors"
	"fmt"

	st "github.com/Freedom645/BoardGame/domain/enum/stone_type"
	"github.com/Freedom645/BoardGame/domain/game/point"
)

const (
	WidthLim  = 8
	HeightLim = 8
)

type Board struct {
	size   point.Point
	stones [][]st.StoneType
}

func NewBoard() *Board {
	b := new(Board)

	b.size = point.NewPoint(WidthLim, HeightLim)

	b.stones = make([][]st.StoneType, b.Height())
	for i := 0; i < b.Height(); i++ {
		b.stones[i] = make([]st.StoneType, b.Width())
	}

	return b
}

/* 盤面の縦幅を取得 */
func (b *Board) Height() int {
	return b.size.Y()
}

/* 盤面の横幅を取得 */
func (b *Board) Width() int {
	return b.size.X()
}

/* 左上 */
func (b *Board) TopLeft() point.Point {
	return point.Zero()
}

/* 右上 */
func (b *Board) TopRight() point.Point {
	return point.NewPoint(WidthLim-1, 0)
}

/* 左下 */
func (b *Board) BottomLeft() point.Point {
	return point.NewPoint(0, HeightLim-1)
}

/* 右下 */
func (b *Board) BottomRight() point.Point {
	return point.NewPoint(WidthLim-1, HeightLim-1)
}

/* 盤面内判定 */
func (b *Board) Contain(p point.Point) bool {
	zero := point.Zero()
	return zero.Less(p) && p.LessThan(b.size)
}

/* 盤面情報を参照 */
func (b *Board) TypeAt(p point.Point) st.StoneType {
	return b.stones[p.Y()][p.X()]
}

/* 反転する護石の座標を取得 */
func (b *Board) PutIf(p point.Point, stone st.StoneType) []point.Point {
	DIR := point.ConstDir8()

	if !b.Contain(p) {
		panic(fmt.Sprintf("out range [%d, %d]", p.X(), p.Y()))
	}

	var res = make([]point.Point, 0)
	for _, v := range DIR {
		list, err := b.searchReversePoint(p, stone, v)
		if err != nil {
			continue
		}

		res = append(res, list...)
	}

	return res
}

/* 一方向への返せる石の座標を探索 */
func (b *Board) searchReversePoint(base point.Point, stone st.StoneType, vector point.Point) ([]point.Point, error) {
	var res []point.Point
	for i := 1; ; i++ {
		p := base.Plus(vector.Product(i))
		if !b.Contain(p) {
			break
		}
		if b.TypeAt(p) == st.None {
			break
		}
		if b.TypeAt(p) == stone {
			// 同色があれば中断
			return res, nil
		}
		res = append(res, p)
	}

	return nil, errors.New("not found pair stone")
}

/* リストの座標に従い石を置く */
func (b *Board) Put(list *[]point.Point, stone st.StoneType) int {
	var res int = 0
	for _, p := range *list {
		if b.PutOne(p, stone) {
			res++
		}
	}

	return res
}

/* 1箇所に石を置く */
func (b *Board) PutOne(p point.Point, stone st.StoneType) bool {
	res := b.TypeAt(p) == st.RevStone(stone)
	b.stones[p.Y()][p.X()] = stone
	return res
}

/* 盤面を埋める */
func (b *Board) Fill(stone st.StoneType) *Board {
	for i := 0; i < b.Height(); i++ {
		for j := 0; j < b.Width(); j++ {
			b.stones[i][j] = stone
		}
	}
	return b
}
