#!/bin/bash

while read line; do
  curl (cd /plugins/ && $line -O)
done < grafana-plugins.txt