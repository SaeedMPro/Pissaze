package util

func NilFixer[T any](slice []T) []T {
    if slice == nil {
        return make([]T, 0)
    }
    return slice
}