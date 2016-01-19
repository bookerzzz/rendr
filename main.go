package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/codegangsta/cli"
)

func resolveArgs(args []string) (dataFile, tmplFile, outFile string) {
	r := []*string{
		&dataFile,
		&tmplFile,
		&outFile,
	}
	for i, a := range args {
		*r[i] = a
	}
	return
}

func render(tmpl *template.Template, d map[string]interface{}, w io.Writer) {
	err := tmpl.Execute(w, d)
	if err != nil {
		fmt.Printf("Unable to execute template with error '%s'\n", err.Error())
		return
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "rendr"
	app.Usage = `
  Rendr a template with a list of datasets contained in a json file and output to files
    rendr data.json template.tmpl "result.{{ .PossibleDataKey }}.out"
  `
	app.EnableBashCompletion = true
	app.Action = func(c *cli.Context) {
		dataFile, tmplFile, outFile := resolveArgs(c.Args())
		js, err := ioutil.ReadFile(dataFile)
		if err != nil {
			fmt.Printf("Unable to read data file '%s' with error '%s'\n", dataFile, err.Error())
			return
		}

		tmpl, err := template.ParseFiles(tmplFile)
		if err != nil {
			fmt.Printf("Unable to read template file '%s' with error '%s'\n", tmplFile, err.Error())
			return
		}

		var dl []map[string]interface{}
		err = json.Unmarshal(js, &dl)
		if err != nil {
			fmt.Printf("Unable to unmarshal json with error '%s'\n", err.Error())
			return
		}

		outFileTmpl, err := template.New("outFileTmpl").Parse(outFile)
		if err != nil {
			fmt.Printf("Unable to parse outfile name as template with error '%s'\n", err.Error())
			return
		}

		fwritten := map[string]bool{}

		for _, d := range dl {
			if outFile != "" {
				ofb := bytes.Buffer{}
				err := outFileTmpl.Execute(&ofb, d)
				if err != nil {
					fmt.Printf("Unable to resolve outfile name with error '%s'\n", err.Error())
					return
				}
				fn := ofb.String()
				flag := os.O_WRONLY | os.O_CREATE | os.O_APPEND
				if !fwritten[fn] {
					flag = flag | os.O_TRUNC
				}
				fp, err := os.OpenFile(fn, flag, 0666)
				if err != nil {
					fmt.Printf("Unable to write to file '%s' with error '%s'\n", fn, err.Error())
					return
				}
				fwritten[fn] = true
				render(tmpl, d, fp)
				fp.Close()
			} else {
				render(tmpl, d, os.Stdout)
			}
		}
	}

	app.Run(os.Args)
}
