#!/bin/sh
find exercises -perm +0111 -type f -exec rm -f {} \;
find examples -perm +0111 -type f -exec rm -f {} \;
