# Korpus

Corpus of data to test backups on

## Included data sets

- [Linux Kernel](github.com/torvalds/linux): Large size, deep directory
    structure, diverse file formats.
- [Linux Firmware](https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/):
    Lots of binary files.
- [LLVM Project](github.com/llvm/llvm): Large source codebase with frequent changes.
- [VS Code](github.com/microsoft/vscode): Modern development project with
    binaries, JSON, and Markdown files.
- [OpenCV](github.com/opencv/opencv): A mix of C++ code, large datasets, and documentation.
- [TensorFlow](github.com/tensorflow/tensorflow): Diverse files including
    Python, protobuf, and binaries.
- [OpenBSD src](https://github.com/openbsd/src): Core source code of the
    OpenBSD operating system, including the kernel, userland utilities, and
    system libraries. Features a security-focused, minimalist design.
- [OpenBSD ports](https://github.com/openbsd/ports): Build scripts and metadata
    for third-party software, tailored to the OpenBSD environment. Includes
    patches and descriptions.
- [OpenBSD www](https://github.com/openbsd/www): The official OpenBSD website
    source, containing documentation, release notes, and man page information.
- [NetBSD src](https://github.com/NetBSD/src): Source code for the NetBSD
    operating system, designed for portability and supporting a wide range of
    hardware platforms.
- [NetBSD pkgsrc](https://github.com/NetBSD/pkgsrc): Portable package
    management framework. Includes build scripts, patches, and metadata for
    thousands of software packages.
- [FreeBSD src](https://github.com/freebsd/freebsd-src): Core source code of
    the FreeBSD operating system, focusing on performance and modern hardware
    support.
- [FreeBSD ports](https://github.com/freebsd/freebsd-ports): The FreeBSD ports
    collection, offering build scripts, metadata, and patches for third-party
    applications.
- [FreeBSD doc](https://github.com/freebsd/freebsd-doc): Documentation and
    website source for FreeBSD, containing guides, release notes, and other
    resources in Markdown, HTML, and other formats.

## How to

Korpus relies on git submodules to populate large source codebases.
These codebases will not be checked out automatically when first cloning the
project.

To initialize submodules and fetch their content, run:

```sh
git submodule update --init --remote --recursive
```

To update existing submodules, run:

```sh
git submodule update --remote --recursive
```
