// Copyright (c) 2015 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package compile

import "github.com/uber/thriftrw-go/ast"

// Module represents a compiled Thrift module. It contains all information
// about all known types, constants, services, and includes from the Thrift
// file.
//
// ThriftPath is the path to the Thrift file from which this module was
// compiled. All includes made by the Thrift file are relative to that path.
//
// The module name is usually just the basename of the ThriftPath.
type Module struct {
	Name       string
	ThriftPath string

	// Mapping from the /Thrift name/ to the compiled representation of
	// different definitions.

	Includes  map[string]*IncludedModule
	Constants map[string]*Constant
	Types     map[string]TypeSpec
	Services  map[string]*ServiceSpec
}

// LookupType for Module.
func (m *Module) LookupType(name string) (TypeSpec, error) {
	if t, ok := m.Types[name]; ok {
		return t, nil
	}

	mname, iname := splitInclude(name)
	if len(mname) == 0 {
		return nil, lookupError{Name: name, ModuleName: m.Name}
	}

	included, ok := m.Includes[mname]
	if !ok {
		return nil, lookupError{
			Name:       name,
			ModuleName: m.Name,
			Reason:     unrecognizedModuleError{Name: mname},
		}
	}

	spec, err := included.Module.LookupType(iname)
	if err != nil {
		return nil, lookupError{
			Name:       name,
			ModuleName: m.Name,
			Reason:     err,
		}
	}

	return spec, nil
}

// lookupEnum looks up an enum with the given name.
//
// Return the enum and true/false indicating whether a matching enum was
// actually found.
func (m *Module) lookupEnum(name string) (*EnumSpec, bool) {
	t, err := m.LookupType(name)
	if err != nil {
		return nil, false
	}

	if enum, ok := t.(*EnumSpec); ok {
		return enum, true
	}
	return nil, false
}

// LookupConstant for Module.
func (m *Module) LookupConstant(name string) (ast.ConstantValue, error) {
	if c, ok := m.Constants[name]; ok {
		return c.Value, nil
	}

	mname, iname := splitInclude(name)
	if len(mname) == 0 {
		return nil, lookupError{Name: name, ModuleName: m.Name}
	}

	// First check if we have an enum that matches
	if enum, ok := m.lookupEnum(mname); ok {
		if item, ok := enum.LookupItem(iname); ok {
			return ast.ConstantInteger(item.Value), nil
		}

		return nil, lookupError{
			Name:       name,
			ModuleName: m.Name,
			Reason: unrecognizedEnumItemError{
				EnumName: mname,
				ItemName: iname,
			},
		}
	}

	// Then check includes.
	included, ok := m.Includes[mname]
	if !ok {
		return nil, lookupError{
			Name:       name,
			ModuleName: m.Name,
			Reason:     unrecognizedModuleError{Name: mname},
		}
	}

	c, err := included.Module.LookupConstant(iname)
	if err != nil {
		return nil, lookupError{
			Name:       name,
			ModuleName: m.Name,
			Reason:     err,
		}
	}

	return c, nil
}

// LookupService for Module.
func (m *Module) LookupService(name string) (*ServiceSpec, error) {
	if s, ok := m.Services[name]; ok {
		return s, nil
	}

	mname, iname := splitInclude(name)
	if len(mname) == 0 {
		return nil, lookupError{Name: name, ModuleName: m.Name}
	}

	included, ok := m.Includes[mname]
	if !ok {
		return nil, lookupError{
			Name:       name,
			ModuleName: m.Name,
			Reason:     unrecognizedModuleError{Name: mname},
		}
	}

	s, err := included.Module.LookupService(iname)
	if err != nil {
		return nil, lookupError{
			Name:       name,
			ModuleName: m.Name,
			Reason:     err,
		}
	}

	return s, nil
}

// IncludedModule represents an included module in the Thrift file.
//
// The name of the IncludedModule is the name under which the module is
// exposed in the Thrift file which included it. This is usually the same as
// the Module name except when our custom include-as syntax is used.
type IncludedModule struct {
	Name   string
	Module *Module
}
