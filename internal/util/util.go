package util

func NilString(str string) *string {
	if len(str) > 0 {
		return &str
	}

	return nil
}

func NilUint64(i uint64) *uint64 {
	if i != 0 {
		return &i
	}

	return nil
}