<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/fiscusproject/brandbook/master/banner/fiscus-readme-tagline-1280x320-dark.png">
    <img src="https://raw.githubusercontent.com/fiscusproject/brandbook/master/banner/fiscus-readme-tagline-1280x320-light.png" alt="Fiscus — free and open-source fiscalization" width="640">
  </picture>
</p>

# fiscus

Status: pre-alpha, service scaffolding only, no fiscalization implemented yet.

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

On October 22, 2026, we decided to fully align with the NLnet Foundation's AI policy requirements. From this date forward, AI will not be used to produce, edit, or manage code in this
repository.

[NLNet Foundation AI Policy Version 2.0](https://nlnet.nl/foundation/policies/generativeAI/2.0/)


### AI Policy Disclaimer

The code in this repository prior to October 22 was already aligned with Version 2.0 of the NLnet Foundation's AI policy requirements.
