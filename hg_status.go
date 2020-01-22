package main

import (
    "fmt"
    "os/exec"
    "strings"
)

func main() {
    var modified, added, removed, clean, missing, unknown, ignored int
    cmd := exec.Command("hg", "status")
    out, err := cmd.Output()
    if err != nil {
        return
    }
    lines := strings.Split(string(out), "\n")
    for _, line := range(lines) {
        e := strings.Split(line, " ")
        if len(e) != 2 {
            continue
        }
        switch e[0] {
        case "M":
            modified++
        case "A":
            added++
        case "R":
            removed++
        case "C":
            clean++
        case "!":
            missing++
        case "?":
            unknown++
        case "I":
            ignored++
        }
    }
    fmt.Printf("%d|%d|%d|%d|%d|%d|%d", modified, added, removed, clean, missing, unknown, ignored)
}
