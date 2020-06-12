package tools

func IsSet(x map[interface{}]interface{}, k interface{}) bool {
	if _, ok := x[k]; !ok {
		return false
	} else {
		return true
	}
}