# Extension Utilities (By TitanPl Extension Support)

This directory contains utility functions for bridging Go and JavaScript.

## Go Utilities (`utils/*.go`)

The Go utilities help with data extraction and function registration.

### `native.go`
- `Register(name string, fn func(map[string]any) (any, error))`: Registers a Go function to be callable from JavaScript.
- `titan_invoke`: The internal entry point for TitanPL. Do not call this directly.

### `input.go`
- `GetString(m map[string]any, key string) (string, error)`: Safely extracts a string from the input map.
- `GetInt(m map[string]any, key string) (int, error)`: Safely extracts an integer (handles float64 from JSON).
- `GetBool(m map[string]any, key string, def bool) bool`: Safely extracts a boolean with a default value.

## JavaScript Utilities (`utils/native.js`)

### `createExt(extName)`
Initializes the bridge for a specific extension.

```javascript
import { createExt } from "./utils/native";
export const ext = createExt("@ext/your-extension-name");
```

- `ext.call(fnName, payload)`: Calls a registered Go function.
- `ext.require(fnName)`: Gets a direct reference to a native function.
