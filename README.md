# seqdiff

[![CI](https://github.com/headlesslab/seqdiff/actions/workflows/ci.yml/badge.svg)](https://github.com/headlesslab/seqdiff/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/headlesslab/seqdiff.svg)](https://pkg.go.dev/github.com/headlesslab/seqdiff)

Longest common subsequence of two sequences of characters, words or lines: the diff core of the got test framework, with an O(M log L) algorithm (`YadLCS`) and the textbook one for reference (`StandardLCS`).


## Provenance

`seqdiff` is a Satellite module of [wand](https://github.com/headlesslab/wand). It is a snapshot of the `lib/lcs` package of [got](https://github.com/ysmood/got) `v0.40.0` (commit `abba87a`), imported without git history and maintained here by headlesslab; see [NOTICE](NOTICE). The module path changed, the exported identifiers did not.

Replace `github.com/ysmood/got/lib/lcs` with `github.com/headlesslab/seqdiff` and `lcs.` with `seqdiff.`; `YadLCS`, `NewChars`, `NewLines`, `Sequence`, ... keep their names.

## Usage

```go
import "github.com/headlesslab/seqdiff"
```

The tests are the documentation: see [seqdiff_test.go](seqdiff_test.go).

## Versioning

`v0.x`: no breaking change without a major version bump. Releases are cut by hand when the module changes; wand pins each release exactly.

## License

MIT, see [LICENSE](LICENSE).
