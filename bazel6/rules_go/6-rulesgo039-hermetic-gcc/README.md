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

Output shows a large json blob that, after formatting looks somewhat like this:

```
{
  "NotHandled": false,
  "Sizes": {
    "WordSize": 8,
    "MaxAlign": 8
  },
  "Roots": [
    "@//example/capitalize:capitalize"
  ],
  "Packages": [
    {
      "ID": "@io_bazel_rules_go//stdlib:internal/goarch",
      "Name": "goarch",
      "PkgPath": "internal/goarch",
      "GoFiles": [
        "/home/juanique/.cache/bazel/_bazel_juanique/6cc4f7b2f16472640a12b38cd00d1265/external/go_sdk/src/internal/goarch/goarch.go",
        "/home/juanique/.cache/bazel/_bazel_juanique/6cc4f7b2f16472640a12b38cd00d1265/external/go_sdk/src/internal/goarch/goarch_amd64.go",
        "/home/juanique/.cache/bazel/_bazel_juanique/6cc4f7b2f16472640a12b38cd00d1265/external/go_sdk/src/internal/goarch/zgoarch_amd64.go"
      ],
      "CompiledGoFiles": [
        "/home/juanique/.cache/bazel/_bazel_juanique/6cc4f7b2f16472640a12b38cd00d1265/external/go_sdk/src/internal/goarch/goarch.go",
        "/home/juanique/.cache/bazel/_bazel_juanique/6cc4f7b2f16472640a12b38cd00d1265/external/go_sdk/src/internal/goarch/goarch_amd64.go",
        "/home/juanique/.cache/bazel/_bazel_juanique/6cc4f7b2f16472640a12b38cd00d1265/external/go_sdk/src/internal/goarch/zgoarch_amd64.go"
      ],
      "ExportFile": "/home/juanique/.cache/bazel/_bazel_juanique/6cc4f7b2f16472640a12b38cd00d1265",
      "Standard": true
    },
    {
      "ID": "@io_bazel_rules_go//stdlib:runtime/internal/math",
      "Name": "math",
      "PkgPath": "runtime/internal/math",
      "GoFiles": [
        "/home/juanique/.cache/bazel/_bazel_juanique/6cc4f7b2f16472640a12b38cd00d1265/external/go_sdk/src/runtime/internal/math/math.go"
      ],
      "CompiledGoFiles": [
        "/home/juanique/.cache/bazel/_bazel_juanique/6cc4f7b2f16472640a12b38cd00d1265/external/go_sdk/src/runtime/internal/math/math.go"
      ],
      "ExportFile": "/home/juanique/.cache/bazel/_bazel_juanique/6cc4f7b2f16472640a12b38cd00d1265",
      "Imports": {
        "internal/goarch": "@io_bazel_rules_go//stdlib:internal/goarch"
      },
      "Standard": true
    },
    ...
```
