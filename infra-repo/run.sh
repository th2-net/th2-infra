#!/bin/bash

cd plugins/ 

while read line; do
  curl $line -O
done < grafana-plugins.txt