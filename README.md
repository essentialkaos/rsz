<p align="center"><a href="#readme"><img src=".github/images/card.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/y/rsz"><img src="https://kaos.sh/y/ccb1a82d38264e48bd7d2238e493e29e.svg" alt="Codacy" /></a>
  <a href="https://kaos.sh/w/rsz/ci"><img src="https://kaos.sh/w/rsz/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/w/rsz/codeql"><img src="https://kaos.sh/w/rsz/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="#license"><img src=".github/images/license.svg"/></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#command-line-completion">Command-line completion</a> • <a href="#man-documentation">Man documentation</a> • <a href="#usage">Usage</a> • <a href="#ci-status">CI Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

`rsz` is a simple utility for image resizing.

### Installation

#### From source

To build the `rsz` from scratch, make sure you have a working Go 1.23+ workspace (_[instructions](https://go.dev/doc/install)_), then:

```
go install github.com/essentialkaos/rsz@latest
```

#### Prebuilt binaries

You can download prebuilt binaries for Linux and macOS from [EK Apps Repository](https://apps.kaos.st/rsz/latest):

```bash
bash <(curl -fsSL https://apps.kaos.st/get) rsz
```

#### Container Image

The latest version of `rsz` also available as container image on [Docker Hub](https://kaos.sh/d/rsz) and [GitHub Container Registry](https://kaos.sh/p/rsz):

```bash
podman run --rm -it ghcr.io/essentialkaos/rsz:latest image.png 0.55 thumbnail.png
# or
docker run --rm -it ghcr.io/essentialkaos/rsz:latest image.png 0.55 thumbnail.png
```

### Upgrading

Since version `1.2.0` you can update `rsz` to the latest release using [self-update feature](https://github.com/essentialkaos/.github/blob/master/APPS-UPDATE.md):

```bash
rsz --update
```

This command will runs a self-update in interactive mode. If you want to run a quiet update (_no output_), use the following command:

```bash
rsz --update=quiet
```

### Command-line completion

You can generate completion for `bash`, `zsh` or `fish` shell.

Bash:
```bash
sudo rsz --completion=bash 1> /etc/bash_completion.d/rsz
```

ZSH:
```bash
sudo rsz --completion=zsh 1> /usr/share/zsh/site-functions/rsz
```

Fish:
```bash
sudo rsz --completion=fish 1> /usr/share/fish/vendor_completions.d/rsz.fish
```

### Man documentation

You can generate man page using next command:

```bash
rsz --generate-man | sudo gzip > /usr/share/man/man1/rsz.1.gz
```

### Usage

<img src=".github/images/usage.svg" />

### CI Status

| Branch | Status |
|--------|--------|
| `master` | [![CI](https://kaos.sh/w/rsz/ci.svg?branch=master)](https://kaos.sh/w/rsz/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/rsz/ci.svg?branch=develop)](https://kaos.sh/w/rsz/ci?query=branch:develop) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/.github/blob/master/CONTRIBUTING.md).

### License

[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://kaos.dev"><img src="https://raw.githubusercontent.com/essentialkaos/.github/refs/heads/master/images/ekgh.svg"/></a></p>
