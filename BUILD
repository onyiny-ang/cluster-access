load("@io_bazel_rules_go//go:def.bzl", "gazelle", "go_prefix")

exports_files(["LICENSES"])

gazelle(
    name = "gazelle",
    external = "vendored"
)

go_prefix("k8s.io/cluster-access")

genrule(
        name = "sources",
        outs = ["SOURCES.md"],
        stamp = 1,
        cmd = "VERSION=\"$$(awk '/STABLE_BUILD_GIT_COMMIT/ {print $$2}' bazel-out/stable-status.txt)\"; " +
          "echo \"The sources for this release can be found on GitHub:\n" +
          "> $@",
    visibility = [
      "//cmd:__subpackages__",
     ]
)
