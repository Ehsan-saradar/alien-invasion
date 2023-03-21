package alieninvasion

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestInvasion_Run(t *testing.T) {
	type fields struct {
		mapStr string
		n      int
		rnd    *rand.Rand
	}
	tests := []struct {
		name         string
		fields       fields
		outputMapStr string
	}{
		{
			name: "1 Aliens",
			fields: fields{
				mapStr: "Foo north=Bar south=Qu-ux west=Baz\nBar south=Foo west=Bee\nBaz east=Foo\nQu-ux north=Foo\nBee east=Bar",
				n:      1,
				rnd:    rand.New(rand.NewSource(1)),
			},
			outputMapStr: "Foo north=Bar south=Qu-ux west=Baz\nBar south=Foo west=Bee\nBaz east=Foo\nQu-ux north=Foo\nBee east=Bar",
		},
		{
			name: "3 Aliens",
			fields: fields{
				mapStr: "Foo north=Bar south=Qu-ux west=Baz\nBar south=Foo west=Bee\nBaz east=Foo\nQu-ux north=Foo\nBee east=Bar",
				n:      3,
				rnd:    rand.New(rand.NewSource(3)),
			},
			outputMapStr: "Bar west=Bee\nBaz\nQu-ux\nBee east=Bar",
		},
		{
			name: "10 Aliens",
			fields: fields{
				mapStr: "Foo north=Bar south=Qu-ux west=Baz\nBar south=Foo west=Bee\nBaz east=Foo\nQu-ux north=Foo\nBee east=Bar",
				n:      10,
				rnd:    rand.New(rand.NewSource(1)),
			},
			outputMapStr: "Qu-ux\nBee",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var m Map
			err := m.UnmarshalText([]byte(tt.fields.mapStr))
			if err != nil {
				t.Errorf("Map.UnmarshalText() error = %v", err)
			}

			inv := NewInvasion(&m, tt.fields.n, tt.fields.rnd)
			inv.Run()

			if !reflect.DeepEqual(inv.Map.String(), tt.outputMapStr) {
				t.Errorf("Invasion.Run() = %v, want %v", inv.Map.String(), tt.outputMapStr)
			}
		})
	}
}
