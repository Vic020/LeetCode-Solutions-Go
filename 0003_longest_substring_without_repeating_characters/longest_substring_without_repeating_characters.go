package main

import "fmt"

// O(n^2) 386ms
func lengthOfLongestSubstring1(s string) int {
	l, p, q, ml, nl := len(s), 0, 0, 0, 0

	d := map[uint8]bool{}

	for p < l && q < l {

		_, ok := d[s[q]]

		if ok {
			if nl > ml {
				ml = nl
			}
			p++
			q = p
			nl = 0
			d = map[uint8]bool{}
		} else {
			d[s[q]] = true
			q++
			nl++
		}
	}

	if nl > ml {
		ml = nl
	}

	return ml
}

//O(n) + O(n-1) 16ms
func lengthOfLongestSubstring2(s string) int {
	l, p, q, ml := len(s), 0, 0, 0

	d := map[uint8]bool{}

	for p < l && q < l {

		_, ok := d[s[q]]

		if ok {
			delete(d, s[p])
			p++
		} else {
			d[s[q]] = true
			q++
			if q-p > ml {
				ml = q - p
			}
		}
	}

	return ml
}

//O(n) 12ms
func lengthOfLongestSubstring3(s string) int {
	l, p, q, ml := len(s), 0, 0, 0

	d := map[uint8]int{}

	for ; q < l; q++ {

		a, ok := d[s[q]]

		if ok && p <= a {
			p = a + 1
		}

		if q-p+1 > ml {
			ml = q - p + 1
		}

		d[s[q]] = q

	}

	return ml
}

//O(n) 4ms
func lengthOfLongestSubstring(s string) int {
	l, p, q, ml, d := len(s), 0, 0, 0, [128]int{0}

	var a int

	for ; q < l; q++ {

		a = d[s[q]]

		if p <= a {
			p = a
		}

		if q-p+1 > ml {
			ml = q - p + 1
		}

		d[s[q]] = q + 1

	}

	return ml
}

func main() {
	fmt.Println(lengthOfLongestSubstring("bbbb"))
	fmt.Println(lengthOfLongestSubstring("abba"))
	fmt.Println(lengthOfLongestSubstring("a"))
	fmt.Println(lengthOfLongestSubstring("abcbef"))
}
