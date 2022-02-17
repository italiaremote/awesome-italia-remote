# Awesome Italia Remote

A list of remote-friendly or full-remote companies that targets Italian talents.  
Companies can be both based in Italy or around the world but with operations in Italy.

Please read the contribution guidelines before opening a pull request or contributing to this repository.

## Companies

Tech stacks are usually incomplete check the careers page of each company to get more details.

### Software and Cloud

Name | Career Page | Company Type | Remote Policy | Hiring Policy | Stack
------------ | -- | ------- | ------- |---------------| -------

{{- range .cloud_software}}
[{{.Name}}]({{.URL}}) | [Career Page]({{.CareerPageURL}}) | {{.Type}} | {{.RemotePolicy}} | {{.HiringPolicy}} | {{.GetTagsString}}
{{- end}}

### Marketing and Writing

Name | Career Page | Remote Policy| Hiring Policy
------------ | -- | ------- | -------
{{- range .marketing_writing}}
[{{.Name}}]({{.URL}}) | [Career Page]({{.CareerPageURL}}) | {{.RemotePolicy}} | {{.HiringPolicy}}
{{- end}}

### HR

Name | Career Page | Remote Policy| Hiring Policy
------------ | -- | ------- | -------
{{- range .hr}}
[{{.Name}}]({{.URL}}) | [Career Page]({{.CareerPageURL}}) | {{.RemotePolicy}} | {{.HiringPolicy}}
{{- end}}

### Design and UX

Name | Career Page | Remote Policy| Hiring Policy
------------ | -- | ------- | -------
{{- range .design_ux}}
[{{.Name}}]({{.URL}}) | [Career Page]({{.CareerPageURL}}) | {{.RemotePolicy}} | {{.HiringPolicy}}
{{- end}}


