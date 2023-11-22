package main

import (
  "log"
  "klog.co/plantz/webServer"
)

func main() {
  log.Println("Starting web server")
  webServer.Start()
}
