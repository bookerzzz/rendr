# Rendr
Execute any (Go) template with a given dataset list (JSON) and output results to stdout or to file(s)

The intention of this tool is to output multiple similar files based on a given template with differences provided as
attributes in the array of data objects in the given data.json file.

# Install
go install github.com/bookerzzz/rendr

# Usage
See the example folder for a complete overview.
```bash
rendr [--global(-g)='{"GlobalKey":"GlobalValue"}'] path/to/data.json path/to/template.tmpl path/to/output/dir/filename.ext
```

You may also use data values from the data.json file entries so if your data contains
```js
[
  {
    "Dir", "my/output/dir",
    "Filename": "myoutputfile.txt"
  }
]
```
you can run
```bash
rendr path/to/data.json path/to/template.tmpl "{{ .Dir }}/{{ .Filename }}"
```
to create `./my/output/dir/myoutputfile.txt` with results of `template.tmpl` after rendering.
If the path to the output file doesn't exists, Rendr will attempt to create it.

You may also use/inject the global values into the data.json file, so render will first apply global values to the given data.json file before parsing it and providing the resolved set of data to the template for execution.
```bash
rendr [--global(-g)='{"FileExt":"txt"}'] data.json path/to/template.tmpl path/to/output/dir/filename.ext
```
```js
// data.json
[
  {
    "Dir", "my/output/dir",
    "Filename": "myoutputfile.{{ .FileExt }}"
  }
]
```
