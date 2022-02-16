# Contribution Guidelines

You can contribute to awesome-italia-remote with issues and PRs.

## Issues

Issues can be used to ask questions, report errors, or start discussions.  

## Pull Requests

Pull requests can be used to add/edit/delete companies from the list.

### How To Contribute with PRs

The new companies must be added to the `/data` folder using a dedicated file in `snake_case.json` format. The files are divided by category into sub-folders (software, marketing, ecc...).
Subsequently, the `README.md` file is automatically generated, so you don't have to edit it manually.

1. Open `./data` directory
2. Choose the correct company category (software, marketing, ecc...) and enter that directory
3. Add a new JSON file for the new company (file name should be a slugified version of the company name)
4. File content should respect the following format:

```JSON
{
    "name": "CompanyName",
    "career_page_url": "https://companyname.companytld/jobs",
    "url": "https://www.companyname.companytld/",
    "type": "B2B",
    "tags": [
        "PHP",
        "Go",
        "AWS"
    ]
}
```

#### Allowed Company Types

1. B2B
2. B2C
3. Consulting
4. Product

To suggest a new company type, please open an issue.
