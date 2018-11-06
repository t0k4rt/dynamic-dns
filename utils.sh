#!/usr/bin/env bash
set -e

# get real full path
realpath () {
  local path=$1
  cd $path && pwd
}

# compute gopath
gopath () {
  local FULLPATH=$1
  local PATH=""
  local valid=false
  OIFS=$IFS
  IFS='/'
  for x in $FULLPATH
  do
      if [[ $x == "src" ]]; then
        valid=true
        break
      fi
      if [ ! -e $x ]; then
        PATH="${PATH}/${x}"
      fi
  done
  IFS=$OIFS
  if [ $valid = true ]; then
    echo $PATH
  else
    echo "ERROR: could not infer gopath" 1>&2
    exit 1
  fi
}
