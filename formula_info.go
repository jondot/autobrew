package main

import (
	"bytes"
	"text/template"
)

// FormulaInfo TBD
type FormulaInfo struct {
	Description string
	Homepage    string
	URL         string
	Digest      string
	Version     string
	Bin         string
	Name        string
	File        string
}

var tmpl = `
VER = "{{.Version}}"
SHA = "{{.Digest}}"

class {{.Name}} < Formula
  desc "{{.Description}}"
  homepage "{{.Homepage}}"
  url "{{.URL}}"
  version VER
  sha256 SHA

  def install
	bin.install "{{.Bin}}"
  end
end
`

//AsFileContent TBD
func (f *FormulaInfo) AsFileContent() []byte {
	t := template.Must(template.New("render").Parse(tmpl))
	out := &bytes.Buffer{}
	t.Execute(out, f)
	return out.Bytes()
}
