load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def install_mingw64_toolchain():
    http_archive(
        name = "mingw64_gcc",
        urls = ["https://github.com/brechtsanders/winlibs_mingw/releases/download/13.1.0-16.0.2-11.0.0-ucrt-r1/winlibs-x86_64-mcf-seh-gcc-13.1.0-llvm-16.0.2-mingw-w64ucrt-11.0.0-r1.zip"],
        strip_prefix = "mingw64",
        build_file = "//toolchain_mingw64_gcc:migwin64_gcc.BUILD",
        sha256 =  "24710ca00ae72f0197f98bd7f22f9ab9958c20ebcb0997c7eab186784815133e",
    )

    native.register_toolchains(
        "@mingw64_gcc//:toolchain",
    )