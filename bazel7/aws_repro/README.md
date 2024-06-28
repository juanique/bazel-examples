# Attempt to build AWS C++ SDO

Requires a working C++ toolchain:

```
sudo apt-get install -y gcc build-essential
```

To attempt the build:

```
bazel build @aws_sdk//:aws_sdk_cpp --spawn_strategy=processwrapper-sandbox --copt=-w
```

It fails:

```
CMake Error at cmake/external_dependencies.cmake:46 (message):
  Could not find openssl
Call Stack (most recent call first):
  CMakeLists.txt:112 (include)
```
