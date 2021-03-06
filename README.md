<img src="https://raw.githubusercontent.com/scraly/gophers/main/harry-gopher.png" alt="jedi-gopher" width=300 >

<p align="Left">
  <p align="Left">
    <a href="https://github.com/borischen0203/shoris/releases/latest"><img alt="GitHub release" src="https://img.shields.io/github/v/release/borischen0203/shoris.svg?logo=github&style=flat-square"></a>
    <a href="https://github.com/borischen0203/shoris/actions/workflows/go.yml"><img alt="GitHub release" src="https://github.com/borischen0203/shoris/actions/workflows/go.yml/badge.svg?logo=github&style=flat-square"></a>
    <a href="https://goreportcard.com/report/github.com/borischen0203/shoris"><img src="https://goreportcard.com/badge/github.com/borischen0203/shoris" alt="Code Status" /></a>
  </p>
</p>

# shoris
This `shoris` command-line tool is able to shorten a long url.


## Features
- Be able to shorten a long url.
- Be able to custom alias.
- Auto write short URL to clipboard.

## Installation

### On macOS via Homebrew
Step1:
```bash
brew tap borischen0203/shoris
```
Step2:
```bash
brew install shoris
```
Or

One line command
```bash
brew install borischen0203/shoris/shoris
```

## How to use

### Demo example
```bash
# without alias
$ shoris https://www.youtube.com/watch?v=072tU1tamd0
( *・ω・)✄╰ひ╯ https://tiny.one/pm8ap5fx
Short URL has been saved to clipboard!

# with alias
$ shoris https://www.youtube.com/watch?v=072tU1tamd0 hotpot
( *・ω・)✄╰ひ╯ https://tiny.one/hotpot
Short URL has been saved to clipboard!

```

## Tech stack
- Golang
- Cobra



### Todo:
- [X] Custom alias.
- [X] Auto write to clipboard.

