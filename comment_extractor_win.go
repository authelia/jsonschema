//go:build windows

package jsonschema

import (
	"path"
	"strings"
)

func pathJoinModule(elem ...string) string {
	size := 0

	for i, e := range elem {
		size += len(e)

		if strings.ContainsRune(e, '\\') {
			elem[i] = strings.ReplaceAll(e, `\`, `/`)
		}
	}

	if size == 0 {
		return ""
	}

	buf := make([]byte, 0, size+len(elem)-1)

	for _, e := range elem {
		if len(buf) > 0 || e != "" {
			if len(buf) > 0 {
				buf = append(buf, '/')
			}
			buf = append(buf, e...)
		}
	}

	return path.Clean(string(buf))
}
