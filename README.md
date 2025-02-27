# iterex

[![Tests][test-badge-img]][test-badge-url]
[![Coverage][codecov-badge-img]][codecov-badge-url]

```
iter
···re·
    ex
```

Iterator-based regular expressions for Go.

## TL;DR

1. Change `regexp` to `iterex`
2. Change `All` to `Each`
3. Iterate!

This package is like the standard `regexp` with two differences:

- it returns iterators (instead of all matches at once)
- limits are optional (instead of having to specify `-1`)

## Requirements

Go 1.23 or higher - since the `range` keyword accepts iterator functions:

- https://tip.golang.org/doc/go1.23#language

## Documentation

Detailed documentation:

- https://pkg.go.dev/github.com/rsp/iterex

## Quick start

### Install

`go get github.com/rsp/iterex`

### Import

```go
import "github.com/rsp/iterex"`
```

### Use

Instead of:

```go
re := regexp.MustCompile(pattern)
```

you call:

```go
ir := iterex.MustCompile(pattern)
```

And instead of:

```go
slice := re.FindAllString(str, -1)
```

you call:

```go
iterator := re.FindAllString(str)
```

(limit is optional, defaults to `-1`)

Then you can iterate:

```go
for s := range iterator {
  fmt.Println(s)
}
```

## Introduction

`iterex` provides lazy version of "All" receiver functions from
standard `regexp` package.

Instead of "All" they have "Each" in their names because instead of
returing all results at once, they return iterators that iterate over each result.

## Constructors

This package provides 4 ways to compile a pattern,
just like the standard `regexp` package:

- Return error:
  - `ir, err := iterex.Compile("...")`
  - `ir, err := iterex.CompilePOSIX("...")`
- Panic on error:
  - `ir := iterex.MustCompile("...")`
  - `ir := iterex.MustCompilePOSIX("...")`

## Receiver functions

For every `regexp` receiver function with `All` in the name,
it provides a function with `Each` that returns an iterator instead of all results at once.

### Byte slice functions

`var b []byte`

- `ir.FindEach(b, n)` - iterator version of `re.FindAll(b, n)`
- `ir.FindEachIndex(b, n)` - iterator version of `re.FindAllIndex(b, n)`
- `ir.FindEachSubmatch(b, n)` - iterator version of `re.FindAllSubmatch(b, n)`
- `ir.FindEachSubmatchIndex(b, n)` - iterator version of `re.FindAllSubmatchIndex(b, n)`

Note that unlike in `regexp` the `n` is optional:

- `ir.FindEach(b)` is the same as `ir.FindEach(b, -1)`

### String functions

`var s string`

- `ir.FindEachString(s, n)` - iterator version of `re.FindAllString(s, n)`
- `ir.FindEachStringIndex(s, n)` - iterator version of `re.FindAllStringIndex(s, n)`
- `ir.FindEachStringSubmatch(s, n)` - iterator version of `re.FindAllStringSubmatch(s, n)`
- `ir.FindEachStringSubmatchIndex(s, n)` - iterator version of `re.FindAllStringSubmatchIndex(s, n)`

Note that unlike in `regexp` the `n` is optional:

- `ir.FindEachString(s)` is the same as `ir.FindAllString(s, -1)`

<!-- ## Issues

For any bug reports or feature requests please
[post an issue on GitHub][issues-url]. -->

## Author

Rafał Pocztarski - <img src="https://cdn.jsdelivr.net/npm/simple-icons@v13/icons/github.svg" alt="GitHub" width="16" height="16" style="vertical-align:middle"> [rsp][github-follow-url]

## License

MIT License (Expat). See [LICENSE.md](LICENSE.md) for details.

[github-url]: https://github.com/rsp/iterex
[github-logo]: https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png
[readme-url]: https://github.com/rsp/iterex#readme
[issues-url]: https://github.com/rsp/iterex/issues
[license-url]: https://github.com/rsp/iterex/blob/master/LICENSE.md
[actions-url]: https://github.com/rsp/iterex/actions
[license-img]: https://img.shields.io/npm/l/ende.svg
[github-follow-url]: https://github.com/rsp
[github-follow-img]: https://img.shields.io/github/followers/rsp.svg?style=social&logo=github&label=Follow
[twitter-follow-url]: https://twitter.com/intent/follow?screen_name=pocztarski
[twitter-follow-img]: https://img.shields.io/twitter/follow/pocztarski.svg?style=social&logo=twitter&label=Follow
[stackoverflow-url]: https://stackoverflow.com/users/613198/rsp
[stackexchange-url]: https://stackexchange.com/users/303952/rsp
[stackexchange-img]: https://stackexchange.com/users/flair/303952.png
[test-badge-img]: https://github.com/rsp/iterex/actions/workflows/test.yml/badge.svg
[test-badge-url]: https://github.com/rsp/iterex/actions/workflows/test.yml
[codecov-badge-img]: https://codecov.io/gh/rsp/iterex/graph/badge.svg?token=BBAWZZM6Q2
[codecov-badge-url]: https://codecov.io/gh/rsp/iterex
