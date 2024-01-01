def oci_image_loader_impl(ctx):
    script = "#!/bin/bash\n./{} {} {}".format(
        ctx.executable._loader_bin.short_path,
        ctx.file.image.short_path,
        " ".join(ctx.attr.repo_tags)
    )

    output = ctx.actions.declare_file(ctx.label.name + ".sh")
    ctx.actions.write(
        output = output,
        content = script,
        is_executable = True,
    )

    runfiles = ctx.runfiles(files = [ctx.file.image])
    runfiles = runfiles.merge(ctx.attr.image[DefaultInfo].default_runfiles)
    runfiles = runfiles.merge(ctx.attr._loader_bin[DefaultInfo].default_runfiles)

    return [
        DefaultInfo(files = depset([output]), runfiles = runfiles, executable = output)
    ]

oci_image_loader = rule(
    implementation = oci_image_loader_impl,
    attrs = {
        "image": attr.label(
            mandatory = True,
            allow_single_file = True, 
            doc = "Label of a directory containing an OCI layout, e.g. `oci_image`",
        ),
        "repo_tags": attr.string_list(
            mandatory = True,
            doc = "Tags to set to the image",
        ),
        "_loader_bin": attr.label(
            default = "//oci/loader",
            allow_single_file = True,
            executable = True,
            cfg = "target",
            doc = "Binary tool that can take in a directory containing an OCI layout and loade it."
        ),
    },
    executable = True,
)