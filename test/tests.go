package test

// import (
// 	"reflect"
// 	"testing"
// 	"time"

// 	. "github.com/OZahed/reflectutil"
// )

// func TestTypeCast(t *testing.T) {
// 	type args struct {
// 		src interface{}
// 		dst interface{}
// 	}
// 	tests := []struct {
// 		args    args
// 		name    string
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := TypeCast(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
// 				t.Errorf("TypeCast() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// type Enum string

// func TestScanValue(t *testing.T) {

// 	t.Run("Enum reading", func(t *testing.T) {
// 		defaultValue := reflect.ValueOf("Hello World")
// 		wantValue := Enum("Hello World")
// 		gotRes, err := ScanValue[Enum](defaultValue)
// 		if (err != nil) != false {
// 			t.Errorf("ScanValue() error = %v, wantErr %v", err, false)
// 			return
// 		}
// 		if !reflect.DeepEqual(gotRes, wantValue) {
// 			t.Errorf("ScanValue() = %#v, want %#v", gotRes, wantValue)
// 		}
// 	})

// 	t.Run("reading int", func(t *testing.T) {
// 		now := time.Now()
// 		defaultValue := reflect.ValueOf(now)
// 		wantValue := now
// 		gotRes, err := ScanValue[time.Time](defaultValue)
// 		if (err != nil) != false {
// 			t.Errorf("ScanValue() error = %v, wantErr %v", err, false)
// 			return
// 		}
// 		if !reflect.DeepEqual(gotRes, wantValue) {
// 			t.Errorf("ScanValue() = %#v, want %#v", gotRes, wantValue)
// 		}
// 	})

// 	t.Run("error in reading uint634", func(t *testing.T) {
// 		now := time.Now()
// 		defaultValue := reflect.ValueOf(now)
// 		wantValue := uint64(0)
// 		gotRes, err := ScanValue[uint64](defaultValue)
// 		if (err != nil) != true {
// 			t.Errorf("ScanValue() error = %v, wantErr %v", err, true)
// 			return
// 		}
// 		if !reflect.DeepEqual(gotRes, wantValue) {
// 			t.Errorf("ScanValue() = %#v, want %#v", gotRes, wantValue)
// 		}
// 	})

// }
