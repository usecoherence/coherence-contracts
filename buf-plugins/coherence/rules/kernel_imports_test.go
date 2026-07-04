package rules

import (
	"testing"

	"buf.build/go/bufplugin/check"
	"buf.build/go/bufplugin/check/checktest"
)

func validSpec() *checktest.CheckTest {
	return &checktest.CheckTest{
		Request: &checktest.RequestSpec{
			Files: &checktest.ProtoFileSpec{
				DirPaths:  []string{"testdata/valid"},
				FilePaths: []string{"coherence/kernel/v1/kernel_valid.proto"},
			},
		},
		Spec: &check.Spec{
			Rules: []*check.RuleSpec{KernelImportsRuleSpec},
		},
		ExpectedAnnotations: nil,
	}
}

func invalidSpec() *checktest.CheckTest {
	return &checktest.CheckTest{
		Request: &checktest.RequestSpec{
			Files: &checktest.ProtoFileSpec{
				DirPaths:  []string{"testdata/invalid"},
				FilePaths: []string{"coherence/kernel/v1/kernel_bad_import.proto"},
			},
		},
		Spec: &check.Spec{
			Rules: []*check.RuleSpec{KernelImportsRuleSpec},
		},
		ExpectedAnnotations: []checktest.ExpectedAnnotation{
			{
				RuleID: KernelImportsRuleID,
				FileLocation: &checktest.ExpectedFileLocation{
					FileName: "coherence/kernel/v1/kernel_bad_import.proto",
				},
			},
		},
	}
}

func TestKernelImportsValid(t *testing.T) {
	t.Parallel()
	validSpec().Run(t)
}

func TestKernelImportsInvalid(t *testing.T) {
	t.Parallel()
	invalidSpec().Run(t)
}
