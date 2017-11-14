package reports

import (
	"os"
	"text/template"
)

// Report generates
type Reports struct {
	TotalCount int
	Pass int
	TestName []map[string][]string
}

func (rep *Reports) FaildCount() int{
	return rep.TotalCount - rep.Pass
}

const report_tesmp = `
{{.Pass}} Passed
{{.FaildCount}} Failed
------------
{{.TotalCount}} Requests
------------

---------------------------------------------------------------------

{{range $test := .TestName}} {{range $key, $value := $test}}{{ index $test $key 1}} : {{$key}} => {{ index $test $key 0}}{{end}}
{{end}}
`

func GeneratesReport(report *Reports) {

	tmpl, err := template.New("test").Parse(report_tesmp)
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, report)
	if err != nil { panic(err) }

}