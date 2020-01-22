#!/usr/bin/env bash

script_name=$0

function start() {
    echo "Terminal sharing start. tail this file.---> log.session"
    echo "if Terminal sharing finished, do exit command."
    LANG=C script -fqt 2> log.time >(awk '{print strftime("%F %T ") $0}{fflush() }'> log.session)
}

function play() {
    echo "Terminal play start."
    scriptreplay -t \
        <(awk 'BEGIN{i=0}{cmd="cat log.session|cut -c 21-|tail -c +"i" \
            |head -c "$2"|wc -l";cmd|getline c;i+=$2;print $1,$2+c*20}' log.time) log.session
}

function usage() {
    echo "${script_name} s|p"
    echo "s: start process"
    echo "p: play terminal"
    exit 1
}


case $1 in
    "s" ) start;;
    "p" ) play;;
     *  ) usage;;
esac
