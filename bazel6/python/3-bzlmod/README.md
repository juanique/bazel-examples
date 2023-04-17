# DOES NOT WORK
This example does not work.

```
ERROR: /home/juanique/bazel-examples/bazel6/python/3-bzlmod/BUILD.bazel:12:10: no such package '@[unknown repo 'pip_rich' requested from @]//': The repository '@[unknown repo 'pip_rich' requested from @]' could not be resolved: No repository visible as '@pip_rich' from main repository and referenced by '//:version'
```

# Run example

##  Run a simple python binary that makes use of a PIP dependency

```
bazel run :version
```
