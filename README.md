# 📝 Jot

A small Go program to easily be able to jot down thoughts and stuff.

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

Running `jot todo` will open a `TODO.md` file in your `JOT_HOME` directory.
This file is meant to be used as an easy way of keeping track of things that need doing.

## 🏗️ Setup

**Home directory**

The home directory for the notes needs to be specified.
This can be done by either using the `--home` flag (has precedence), or setting the `JOT_HOME` environment variable.

**Editor**

Which editor to use to edit the files can be specified using either the `--editor` flag (has precedence)
or setting the `EDITOR` environment variable.

If neither of these are set, it defaults to `vi`.

### Environment Variables

**Optional**

- `JOT_HOME`: the path to the directory where files should be created
- `EDITOR`: is the editor of choice to open the files with

## 🐶 Usage

`jot`: create and open a new file `$JOT_HOME/YYYY/MM/DD.md` with header `## HH:MM`

`jot view`: concats all the files in the last month in a temporary file and opens it

`jot todo`: open the file `$JOT_HOME/TODO.md`

---

To make jotting outside of a terminal easier, it's recommended to create a keyboard shortcut
to open up a new terminal and run `jot` (it might be necessary to specify the home directory and editor using the flags,
as envvars might not be loaded).
The setup will vary depending on OS, desktop environment and shell used, so figuring out how to do this is an exercise
left for the reader.

## 📝 Development

> It's recommended setting the `JOT_HOME` envvar to a different path than you usually use to avoid messing up "production" notes.

`make run`: run the code

`make build` or `make`: build the code

## 🤝 Contribution

Feel free to create an [issue](https://github.com/AuStien/jot/issues) or [PR](https://github.com/AuStien/jot/pulls).

As the project is still quite small there are no requirements for issues or PRs.

