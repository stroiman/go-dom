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

## Completeness

This is not a complete representation of all data, but strives to be the place
to look.

Features are added as 

## Packages

This package is divided into subpackages:

- `html` contains mapping from HTML element tag names to IDL interface name.
- `idl` contains the Web IDL specifications.

[^1]: This can, and will probably be shrunk by stripping irrelevant data from
    the JSON files. E.g., duplication of original Web IDL files, as well as a
significant amount of whitespace.

## Coding guidelines

This codebase is tested in general terms by inspecting properties on select
values, and comparing them against expectations. E.g., the `URL` interface from
the `url` specification should have a non-static `toJSON` operation, and a
static `parse` operation.

When a new feature is supported, add a test for a type that uses the feature,
showing.

### Historic code

The first version of the code for the idl module consisted of structures
"reverse engineered" from the JSON data, not knowing exactly what they
represented. This data _should_ be complete, but exposes a less useful model.

Later, I started reading the specs for Web IDL itself, and started a new set of
types that has a model reflecting the standard itself. This is not complete
however.

Eventually, the old model will be removed (or unexported), leaving only the new
model.
