# LGTM

[![Build Status](http://drone.gitea.io/api/badges/go-gitea/lgtm/status.svg)](http://drone.gitea.io/go-gitea/lgtm)
[![Join the chat at https://gitter.im/go-gitea/lgtm](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/go-gitea/lgtm?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![](https://images.microbadger.com/badges/image/gitea/lgtm.svg)](http://microbadger.com/images/gitea/lgtm "Get your own image badge on microbadger.com")
[![Coverage Status](https://coverage.gitea.io/badges/go-gitea/lgtm/coverage.svg)](https://coverage.gitea.io/go-gitea/lgtm)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-gitea/lgtm)](https://goreportcard.com/report/github.com/go-gitea/lgtm)
[![GoDoc](https://godoc.org/github.com/go-gitea/lgtm?status.svg)](https://godoc.org/github.com/go-gitea/lgtm)

LGTM is a simple pull request approval system using GitHub protected branches
and maintainers files or maintainers groups. Pull requests are locked and cannot
be merged until the minimum number of approvals are received. Project
maintainers can indicate their approval by commenting on the pull request and
including LGTM (looks good to me) in their approval text.

## Install

You can download prebuilt binaries from the GitHub releases or from our
[download site](http://dl.gitea.io.de/lgtm). You are a Mac user? Just take
a look at our [homebrew formula](https://github.com/go-gitea/homebrew-gitea).
If you are missing an architecture just write us on our pretty active
[Gitter](https://gitter.im/go-gitea/lgtm) chat. If you find a security issue
please contact security@gitea.io first.

## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions](http://golang.org/doc/install.html).
As this project relies on vendoring of the dependencies and we are not
exporting `GO15VENDOREXPERIMENT=1` within our makefile you have to use a Go
version `>= 1.6`. It is also possible to just simply execute the
`go get github.com/go-gitea/lgtm` command, but we prefer to use our `Makefile`:

```bash
go get -d github.com/go-gitea/lgtm
cd $GOPATH/src/github.com/go-gitea/lgtm
make clean build

bin/lgtm -h
```

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

* [Maintainers](https://github.com/orgs/go-gitea/people)
* [Contributors](https://github.com/go-gitea/gitea/graphs/contributors)

## License

This project is under the Apache-2.0 License. See the [LICENSE](LICENSE) file
for the full license text.

## Copyright

```
Copyright (c) 2016 The Gitea Authors <https://gitea.io>
```
