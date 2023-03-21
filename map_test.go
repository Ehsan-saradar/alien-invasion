package alieninvasion

import (
	"reflect"
	"testing"
)

func TestMap_MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		m       Map
		want    []byte
		wantErr bool
	}{
		{
			name:    "EmptyMap",
			m:       Map{},
			want:    nil,
			wantErr: false,
		},
		{
			name: "NormalMap",
			m: Map{
				cities: []*City{
					{
						name: "A",
						north: &City{
							name: "NA",
						},
						south: &City{
							name: "SA",
						},
						east: &City{
							name: "EA",
						},
						west: &City{
							name: "WA",
						},
					},
					{
						name: "B",
						north: &City{
							name: "NB",
						},
						south: &City{
							name: "SB",
						},
						east: &City{
							name: "EB",
						},
						west: &City{
							name: "WB",
						},
					},
				},
			},
			want:    []byte("A north=NA south=SA east=EA west=WA\nB north=NB south=SB east=EB west=WB"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("Map.MarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map.MarshalText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap_UnmarshalText(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name             string
		m                Map
		args             args
		wantErr          bool
		checkConnections bool
	}{
		{
			name: "EmptyMap",
			m: Map{
				cities: make([]*City, 0, 1),
			},
			args: args{
				text: []byte{},
			},
			wantErr:          false,
			checkConnections: false,
		},
		{
			name: "NormalMap",
			m: Map{
				cities: []*City{
					{
						name: "A",
						north: &City{
							name: "B",
						},
					},
					{
						name: "B",
						south: &City{
							name: "A",
						},
					},
				},
			},
			args: args{
				text: []byte("A north=B\nB south=A"),
			},
			wantErr:          false,
			checkConnections: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var m Map
			if err := m.UnmarshalText(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("Map.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.checkConnections {
				for _, city := range m.cities {
					if city.north != nil {
						if city.north.south != city {
							t.Errorf("Map.UnmarshalText() = %v, want %v", city.north.south, city)
						}
					}
					if city.south != nil {
						if city.south.north != city {
							t.Errorf("Map.UnmarshalText() = %v, want %v", city.south.north, city)
						}
					}
					if city.east != nil {
						if city.east.west != city {
							t.Errorf("Map.UnmarshalText() = %v, want %v", city.east.west, city)
						}
					}
					if city.west != nil {
						if city.west.east != city {
							t.Errorf("Map.UnmarshalText() = %v, want %v", city.west.east, city)
						}
					}
				}
			}
		})
	}
}
