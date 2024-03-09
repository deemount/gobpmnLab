package utils

import "reflect"

func MergeStructs(dst, src interface{}) {
	dstVal := reflect.ValueOf(dst)
	srcVal := reflect.ValueOf(src)

	// Ensure dst is a pointer and src is not nil
	if dstVal.Kind() != reflect.Ptr || dstVal.IsNil() {
		panic("dst must be a non-nil pointer")
	}
	if srcVal.Kind() != reflect.Ptr || srcVal.IsNil() {
		panic("src must be a non-nil pointer")
	}

	// Dereference pointers to get to the actual struct values
	dstVal = dstVal.Elem()
	srcVal = srcVal.Elem()

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		dstField := dstVal.FieldByName(srcVal.Type().Field(i).Name)

		if dstField.IsValid() && dstField.CanSet() {
			dstField.Set(srcField)
		}
	}
}
