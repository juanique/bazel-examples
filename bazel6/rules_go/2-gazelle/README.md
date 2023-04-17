# Run example


## Run a simple go binary

```
bazel run //example/cli:cli
```

## Regenerate BUILD.bazel files using gazelle

```
# Delete existing build files
rm example/cli/BUILD.bazel example/capitalize/BUILD.bazel

# Regenerate them
bazel run //:gazelle
```


## Windows notes

### Bash dependency

On windows, `bazel run //:gazelle` can fail with:

```
/bin/bash: C:\...\_bazel_\uqpew3t7\execroot\bazel-example\bazel-out\x64_windows-fastbuild\bin\gazelle: command not found
```
