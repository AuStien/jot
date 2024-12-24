# 📝 Jot

A small Go program to simplify jotting down thoughts.

It consists of two main components; the [binder](#binder) and the [journal](#journal).

## 🐶 Usage

`jot binder`: organize files in binders (directories), see [binder](#binder)

`jot journal`: keep a journal, see [journal](#journal)

`jot todo`: edit a global TODO file (opens `$JOT_HOME/TODO.md`)

### Journal

`jot journal view`: concats all the files in the last month in a temporary file and opens it

`jot journal add`: add a new entry in the journal (creates and opens a new file `$JOT_HOME/YYYY/MM/DD.md` with header `## HH:MM`)

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
# Tuesday 31/12/24

## 10:43

Some thoughts written down.

## 12:01

More thoughts.
```


### Binder

`jot binder <path-to-file>`: opens the specified file, creating the directories and file if necessary

When using autocompletion `source <(jot completion <shell>)`, double pressing tab will show suggestions
of the files/directories that already exists.

---

To make jotting outside of a terminal easier, it's recommended to create a keyboard shortcut
to open up a new terminal and run `jot` (it might be necessary to specify the home directory and editor using the flags,
as envvars might not be loaded).
The setup will vary depending on OS, desktop environment and shell used, so figuring out how to do this is an exercise
left for the reader.

## 🏗️ Setup

**Home directory**

The home directory for the notes needs to be specified.  
This can be done by either using the `--home` flag (has precedence), or setting the `JOT_HOME` environment variable.

Defaults to `$XDG_DATA_HOME`. If undefined, uses `$HOME/.local/share/jot`.

**Editor**

Which editor to use to edit the files can be specified using either the `--editor` flag (has precedence)
or setting the `EDITOR` environment variable.

Defaults to `vi`.

### Environment Variables

**Optional**

- `JOT_HOME`: the path to the directory where files should be created
- `EDITOR`: is the editor of choice to open the files with

## 📝 Development

> It's recommended setting the `JOT_HOME` envvar to a different path than you usually use to avoid messing up "production" notes.

`make run`: run the code

`make build` (or just `make)`: build the code

## 🤝 Contribution

Feel free to create an [issue](https://github.com/AuStien/jot/issues) or [PR](https://github.com/AuStien/jot/pulls).

As the project is still quite small there are no requirements for issues or PRs.

