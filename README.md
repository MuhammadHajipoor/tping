# tping
## _A CLI wrapper for the Windows `ping` command that adds timestamps to each line of output._

![tping logo](https://raw.githubusercontent.com/muhammadhajipoor/tping/main/assets/tping-logo.png)

## 📝 Description

`tping` is a simple command-line tool written in Go. It wraps the native `ping` command and prepends a timestamp to each line of its output.

This tool is particularly useful in network diagnostics when you need to know **when** a device responded or failed to respond. It's great for monitoring unstable connections like routers or switches that may go up and down during the day.

## 📦 Features

- Adds current date and time to every `ping` response
- Colored output:
  - 🟡 Timestamps in yellow
  - 🟢 Successful replies in green
  - 🔴 Errors like "Request timed out" or "General failure" in red
  - ⚪ Other output lines in white
- Fully compatible with all arguments of the Windows `ping` command

## 💡 Why I Created This

I had a network switch that was randomly disconnecting and reconnecting. The built-in `ping` command didn’t show the time of each response. I needed a way to know **exactly when** the connection failed, so I wrote `tping` to log timestamps along with ping results.

## 🔧 Installation
1. Install Go (https://go.dev/dl/)
2. Clone the repository:
   ```bash
   git clone https://github.com/muhammadhajipoor/tping.git
   cd tping.
   go build -o tping.exe
    ```
3. Or compile for Windows:
    ```bash
    GOOS=windows GOARCH=amd64 go build -o tping.exe main.go
    ```
## 🚀 Usage
```cmd
./tping google.com
./tping -n 10 8.8.8.8
./tping -t 192.168.1.1
```
## 🔍 Example Output
```cmd
[2025-06-29 23:45:01] Reply from 8.8.8.8: bytes=32 time=18ms TTL=116
[2025-06-29 23:45:02] Request timed out.
[2025-06-29 23:45:03] Reply from 8.8.8.8: bytes=32 time=19ms TTL=116
```
## 🧾 License
This project is licensed under the MIT License.
Feel free to use, modify, and distribute it.

Built with ❤️ using Go.
