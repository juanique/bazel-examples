# Attempt to build AWS C++ SDO

Requires a working C++ toolchain:

```
sudo apt-get install -y gcc build-essential
```

To attempt the build:

```
bazel build @alsa//:alsa_with_rules_foreign
```

It fails with:

```
configure.ac:49: error: possibly undefined macro: AC_CHECK_ATTRIBUTE_SYMVER
      If this token and others are legitimate, please use m4_pattern_allow.
      See the Autoconf documentation.
make: *** [Makefile:384: /home/juan-munoz/.cache/bazel/_bazel_juan-munoz/95fba053af938229db5bffcfc16c50fe/sandbox/linux-sandbox/3/execroot/bazel-example/external/alsa/configure] Error 1
```
