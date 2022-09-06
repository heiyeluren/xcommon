package xcommon

import "strconv"

func ToInt(s string) int {
	if s == "" {
		return 0
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return v
}

func ToInt32(s string) int32 {
	if s == "" {
		return 0
	}

	v, err := strconv.ParseInt(s, 0, 0)
	if err != nil {
		return 0
	}

	return int32(v)
}

func ToUint64(s string) uint64 {
	if s == "" {
		return 0
	}

	v, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		return 0
	}

	return v
}

func ToFloat32(s string) float32 {
	if s == "" {
		return 0
	}

	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0
	}

	return float32(v)
}
