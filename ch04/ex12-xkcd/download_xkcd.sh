#!/bin/sh

# create dir
rm -fr archives
mkdir archives

# get latest commic NUMBER
latest=`curl -s 'http://xkcd.com/info.0.json' | grep -oE '\"num\"\s*:\s*[0-9]+' | grep -oE '[0-9]+'`
echo "latest = ${latest}"

# downlaod json of archived commic from http://xkcd.com
for i in `seq 1 ${latest}`
do
  n=`printf %04d ${i}`
  URL="http://xkcd.com/${i}/info.0.json"
  OUT="archives/xkcd${n}.json"
  echo "download ${URL} to ${OUT}"
  curl -s ${URL} > ${OUT}
done