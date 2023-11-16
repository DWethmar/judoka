package assets

import (
	"image"
	"reflect"
	"testing"
)

func TestCreateCells(t *testing.T) {
	type args struct {
		columns int
		rows    int
		width   int
		height  int
	}
	tests := []struct {
		name string
		args args
		want [][]image.Rectangle
	}{
		{
			name: "create cells",
			args: args{
				columns: 6,
				rows:    1,
				width:   64,
				height:  64,
			},
			want: [][]image.Rectangle{
				{
					image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: 64, Y: 64}},
				},
				{
					image.Rectangle{Min: image.Point{X: 64, Y: 0}, Max: image.Point{X: 128, Y: 64}},
				},
				{
					image.Rectangle{Min: image.Point{X: 128, Y: 0}, Max: image.Point{X: 192, Y: 64}},
				},
				{
					image.Rectangle{Min: image.Point{X: 192, Y: 0}, Max: image.Point{X: 256, Y: 64}},
				},
				{
					image.Rectangle{Min: image.Point{X: 256, Y: 0}, Max: image.Point{X: 320, Y: 64}},
				},
				{
					image.Rectangle{Min: image.Point{X: 320, Y: 0}, Max: image.Point{X: 384, Y: 64}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateCells(tt.args.columns, tt.args.rows, tt.args.width, tt.args.height)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCells() = %v, want %v", got, tt.want)
			}
		})
	}
}

// util_test.go:45: CreateCells() =
// [[(0,0)-(64,64)] [(64,0)-(128,64)] [(128,0)-(192,64)] [(192,0)-(256,64)] [(256,0)-(320,64)] [(320,0)-(384,64)]], want
// [[(0,0)-(64,64) (64,0)-(128,64) (128,0)-(192,64) (192,0)-(256,64) (256,0)-(320,64) (320,0)-(384,64)]]
