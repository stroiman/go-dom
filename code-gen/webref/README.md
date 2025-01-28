# Webref - Go types representing Web IDL standards

This library exposes data from Web IDL specifications as Go types.

Data is sourced from [github.com/w3c/webref](https://github.com/w3c/webref)

> [!IMPORTANT]
>
> This tool is not completed, in that not all information in the IDL specs are
> exposed. If you find that you need something not exposed, please file an
> issue. Or even better, make a PR.

## About the compiled file size

The intention of this is to be a tool for the build process of other go tools
targeting the web, not runtime use in web applications.

For that reason, minimising the size of compiled library has not been a
priority. The compiled library takes up about 16Mb on disk, which is mainly
embedded data files, which of course takes up space in source code too.[^1]

## Packages

This package is divided into subpackages:

- `html` contains mapping from HTML element tag names to IDL interface name.
- `IDL` contains the Web IDL specifications.

[^1]: This can, and will probably be shrunk by stripping irrelevant data from
    the JSON files. E.g., duplication of original Web IDL files, as well as a
significant amount of whitespace.
