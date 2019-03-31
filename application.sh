#!/bin/bash

function dep {
    docker-compose -f docker-compose-dev.yml up dep
}

function build {
    docker build -t go-tools .
}

function serve {
    build && dep && run
}

function run {
    docker-compose -f docker-compose-dev.yml up app
}

"$@"