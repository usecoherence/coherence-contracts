// Package rules contains custom Buf lint rules for Coherence protobuf contracts.
package rules

import (
	"context"
	"slices"
	"strings"

	"buf.build/go/bufplugin/check"
	"buf.build/go/bufplugin/check/checkutil"
	"buf.build/go/bufplugin/descriptor"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	KernelImportsRuleID = "COHERENCE_KERNEL_IMPORTS"

	kernelPrefix       = "coherence/kernel/"
	googleProtoPrefix  = "google/protobuf/"
)

var (
	allowedKernelImportPrefixes = []string{
		kernelPrefix,
		googleProtoPrefix,
	}

	allowedKernelImportsText = strings.Join(allowedKernelImportPrefixes, ", ")
	kernelPathText           = strings.TrimSuffix(kernelPrefix, "/")

	KernelImportsRuleSpec = &check.RuleSpec{
		ID:      KernelImportsRuleID,
		Default: true,
		Type:    check.RuleTypeLint,
		Purpose: "Checks that files in " + kernelPathText + " only import " + allowedKernelImportsText + ".",
		Handler: checkutil.NewFileRuleHandler(checkKernelImports, checkutil.WithoutImports()),
	}
)

func isAllowedKernelImport(s string) bool {
	return slices.ContainsFunc(allowedKernelImportPrefixes, func(p string) bool {
		return strings.HasPrefix(s, p)
	})
}

func invalidImportPaths(imports protoreflect.FileImports) []string {
	var paths []string
	for i := 0; i < imports.Len(); i++ {
		path := imports.Get(i).Path()
		if !isAllowedKernelImport(path) {
			paths = append(paths, path)
		}
	}
	return paths
}

func checkKernelImports(
	_ context.Context,
	responseWriter check.ResponseWriter,
	_ check.Request,
	fileDescriptor descriptor.FileDescriptor,
) error {
	filePath := fileDescriptor.ProtoreflectFileDescriptor().Path()
	if !strings.HasPrefix(filePath, kernelPrefix) {
		return nil
	}
	for _, importPath := range invalidImportPaths(fileDescriptor.ProtoreflectFileDescriptor().Imports()) {
		responseWriter.AddAnnotation(
			check.WithMessagef(
				"File %q imports %q which is not allowed. Files in %s may only import %s.",
				filePath,
				importPath,
				kernelPathText,
				allowedKernelImportsText,
			),
			check.WithDescriptor(fileDescriptor.ProtoreflectFileDescriptor()),
		)
	}
	return nil
}
