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
