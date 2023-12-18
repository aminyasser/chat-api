#!/bin/sh
go run main.go &
go run queue/consumer/consume.go