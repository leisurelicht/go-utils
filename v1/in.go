package utils

func InSlice(val any, s any) bool {
	switch s.(type) {
	case []string:
		return InStringSlice(val.(string), s.([]string))
	case []int:
		return InIntSlice(val.(int), s.([]int))
	case []int64:
		return InInt64Slice(val.(int64), s.([]int64))
	case []float64:
		return InFloat64Slice(val.(float64), s.([]float64))
	}

	return false
}

func InStringSlice(val string, s []string) bool {
	if len(s) == 0 {
		return false
	} else if len(s) < 100 {
		for _, item := range s {
			if item == val {
				return true
			}
		}
	} else {
		i, j := 0, len(s)-1
		for i <= j {
			if s[i] == val || s[j] == val {
				return true
			}
			i++
			j--
		}
	}
	return false
}

func InIntSlice(val int, s []int) bool {
	if len(s) == 0 {
		return false
	} else if len(s) < 100 {
		for _, item := range s {
			if item == val {
				return true
			}
		}
	} else {
		i, j := 0, len(s)-1
		for i <= j {
			if s[i] == val || s[j] == val {
				return true
			}
			i++
			j--
		}
	}
	return false
}

func InInt64Slice(val int64, s []int64) bool {
	if len(s) == 0 {
		return false
	} else if len(s) < 100 {
		for _, item := range s {
			if item == val {
				return true
			}
		}
	} else {
		i, j := 0, len(s)-1
		for i <= j {
			if s[i] == val || s[j] == val {
				return true
			}
			i++
			j--
		}
	}
	return false
}

func InFloat64Slice(val float64, s []float64) bool {
	if len(s) == 0 {
		return false
	} else if len(s) < 100 {
		for _, item := range s {
			if item == val {
				return true
			}
		}
	} else {
		i, j := 0, len(s)-1
		for i <= j {
			if s[i] == val || s[j] == val {
				return true
			}
			i++
			j--
		}
	}
	return false
}
