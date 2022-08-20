#!/usr/bin/env bash
export AUTH_PORT=":4004"
export TOKEN_CLEAR_DURATION=36000000 #milliseconds - 10 hours

go run main.go