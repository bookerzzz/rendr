
example:
	rendr --global='{"GlobName":"myglobal","GlobValue":"is the same everywhere"}' example/data.json example/template.tmpl "example/{{ .Name }}.go"

install:
	go install .

.PHONY: example install
