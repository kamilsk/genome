> # 🧬 genome
>
> Build your app by DNA.

[![Build][build.icon]][build.page]
[![Template][template.icon]][template.page]

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

Full description of the idea is available [here][design.page].

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
$ curl -sSfL https://bit.ly/install-genome | sh
# or
$ wget -qO-  https://bit.ly/install-genome | sh
```

### Source

```bash
# use standard go tools
$ go get github.com/kamilsk/genome@latest
# or use egg tool
$ egg tools add github.com/kamilsk/genome@latest
```

> [egg][]<sup id="anchor-egg">[1](#egg)</sup> is an `extended go get`.

### Bash and Zsh completions

```bash
$ genome completion bash > /path/to/bash_completion.d/genome.sh
$ genome completion zsh  > /path/to/zsh-completions/_genome.zsh
```

<sup id="egg">1</sup> The project is still in prototyping.[↩](#anchor-egg)

---

made with ❤️ for everyone

[build.page]:       https://travis-ci.org/kamilsk/genome
[build.icon]:       https://travis-ci.org/kamilsk/genome.svg?branch=master
[design.page]:      https://www.notion.so/octolab/genome-794781effdae4278ac6bc92637cbf74c?r=0b753cbf767346f5a6fd51194829a2f3
[promo.page]:       https://github.com/kamilsk/genome
[template.page]:    https://github.com/octomation/go-tool
[template.icon]:    https://img.shields.io/badge/template-go--tool-blue

[egg]:              https://github.com/kamilsk/egg
