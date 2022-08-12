# Team divider [![Release](https://img.shields.io/github/v/release/Aldebaranoro/team-divider?color=cyan&label=%20)](https://github.com/Aldebaranoro/team-divider/releases) [![](https://github.com/Aldebaranoro/team-divider/workflows/Tests/badge.svg)](https://github.com/Aldebaranoro/team-divide/actions)

`team-divider` lets you to randomly distribute participants into different teams.

![cli](https://user-images.githubusercontent.com/48175755/180143638-24e80789-c598-430e-b789-3d28f3af9d3a.png)

## Installation

### Download

Download the latest stable [release](https://github.com/Aldebaranoro/team-divider/releases) for use on your desktop or
server.

### Go

```shell
go install github.com/Aldebaranoro/team-divider/cmd/team-divider@latest
```

### 🍺 Homebrew

```bash
$ brew tap Aldebaranoro/tap
$ brew install team-divider
```

or

```bash
$ brew install Aldebaranoro/tap/team-divider
```

### 🍨 Scoope

```shell
> scoop bucket add Aldebaranoro https://github.com/Aldebaranoro/scoop-bucket
> scoop install Aldebaranoro/team-divider
```

## Usage

### Divider

```bash
$ team-divider player_name... [flags]
```

For example, to divide people into two teams, you can do:

```bash
$ team-divider Nobby Samuel Fred "Havelock Vetinari"
```

If you want to specify your number of teams, use the `-n` or `--teams` flag:

```bash
$ team-divider Nobby Samuel Fred "Havelock Vetinari" -n 3
```

### Indexer

Indexer will allow you to determine the order of the player's move or compliance with any game conditions.

![image](https://user-images.githubusercontent.com/48175755/182078029-c913525b-9598-470a-8186-6f4225c09c7a.png)

## Credits

* [Anatoly Karas](https://github.com/Aldebaranoro)

## License

The MIT License (MIT) - see [`LICENSE`](https://github.com/Aldebaranoro/team-divider/blob/main/LICENSE) for more details
