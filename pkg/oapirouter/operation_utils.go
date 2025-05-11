package oapirouter

import v3 "github.com/pb33f/libopenapi/datamodel/high/v3"

const (
	ParameterInPath   = "path"
	ParameterInQuery  = "query"
	ParameterInHeader = "header"
)

func HasNoParameters(op *v3.Operation) bool {
	if op.Parameters == nil || len(op.Parameters) == 0 {
		return true
	}
	return false
}

func HasPathParameters(op *v3.Operation) bool {
	if op.Parameters == nil || len(op.Parameters) == 0 {
		return false
	}
	for _, param := range op.Parameters {
		if param.In == ParameterInPath {
			return true
		}
	}
	return false
}

func HasQueryParameters(op *v3.Operation) bool {
	if op.Parameters == nil || len(op.Parameters) == 0 {
		return false
	}
	for _, param := range op.Parameters {
		if param.In == ParameterInQuery {
			return true
		}
	}
	return false
}

func HasHeaderParameters(op *v3.Operation) bool {
	if op.Parameters == nil || len(op.Parameters) == 0 {
		return false
	}
	for _, param := range op.Parameters {
		if param.In == ParameterInHeader {
			return true
		}
	}
	return false
}

func HasRequiredHeaderParameters(op *v3.Operation) bool {
	if op.Parameters == nil || len(op.Parameters) == 0 {
		return false
	}
	for _, param := range op.Parameters {
		if param.In == ParameterInHeader && *param.Required {
			return true
		}
	}
	return false
}
