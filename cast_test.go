package typecast

import "testing"

func TestTypeCast(t *testing.T) {
	type args struct {
		src interface{}
		dst interface{}
	}
	tests := []struct {
		args    args
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := TypeCast(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("TypeCast() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
