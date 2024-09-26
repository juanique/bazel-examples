load("@bazel_tools//tools/cpp:windows_cc_toolchain_config.bzl", "cc_toolchain_config")

cc_toolchain_config(
    name = "msys_x64_mingw",
    cpu = "x64_windows",
    compiler = "mingw",
    toolchain_identifier = "mingw32",
    host_system_name = "windows_x86_64",
    target_system_name = "windows_x86_64",
    target_libc = "mingw",
    abi_version = "local",
    abi_libc_version = "local",
    msvc_cl_path = "unknown",
    msvc_ml_path = "unknown",
    msvc_link_path = "unknown",
    msvc_lib_path = "unknown",
    tool_paths = {
        "ar": "bin/ar",
        "compat-ld": "bin/compat-ld",
        "cpp": "bin/cpp",
        "dwp": "bin/dwp",
        "gcc": "bin/gcc",
        "gcov": "bin/gcov",
        "ld": "bin/ld",
        "nm": "bin/nm",
        "objcopy": "bin/objcopy",
        "objdump": "bin/objdump",
        "strip": "bin/strip",
    },
)

filegroup(
    name = "all",
    srcs = glob(["**/*",])
)

toolchain(
    name = "toolchain",
    exec_compatible_with = [
        "@platforms//os:windows",
        "@platforms//cpu:x86_64",
    ],
    target_compatible_with = [
        "@platforms//os:windows",
        "@platforms//cpu:x86_64",
    ],
    toolchain_type = "@rules_cc//cc:toolchain_type",
    toolchain = ":cc_toolchain",
)


cc_toolchain(
    name = "cc_toolchain",
    all_files =  ":all",
    ar_files = ":all",
    as_files = ":all",
    compiler_files = ":all",
    dwp_files = ":all",
    linker_files = ":all",
    objcopy_files = ":all",
    strip_files = ":all",
    dynamic_runtime_lib = ":all",
    static_runtime_lib = ":all",
    toolchain_config = ":msys_x64_mingw",
)