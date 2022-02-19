# Awesome Italia Remote

A list of remote-friendly or full-remote companies that targets Italian talents.  
Companies can be both based in Italy or around the world but with operations in Italy.

Please read the contribution guidelines before opening a pull request or contributing to this repository.

## Legenda

### Remote Policy

Label | Meaning
--- | ---
Full | Company doesn't have physical offices, so you'll always work remotely.
Hybrid | Company allows remote but only for some days.
Optional | Company allows you to choose when work remotely or in office, but can ask you to go sometimes.

### Hiring Policy

Label | Meaning
--- | ---
Direct | Company is hiring directly with a legal entity in Italy.
Contract | Company is hiring contractors in Italy, VAT Number is required.
Intermediary | Company is hiring using a payroll intermediary in Italy.

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




