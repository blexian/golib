package leetcode

func pushDominoes(dominoes string) string {
	n := len(dominoes)
	res := make([]uint8, n)
	for i := 0; i < n; i++ {
		res[i] = dominoes[i]
		switch dominoes[i] {
		case 'L': // 往回找.
			for j := i - 1; j >= 0; j-- {
				if dominoes[j] == '.' {
					res[j] = 'L'
				}
				if dominoes[j] == 'L' || dominoes[j] == 'R' {
					break
				}
			}
		case 'R':
			// 向右找到第一个L或者R或者数组尾巴
			t := '.'
			j := i + 1
			for ; j < n; j++ {
				if dominoes[j] == 'R' {
					res[j] = 'R'
					t = 'R'
					break
				}
				if dominoes[j] == 'L' {
					res[j] = 'L'
					t = 'L'
					break
				}
			}
			if t == 'R' {
				for k := i + 1; k < j; k++ {
					res[k] = 'R'
				}
				i = j - 1
			}
			if t == 'L' {
				mid := (i + j) / 2
				mod := (j - i + 1) % 2
				rt := mid
				lh := mid
				if mod == 1 {
					rt = mid - 1
					lh = mid + 1
					res[mid] = '.'
				} else {
					rt = mid
					lh = mid + 1
				}
				for k := i + 1; k <= rt; k++ {
					res[k] = 'R'
				}
				for k := lh; k < j; k++ {
					res[k] = 'L'
				}
				i = j
			}
			if t == '.' { // 数组尾巴了，全变R
				for k := i + 1; k < j; k++ {
					res[k] = 'R'
				}
				i = j - 1
			}
		case '.':
		}
	}
	return string(res[:n])
}
