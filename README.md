# pomerium-config-generator

Config generator for Pomerium with Role-based Access Control locally.

Pomerium supports filtering access with IdP-defined group, but we want to define roles locally on the file.

## Usage

```shell
$ docker run -v $(pwd):/data docker.pkg.github.com/lifememoryteam/pomerium-config-generator/pomerium-config-generator:latest pomerium-config-generator --config=/data/testdata/config.yaml.tmpl --policy=/data/testdata/policy.yaml --team=/data/testdata/team.yaml --output=/data/testdata/output.yaml
```

## File structure
- config.yaml.tmpl
  - Template with Pomerium Configuration
  - To merge policy.yaml and team.yaml automatically.
- policy.yaml
  - Put Pomerium Proxy Policy with `pomerium`
  - Put pomerium-config-generator Configuration with `poerium-config-generator`
    - Supports only `group`
- team.yaml
  - Role / Group Definition
