[![Go Report Card](https://goreportcard.com/badge/github.com/yitsushi/advent-of-code-2020)](https://goreportcard.com/report/github.com/yitsushi/advent-of-code-2020)
[![Coverage Status](https://coveralls.io/repos/github/yitsushi/advent-of-code-2020/badge.svg?branch=main)](https://coveralls.io/github/yitsushi/advent-of-code-2020?branch=main)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/yitsushi/advent-of-code-2020)


[Last year][2019], I tried to extend my knowledge and tried to solve
all puzzles with Haskell. Sadly, I had no time to learn more and more
because I don't use it daily as part of my job.

This is one of the reasons why I do it in Go this year. I have some
other reasons too like:

* I planned to do it with Go since February
* I hope I can face some beast this year and I can push my limits
* I want to play a bit safer this year as I have even less time
    than last year

Still curious and looking for the opportunity to work with Haskell
again more and more, but for now, I'll stick with Go.

**TODO**: Extend this document.

# How to use

> You can also get the [JSON] for this private leaderboard. Please
> don't make frequent automated requests to this service - avoid
> sending requests more often than once every 15 minutes (900 seconds).
> If you do this from a script, you'll have to provide your session
> cookie in the request; a fresh session cookie lasts for about a month.
> Timestamps use Unix time.
>
> Source: adventofcode.com

Session is required to download your input file and submit your solution:

```
export AOC_SESSION="................."
```

### Generate basic file and directory structure for a day

```
❯ go run ./cmd/aoc2020  scaffold --day=4
```

### Download Input File

```
❯ aoc2020 download --day=3
Done... input/day03/part1
```

Note: The filename is `part1` because it's easier to manage the input file
      in case we have different input for part 1 and part 2.

### Execute solver

```
# Simple
❯ aoc2020 solve --day=1 --part=1 --input-file=input/day01/part1
805731

# With Info messages
❯ aoc2020 solve --day=1 --part=1 --input-file=input/day01/part1 --log-level=info
INFO[0000] Part1: 11.047µs
805731

# Debug
❯ aoc2020 solve --day=1 --part=1 --input-file=input/day01/part1 --log-level=debug
DEBU[0000] Calling *day01.Solver => Part1
INFO[0000] Part1: 7.1µs
805731
```

### Submit Solution

```
❯ aoc2020 submit --day=3 --part=2 --input-file input/day03/part1 --log-level=info
INFO[0000] Part2: 10.1µs
INFO[0000] Did you already complete it?
Done \o/
```

# History

* [2019] => Haskell (~75%)
* [2018] => Python (100%)
* [2017] => [aoc-random-language] (100%)
* [2016] => C++ (100%)

[aoc-random-language]: https://github.com/Yitsushi/aoc-random-language
[2019]: https://github.com/yitsushi/advent-of-code-2019
[2018]: https://github.com/yitsushi/advent-of-code-2018
[2017]: https://github.com/yitsushi/advent-of-code-2017
[2016]: https://github.com/yitsushi/advent-of-code-2016
