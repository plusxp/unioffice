#!/bin/bash

find _examples/ -maxdepth 2 -mindepth 2 -exec sh -c "cd {}; echo building {}; go run main.go" \;
