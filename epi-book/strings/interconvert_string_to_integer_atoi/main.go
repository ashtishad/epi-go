package epi

const (
	MaxInt32 = 1<<31 - 1
	MinInt32 = -1 << 31
)

func myAtoi(s string) int {
	var (
		symbolRead bool
		multiplier int64 = 1
		result     int64
	)
	for _, c := range s {
		// skip leading spaces
		if c == ' ' {
			if symbolRead {
				break
			}
			continue
		}
		// check for negative sign
		if c == '+' || c == '-' {
			if symbolRead {
				return int(result)
			}
			symbolRead = true
			if c == '-' {
				multiplier = -1
			}
			continue
		}

		// check for invalid character
		if c < '0' || c > '9' {
			break
		}

		symbolRead = true
		value := int64(c - '0')
		result = (result * 10) + (value * multiplier)

		// check for overflow and underflow
		if result > MaxInt32 {
			return MaxInt32
		}
		if result < MinInt32 {
			return MinInt32
		}
	}
	return int(result)
}

/*
	I only need to handle four cases:
		discards all leading whitespaces (input: "   -42)
		sign of the number (sometime not in s[0])
		overflow and underflow
		invalid input
*/
