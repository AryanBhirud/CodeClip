# CodeClip 📋

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go)
[![License](https://img.shields.io/badge/License-GPLv3-blue.svg?style=for-the-badge)](LICENSE)
![Platform](https://img.shields.io/badge/Platform-Windows-blue?style=for-the-badge&logo=windows)

A lightning-fast tool to capture and share your entire codebase with LLMs.

</div>

## 🚀 Features

- **Intelligent Scanning**: Automatically excludes non-essential directories (`node_modules`, `.git`, etc.)
- **Smart Filtering**: Ignores binary files, images, and other non-code assets
- **Dual Output**: Generates both file output and clipboard content
- **Path Preservation**: Maintains file paths in the output for better context
- **Zero Configuration**: Works out of the box with sensible defaults

## 🎯 Perfect for

- Sharing codebase context with AI tools (ChatGPT, Claude, etc.)
- Quick code documentation and analysis
- Creating comprehensive code snapshots
- Code review preparation

## 📦 Installation

### Windows Users
1. Download the latest `cclip_win64.zip` from the [releases page](https://github.com/yourusername/codeclip/releases)
2. Extract the archive to your desired location
3. Add the extraction path to your system's `PATH` environment variable

## Verify Installation
Run `cclip --version` to check is codeclip has been installed correctly.


## 🛠️ Usage

Navigate to your project directory and run `cclip`

This will:
1. Scan your codebase
2. Generate `codebase_dump.txt` in the current directory
3. Copy the content to your clipboard automatically

## 🔍 What Gets Captured?

CodeClip intelligently processes your codebase while ignoring:

### Directories
- `node_modules`
- `.git`
- `.env`
- `vendor`
- `__pycache__`

### File Types
- Executables (`.exe`, `.dll`, `.bin`)
- Archives (`.zip`, `.tar`, `.gz`)
- Images (`.png`, `.jpg`, `.svg`, etc.)
- Log files (`.log`)
- Lock files (`.lock`)

## 💡 Best Practices

- Run CodeClip from your project's root directory
- Review the generated `codebase_dump.txt` before sharing
- Consider adding `codebase_dump.txt` to your `.gitignore`

## 🔧 Technical Details

CodeClip is built with Go, leveraging:
- Efficient file system traversal
- Concurrent processing
- Cross-platform clipboard support via [atotto/clipboard](https://github.com/atotto/clipboard)

## 🤝 Contributing

Contributions are welcome! Feel free to:
- Report bugs
- Suggest features
- Submit pull requests

## 📄 License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.
```
CodeClip - Copy your entire codebase instantly
Copyright (C) 2024 Aryan Bhirud

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
```

### What this means:
- ✔️ You can use this software commercially
- ✔️ You can modify this software
- ✔️ You can distribute this software
- ✔️ You can use and modify the software in private
- ✔️ You can patent your modifications
- ❗ You MUST open source your modifications
- ❗ You MUST include the original license
- ❗ You MUST state your changes
- ❗ You MUST disclose source

---

Made with ❤️ for developers who love clean workflows.

