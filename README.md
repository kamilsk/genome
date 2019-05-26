> # 🧬 genome
>
> Build your app by DNA.

[![Build Status][icon_build]][page_build]

## 💡 Idea

```bash
$ genome clone service internal/service/dolly <<YAML
kind: sheep
code: 6LLS
life:
  born: 5 July 1996
  died: 14 February 2003
YAML
```

Full description of the idea is available
[here](https://www.notion.so/octolab/genome-794781effdae4278ac6bc92637cbf74c?r=0b753cbf767346f5a6fd51194829a2f3).

## 🏆 Motivation

...

## 🤼‍♂️ How to

...

## 🧩 Installation

### Homebrew

```bash
$ brew install kamilsk/tap/genome
```

### Binary

```bash
$ REQ_VER=0.0.1  # all available versions are on https://github.com/kamilsk/genome/releases/
$ REQ_OS=Linux   # macOS is also available
$ REQ_ARCH=64bit # 32bit is also available
# wget -q -O genome.tar.gz
$ curl -sL -o genome.tar.gz \
       https://github.com/kamilsk/genome/releases/download/"${REQ_VER}/genome_${REQ_VER}_${REQ_OS}-${REQ_ARCH}".tar.gz
$ tar xf genome.tar.gz -C "${GOPATH}"/bin/ && rm genome.tar.gz
```

### From source code

```bash
# using standard go tools
$ go get -u github.com/kamilsk/genome
# or using egg tool
$ egg github.com/kamilsk/genome -- go install .
# with mirror
$ egg bitbucket.org/kamilsk/genome -- go install .
```

> [egg][page_egg]<sup id="anchor-egg">[1](#egg)</sup> is an `extended go get`.

### Bash and Zsh completions

```bash
$ genome completion -f bash > /path/to/bash_completion.d/genome.sh
$ genome completion -f zsh  > /path/to/zsh-completions/_genome.zsh
```

<sup id="egg">1</sup> The project is still in prototyping.[↩](#anchor-egg)

---

made with ❤️ for everyone

[icon_build]:      https://travis-ci.org/kamilsk/genome.svg?branch=master

[page_build]:      https://travis-ci.org/kamilsk/genome
[page_promo]:      https://github.com/kamilsk/genome
[page_egg]:        https://github.com/kamilsk/egg
