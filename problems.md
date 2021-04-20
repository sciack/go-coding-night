# Problems

## Roman number
Write a function that convert roman number to integer, the function will always have a correct input.
|Roman|Integer|
|-----|------|
|I | 1|
|V | 5|
|IV| 4|
|X|10|

# Fibonacci series
Write a function that calculate the fibonacci series (hint the recursive one will be too slow if n is "big")

Fibonacci numbers:
```
a(0) = 0
a(1) = 1
a(n) = a(n-1) + a(n-2)
```

so
|n|fibonacci|
|-|---------|
|0|0|
|1|1|
|2|1|
|3|2|
|4|3|
|5|5|
|6|8|


## Game of life
This easy game is a turn based game with the following rules [Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life):

1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
1. Any live cell with two or three live neighbours lives on to the next generation.
1. Any live cell with more than three live neighbours dies, as if by overpopulation.
1. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

This could be printed on console using this library `github.com/gdamore/tcell/v2` [github](github.com/gdamore/tcell)

```
go get github.com/gdamore/tcell/v2 v2.1.0
```

but the documentation is quite poor.