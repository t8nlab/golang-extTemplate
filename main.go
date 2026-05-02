package main

import ext "github.com/ext/go-extension/utils"

/**
 * 1. Define your function
 * It must accept (map[string]any) and return (any, error)
 */
func add(input map[string]any) (any, error) {
	// 2. Extract inputs using type helpers from utils
	// GetInt, GetString, GetBool help handle JSON type conversions
	n1, _ := ext.GetInt(input, "n1")
	n2, _ := ext.GetInt(input, "n2")

	// 3. Perform your logic
	sum := (n1 + n2)

	// 4. Return the result (will be auto-serialized to JSON for JS)
	return sum, nil
}

// 5. Register your functions in the init() block
// This makes them available to be called via ext.call("name", payload) in JS
func init() {
	ext.Register("add_number", add)
}

/**
 * The main function is required for c-shared build mode 
 * but should remain empty as the logic is handled by exports.
 */
func main() {}