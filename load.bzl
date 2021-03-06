load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

def repositories():
    http_archive(
        name = "io_bazel_rules_go",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.21.1/rules_go-v0.21.1.tar.gz",
            "https://github.com/bazelbuild/rules_go/releases/download/v0.21.1/rules_go-v0.21.1.tar.gz",
    ],
        sha256 = "b34cbe1a7514f5f5487c3bfee7340a4496713ddf4f119f7a225583d6cafd793a",
    )

    http_archive(
        name = "bazel_gazelle",
        urls = [
            "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/v0.19.1/bazel-gazelle-v0.19.1.tar.gz",
            "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.19.1/bazel-gazelle-v0.19.1.tar.gz",
        ],
        sha256 = "86c6d481b3f7aedc1d60c1c211c6f76da282ae197c3b3160f54bd3a8f847896f",
    )

    http_archive(
        name = "io_bazel_rules_docker",
        sha256 = "14ac30773fdb393ddec90e158c9ec7ebb3f8a4fd533ec2abbfd8789ad81a284b",
        strip_prefix = "rules_docker-0.12.1",
        urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.12.1/rules_docker-v0.12.1.tar.gz"],
    )

    http_archive(
        name = "io_bazel_rules_docker",
        sha256 = "14ac30773fdb393ddec90e158c9ec7ebb3f8a4fd533ec2abbfd8789ad81a284b",
        strip_prefix = "rules_docker-0.12.1",
        urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.12.1/rules_docker-v0.12.1.tar.gz"],
    )

    git_repository(
        name = "com_google_protobuf",
        commit = "09745575a923640154bcf307fba8aedff47f240a",
        remote = "https://github.com/protocolbuffers/protobuf",
        shallow_since = "1558721209 -0700",
    )
