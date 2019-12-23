package policy

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/lifememoryteam/pomerium-config-generator/pkg/team"
)

type Data struct {
	Policy []RawPolicy `yaml:"policy"`
}

type RawPolicy struct {
	Pomerium map[string]interface{}     `yaml:"pomerium"`
	Policy   RawPomeriumConfigGenerator `yaml:"pomerium_config_generator"`
}

type RawPomeriumConfigGenerator struct {
	Group []string `yaml:"group"`
}

type Value map[string]interface{}

type Policy struct {
	Policy []Value `yaml:"policy"`
}

func Load(file string) (*Data, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var d Data
	if err := yaml.Unmarshal(buf, &d); err != nil {
		return nil, err
	}
	return &d, err
}

func (p *Data) Generate(team *team.Team) *Policy {
	po := Policy{}

	for _, policy := range p.Policy {
		val := Value{}
		for k, v := range policy.Pomerium {
			val[k] = v
		}

		var users []string
		for _, t := range policy.Policy.Group {
			if v, ok := (*team)[t]; ok {
				users = append(users, v.Member...)
			}
		}

		val["allowed_users"] = users
		po.Policy = append(po.Policy, val)
	}

	return &po
}

func (p *Policy) Yaml() ([]byte, error) {
	v, err := yaml.Marshal(p)
	if err != nil {
		return nil, err
	}
	return v, nil
}
