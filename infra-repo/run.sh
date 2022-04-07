#!/bin/bash

while read line; do
  curl (cd /app/repo/plugins/ && $line -O)
done < grafana-plugins.txt