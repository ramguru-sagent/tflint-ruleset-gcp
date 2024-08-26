package main

import (
    "github.com/yourusername/tflint-ruleset-gcp/rules"
    "github.com/terraform-linters/tflint-plugin-sdk/plugin"
    "github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        RuleSet: &CustomRuleSet{},
    })
}

type CustomRuleSet struct{}

func (r *CustomRuleSet) Name() string {
    return "custom"
}

func (r *CustomRuleSet) Version() string {
    return "0.1.0"
}

func (r *CustomRuleSet) Rules() []tflint.Rule {
    return []tflint.Rule{
        rules.NewGCPResourceMissingLabelsRule(),
    }
}

