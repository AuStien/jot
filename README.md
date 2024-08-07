# 📖 Logbook

A small Go program to easily be able to log thoughts and stuff.

Will follow the following directory structure based on the current date (where December 31, 2024 would be `2024/12/31.md`):

```
2024/
  01/
    01.md
    02.md
  02/
    01.md
```

Will follow the following file structure, with headers being automatically generated:

```Markdown
# Thursday 18/07/24

## 10:43

Some thoughts written down.

## 12:01

More thoughts.
```

Running `log todo` will open a `TODO.md` file in your `LOGBOOK_HOME` directory.
This file is meant to be used as an easy way of keeping track of things that need doing.

## 🏗️ Setup

### Environment Variables

**Required**

- `LOGBOOK_HOME`: the path to the directory where files should be created

**Optional**

- `EDITOR`: is the editor of choice to open the files with. Defaults to `vim` if not set

## 🐶 Usage

`log`: create a new file `$LOGBOOK_HOME/YYYY/MM/DD.md`, with header `## HH:MM`, and open with `$EDITOR` (or `vim`)

`log view`: opens the most current file (TODO: check further back than the first of each month and year)

`log todo`: open the file `$LOGBOOK_HOME/TODO.md` with `$EDITOR` (or `vim`)

## 📝 Development

> It's recommended setting the `LOGBOOK_HOME` envvar to a different path than you usually use to avoid messing up "production" notes.

`make run`: run the code

`make build` or `make`: build the code

## 🤝 Contribution

Feel free to create an [issue](https://github.com/AuStien/logbook/issues) or [PR](https://github.com/AuStien/logbook/pulls).

As the project is still quite small there are no requirements for issues or PRs.

## 🌈 Future possibilites

- Add tooling for grouping together notes
- Add possiblity of storing files remotely using `sftp` and/or FUSE
