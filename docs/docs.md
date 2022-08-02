# get

## Usage
> 一个下载器.

get

## Description

```
这是一个下载 "CLI" 应用程序,它可以下载一些音乐或其他一些东西.
```
## Examples

```bash
get mc 稻香
get clone fzdwx/get

```

## Flags
|Flag|Usage|
|----|-----|
|`--debug`|启用调试消息|
|`--disable-update-checks`|禁用更新检查|
|`--raw`|打印无样式的原始输出（如果将输出写入文件，则设置它）|

## Commands
|Command|Usage|
|-------|-----|
|`get clone`|这是一个克隆项目的子命令,目前只支持克隆github项目.|
|`get completion`|Generate the autocompletion script for the specified shell|
|`get help`|Help about any command|
|`get msc`|下载音乐.|
# ... clone
`get clone`

## Usage
> 这是一个克隆项目的子命令,目前只支持克隆github项目.

get clone [owner]/[repo]
## Examples

```bash
get clone fzdwx/get
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
> 下载音乐.

get msc [song name]

## Description

```
这是一个下载音乐的子命令,目前只支持酷我以及网易云平台,默认使用酷我.
```
## Examples

```bash
get mc 不能说的秘密
get msc 我的天空
get msc 平凡之路 -w # 使用网易云
get msc 七里香 -k   # 酷我
```

## Flags
|Flag|Usage|
|----|-----|
|`-k, --kuwo`|使用酷我平台下载音乐 (default true)|
|`-w, --neteasy`|使用网易云平台下载音乐|


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 02 August 2022**
