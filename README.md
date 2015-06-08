# visa

Go-based VISA resource manager.

[![GoDoc][godoc image]][godoc link]

## Installation

```bash
$ go get github.com/gotmc/visa
```

## Documentation

Documentation can be found at either:

- <https://godoc.org/github.com/gotmc/visa>
- <http://localhost:6060/pkg/github.com/gotmc/visa/> after running `$
  godoc -http=:6060`

## Contributing

[visa][] is developed using [Scott Chacon][]'s [GitHub Flow][]. To
contribute, fork [visa][], create a feature branch, and then
submit a [pull request][].  [GitHub Flow][] is summarized as:

- Anything in the `master` branch is deployable
- To work on something new, create a descriptively named branch off of
  `master` (e.g., `new-oauth2-scopes`)
- Commit to that branch locally and regularly push your work to the same
  named branch on the server
- When you need feedback or help, or you think the branch is ready for
  merging, open a [pull request][].
- After someone else has reviewed and signed off on the feature, you can
  merge it into master.
- Once it is merged and pushed to `master`, you can and *should* deploy
  immediately.

## Testing

Prior to submitting a [pull request][], please run:

```bash
$ gofmt
$ golint
$ go vet
$ go test
```

To update and view the test coverage report:

```bash
$ go test -coverprofile coverage.out
$ go tool cover -html coverage.out
```

## License

[visa][] is released under the MIT license. Please see the
[LICENSE.txt][] file for more information.

[GitHub Flow]: http://scottchacon.com/2011/08/31/github-flow.html
[godoc image]: https://godoc.org/github.com/gotmc/visa?status.svg
[godoc link]: https://godoc.org/github.com/gotmc/visa
[LICENSE.txt]: https://github.com/gotmc/visa/blob/master/LICENSE.txt
[pull request]: https://help.github.com/articles/using-pull-requests
[Scott Chacon]: http://scottchacon.com/about.html
[visa]: https://github.com/gotmc/visa
