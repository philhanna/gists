# Gists
[![Go Report Card](https://goreportcard.com/badge/github.com/philhanna/gists)][idGoReportCard]
[![PkgGoDev](https://pkg.go.dev/badge/github.com/philhanna/gists)][idPkgGoDev]


A program to download all the user's gists from Github

## Installation

Download the Git repository and install the executables
```bash
git clone git@github.com:/philhanna/gists
cd gists
go install cmd/download/gists_download.go
go install cmd/fs/gists_to_files.go
```

Create the configuration file:
- On Linux/MacOS: `$HOME/.config/gists/config.yaml`
- On Windows: `%APPDATA%\gists\config.yaml`
with these parameters:
```yaml
username: <your userID>
token: <a github token for this user>
```

See the [Github documentation](https://docs.github.com/en/enterprise-server@3.9/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens) for
details about how to create a Github access token.

[idGoReportCard]: https://goreportcard.com/report/github.com/philhanna/gists
[idPkgGoDev]: https://pkg.go.dev/github.com/philhanna/gists
