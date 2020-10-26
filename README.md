# Fetch Challenges

Two engineering challenges in one.

This Go application is a microservice with two api endpoints:

1. `/api/v1/pyramid/:word`
2. `/api/v1/versions/:versionA/:versionB`

These correspond to the challenges detailed further down.

## Two? Why?

I had missed the emails sent by Fetch's recruiter. @kjkondratuk reached out to see what was up, and forwarded to me one of the engineering challenges: the pyramid word challenge. Shortly after I had started working on it in a REPL, the recruiter reached out, again. I found the buried email with the original challenge, and decided that I'd complete both of them.

You can find the repl where I had started my work, before moving it to a github repo, here: https://repl.it/@LiamConrad/PyramidWord.

## Getting Started

You'll need to have Go 1.14 or later installed to compile this project. (I'm using Go 1.15.)

### Installing Go

Please refer Go's installation and getting started documentation, which you can find here: https://golang.org/doc/install.

On a Mac, you can install it using Homebrew a la `brew install go`.

Verify you have Go installed by bringing up a terminal, and entering `go version`.

### Installing this Repo

With Go installed, simply run `go get github.com/storskegg/fetch_challenges`. You should find the repo cloned into `$HOME/go/src/github.com/storskegg/fetch/challenges`.

In that directory, run the following commands:
1. `go mod vendor`
2. `go test`
3. `go build`

### Starting the Server

To start the server, simply run the new executable in that directory.

- `fetch_challenges` for Linux and Mac
- `fetch_challenges.exe` for Windows

## Challenges

### Pyramid Word Test

`/api/v1/pyramid/:word`

This endpoint determines whether a word is a pyramid word, that is, a word whose letter counts increment cardinally without skipping any numbers. For example:

- `a` is technically a pyramid word, because it has a single letter with a count of 1
- `app` is also a pyramid word, with letter counts in ascending order: `a: 1, p: 2`
- `banana` is a pyramid word, because its letter counts, in ascending order, are `b: 1, n: 2, a: 3`
- `bandana` is _not_ a pyramid word, because its letter counts, in ascending order, are `b: 1, d: 1, n: 2, a: 3`
- `momoo` is _not_ a pyramid word, because its letter counts, in ascending order, are `m: 2, o: 3`, and there is no letter with a count of `1`

The endpoint returns a JSON object containing the tested word, and whether the word is a pyramid word.

Example Responses:
```
{
  "word": "banana",
  "isPyramid": true
}
```

```
{
  "word": "apple",
  "isPyramid": false
}
``` 

With the server running, and using your favorite web client for API testing, send a request to `http://localhost:3000/api/v1/pyramid/:word`, where `:word` is any word you wish.

### Versions Comparison Test

`/api/v1/versions/:versionA/:versionB`

This endpoint evaluates two versions, `versionA` and `versionB`, and determines `versionA`'s relationship to `versionB`, and returns the results as JSON object. It only accepts version strings in formats like `1.0`, `2020.10.15` or `1.2.3.4.5.6.7`; but _not_ formats like `v3.0.1`, `1.0.5b` or `2.2.0-beta`. It will error if either or both version strings are not in a supported format.

Example Responses:
```
{
  "versionA": "1.7.8",
  "versionB": "1.7.9",
  "relationship": "1.7.8 is before 1.7.9"
}
```

```
{
  "versionA": "2.0",
  "versionB": "2.0.0",
  "relationship": "2.0 is equal to 2.0.0"
}
```

```
{
  "versionA": "2.0",
  "versionB": "2.0.1",
  "relationship": "2.0 is before 2.0.1"
}
```

```
{
  "versionA": "2.0.1",
  "versionB": "2.0",
  "relationship": "2.0.1 is after 2.0"
}
```

With the server running, and using your favorite web client for API testing, send a request to `http://localhost:3000/api/v1/versions/:versionA/:versionB`, where `:versionA` and `:versionB` are version strings.
