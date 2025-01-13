package gopointer

func ToPointer[T any](value T) *T {
	return &value
}

func ToValue[T any](pointer *T) T {
	return *pointer
}
