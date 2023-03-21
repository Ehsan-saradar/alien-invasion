package alieninvasion

import (
	"reflect"
	"testing"
)

func TestCity_MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		c       City
		want    []byte
		wantErr bool
	}{
		{
			name: "TrappedCity",
			c: City{
				name: "TrappedCity",
			},
			want:    []byte("TrappedCity"),
			wantErr: false,
		},
		{
			name: "CityWithNeighbors",
			c: City{
				name: "CityWithNeighbors",
				north: &City{
					name: "North",
				},
				south: &City{
					name: "South",
				},
				east: &City{
					name: "East",
				},
				west: &City{
					name: "West",
				},
			},
			want:    []byte("CityWithNeighbors north=North south=South east=East west=West"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("City.MarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("City.MarshalText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCity_UnmarshalText(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name    string
		c       City
		args    args
		wantErr bool
	}{
		{
			name: "EmptyString",
			c:    City{},
			args: args{
				text: []byte(""),
			},
			wantErr: true,
		},
		{
			name: "TrappedCity",
			c: City{
				name: "TrappedCity",
			},
			args: args{
				text: []byte("TrappedCity"),
			},
			wantErr: false,
		},
		{
			name: "CityWithNeighbors",
			c: City{
				name: "CityWithNeighbors",
				north: &City{
					name: "North",
				},
				south: &City{
					name: "South",
				},
				east: &City{
					name: "East",
				},
				west: &City{
					name: "West",
				},
			},
			args: args{
				text: []byte("CityWithNeighbors north=North south=South east=East west=West"),
			},
			wantErr: false,
		},
		{
			name: "InvalidDirectionFormat",
			c: City{
				name: "City",
			},
			args: args{
				text: []byte("City north North"),
			},
			wantErr: true,
		},
		{
			name: "UnknownDirection",
			c: City{
				name: "City",
			},
			args: args{
				text: []byte("City unknown=Unknown"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c City
			if err := c.UnmarshalText(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("City.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(c, tt.c) {
				t.Errorf("City.UnmarshalText() = %v, want %v", c, tt.c)
			}
		})
	}
}
