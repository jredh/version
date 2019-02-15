# Golang Versioning Tool

## What it does

- Updates your version file (it will not create one since this operates upon working directories and that'd be weird)
- Adds the version tag to the commit to enable automatic release detection (you will need to push tags)

## Instructions

1. Have a `version` file in your working directory
2. Install the package: `$ go install github.com/thejaredhooper/version`
3. Execute the command against your repo or add it to your githook
4. Usage:
```
$ version -help
```

Note: In order for this to work properly you need to push `tags` as well. To do this you can either

`git push --follow-tags`

or enable it permanently

`git config --global push.followTags true`
