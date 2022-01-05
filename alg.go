package main

// 56
func Merge(intervals [][]int) [][]int {
	quick_sort(intervals, 0, len(intervals)-1)
	result := [][]int{}

	for _, v := range intervals {
		if len(result) == 0 {
			result = append(result, v)
		} else if v[0] > result[len(result)-1][1] {
			result = append(result, v)
		} else {
			tm := result[len(result)-1]
			if v[1] > tm[1] {
				tm[1] = v[1]
			}
		}
	}
	return result
}

func quick_sort(res [][]int, l, h int) {
	if h-l <= 0 {
		return
	}

	if h-l == 1 {
		if res[l][0] > res[h][0] {
			res[l], res[h] = res[h], res[l]
		}
		return
	}

	mid := (l+h)/2 + (l+h)%2
	midVal := res[mid]
	left, right := l, h
	res[mid] = res[right]
	for {
		for left < right && res[left][0] < midVal[0] {
			left++
		}
		if left >= right {
			break
		}
		res[right] = res[left]
		right--

		for left < right && res[right][0] > midVal[0] {
			right--
		}
		if left >= right {
			break
		}
		res[left] = res[right]
		left++
	}

	res[left] = midVal
	quick_sort(res, l, left-1)
	quick_sort(res, left+1, h)
}
