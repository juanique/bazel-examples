# Run example

##  Run a simple python binary that makes use of a PIP dependency

```
bazel run :version
```

## Update pip dependencies

```
# Add a new pip library to the requirements.txt file
echo "spinners==0.0.24" >> requirements.txt

# Update the lock file
bazel run //:compile_pip.update
```

The new requirement is now available as `@pip_spinners//:pkg`
