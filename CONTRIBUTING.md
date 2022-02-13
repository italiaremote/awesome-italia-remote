# Contribution Guidelines

You can contribute to awesome-italia-remote with issues and PRs.

## Issues

Issues can be used to ask questions, report errors, or start discussions.  

## Pull Requests

Pull requests can be used to add/edit/delete companies from the list.

### How To Contribute with PRs

Companies must be added to `./src/data.json` file using the JSON format.  
The `README.md` file is automatically generated, so it should be left untouched.  

1. Open `./src/data.json`
2. Add the company

```JSON
{
    "name": "CompanyName",
    "career_page_url": "https://companyname.companytld/jobs",
    "url": "https://www.companyname.companytld/",
    "type": "Product",
    "tags": [
        "PHP",
        "Go",
        "AWS"
    ]
}
```

#### Allowed Company Types

1. Product
2. Consulting

To suggest a new company type, please open an issue.
