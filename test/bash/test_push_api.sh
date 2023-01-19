#!/usr/bin/bash

curl -X POST \
  -d '{"token": "fsdfljs3123l1k23klj12k","channel": "test","content": "test rumia"}' \
  -H 'Content-Type:application/json' \
  http://127.0.0.1:8081/push