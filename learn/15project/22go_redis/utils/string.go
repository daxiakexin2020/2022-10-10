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
