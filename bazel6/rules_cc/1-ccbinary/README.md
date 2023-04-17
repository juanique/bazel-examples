# Run example

```
bazel run :cc_version
```

# General notes

Requires a working C++ toolchain to work.

## Windows notes

If Visual C++ build tools are not installed is not available, you get the following error:

```
The target you are compiling requires Visual C++ build tools.
Bazel couldn't find a valid Visual C++ build tools installation on your machine.
Please check your installation following https://bazel.build/docs/windows#using
```

After installing Visual C++ (follow link in the error message) you may need to set the `BAZEL_VC` environment variable, using Powershell:

```
$Env:BAZEL_VC="C:\Program Files\Microsoft Visual Studio\2022\Community\VC"
```

Adjust to the right installation path.
