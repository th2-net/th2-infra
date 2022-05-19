#!/bin/sh

mkdir ./plugins
while read url; do
    wget "$url" --directory-prefix ./plugins || exit 1
done < grafana-plugins.txt