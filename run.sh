#!/usr/bin/env sh

./bin/server &
PID1="$!"
trap "kill $PID1; exit" SIGHUP SIGINT SIGQUIT SIGTERM
./bin/client


