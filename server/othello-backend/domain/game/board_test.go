package game

import (
	"reflect"
	"testing"

	st "github.com/Freedom645/BoardGame/domain/enum/stone_type"
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
		p Point
	}
	tests := []struct {
		name string
		b    *Board
		args args
		want bool
	}{
		{"top left is true", board, args{Zero()}, true},
		{"top right is true", board, args{board.TopRight()}, true},
		{"bottom left is true", board, args{NewPoint(0, HeightLim-1)}, true},
		{"bottom right is true", board, args{NewPoint(WidthLim-1, HeightLim-1)}, true},

		{"top left +up is false", board, args{Up()}, false},
		{"top left +left is false", board, args{Left()}, false},

		{"top right +up is false", board, args{Add(NewPoint(WidthLim-1, 0), Up())}, false},
		{"top right +right is false", board, args{Add(NewPoint(WidthLim-1, 0), Right())}, false},

		{"bottom left +down is false", board, args{Add(NewPoint(0, HeightLim-1), Down())}, false},
		{"bottom left +left is false", board, args{Add(NewPoint(0, HeightLim-1), Left())}, false},

		{"bottom right +down is false", board, args{Add(NewPoint(WidthLim-1, HeightLim-1), Down())}, false},
		{"bottom right +right is false", board, args{Add(NewPoint(WidthLim-1, HeightLim-1), Right())}, false},
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
		p Point
	}
	tests := []struct {
		name string
		b    *Board
		args args
		want st.StoneType
	}{
		{"top left is true", board, args{Zero()}, st.None},
		{"top right is true", board, args{NewPoint(WidthLim-1, 0)}, st.None},
		{"bottom left is true", board, args{NewPoint(0, HeightLim-1)}, st.None},
		{"bottom right is true", board, args{NewPoint(WidthLim-1, HeightLim-1)}, st.None},
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
	trAndbl.Put(&[]Point{trAndbl.TopRight(), trAndbl.BottomLeft()}, st.Black)

	type args struct {
		p     Point
		stone st.StoneType
	}
	tests := []struct {
		name string
		b    *Board
		args args
		want []Point
	}{
		{"put stone on [0,0] of empty board is noting", empty, args{Zero(), st.Black}, []Point{}},

		{"put stone on [2,0] of topLeft board is [1,0]", tl, args{NewPoint(2, 0), st.Black}, []Point{NewPoint(1, 0)}},
		{"put stone on [0,2] of topLeft board is [0,1]", tl, args{NewPoint(0, 2), st.Black}, []Point{NewPoint(0, 1)}},
		{"put stone on [2,2] of topLeft board is [1,1]", tl, args{NewPoint(2, 2), st.Black}, []Point{NewPoint(1, 1)}},

		{"put stone on TR +down*2 of TR and BL board is TR +down", trAndbl, args{Add(trAndbl.TopRight(), Down(), Down()), st.Black}, []Point{Add(trAndbl.TopRight(), Down())}},
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
		base   Point
		stone  st.StoneType
		vector Point
	}
	tests := []struct {
		name    string
		b       *Board
		args    args
		want    []Point
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
		list  *[]Point
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
		p     Point
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
