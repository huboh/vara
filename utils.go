package vara

import "fmt"

// GetToken generates a unique token for the given value based on its type.
func GetToken(v any) string {
	return fmt.Sprintf("%T", v)
}
