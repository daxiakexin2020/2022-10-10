package helper

type Num interface {
	~int | ~int32 | ~int64 | ~uint | ~uint32 | ~uint64
}

type Str interface {
	string
}

type NumStr interface {
	Num | Str
}

func InArray[T NumStr](arr []T, dst T) bool {
	for _, v := range arr {
		if v == dst {
			return true
		}
	}
	return false
}

func InMap[T NumStr](m map[T]T, dst T) bool {
	_, ok := m[dst]
	return ok
}
