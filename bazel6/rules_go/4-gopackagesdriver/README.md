# Run example


## Run a simple go binary

```
bazel run //example/cli:cli
```

## Add a new external dependency

```
# update go.mod/go.sum files
bazel run @io_bazel_rules_go//go get github.com/spf13/cobra

# Delete existing build files
rm example/cli/BUILD.bazel example/capitalize/BUILD.bazel


# Update bazel external repositories
bazel run //:gazelle-update-repos
```

The external dependency can now be imported and `gazelle` will automatically add the `deps` to the corresponding target.
