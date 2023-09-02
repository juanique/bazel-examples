# Codelab 0 : Installing bazel and setting up a workspace

In this codelab we will initialize an empty bazel workspace and verify the installation. We'll also learn about some core bazel concepts.

## Part 1: Bazelisk

[Bazelisk](https://github.com/bazelbuild/bazelisk) is a user-friendly launcher for bazel. Instead of installing `bazel` directly, we will install `bazelisk` which will download and execute the correct `bazel` version for us.

You can download bazelisk from its [releases page on GitHub](https://github.com/bazelbuild/bazelisk/releases). Simply download the appropriate release for your OS and architecture and copy it to a place in your `$PATH`. Also, we will rename the downloaded file to `bazel` because its interface is effectively the same as bazel. The following example is for Linux AMD64 and assumes `~/bin` is in the system `$PATH`:

```
$ wget https://github.com/bazelbuild/bazelisk/releases/download/v1.18.0/bazelisk-linux-amd64
# mv bazelisk-linux-amd64 ~/bin/bazel
```

We can now check the installation (output may change depending on when you run this):

```
$ bazel --version

2023/09/02 12:17:38 Downloading https://releases.bazel.build/6.3.2/release/bazel-6.3.2-linux-x86_64...
bazel 6.3.2
```

Here, bazelisk downloaded `bazel` version `6.3.2` for us.

## Part 2: Setting up a workspace

Our bazel workspace is the directory in which our Bazel compiled source code will live. A workspace is identified by having a file named `WORKKSPACE` at the root. Let's create a workspace, we'll call it `monorepo`:

```
$ mkdir monorepo
$ cd monorepo
$ echo 'workspace(name = "monorepo")' > WORKSPACE
```

Our `WORKSPACE` file now contains a single line that contains the name of the workspace. We can verify that our workspace is properly setup by running the `bazel info` command which provides information about the workspace:

```
$ bazel info

Starting local Bazel server and connecting to it...
bazel-bin: /home/juanique/.cache/bazel/_bazel_juanique/1f8e0e029d65256e712d13a7e248f948/execroot/monorepo/bazel-out/k8-fastbuild/bin
bazel-genfiles: /home/juanique/.cache/bazel/_bazel_juanique/1f8e0e029d65256e712d13a7e248f948/execroot/monorepo/bazel-out/k8-fastbuild/bin
bazel-testlogs: /home/juanique/.cache/bazel/_bazel_juanique/1f8e0e029d65256e712d13a7e248f948/execroot/monorepo/bazel-out/k8-fastbuild/testlogs
character-encoding: file.encoding = ISO-8859-1, defaultCharset = ISO-8859-1
command_log: /home/juanique/.cache/bazel/_bazel_juanique/1f8e0e029d65256e712d13a7e248f948/command.log
committed-heap-size: 524MB
execution_root: /home/juanique/.cache/bazel/_bazel_juanique/1f8e0e029d65256e712d13a7e248f948/execroot/monorepo
gc-count: 3
gc-time: 15ms
install_base: /home/juanique/.cache/bazel/_bazel_juanique/install/a09dbb90c658248f08f9aa0eba11997d
java-home: /home/juanique/.cache/bazel/_bazel_juanique/install/a09dbb90c658248f08f9aa0eba11997d/embedded_tools/jdk
java-runtime: OpenJDK Runtime Environment (build 11.0.6+10-LTS) by Azul Systems, Inc.
java-vm: OpenJDK 64-Bit Server VM (build 11.0.6+10-LTS, mixed mode) by Azul Systems, Inc.
max-heap-size: 8378MB
output_base: /home/juanique/.cache/bazel/_bazel_juanique/1f8e0e029d65256e712d13a7e248f948
output_path: /home/juanique/.cache/bazel/_bazel_juanique/1f8e0e029d65256e712d13a7e248f948/execroot/monorepo/bazel-out
package_path: %workspace%
release: release 6.3.2
repository_cache: /home/juanique/.cache/bazel/_bazel_juanique/cache/repos/v1
server_log: /home/juanique/.cache/bazel/_bazel_juanique/1f8e0e029d65256e712d13a7e248f948/java.log.hyperion.juanique.log.java.20230902-122239.29552
server_pid: 29552
used-heap-size: 88MB
workspace: /home/juanique/monorepo
```

## Part 3: Changing the bazel version

There are a few ways of changing the bazel version with bazelisk. One of them is to include a file named `.bazelversion` that contains the version we want to use. Earlier we saw that bazelisk automatically selected for us version `6.3.2` but we may want to use a different version for multiple reasons. Let's pin our version to `6.2.0`:

```
$ echo "6.2.0" > .bazelversion

$ bazel --version

2023/09/02 12:25:59 Downloading https://releases.bazel.build/6.2.0/release/bazel-6.2.0-linux-x86_64...
bazel 6.2.0
```

Bazelisk automatically downloaded the desired version for us. We can also override this by using the `USE_BAZEL_VERSION` environment variable:

```
$ USE_BAZEL_VERSION=6.2.1 bazel --version

2023/09/02 12:28:12 Downloading https://releases.bazel.build/6.2.1/release/bazel-6.2.1-linux-x86_64...
bazel 6.2.1
```

## Conclusion

In this codelab you learned:
- How to install `bazel` using `bazelisk`.
- How create and name a new bazel workspace.
- How to pin the bazel version using `.bazelversion`
- How to change the bazel version using `USE_BAZEL_VERSION`

Results for this codelab are in the subdirectories of this document.