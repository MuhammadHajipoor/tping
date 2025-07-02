package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"golang.org/x/sys/windows"
	"golang.org/x/term"
)

// Prints usage instructions for the tping command-line tool
func printUsage() {
	fmt.Println(`Usage: tping [-t] [-a] [-n count] [-l size] [-f] [-i TTL] [-v TOS]
            [-r count] [-s count] [[-j host-list] | [-k host-list]]
            [-w timeout] [-R] [-S srcaddr] [-c compartment] [-p]
            [-4] [-6] target_name

Options:
    -t             Ping the specified host until stopped.
                   To see statistics and continue - type Control-Break;
                   To stop - type Control-C.
    -a             Resolve addresses to hostnames.
    -n count       Number of echo requests to send.
    -l size        Send buffer size.
    -f             Set Don't Fragment flag in packet (IPv4-only).
    -i TTL         Time To Live.
    -v TOS         Type Of Service (IPv4-only. This setting has been deprecated
                   and has no effect on the type of service field in the IP
                   Header).
    -r count       Record route for count hops (IPv4-only).
    -s count       Timestamp for count hops (IPv4-only).
    -j host-list   Loose source route along host-list (IPv4-only).
    -k host-list   Strict source route along host-list (IPv4-only).
    -w timeout     Timeout in milliseconds to wait for each reply.
    -R             Use routing header to test reverse route also (IPv6-only).
                   Per RFC 5095 the use of this routing header has been
                   deprecated. Some systems may drop echo requests if
                   this header is used.
    -S srcaddr     Source address to use.
    -c compartment Routing compartment identifier.
    -p             Ping a Hyper-V Network Virtualization provider address.
    -4             Force using IPv4.
    -6             Force using IPv6.`)
}

// Enables ANSI escape code support on Windows terminals
func enableVirtualTerminalWindows() {
	if runtime.GOOS == "windows" {
		handle := windows.Stdout
		var mode uint32
		windows.GetConsoleMode(handle, &mode)
		windows.SetConsoleMode(handle, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
	}
}

func main() {
	// Get command-line arguments (excluding the program name)
	args := os.Args[1:]

	// If no arguments were provided, show usage and exit
	if len(args) == 0 {
		printUsage()
		os.Exit(0)
	}

	// Enable ANSI color support on Windows
	enableVirtualTerminalWindows()

	// Check if the output is a terminal (TTY)
	isTTY := term.IsTerminal(int(os.Stdout.Fd()))

	// ANSI color codes (used only if output is a TTY)
	green, red, yellow, white, reset := "", "", "", "", ""
	if isTTY {
		green = "\033[32m"
		red = "\033[31m"
		yellow = "\033[33m"
		white = "\033[37m"
		reset = "\033[0m"
	}

	// Create the command to execute (ping with all given arguments)
	cmd := exec.Command("ping", args...)

	// Get a pipe to read stdout (standard output) of the ping command
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Failed to get command output:", err)
		return
	}

	// Start executing the ping command
	if err := cmd.Start(); err != nil {
		fmt.Println("Failed to start command:", err)
		return
	}

	// Read the command output line-by-line in real-time
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()

		// Skip empty lines
		if strings.TrimSpace(line) == "" {
			fmt.Println()
			continue
		}

		// Format current timestamp in yellow
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		timestampStr := fmt.Sprintf("[%s%s%s]", yellow, timestamp, reset)

		// Determine color based on ping output line content
		var color string
		lower := strings.ToLower(line)

		// If it's a reply, show in green
		if strings.Contains(line, "TTL") {
			color = green
		} else if strings.Contains(lower, "request timed out") ||
			strings.Contains(lower, "general failure") ||
			strings.Contains(lower, "unreachable") {
			// If it's a known error, show in red
			color = red
		} else {
			// Default to white for other messages
			color = white
		}

		// Print the formatted output: timestamp + colored ping line
		fmt.Printf("%s %s%s%s\n", timestampStr, color, line, reset)
	}

	// Wait for the ping command to finish
	if err := cmd.Wait(); err != nil {
		fmt.Println("Command finished with error:", err)
	}
}
