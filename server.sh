#!/bin/bash

### Create the response FIFO
rm -f response
mkfifo response

function handleRequest() {
    while read line; do
        echo $line
        trline=`echo $line | tr -d '[\r\n]'`

        [ -z "$trline" ] && break

        HEADLINE_REGEX='(.*?)\s(.*?)\sHTTP.*?'

        [[ "$trline" =~ $HEADLINE_REGEX ]] &&
            REQUEST=$(echo $trline | sed -E "s/$HEADLINE_REGEX/\1 \2/")
    done

  echo -e 'HTTP/1.1 200\r\n\r\n\r\n</h1>PONG</h1>' > response
}

case "$REQUEST" in
  "GET /") RESPONSE="HTTP/1.1 200 OK\r\nContent-Type: text/html\r\n\r\n</h1>PONG</h1>" ;;
        *) RESPONSE="HTTP/1.1 404 NotFound\r\n\r\n\r\nNot Found" ;;
esac

echo -e $RESPONSE > response