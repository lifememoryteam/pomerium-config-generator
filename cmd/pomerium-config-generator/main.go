package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/lifememoryteam/pomerium-config-generator/pkg/policy"
	"github.com/lifememoryteam/pomerium-config-generator/pkg/team"
	"github.com/lifememoryteam/pomerium-config-generator/pkg/template"
)

var (
	usage = `Usage:
  ./pomerium-config-generator --config=<config.yaml.tmpl> --policy=<policy.yaml> --team=<team.yaml> --output=<output_file>
`

	configPath = flag.String("config", "config.yaml.tmpl", "template for config.yaml")
	policyPath = flag.String("policy", "policy.yaml", "policy file")
	teamPath   = flag.String("team", "team.yaml", "team file")
	outputPath = flag.String("output", "output.yaml", "output file")
	stdoutFlag = flag.Bool("stdout", false, "only output for stdout")
)

func main() {
	if err := generateConfig(); err != nil {
		fmt.Fprintln(os.Stderr, usage)
		log.Fatal(err)
	}
}

func generateConfig() error {
	flag.Parse()

	po, err := policy.Load(*policyPath)
	if err != nil {
		return err
	}

	te, err := team.Load(*teamPath)
	if err != nil {
		return err
	}

	p := po.Generate(te)

	policyBody, err := p.Yaml()
	if err != nil {
		return err
	}

	out := template.Concat(string(policyBody), *configPath)

	if !*stdoutFlag {
		err = ioutil.WriteFile(*outputPath, out, 0666)
		if err != nil {
			return err
		}
	} else {
		fmt.Println(string(out))
	}

	return nil
}
