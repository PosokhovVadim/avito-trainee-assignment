package utils

import "strconv"

func ConvertToInt(args ...string) ([]int64, error) {
	var result []int64
	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			return nil, err
		}
		result = append(result, int64(num))
	}
	return result, nil
}
