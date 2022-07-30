# Quick Start - Install get

> [!TIP]
> get is installable via [instl.sh](https://instl.sh).\
> You just have to run the following command and you're ready to go!

<!-- tabs:start -->

#### ** Windows **

### Windows Command

```powershell
iwr instl.sh/fzdwx/get/windows | iex
```

#### ** Linux **

### Linux Command

```bash
curl -sSL instl.sh/fzdwx/get/linux | bash
```

#### ** macOS **

### macOS Command

```bash
curl -sSL instl.sh/fzdwx/get/macos | bash
```

#### ** Compile from source **

### Compile from source with Golang

?> **NOTICE**
To compile get from source, you have to have [Go](https://golang.org/) installed.

Compiling get from source has the benefit that the build command is the same on every platform.\
It is not recommended to install Go only for the installation of get.

```command
go install github.com/fzdwx/get@latest
```

<!-- tabs:end -->
