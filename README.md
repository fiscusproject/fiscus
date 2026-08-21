<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/fiscusproject/brandbook/master/banner/fiscus-readme-tagline-1280x320-dark.png">
    <img src="https://raw.githubusercontent.com/fiscusproject/brandbook/master/banner/fiscus-readme-tagline-1280x320-light.png" alt="Fiscus — free and open-source fiscalization" width="640">
  </picture>
</p>

# fiscus

## Dev Environment Setup

### In a dev container

`.devcontainer/` defines a [DevPod](https://devpod.sh) that
installs the pinned toolchain automatically.

```sh
devpod up . --ide vscode
```

### On your machine

Install [mise](https://mise.jdx.dev), then from the repository root run:

```sh
mise install
```

This installs the toolchain versions pinned in `mise.toml`.

## Contributing

### Do not commit tainted code to the repository

If you commit code that wasn't written by yourself, double-check that the license on that code permits import into the Fiscus source code repository, and permits free distribution.

## AI Policy

The code in this repository was produced with AI assistance. All decisions were made by the project maintainers, every line of code was human-reviewed, and the maintainers remain accountable for the accuracy and originality of the work.

