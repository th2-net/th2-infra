#!/bin/sh

while read line; do
wget "$line"
done < grafana-plugins.txt