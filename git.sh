#!/bin/bash

doGit() {
    local comment=$1
    if [ -z $comment ]; then
	echo "Please provide a comment"
	exit 1
    fi
    git add .
    git commit -m $comment
    git push origin master
}

doGit $1

