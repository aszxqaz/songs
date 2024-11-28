package helpers

func Coalesce[T any](r *T, fallback T) T {
	if r == nil {
		return fallback
	}
	return *r
}
