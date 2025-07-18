# gyu-dom
DOM made by WebAssembly component and Go lang.

- Documents
  - https://component-model.bytecodealliance.org/language-support/go.html

### init project
```
go mod init [module-path]
```

### update go mod dependency
```
go mod tidy
```

### Resolve wit import and build package -> .wasm
```
wkg wit build
```

### Generate bindings wasm <-> go
```
go tool wit-bindgen-go generate --world world-domparser --out internal ./pews:dom@0.0.1.wasm
```

### Build component
```
tinygo build -target=wasip2 -o domparser.wasm --wit-package pews:dom@0.0.1.wasm --wit-world world-domparser main.go
```

### Exam wasm component exports
```
wasm-tools component wit domparser.wasm
```
