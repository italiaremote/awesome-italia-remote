# Awesome Italia Remote

A list of remote-friendly or full-remote companies that targets Italian talents.  
Companies can be both based in Italy or around the world but with operations in Italy.

Please read the contribution guidelines before opening a pull request or contributing to this repository.

## Companies

Tech stacks are usually incomplete check the careers page of each company to get more details.

### Software and Cloud

Name | Website | Company Type | Remote Policy | Hiring Policy | Stack
------------ | ------- | ------- | ------- |---------------| -------

{{- range .cloud_software}}
[{{.Name}}]({{.CareerPageURL}}) | {{.URL}} | {{.Type}} | {{.RemotePolicy}} | {{.HiringPolicy}} | {{.GetTagsString}}
{{- end}}

### Marketing and Writing

Name | Website | Remote Policy| Hiring Policy
------------ | ------- | ------- | -------
{{- range .marketing_writing}}
[{{.Name}}]({{.CareerPageURL}}) | {{.URL}} | {{.RemotePolicy}} | {{.HiringPolicy}}
{{- end}}

### HR

Name | Website | Remote Policy| Hiring Policy
------------ | ------- | ------- | -------
{{- range .hr}}
[{{.Name}}]({{.CareerPageURL}}) | {{.URL}} | {{.RemotePolicy}} | {{.HiringPolicy}}
{{- end}}

### Design and UX

Name | Website | Remote Policy| Hiring Policy
------------ | ------- | ------- | -------
{{- range .design_ux}}
[{{.Name}}]({{.CareerPageURL}}) | {{.URL}} | {{.RemotePolicy}} | {{.HiringPolicy}}
{{- end}}
