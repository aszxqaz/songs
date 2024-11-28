package helpers

func DerefOrDefault[T any](v *T, defaultVal ...T) T {
	if v == nil {
		var t T
		if len(defaultVal) > 0 {
			t = defaultVal[0]
		}
		return t
	}
	return *v
}
