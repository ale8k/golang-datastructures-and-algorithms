package linearsearch

import (
	"fmt"
	"reflect"
)

// Searches for the provided param in a linear fashion
// returns -1 if cannot be found and nil for the value
func PrimitiveSearch(param interface{}, slice interface{}) int {
	sliceType := reflect.TypeOf(slice)
	sliceKind := sliceType.Kind()
	sliceVal := reflect.ValueOf(slice)

	fmt.Println(sliceType, sliceKind, sliceVal)

	if sliceKind == reflect.Slice && sliceVal.Len() > 0 {
		for i := 0; i < sliceVal.Len(); i++ {
			match := reflect.DeepEqual(param, sliceVal.Index(i).Interface())
			if match {
				return i
			}
		}
	}
	return -1
}
