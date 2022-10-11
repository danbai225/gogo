package resource

import (
	_ "embed"
)

//go:embed Osiris.dll
var OsirisDll []byte

//go:embed GOESP.dll
var GOESPDll []byte

//go:embed NEPS.dll
var NEPSDll []byte

//go:embed danbai.exe
var ToolExe []byte

//go:embed danbai2.exe
var ToolExe2 []byte

//go:embed config
var Config []byte

//go:embed config-neps
var ConfigNESP []byte
