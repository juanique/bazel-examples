# What's in here

- `bazel v6.3.2``
- `rules_go v0.41.0`
- `aspect_gcc_toolchain v0.4.2`
- A simple golang library in `//example/capitalize`
- A `gopackagesdriver` script in `//bazel/scripts:gopackagesdriver`

# What to do:

```
echo {} | bazel run @io_bazel_rules_go//go/tools/gopackagesdriver -- file=example/capitalize/capitalize.go
```

- *Expected*: Some JSON output for the language server.
- *Actual*: `cgo-builtin-prolog:1:10: fatal error: stddef.h: No such file or directory`