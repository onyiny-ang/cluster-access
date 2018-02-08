# These dependencies' versions are pulled from the k/k WORKSPACE.
# https://github.com/kubernetes/kubernetes/blob/77ac663df427d1ae0cb45adb0a3eba263809c837/build/root/WORKSPACE
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "441e560e947d8011f064bd7348d86940d6b6131ae7d7c4425a538e8d9f884274",
    strip_prefix = "rules_go-c72631a220406c4fae276861ee286aaec82c5af2",
    urls = ["https://github.com/bazelbuild/rules_go/archive/c72631a220406c4fae276861ee286aaec82c5af2.tar.gz"],
)

http_archive(
    name = "io_kubernetes_build",
    sha256 = "8e49ac066fbaadd475bd63762caa90f81cd1880eba4cc25faa93355ef5fa2739",
    strip_prefix = "repo-infra-e26fc85d14a1d3dc25569831acc06919673c545a",
    urls = ["https://github.com/kubernetes/repo-infra/archive/e26fc85d14a1d3dc25569831acc06919673c545a.tar.gz"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains", "go_repository")
go_rules_dependencies()
go_register_toolchains(go_version = "1.9.1")
