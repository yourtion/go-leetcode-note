# {{.start}} - {{.end}}

## 本周收获

{{range .data}}{{.Harvest}}{{end}}
{{range .data}}
## {{.Problem.Pid}}. {{.Problem.Name}}

[{{.Problem.Url}}]({{.Problem.Url}})

### 思考

{{.Solution}}

{{.Submissions}}

[{{.Submissions}}]({{.Submissions}})

### 反思

{{.Rethink}}

{{end}}