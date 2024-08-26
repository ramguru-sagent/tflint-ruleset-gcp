package rules

import (
    "github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type GCPResourceMissingLabelsRule struct{}

func NewGCPResourceMissingLabelsRule() *GCPResourceMissingLabelsRule {
    return &GCPResourceMissingLabelsRule{}
}

func (r *GCPResourceMissingLabelsRule) Name() string {
    return "gcp_resource_missing_labels"
}

func (r *GCPResourceMissingLabelsRule) Enabled() bool {
    return true
}

func (r *GCPResourceMissingLabelsRule) Severity() string {
    return tflint.ERROR
}

func (r *GCPResourceMissingLabelsRule) Link() string {
    return ""
}

func (r *GCPResourceMissingLabelsRule) Check(runner tflint.Runner) error {
    return runner.WalkResources("google_*", func(resource *tflint.Resource) error {
        if labels := resource.GetAttribute("labels"); labels == nil {
            runner.EmitIssue(
                r,
                "The resource is missing required labels.",
                resource.DeclRange,
            )
        }
        return nil
    })
}

