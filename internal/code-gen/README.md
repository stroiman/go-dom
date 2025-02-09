# gost-dom code generator

This is part of the [Gost-DOM](https://github.com/gost-dom/browser) project.

This repository contains code to generate code for the code-gen browser based on
web specifications.

### Building the code generator.

The code is generated from specifications from the
[webref](https://github.com/w3c/webref) repository, which is added as a
submodule to this project.

To build the code generator, you need to fetch the submodule and a _curated_ set
of files.

Prerequisites: Node.js and npm (or compatible alternatives)

```sh
$ git submodul update --init
$ cd webref
$ npm install # Or your favourite node package manager
$ npm run curate
```

This build a set of files in the `curated/` subfolder.

> [!NOTE]
>
> The webref is in the process of being moved to a new self-contained repo, so
> no custom steps are needed.
