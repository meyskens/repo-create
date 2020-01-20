Repo Creator
============

This is a small tool to quickly prepare GitHub repositories, bootstrapped with files, projects and labels.
It is meant to be used in education to let students work under a specific GitHub org with some files, labels and settings pre-setup.

## Usage
This is a CLI tool built in Go. Currently there are no binaries available due to time shortage. PRs to add Go Releaser very welcome!

## Auth
This tool uses a GitHub personal token, set in env vars (or `--auth-token` flag).
```console
$ export AUTH_TOKEN=<token>
$ repo-create
```

### Create repositories
This command will create empty repositories in a given org with a given prefix, these can also be set to private.

Example:
```console
$ repo-create create --org itfactory-tm -n 22 --prefix Keuzeproject1-MIN  --private
```

### Clone to repositories
Clone will clone a repo from a given source and will automatically push the content to the empty repositories.

Example:
```console
$ repo-create clone --org itfactory-tm -n 22 --prefix Keuzeproject1-MIN --source https://github.com/itfactory-tm/Keuzeproject1_MIN.git
```

### Add labels
This will create a new label in all repositories with a given color and name

Example:
```console
$ repo-create label --org itfactory-tm -n 22 --prefix Keuzeproject1-MIN --name "21:45 - 23:15" --color "ffffff"
```

### Remove labels
This will remove a label in all repositories with a given name, meant to remove unneeded defaults

Example:
```console
$ repo-create rm-label --org itfactory-tm -n 22 --prefix Keuzeproject1-MIN --name "bug"
```


### Branch protection
This will enable branch protection to enforce code review

Example:
```console
$ repo-create protect --org itfactory-tm -n 22 --prefix Keuzeproject1-MIN
```

### Add projects
This will add a project to the repo prefilled with empty collumns

Example:
```console
$ repo-create project --org itfactory-tm -n 22 --prefix Keuzeproject1-MIN --name Kanban --collumns todo,doing,done
```