#!/bin/bash
service redis-server start
sleep 5
service rabbitmq-server start
sleep 5
./main
