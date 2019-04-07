#!/bin/bash

doGit() {
    local comment=$1
    git add .
    git commit -m $comment
    git push origin master
}

doGit $1

