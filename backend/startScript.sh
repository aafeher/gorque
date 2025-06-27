#!/usr/bin/env bash

log() {
  echo "STARTSCRIPT: $1"
}

buildServer() {
  log "Building server binary"
  go build -buildvcs=false -gcflags "all=-N -l" -o /gorque-backend /go/src/gorque/backend
  log "Server binary built"
}

runServer() {
  log "Run server"

  log "Killing old server"
  killall dlv
  killall gorque-backend
  log "Run in debug mode"
  dlv exec /gorque-backend --listen=:40000 --headless=true --api-version=2 --accept-multiclient --continue &
}

rerunServer() {
  log "Rerun server"
  buildServer
  runServer
}

liveReloading() {
  log "Run liveReloading"
  inotifywait -e "MODIFY,DELETE,MOVED_TO,MOVED_FROM" -m -r --include '.go$' /go/src/gorque | (
    # read changes from inotify, batch results to a second (read -t 1)
    while true; do
      read path action file
      ext=${file: -3}
      if [[ "$ext" == ".go" ]]; then
        echo "$file"
      fi
    done
  ) | (
    WAITING=""
    while true; do
      file=""
      read -t 1 file
      if test -z "$file"; then
        if test ! -z "$WAITING"; then
          echo "CHANGED"
          WAITING=""
        fi
      else
        log "File ${file} changed" >>/tmp/filechanges.log
        WAITING=1
      fi
    done
  ) | (
    # read statement release when some file has been changed
    while true; do
      read TMP
      log "File Changed. Reloading..."
      rerunServer
    done
  )
}

initializeFileChangeLogger() {
  echo "" > /tmp/filechanges.log
  tail -f /tmp/filechanges.log &
}

main() {
  initializeFileChangeLogger
  buildServer
  runServer
  liveReloading
}

main
