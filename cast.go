/*
reflectutil is a package to cast Type Values fot data types that are not really straight forward to cast like
Structs, Maps, Slices of structs and ...
*/
package reflectutil

import (
	"fmt"
	"reflect"
)

type CastMap map[reflect.Type]reflect.Value

/*
	Caster type is here to help casting types into possible types for a datatype

example:

In Repository package Storage.go:

	type Storage struct {
		storageName  string
		storageSize  int
	}

	func (s Storage) CastTo() CastFuncMap {
		funcsMap := make(CastFuncMap)
		funcsMap[reflect.TypeOf("")] =reflect.ValueOf(s.storageName)
		funcsMap[reflect.TypeOf(0)] = reflect.ValueOf(s.storageSize)

		return funcMap
	}
*/
type Caster interface {
	// CasteTo defines a map of possible types that a specific type can be casted into
	// If the destination type was not provided the default type cast procedure would be executed
	CastTo() CastMap
}

/*
Castee interface is used for when you want to make control the casting processes on destination value

example:

	type Slug struct {
		name  string
	}

	func(s *Storage)ScanValue(value reflect.Value) error {
		return reflectutil.Scan[string](value, &s.storageName)
	}
*/
type Scanner interface {
	// ScanValue function, can change reflect rules, let's say Slug field in destination type is an struct like above
	// but in source type is an string, using ScanValue function helps with bypassing the default reflection behaviors
	ScanValue(value reflect.Value) error
}

// TypeCast is a simple implementation for casting one type into another similar data type
func TypeCast(src, dst interface{}) error {
	if src == nil || dst == nil {
		return NewError(reflect.ValueOf(src), reflect.ValueOf(dst), "source or destination is nil")
	}

	dstType := reflect.TypeOf(dst)
	dstValue := reflect.ValueOf(dst)
	srcValue := reflect.ValueOf(src)
	srcType := reflect.TypeOf(src)

	if dstType.Kind() != reflect.Pointer {
		return NewError(srcValue, dstValue, "destination is not a pointer and can not be assigned")
	}

	if srcType.Kind() == reflect.Pointer {
		srcType = srcType.Elem()
		srcValue = srcValue.Elem()
	}

	if dstValue.CanSet() {
		return NewError(srcValue, dstValue, "destination value can not be set")
	}

	if !srcValue.IsValid() {
		return NewError(srcValue, dstValue, "source does not represent a value")
	}

	if dstType.Kind() == reflect.Pointer {
		dstValue = dstValue.Elem()
		dstType = dstType.Elem()
	}

	// caster value should be checked before Kind check, cause some kinds in caster ma do not match
	if srcType.AssignableTo(reflect.TypeOf((Caster)(nil))) {
		return callCaster(srcValue, dstValue)
	}

	if dstType.AssignableTo(reflect.TypeOf((Scanner)(nil))) {
		return callScanner(srcValue, dstValue)
	}

	if srcType.Kind() != dstType.Kind() {
		return NewError(srcValue, dstValue, "source and destination are not of same reflect kind")
	}

	if srcType == dstType {
		dstValue.Set(srcValue)

		return nil
	}

	switch {
	case srcType.Kind() == reflect.Slice:
		return CopySlice(srcValue, dstValue)
	case srcType.Kind() == reflect.Map:
		return CopyMap(srcValue, dstValue)
	case srcType.Kind() == reflect.Interface:
		return CopyInterface(srcValue, dstValue)
	case srcType.Kind() == reflect.Func:
		return NewError(srcValue, dstValue, "casting function types is not supported")
	// we should check type equal before checking struct
	case dstType == srcType:
		return copyValue(srcValue, dstValue)
	case srcType.Kind() == reflect.Struct:
		return CopyStruct(srcValue, dstValue)
	}

	return nil
}

func CopyMap(srcValue, dstValue reflect.Value) error {
	srcKType := srcValue.Type().Key()
	srcVType := srcValue.Type().Elem()

	dstKType := srcValue.Type().Key()
	dstVType := srcValue.Type().Elem()

	if srcKType != dstKType || srcVType != dstVType {
		return NewError(srcValue, dstValue, "key value types on maps are not compatible")
	}

	if dstValue.IsZero() {
		dstValue.Set(reflect.MakeMap(dstValue.Type()))
	}

	for _, k := range srcValue.MapKeys() {
		dstValue.SetMapIndex(k, srcValue.MapIndex(k))
	}

	return nil

}

