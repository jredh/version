# Golang Versioning Tool

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