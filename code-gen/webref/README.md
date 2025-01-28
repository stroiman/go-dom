# Webref - Go types representing Web IDL standards

This library exposes data from Web IDL specifications as Go types.

Data is sourced from [github.com/w3c/webref](https://github.com/w3c/webref)

> [!IMPORTANT]
>
> This tool is not completed, in that not all information in the IDL specs are
> exposed. If you find that you need something not exposed, please file an
> issue. Or even better, make a PR.

## A tool for code-gen

The intention of this is to be a tool for the build process of other go tools
targeting the web.

## Packages

This package is divided into two subpackages

- `html` contains mapping from HTML element tag names to IDL interface name.
- `IDL` contains the Web IDL specifications.
