package structs

import (
	"reflect"
	"testing"
)

type Users struct {
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age,omitempty"`
	Parent *Users `json:"parent,omitempty"`
}

func TestDeepUpdateStruct(t *testing.T) {
	type args struct {
		b   interface{}
		u   interface{}
		out interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantOut interface{}
	}{
		{name: "update depth one", wantErr: false, args: args{
			b: &Users{
				Name: "Peter",
				Age:  10,
			},
			u: &Users{
				Name: "Peter",
				Age:  30,
			},
			out: &Users{},
		}, wantOut: &Users{
			Name: "Peter",
			Age:  30,
		}},
		{name: "update depth two", wantErr: false, args: args{
			b: &Users{
				Name: "Peter",
				Age:  10,
			},
			u: &Users{
				Name: "Sam",
				Age:  15,
				Parent: &Users{
					Name: "peter dad",
					Age:  30,
				},
			},
			out: &Users{},
		}, wantOut: &Users{
			Name: "Sam",
			Age:  15,
			Parent: &Users{
				Name: "peter dad",
				Age:  30,
			},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeepUpdateStruct(tt.args.b, tt.args.u, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("%v, get error: %v, wantErr: %v", tt.name, err, tt.wantErr)
			} else if !reflect.DeepEqual(tt.args.out, tt.wantOut) {
				t.Errorf("%v, out put not equal with want out.", tt.name)
			}
		})
	}
}
