<p align="center"><a href="#readme"><img src="https://gh.kaos.st/rsz.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/w/rsz/ci"><img src="https://kaos.sh/w/rsz/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/r/rsz"><img src="https://kaos.sh/r/rsz.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/b/rsz"><img src="https://kaos.sh/b/b1546369-70e1-4a1d-9229-8df3c0e4aabd.svg" alt="Codebeat badge" /></a>
  <a href="https://kaos.sh/w/rsz/codeql"><img src="https://kaos.sh/w/rsz/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="#license"><img src="https://gh.kaos.st/apache2.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#command-line-completion">Command-line completion</a> • <a href="#man-documentation">Man documentation</a> • <a href="#usage">Usage</a> • <a href="#build-status">Build Status</a> • <a href="#contributing">Contributing</a> • <a href="#license">License</a></p>

<br/>

`rsz` is a simple utility for image resizing.

### Installation

#### From source

To build the `rsz` from scratch, make sure you have a working Go 1.16+ workspace (_[instructions](https://golang.org/doc/install)_), then:

```
go install github.com/essentialkaos/rsz
```

#### Prebuilt binaries

You can download prebuilt binaries for Linux and macOS from [EK Apps Repository](https://apps.kaos.st/rsz/latest):

```bash
bash <(curl -fsSL https://apps.kaos.st/get) rsz
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

```
Usage: rsz {options} src-image size output-image

Options

  --filter, -f name     Resampling filter name
  --list-filters, -F    Print list of supported resampling filters
  --no-color, -nc       Disable colors in output
  --help, -h            Show this help message
  --version, -v         Show version

Examples

  rsz image.png 256x256 thumbnail.png
  Convert image to exact size

  rsz -f Lanczos image.png 256x256 thumbnail.png
  Convert image to exact size using Lanczos resampling filter

  rsz image.png 25% thumbnail.png
  Convert image to relative size (25% of original)

  rsz image.png 0.55 thumbnail.png
  Convert image to relative size (55% of original)

```

### Build Status

| Branch | Status |
|--------|----------|
| `master` | [![CI](https://kaos.sh/w/rsz/ci.svg?branch=master)](https://kaos.sh/w/rsz/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/rsz/ci.svg?branch=develop)](https://kaos.sh/w/rsz/ci?query=branch:develop) |

### Contributing

Before contributing to this project please read our [Contributing Guidelines](https://github.com/essentialkaos/contributing-guidelines#contributing-guidelines).

### License

[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
