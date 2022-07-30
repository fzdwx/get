# get

## Usage
> This is a downloader.

get

## Description

```
This is a downloader CLI application, it can download some music or some other stuff.
```
## Examples

```bash
get mc 稻香
get clone fzdwx/get
get date
get date --format 20060102
get time
get time --live
```

## Flags
|Flag|Usage|
|----|-----|
|`--debug`|enable debug messages|
|`--disable-update-checks`|disables update checks|
|`--raw`|print unstyled raw output (set it if output is written to a file)|

## Commands
|Command|Usage|
|-------|-----|
|`get clone`|clone github project|
|`get completion`|Generate the autocompletion script for the specified shell|
|`get date`|Prints the current date.|
|`get help`|Help about any command|
|`get msc`|download music.|
|`get time`|Prints the current time|
# ... clone
`get clone`

## Usage
> clone github project

get clone [owner]/[repo]
## Examples

```bash
get git fzdwx/get
```
# ... completion
`get completion`

## Usage
> Generate the autocompletion script for the specified shell

get completion

## Description

```
Generate the autocompletion script for get for the specified shell.
See each sub-command's help for details on how to use the generated script.

```

## Commands
|Command|Usage|
|-------|-----|
|`get completion bash`|Generate the autocompletion script for bash|
|`get completion fish`|Generate the autocompletion script for fish|
|`get completion powershell`|Generate the autocompletion script for powershell|
|`get completion zsh`|Generate the autocompletion script for zsh|
# ... completion bash
`get completion bash`

## Usage
> Generate the autocompletion script for bash

get completion bash

## Description

```
Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(get completion bash)

To load completions for every new session, execute once:

#### Linux:

	get completion bash > /etc/bash_completion.d/get

#### macOS:

	get completion bash > $(brew --prefix)/etc/bash_completion.d/get

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion fish
`get completion fish`

## Usage
> Generate the autocompletion script for fish

get completion fish

## Description

```
Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	get completion fish | source

To load completions for every new session, execute once:

	get completion fish > ~/.config/fish/completions/get.fish

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion powershell
`get completion powershell`

## Usage
> Generate the autocompletion script for powershell

get completion powershell

## Description

```
Generate the autocompletion script for powershell.

To load completions in your current shell session:

	get completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion zsh
`get completion zsh`

## Usage
> Generate the autocompletion script for zsh

get completion zsh

## Description

```
Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(get completion zsh); compdef _get get

To load completions for every new session, execute once:

#### Linux:

	get completion zsh > "${fpath[1]}/_get"

#### macOS:

	get completion zsh > $(brew --prefix)/share/zsh/site-functions/_get

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... date
`get date`

## Usage
> Prints the current date.

get date

## Flags
|Flag|Usage|
|----|-----|
|`-f, --format string`|specify a custom date format (default "02 Jan 06")|
# ... help
`get help`

## Usage
> Help about any command

get help [command]

## Description

```
Help provides help for any command in the application.
Simply type get help [path to command] for full details.
```
# ... msc
`get msc`

## Usage
> download music.

get msc [song name]

## Description

```
this is a subcommand to download music.currently only supports NetEase Cloud platform
```
## Examples

```bash
get mc 不能说的秘密
get msc 我的天空	
```
# ... time
`get time`

## Usage
> Prints the current time

get time

## Description

```
You can print a live clock with the '--live' flag!
```

## Flags
|Flag|Usage|
|----|-----|
|`-l, --live`|live output|


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 30 July 2022**
