/*	Small Number Theory Library

	Author: Serhat Sevki Dincer, jfcgaussATgmail
*/

package snt

// Returns list of primes less than n
func Lspr(n uint32) (P []uint32) {
	if n < 10 {
		if n < 3 {
			return
		}
		P = append(P, 2)
		for q := uint32(3); q < n; q += 2 {
			P = append(P, q)
		}
		return
	}

	// P(k-1)^2 < q < P(k)^2 is prime if not divisible by primes P(0..k-1)
	P = append(P, 2, 3, 5, 7)
	var p2, q, r uint32 = 49, 11, 1176912450 // P[k]^2, candidate, increment list 2,4,2,4,6,2,6,4

	for k := 3; q < n; q, r = q+r&7, r>>4^r<<28 {
		if !(q < p2 && p2 <= n) { // guard against p2 overflow
			p2 = n
		}

	nextq:
		for ; q < p2; q, r = q+r&7, r>>4^r<<28 { // avoid (multiples of) 2,3,5
			if q>>3 == 0 { // guard against q overflow
				return
			}

			for i := 3; i < k; i++ {
				if q%P[i] == 0 { // try candidates < p2 with P[3:k]
					continue nextq
				}
			}
			P = append(P, q)
		} // here q=p2

		k++
		p2 = P[k] * P[k]
	}
	return
}

// Calculates Jacobi(c/b), returns 2 for even b
func Jacobi(c int64, b uint64) int {
	if b&1 == 0 {
		return 2
	}
	if b == 1 {
		return 1
	}

	n, r := 0, 1
	if c < 0 {
		c = -c
		if b&3 == 3 {
			r = -r // negative c
		}
	}

	a := uint64(c)
start:
	a %= b
	if a == 0 {
		return 0
	}

	for n = 0; a&1 == 0; n++ {
		a >>= 1 // largest power of 2 dividing a
	}

	if n&1 != 0 && (b&7 == 3 || b&7 == 5) {
		r = -r // odd power of 2 divides a
	}

	if a == 1 {
		return r
	}

	// a is odd, 2 < a < b
	if a&3 == 3 && b&3 == 3 {
		r = -r // quadratic reciprocity
	}
	a, b = b, a
	goto start
}
