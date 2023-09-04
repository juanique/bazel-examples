# Run example

```
$ bazel build //:hello_ubuntu_tar
Target //:hello_ubuntu_tar up-to-date:
  bazel-bin/hello_ubuntu_tar/tarball.tar


$ podman load --input $(bazel cquery --output=files //:hello_ubuntu_tar)
...
Loaded image(s): localhost/hello:latest

$ podman run hello
PRETTY_NAME="Ubuntu 22.04.3 LTS"
NAME="Ubuntu"
VERSION_ID="22.04"
VERSION="22.04.3 LTS (Jammy Jellyfish)"
VERSION_CODENAME=jammy
ID=ubuntu
ID_LIKE=debian
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
UBUNTU_CODENAME=jammy
```

In the above example, `podman` can be replaced by `docker` transparently.  The argument for `load --input` is the build `.tar` file which can be obtained by `bazel cquery` or just by pointing to the build output `bazel-bin/hello_ubuntu_tar/tarball.tar`.`