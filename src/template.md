# Awesome Italia Remote

A list of remote-friendly or full-remote companies that targets Italian talents.  
Companies can be both based in Italy or around the world but with operations in Italy.

Do you want to insert/remove/modify an entry, open a PR/Issue or contact me on my socials [Twitter](https://twitter.com/alessmarinoac)

## Companies

Tech stacks are usually incomplete check the careers page of each company to get more details.

### Software and Cloud

Name | Website | Company Type | Stack
------------ | ------- | -------| -------

{{- range .cloud_software}}
[{{.Name}}]({{.CareerPageURL}}) | {{.URL}} | {{.Type}} | {{.TagsString}}
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
