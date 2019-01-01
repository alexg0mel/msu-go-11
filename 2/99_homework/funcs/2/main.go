package main

func showMeTheType(i interface{}) string {
	if _, ok := i.(int); ok {
		return "int"
	}
	if _, ok := i.(uint); ok {
		return "uint"
	}
	if _, ok := i.(int8); ok {
		return "int8"
	}
	if _, ok := i.(string); ok {
		return "string"
	}
	if _, ok := i.(int32); ok {
		return "int32"
	}
	if _, ok := i.([]int); ok {
		return "[]int"
	}
	if _, ok := i.(float64); ok {
		return "float64"
	}

	if _, ok := i.(map[string]bool); ok {
		return "map[string]bool"
	}
	return "?"
}

func main() {

}
