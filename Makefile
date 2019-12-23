.PHONY: build
build:
	bazel build //cmd/pomerium-config-generator

run:
	bazel run //cmd/pomerium-config-generator -- --config=`pwd`/testdata/config.yaml.tmpl --policy=`pwd`/testdata/policy.yaml --team=`pwd`/testdata/team.yaml --output=`pwd`/testdata/output.yaml

push:
	bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 push
