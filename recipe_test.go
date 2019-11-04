package convert

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeRecipe(t *testing.T) {
	tests := []struct {
		name    string
		args    interface{}
		want    Recipe
		wantErr bool
	}{
		{
			"not a func",
			"Hello World",
			Recipe{},
			true,
		},
		{
			"invalid func (no args)",
			func() {},
			Recipe{},
			true,
		},
		{
			"invalid func (invalid first arg)",
			func(int, int, *int) {},
			Recipe{},
			true,
		},
		{
			"invalid func (output not a pointer)",
			func(Converter, int, int) {},
			Recipe{},
			true,
		},
		{
			"invalid func (return not error)",
			func(Converter, int, *int) int { return 0 },
			Recipe{},
			true,
		},
		{
			"invalid func (multiple return)",
			func(Converter, int, *int) (int, int) { return 0, 0 },
			Recipe{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MakeRecipe(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeRecipe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeRecipe() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustMakeRecipe(t *testing.T) {
	require.Panics(t, func() {
		_ = MustMakeRecipe(0)
	})
}
