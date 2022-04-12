#!/bin/bash

file = 'grafana-plugins.txt'

while read line; do
  curl $line -O
done < $file