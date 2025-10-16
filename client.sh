#!/usr/bin/env bash

HOST="${1:-127.0.0.1}"
PORT="${2:-8080}"
FIFO="${3:-response}"

# Проверяем, что FIFO существует
if [[ ! -p "$FIFO" ]]; then
  echo "Нет FIFO '$FIFO'. Запустите сервер, чтобы он создал его" >&2
  exit 1
fi

# Открываем FIFO для чтения заранее, чтобы сервер мог в него писать
cat "$FIFO" & READER_PID=$!

# Отправляем минимальный HTTP GET /
printf 'GET / HTTP/1.1\r\nHost: %s\r\nConnection: close\r\n\r\n' "$HOST" | nc "$HOST" "$PORT"

# Ждём завершения чтения ответа
wait "$READER_PID"