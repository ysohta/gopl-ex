#!/bin/sh

# extract top 1000 web sites from data of alexa
curl -s -O http://s3.amazonaws.com/alexa-static/top-1m.csv.zip ; unzip -q -o top-1m.csv.zip top-1m.csv ; head -1000 top-1m.csv > top-1000.csv