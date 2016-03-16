
example:
	rendr --global='{"DefaultUnit":"cm"}' example/data.json example/template.tmpl "example/{{ .FileName }}"

install:
	go install .

.PHONY: example install
