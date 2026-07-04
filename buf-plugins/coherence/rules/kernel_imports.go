package rules

import (
	"context"
	"strings"

	"buf.build/go/bufplugin/check"
	"buf.build/go/bufplugin/check/checkutil"
	"buf.build/go/bufplugin/descriptor"
)

const KernelImportsRuleID = "COHERENCE_KERNEL_IMPORTS"

var KernelImportsRuleSpec = &check.RuleSpec{
	ID:      KernelImportsRuleID,
	Default: true,
	Purpose: "Checks that files in coherence/kernel/ only import coherence/kernel/ or google/protobuf/.",
	Type:    check.RuleTypeLint,
	Handler: checkutil.NewFileRuleHandler(checkKernelImports, checkutil.WithoutImports()),
}

func checkKernelImports(
	_ context.Context,
	responseWriter check.ResponseWriter,
	_ check.Request,
	fileDescriptor descriptor.FileDescriptor,
) error {
	filePath := fileDescriptor.ProtoreflectFileDescriptor().Path()
	if !strings.HasPrefix(filePath, "coherence/kernel/") {
		return nil
	}
	imports := fileDescriptor.ProtoreflectFileDescriptor().Imports()
	for i := 0; i < imports.Len(); i++ {
		importPath := imports.Get(i).Path()
		if strings.HasPrefix(importPath, "coherence/kernel/") {
			continue
		}
		if strings.HasPrefix(importPath, "google/protobuf/") {
			continue
		}
		responseWriter.AddAnnotation(
			check.WithMessagef(
				"File %q imports %q which is not allowed. Files in coherence/kernel/ may only import coherence/kernel/ or google/protobuf/.",
				filePath,
				importPath,
			),
			check.WithDescriptor(fileDescriptor.ProtoreflectFileDescriptor()),
		)
	}
	return nil
}