func CopyInterface(srcValue, dstValue reflect.Value) error {
	if !srcValue.Type().AssignableTo(dstValue.Type()) {
		return NewError(srcValue, dstValue, "can not cast interfaces")
	}

	dstValue.Set(reflect.ValueOf(srcValue.Interface()))

	return nil
}

func CopySlice(srcSlice, dstSlice reflect.Value) (err error) {
	srcElm := srcSlice.Type().Elem()
	dstElm := dstSlice.Type().Elem()

	if srcSlice.Len() > dstSlice.Len() {
		dstSlice.Grow(srcSlice.Len() - dstSlice.Len())
	}

	dstSlice.SetLen(srcSlice.Len())

	for i := 0; i < srcSlice.Len(); i++ {
		if srcSlice.Index(i).Type().AssignableTo(reflect.TypeOf((Caster)(nil))) {
			err = callCaster(srcSlice.Index(i), dstSlice.Index(i))
			if err != nil {
				return err
			}

			continue
		}

		if dstSlice.Index(i).Type().AssignableTo(reflect.TypeOf((Scanner)(nil))) {
			err = callScanner(srcSlice.Index(i), dstSlice.Index(i))
			if err != nil {
				return err
			}
			continue
		}

		switch {
		case srcElm.Kind() == reflect.Interface:
			if err := CopyInterface(srcSlice.Index(i), dstSlice.Index(i)); err != nil {
				return err
			}
		case srcElm.Kind() == reflect.Map:
			if err := CopyMap(srcSlice.Index(i), dstSlice.Index(i)); err != nil {
				return err
			}
		case srcElm.Kind() == reflect.Slice:
			if err := CopySlice(srcSlice.Index(i), dstSlice.Index(i)); err != nil {
				return err
			}
		// For built-in or std lib types or common types
		case dstElm == srcElm:
			if err := copyValue(srcSlice, dstSlice); err != nil {
				return err
			}
		case srcElm.Kind() == reflect.Struct:
			if err := CopyStruct(srcSlice.Index(i), dstSlice.Index(i)); err != nil {
				return err
			}
		}
	}

	return nil
}

func CopyStruct(srcValue, dstValue reflect.Value) error {
	var srcStructField reflect.StructField
	var srcFieldValue reflect.Value

	dstType := dstValue.Type()

	for i := 0; i < srcValue.NumField(); i++ {
		srcFieldValue = srcValue.Field(i)
		srcStructField = srcValue.Type().Field(i)

		if !srcStructField.IsExported() {
			continue
		}

		dstStructField, ok := dstType.FieldByName(srcStructField.Name)
		if !ok {
			continue
		}

		dstFieldValue := dstValue.FieldByName(srcStructField.Name)

		// caster value should be checked before Kind check, cause some kinds in caster ma do not match
		if srcFieldValue.Type().AssignableTo(reflect.TypeOf((Caster)(nil))) {
			err := callCaster(srcFieldValue, dstFieldValue)
			if err != nil {
				return err
			}

			continue
		}

		if dstFieldValue.Type().AssignableTo(reflect.TypeOf((Scanner)(nil))) {
			err := callScanner(srcFieldValue, dstFieldValue)
			if err != nil {
				return err
			}

			continue
		}

		if srcStructField.Type.Kind() != dstStructField.Type.Kind() {
			return NewError(srcFieldValue, dstFieldValue, "are not the same kind")
		}

		// For Almost any data type on standard Library, it would work
		// Unless there is a type alias or Enum or an struct

		switch {
		case srcStructField.Type.Kind() == reflect.Slice, srcStructField.Type.Kind() == reflect.Array:
			if err := CopySlice(srcFieldValue, dstFieldValue); err != nil {
				return err
			}
		case srcStructField.Type.Kind() == reflect.Map:
			if err := CopyMap(srcFieldValue, dstFieldValue); err != nil {
				return err
			}
		case srcStructField.Type.Kind() == reflect.Interface:
			if err := CopyInterface(srcFieldValue, dstFieldValue); err != nil {
				return err
			}
		case srcStructField.Type == dstStructField.Type:
			if err := copyValue(srcFieldValue, dstFieldValue); err != nil {
				return err
			}
		case srcStructField.Type.Kind() == reflect.Struct:
			if err := CopyStruct(srcFieldValue, dstFieldValue); err != nil {
				return err
			}
		case srcFieldValue.Type().Kind() == dstFieldValue.Type().Kind():
			if err := copyValue(srcFieldValue, dstFieldValue); err != nil {
				return err
			}
		default:
			return NewError(srcFieldValue, dstFieldValue, "types are incompatible")
		}
	}

	return nil
}

