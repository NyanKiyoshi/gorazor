// This file is generated by gorazor 1.2.1
// DON'T modified manually
// Should edit source file and re-generate: cases/double_quote.gohtml

package cases

import (
	"io"
	"strings"
)

// Double_quote generates cases/double_quote.gohtml
func Double_quote() string {
	var _b strings.Builder
	RenderDouble_quote(&_b)
	return _b.String()
}

// RenderDouble_quote render cases/double_quote.gohtml
func RenderDouble_quote(_buffer io.StringWriter) {
	_buffer.WriteString("<meta charset=\"utf-8\" />")

}
