package main

import (
	"buf.build/go/bufplugin/check"
	"buf.build/go/bufplugin/info"

	"github.com/usecoherence/coherence-contracts/buf-plugins/coherence/rules"
)

var spec = &check.Spec{
	Rules: []*check.RuleSpec{
		rules.KernelImportsRuleSpec,
	},
	Info: &info.Spec{
		Documentation: "Buf check plugin for Coherence protocol contracts.",
		SPDXLicenseID: "Apache-2.0",
		LicenseURL:    "https://github.com/usecoherence/coherence-contracts/blob/main/LICENSE",
	},
}

func main() {
	check.Main(spec)
}
