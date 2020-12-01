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

```
# Simple
go run ./cmd/aoc2020 solve --day=1 --part=1 --input-file=input/day01/part1
805731

# With Info messages
go run ./cmd/aoc2020 solve --day=1 --part=1 --input-file=input/day01/part1 --log-level=info
INFO[0000] Part1: 11.047µs
805731

# Debug
go run ./cmd/aoc2020 solve --day=1 --part=1 --input-file=input/day01/part1 --log-level=debug
DEBU[0000] Calling *day01.Solver => Part1
INFO[0000] Part1: 7.1µs
805731
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
