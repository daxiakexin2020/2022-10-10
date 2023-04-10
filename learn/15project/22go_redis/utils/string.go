package utils

func SliceUniq(dest []interface{}) []interface{} {
	m := make(map[interface{}]struct{})
	var ret []interface{}
	for _, i := range dest {
		if _, ok := m[i]; !ok {
			m[i] = struct{}{}
			ret = append(ret, i)
		}
	}
	return ret
}

func SliceDiff(dest []interface{}, sdata []interface{}) []interface{} {
	var ret []interface{}
	flags := make(map[interface{}]struct{})
	for _, data := range sdata {
		if _, ok := flags[data]; !ok {
			flags[data] = struct{}{}
		}
	}
	for _, data := range dest {
		if _, ok := flags[data]; !ok {
			ret = append(ret, data)
		}
	}
	return ret
}

func InSlice(dest []interface{}, val interface{}) bool {
	for _, data := range dest {
		if data == val {
			return true
		}
	}
	return false
}
