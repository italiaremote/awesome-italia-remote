# Awesome Italia Remote

A list of remote-friendly or full-remote companies that targets Italian talents.  
Companies can be both based in Italy or around the world but with operations in Italy.

Please read the contribution guidelines before opening a pull request or contributing to this repository.

## Companies

Tech stacks are usually incomplete check the careers page of each company to get more details.

### Software and Cloud

Name | Website | Company Type | Stack
------------ | ------- | -------| -------

{{- range .cloud_software}}
[{{.Name}}]({{.CareerPageURL}}) | {{.URL}} | {{.Type}} | {{.GetTagsString}}
{{- end}}

### Marketing and Writing

Name | Website
------------ | -------
{{- range .marketing_writing}}
[{{.Name}}]({{.CareerPageURL}}) | {{.URL}} 
{{- end}}

### HR

Name | Website
------------ | -------
{{- range .hr}}
[{{.Name}}]({{.CareerPageURL}}) | {{.URL}} 
{{- end}}

### Design and UX

Name | Website
------------ | -------
{{- range .design_ux}}
[{{.Name}}]({{.CareerPageURL}}) | {{.URL}} 
{{- end}}
