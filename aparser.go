package aparser

import (
	"flag"
	"github.com/pkg/errors"
)

type Argument struct {
	Flags []string

	// Required 설정되면
	//   - Default = nil
	Required bool

	Description string

	// Default 설정되면
	//   - Required = false
	Default string

	ConfigKey string

	// Mandatory 설정되면
	//  - Required = true
	//  - Flags = nil
	//  - Default = nil
	//  - ConfigKey = nil
	Mandatory string

	EnvironmentKey string
}

type AParser struct {
	arguments             []*Argument
	essentials            []string
	essentialDescriptions []string
}

func (p *AParser) Parse(args []string) (map[string]*string, error) {
	if len(args) == 0 {
		return nil, errors.New("args is empty")
	}

	fs := flag.NewFlagSet(args[0], flag.ExitOnError)
	result := map[string]*string{}

	for _, argument := range p.arguments {
		var value string
		for _, name := range argument.Flags {
			result[name] = &value
			fs.StringVar(&value, name, argument.Default, argument.Description)
		}
	}

	err := fs.Parse(args[1:])
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(p.essentials) > len(fs.Args()) {
		return nil, errors.New("insufficient arguments")
	}

	if len(p.essentials) < len(fs.Args()) {
		return nil, errors.New("excessive arguments")
	}

	for idx, key := range p.essentials {
		result[key] = &fs.Args()[idx]
	}

	// check required
	for _, argument := range p.arguments {
		if argument.Required {
			if *result[argument.Flags[0]] == "" {
				return nil, errors.New("argument is not set")
			}
		}
	}

	return result, nil
}

func NewAParser(arguments []*Argument) *AParser {
	ret := AParser{}
	for _, argument := range arguments {
		if argument.Mandatory == "" {
			ret.arguments = append(ret.arguments, argument)
		} else {
			ret.essentials = append(ret.essentials, argument.Mandatory)
			ret.essentialDescriptions = append(ret.essentialDescriptions, argument.Description)
		}
	}
	return &ret
}
