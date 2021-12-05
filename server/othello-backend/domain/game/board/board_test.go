package board

import (
	"reflect"
	"testing"

	st "github.com/Freedom645/BoardGame/domain/enum/stone_type"
	point "github.com/Freedom645/BoardGame/domain/game/point"
)

func TestNewBoard(t *testing.T) {
	tests := []struct {
		name string
		want *Board
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoard(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_Height(t *testing.T) {
	tests := []struct {
		name string
		b    *Board
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Height(); got != tt.want {
				t.Errorf("Board.Height() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_Width(t *testing.T) {
	tests := []struct {
		name string
		b    *Board
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Width(); got != tt.want {
				t.Errorf("Board.Width() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_Contain(t *testing.T) {
	board := NewBoard()
	type args struct {
		p point.Point
	}
	tests := []struct {
		name string
		b    *Board
		args args
		want bool
	}{
		{"top left is true", board, args{point.Zero()}, true},
		{"top right is true", board, args{board.TopRight()}, true},
		{"bottom left is true", board, args{point.NewPoint(0, HeightLim-1)}, true},
		{"bottom right is true", board, args{point.NewPoint(WidthLim-1, HeightLim-1)}, true},

		{"top left +up is false", board, args{point.Up()}, false},
		{"top left +left is false", board, args{point.Left()}, false},

		{"top right +up is false", board, args{point.Add(point.NewPoint(WidthLim-1, 0), point.Up())}, false},
		{"top right +right is false", board, args{point.Add(point.NewPoint(WidthLim-1, 0), point.Right())}, false},

		{"bottom left +down is false", board, args{point.Add(point.NewPoint(0, HeightLim-1), point.Down())}, false},
		{"bottom left +left is false", board, args{point.Add(point.NewPoint(0, HeightLim-1), point.Left())}, false},

		{"bottom right +down is false", board, args{point.Add(point.NewPoint(WidthLim-1, HeightLim-1), point.Down())}, false},
		{"bottom right +right is false", board, args{point.Add(point.NewPoint(WidthLim-1, HeightLim-1), point.Right())}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Contain(tt.args.p); got != tt.want {
				t.Errorf("Board.Contain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_TypeAt(t *testing.T) {
	board := NewBoard()
	type args struct {
		p point.Point
	}
	tests := []struct {
		name string
		b    *Board
		args args
		want st.StoneType
	}{
		{"top left is true", board, args{point.Zero()}, st.None},
		{"top right is true", board, args{point.NewPoint(WidthLim-1, 0)}, st.None},
		{"bottom left is true", board, args{point.NewPoint(0, HeightLim-1)}, st.None},
		{"bottom right is true", board, args{point.NewPoint(WidthLim-1, HeightLim-1)}, st.None},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.TypeAt(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.TypeAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_PutIf(t *testing.T) {
	empty := NewBoard()
	tl := NewBoard().Fill(st.White)
	tl.PutOne(tl.TopLeft(), st.Black)

	// 右上/左下
	trAndbl := NewBoard().Fill(st.White)
	trAndbl.Put(&[]point.Point{trAndbl.TopRight(), trAndbl.BottomLeft()}, st.Black)

	type args struct {
		p     point.Point
		stone st.StoneType
	}
	tests := []struct {
		name string
		b    *Board
		args args
		want []point.Point
	}{
		{"put stone on [0,0] of empty board is noting", empty, args{point.Zero(), st.Black}, []point.Point{}},

		{"put stone on [2,0] of topLeft board is [1,0]", tl, args{point.NewPoint(2, 0), st.Black}, []point.Point{point.NewPoint(1, 0)}},
		{"put stone on [0,2] of topLeft board is [0,1]", tl, args{point.NewPoint(0, 2), st.Black}, []point.Point{point.NewPoint(0, 1)}},
		{"put stone on [2,2] of topLeft board is [1,1]", tl, args{point.NewPoint(2, 2), st.Black}, []point.Point{point.NewPoint(1, 1)}},

		{"put stone on TR +down*2 of TR and BL board is TR +down", trAndbl, args{point.Add(trAndbl.TopRight(), point.Down(), point.Down()), st.Black}, []point.Point{point.Add(trAndbl.TopRight(), point.Down())}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.PutIf(tt.args.p, tt.args.stone); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.PutIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_searchReversePoint(t *testing.T) {
	type args struct {
		base   point.Point
		stone  st.StoneType
		vector point.Point
	}
	tests := []struct {
		name    string
		b       *Board
		args    args
		want    []point.Point
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.b.searchReversePoint(tt.args.base, tt.args.stone, tt.args.vector)
			if (err != nil) != tt.wantErr {
				t.Errorf("Board.searchReversePoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Board.searchReversePoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_Put(t *testing.T) {
	type args struct {
		list  *[]point.Point
		stone st.StoneType
	}
	tests := []struct {
		name string
		b    *Board
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Put(tt.args.list, tt.args.stone); got != tt.want {
				t.Errorf("Board.Put() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard_PutOne(t *testing.T) {
	type args struct {
		p     point.Point
		stone st.StoneType
	}
	tests := []struct {
		name string
		b    *Board
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.PutOne(tt.args.p, tt.args.stone); got != tt.want {
				t.Errorf("Board.PutOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
