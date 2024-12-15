package t2

func Int32ArrTo64(arr ...int32) []int64 {
	ret := make([]int64, 0)
	for _, i := range arr {
		ret = append(ret, int64(i))
	}

	return ret
}
