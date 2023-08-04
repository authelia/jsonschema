//go:build !windows

package jsonschema

import (
	"path"
	"strings"
)

func pathJoinModule(base, subpath string) string {
	return path.Join(base, strings.ReplaceAll(subpath, `\`, `/`))
}
