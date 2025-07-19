// cmd/regtechcompliancecore/main.go
package main

import (
"flag"
"log"
"os"

"regtechcompliancecore/internal/regtechcompliancecore"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := regtechcompliancecore.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
