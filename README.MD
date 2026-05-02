# TitanPL Golang Extension Starter

A professional starter kit for building native Golang extensions for the TitanPL framework.

## 🚀 Getting Started

This repository serves as a template. Before you begin developing, you **must** rename the extension to match your project.

### 🛠️ Customization (CRITICAL)
Search and replace `@ext/go-extension` with your desired extension name (e.g., `@my-org/my-ext`) in the following files:
- `package.json`: The `"name"` field.
- `titan.json`: The `"name"` field.
- `ext.js`: The name passed to `createExt()`.
- `go.mod`: The module path (e.g., `github.com/user/my-ext`).
- `main.go`: The import path for the `utils` package.
- `index.d.ts`: The `@package` documentation tag.
- `test/tanfig.json`: The extension name in `allowNative`.
- `test/package.json`: The dependency name.

## 📋 Prerequisites

- [Go](https://go.dev/dl/) (version 1.25.0 or later)
- [Node.js & npm](https://nodejs.org/)
- [TitanPL CLI](https://github.com/t8nlab/titanpl)

## 📂 Project Structure

- `main.go`: The Go source code for the native extension.
- `ext.js`: JavaScript bridge to register and interface with the native extension.
- `index.js`: Main entry point for the extension (Class-based).
- `index.d.ts`: TypeScript definitions for the extension.
- `utils/`: Native Go and JS utility functions for data handling (Provided from TitanPL extension support).
- `titan.json`: Extension metadata for TitanPL.

## 🏗️ Build Guide

To build the native binary:

- **Windows**: `npm run build` (outputs `go-extension.dll`)
- **Linux**: `npm run build:linux` (outputs `go-extension.so`)

The output binaries (and their respective `.h` files) are required for the extension to function. You can change the output filenames in `package.json` and `titan.json`.

## 🧪 Testing Your Extension

1. **Initialize a Test Server**:
   Create a new TitanPL server for your test:
   ```bash
   titan init test
   ```

2. **Link the Extension**:
   From the root of this extension directory, run:
   ```bash
   npm link
   ```
   Then in your `test` directory, run:
   ```bash
   npm link @ext/go-extension (replace with your extension name)
   ```
   *Alternatively, you can install it via relative path:*
   ```bash
   npm install ../
   ```

3. **Configure Permissions**:
   Ensure `test/tanfig.json` allows the native extension:
   ```json
   {
     "extensions": {
       "allowNative": [
         "@ext/go-extension"
       ]
     }
   }
   ```

4. **Usage Example**:
   In your test server's actions or logic:
   ```javascript
   import { log } from "@titanpl/native";
   import Extension from '@ext/go-extension';
   
   const ext = new Extension();
   const result = ext.addNumber(10, 20);
   log(result); // 30
   ```

## 💡 Development Tips

- **Add Native Functions**: Add your Go logic in `main.go` and register it in the `init()` function using `ext.Register("name", func)`.
- **Sync Types**: Always update `index.d.ts` when you add new methods to ensure a great developer experience.
- **Data Handling**: Use the helpers in `utils/` (like `ext.GetInt`, `ext.GetString`) to safely parse inputs from JavaScript.