// Code generated by wit-bindgen-go. DO NOT EDIT.

package terminalstderr

import (
	"internal/cm"
)

// This file contains wasmimport and wasmexport declarations for "wasi:cli@0.2.0".

//go:wasmimport wasi:cli/terminal-stderr@0.2.0 get-terminal-stderr
//go:noescape
func wasmimport_GetTerminalStderr(result *cm.Option[TerminalOutput])
