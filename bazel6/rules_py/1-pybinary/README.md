
# Run example

```
bazel run :version
```

Currently fails with:

```
INFO: Repository rules_cc~0.0.2 instantiated at:
  callstack not available
Repository rule http_archive defined at:
  /home/juanique/.cache/bazel/_bazel_juanique/5eb256d0f1ef2af2265f414630a853d3/external/bazel_tools/tools/build_defs/repo/http.bzl:372:31: in <toplevel>
INFO: Repository rules_java~5.5.0 instantiated at:
  callstack not available
Repository rule http_archive defined at:
  /home/juanique/.cache/bazel/_bazel_juanique/5eb256d0f1ef2af2265f414630a853d3/external/bazel_tools/tools/build_defs/repo/http.bzl:372:31: in <toplevel>
ERROR: Analysis of target '//:version' failed; build aborted: in tag at /home/juanique/gg/bazel-examples/bazel6/rules_py/1-pybinary/MODULE.bazel:5:17, unknown attribute name provided
INFO: Elapsed time: 0.576s
INFO: 0 processes.
ERROR: Build failed. Not running target
FAILED: Build did NOT complete successfully (29 packages loaded, 7 targets configured)
    currently loading: @bazel_tools~sh_configure_extension~local_config_sh//
    Fetching https://github.com/bazelbuild/rules_cc/releases/download/0.0.2/rules_cc-0.0.2.tar.gz
    Fetching https://github.com/bazelbuild/rules_java/releases/download/5.5.0/rules_java-5.5.0.tar.gz
```