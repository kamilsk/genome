> # 🧬 genome
>
> Build your app by DNA.

[![Build][icon_build]][page_build]

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

Full description of the idea is available [here][design].

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
$ curl -sfL https://bit.ly/install-genome | bash
```

### Source

```bash
# use standard go tools
$ go get -u github.com/kamilsk/genome
# or use egg tool
$ egg github.com/kamilsk/genome -- go install .
$ egg bitbucket.org/kamilsk/genome -- go install .
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

[icon_build]:       https://travis-ci.org/kamilsk/genome.svg?branch=master

[page_build]:       https://travis-ci.org/kamilsk/genome
[page_promo]:       https://github.com/kamilsk/genome

[egg]:              https://github.com/kamilsk/egg
[design]:           https://www.notion.so/octolab/genome-794781effdae4278ac6bc92637cbf74c?r=0b753cbf767346f5a6fd51194829a2f3
