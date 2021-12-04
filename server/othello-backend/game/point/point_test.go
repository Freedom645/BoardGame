package point

import (
	"reflect"
	"testing"
)

func TestNewPoint(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want Point
	}{
		{"[1, 1]確認", args{1, 1}, Point{1, 1}},
		{"[100, 0]確認", args{100, 0}, Point{100, 0}},
		{"[0, 100]確認", args{0, 100}, Point{0, 100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPoint(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZero(t *testing.T) {
	tests := []struct {
		name string
		want Point
	}{
		{"ゼロ確認", Point{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zero(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_X(t *testing.T) {
	tests := []struct {
		name string
		p    *Point
		want int
	}{
		{"X確認", &Point{0, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.X(); got != tt.want {
				t.Errorf("Point.X() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Y(t *testing.T) {
	tests := []struct {
		name string
		p    *Point
		want int
	}{
		{"Y確認", &Point{0, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Y(); got != tt.want {
				t.Errorf("Point.Y() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Plus(t *testing.T) {
	type args struct {
		v Point
	}
	tests := []struct {
		name string
		p    *Point
		args args
		want Point
	}{
		{"[0,0]+[0,0]", &Point{0, 0}, args{Point{0, 0}}, Point{0, 0}},
		{"[0,0]+[3,9]", &Point{0, 0}, args{Point{3, 9}}, Point{3, 9}},
		{"[3,9]+[3,9]", &Point{3, 9}, args{Point{3, 9}}, Point{6, 18}},
		{"[3,9]+[-3,-9]", &Point{3, 9}, args{Point{-3, -9}}, Point{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Plus(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point.Plus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Product(t *testing.T) {
	type args struct {
		s int
	}
	tests := []struct {
		name string
		p    *Point
		args args
		want Point
	}{
		{"[10,1]*0", &Point{10, 1}, args{0}, Point{0, 0}},
		{"[10,1]*5", &Point{10, 1}, args{5}, Point{50, 5}},
		{"[10,1]*-1", &Point{10, 1}, args{-1}, Point{-10, -1}},
		{"[10,1]*1", &Point{10, 1}, args{1}, Point{10, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Product(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point.Product() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Eq(t *testing.T) {
	type args struct {
		v Point
	}
	tests := []struct {
		name string
		p    *Point
		args args
		want bool
	}{
		{"[0,1]==[0,1]", &Point{0, 1}, args{Point{0, 1}}, true},
		{"[1,0]==[1,0]", &Point{1, 0}, args{Point{1, 0}}, true},
		{"[1,0]==[0,1]", &Point{1, 0}, args{Point{0, 1}}, false},
		{"[0,1]==[1,0]", &Point{0, 1}, args{Point{1, 0}}, false},
		{"[1,0]==[1,1]", &Point{1, 0}, args{Point{1, 1}}, false},
		{"[0,1]==[1,1]", &Point{0, 1}, args{Point{1, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Eq(tt.args.v); got != tt.want {
				t.Errorf("Point.Eq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_LessThan(t *testing.T) {
	type args struct {
		v Point
	}
	tests := []struct {
		name string
		p    *Point
		args args
		want bool
	}{
		{"[0,0]<[0,0] is false", &Point{0, 0}, args{Point{0, 0}}, false},
		{"[0,0]<[1,0] is false", &Point{0, 0}, args{Point{1, 0}}, false},
		{"[0,0]<[0,1] is false", &Point{0, 0}, args{Point{0, 1}}, false},
		{"[0,0]<[1,1] is true", &Point{0, 0}, args{Point{1, 1}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.LessThan(tt.args.v); got != tt.want {
				t.Errorf("Point.LessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Less(t *testing.T) {
	type args struct {
		v Point
	}
	tests := []struct {
		name string
		p    *Point
		args args
		want bool
	}{
		{"[0,0]<=[-1,-1] is false", &Point{0, 0}, args{Point{-1, -1}}, false},
		{"[0,0]<=[0,-1] is false", &Point{0, 0}, args{Point{0, -1}}, false},
		{"[0,0]<=[-1,0] is false", &Point{0, 0}, args{Point{-1, 0}}, false},
		{"[0,0]<=[0,0] is true", &Point{0, 0}, args{Point{0, 0}}, true},
		{"[0,0]<=[1,0] is true", &Point{0, 0}, args{Point{1, 0}}, true},
		{"[0,0]<=[0,1] is true", &Point{0, 0}, args{Point{0, 1}}, true},
		{"[0,0]<=[1,1] is true", &Point{0, 0}, args{Point{1, 1}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Less(tt.args.v); got != tt.want {
				t.Errorf("Point.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstDir8(t *testing.T) {
	tests := []struct {
		name string
		want []Point
	}{
		{"executable", []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {1, -1}, {-1, 1}, {1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConstDir8(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConstDir8() = %v, want %v", got, tt.want)
			}
		})
	}
}
