# What's in here

- `bazel v6.3.2``
- `rules_go v0.39.0`
- `aspect_gcc_toolchain v0.4.2`
- A simple golang library in `//example/capitalize`
- A `gopackagesdriver` script in `//bazel/scripts:gopackagesdriver`

# What to do:

```
echo {} | bazel run @io_bazel_rules_go//go/tools/gopackagesdriver -- file=example/capitalize/capitalize.go
```

Depending on whether you have a configurec C++ toolchain, you may see a large json blob in the output or the following error:

```
Auto-Configuration Error: Cannot find gcc or CC; either correct your path or set the CC environment variable
```
