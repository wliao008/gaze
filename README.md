# Gaze
Maze generating library in GoLang

[![Circle CI](https://circleci.com/gh/wliao008/gaze.png?style=shield)](https://circleci.com/gh/wliao008/gaze)

## Quickstart
```
$ cd cmd/gaze
$ go run main.go
   _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _
| |  _  |   |_ _ _ _     _ _  |_ _    | |  _   _ _ _ _ _|   |
| | | |_ _|_ _ _ _  | |_  |  _ _ _ _| |_  | |_ _  |   |  _| |
| | | |  _   _ _ _| |_  | |_|  _  |_ _  |_ _ _  | | | | |  _|
| | | | | |_ _ _ _ _|_ _|_ _ _  |_  | | |   |_  |_ _| | |_  |
|_ _| | |_   _  |_   _ _ _  | | | | |  _| |_  |  _ _|  _| | |
|  _  |_  |_  |_ _ _| |   | |_ _ _| |_  |   | |_|   |_ _ _| |
| | |  _| |  _|  _ _ _ _| |_ _ _ _ _|  _| | |_ _ _|_ _ _ _  |
|_  |_  |_ _| |_ _  |    _ _ _  |   |_  |_| | |    _ _ _|  _|
|  _| | |  _   _ _| | |_|  _  |_| |_ _ _|  _|  _|_ _ _  | | |
|_ _ _|_ _ _|_ _ _ _|_ _ _ _|_ _ _ _ _ _ _|_ _ _ _ _ _|_ _  |
```

## Development

### Terminologies
1. [board](https://github.com/wliao008/gaze/blob/master/board.go): imagine a `width` by `height` 2D array, this represents the graphical layout of the maze.
2. [cell](https://github.com/wliao008/gaze/blob/master/cell.go): each item in the 2D array is a cell. Imagine a cell in the middle of a board, this cell could open passage(s) to the neighboring cells in any or all 4 of the directions, this is what essentially creates the maze when all the cells are processed.
3. [direction](https://github.com/wliao008/gaze/blob/master/direction.go): the 4 directions that a cell could open a passage into.

### Algorithms
There are many ways to generate a maze, the [algos](https://github.com/wliao008/gaze/tree/master/algos) package will contains all the different implementation. All of them implements the [Mazer](https://github.com/wliao008/gaze/blob/master/mazer.go) interface that contains the `Generate()` method that new  algorithm would need to implement.

### Solvers
There are also many ways to solve a maze, these are collected in [solvers](https://github.com/wliao008/gaze/tree/master/solvers).



## Contributing
Pull requests are welcome! You could also open a [Github issue](https://github.com/wliao008/gaze/issues) to discuss and get feedback first.

## License
MIT