func copyValue(srcFieldValue, dstFieldValue reflect.Value) error {
	switch {
	// these switch types are required to handle type aliases like Enums
	case srcFieldValue.Type().Kind() == reflect.Pointer:
		// Preventing direct copy of pointer
		return copyValue(srcFieldValue.Elem(), dstFieldValue.Elem())
	case srcFieldValue.Type() == dstFieldValue.Type():
		dstFieldValue.Set(srcFieldValue)
	case srcFieldValue.Type().Kind() == dstFieldValue.Type().Kind():
		switch dstFieldValue.Type().Kind() {
		case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
			dstFieldValue.SetInt(srcFieldValue.Int())
		case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
			dstFieldValue.SetUint(srcFieldValue.Uint())
		case reflect.Float32, reflect.Float64:
			dstFieldValue.SetFloat(srcFieldValue.Float())
		case reflect.String:
			dstFieldValue.SetString(srcFieldValue.String())
		}
	default:
		return NewError(srcFieldValue, dstFieldValue, "could not copy value")
	}

	return nil
}

func callScanner(srcValue, dstValue reflect.Value) error {
	meth := dstValue.MethodByName("ScanValue")
	if !meth.IsValid() {
		return NewError(srcValue, dstValue, "destination does not implement Scanner interface")
	}

	res := meth.Call([]reflect.Value{srcValue})

	if len(res) != 1 {
		return NewError(srcValue, dstValue, "destination did not properly implement ScanValue method")
	}

	err, ok := res[0].Interface().(error)

	if !ok {
		return NewError(srcValue, dstValue, "destination did not properly implement ScanValue method")
	}

	return err
}

// if callCaster is called, the type is already a Caster
func callCaster(srcValue, dstValue reflect.Value) error {
	meth := srcValue.MethodByName("CastTo")
	if !meth.IsValid() {
		return NewError(srcValue, dstValue, "source does not implement Caster interface")
	}

	res := meth.Call([]reflect.Value{})

	if len(res) != 1 {
		return NewError(srcValue, dstValue, "source does not implement Caster interface")
	}

	castMap, ok := res[0].Interface().(CastMap)
	if !ok {
		return NewError(
			srcValue,
			dstValue, "expected the source to be of reflectutil.Caster type",
		)
	}

	if castMap == nil {
		return nil
	}

	val, ok := castMap[dstValue.Type()]
	if !ok {
		return nil
	}

	dstValue.Set(val)

	return nil
}

// MakePrserFrom Usage:
//
//	type MyType struct {
//	    slug string
//	}
func ScanValue[T any](srcValue reflect.Value) (res T, err error) {
	// requiredType := r.TypeFor[T]()
	destType := reflect.TypeFor[T]()
	destValue := reflect.New(destType).Elem()

	res = destValue.Interface().(T)

	if destType.Kind() == reflect.Pointer {
		destType = destType.Elem()
	}

	// If source is already of destination Type
	if destType == srcValue.Type() {
		res = srcValue.Interface().(T)
		return res, nil
	}

	if !destValue.CanSet() {
		return res, NewError(srcValue, destValue, "destination can not be set")
	}

	if destType == srcValue.Type() {
		return srcValue.Interface().(T), nil
	}

	if destType.Kind() == reflect.String {
		if srcValue.Type().Kind() == reflect.String {
			destValue.SetString(srcValue.String())
			return destValue.Interface().(T), nil
		}

		if srcValue.Type().AssignableTo(reflect.TypeOf((fmt.Stringer)(nil))) {
			results := srcValue.MethodByName("String").Call([]reflect.Value{})

			return results[0].Interface().(T), nil
		}
	}

	// if the source value expects an interface,Array, Map or Struct
	// Side effect Call
	err = copyValue(srcValue, destValue)
	if err != nil {
		return res, err
	}

	// destValue is a new value of type T so this type cast won't panic
	return destValue.Interface().(T), nil
}
