#!/bin/bash
go test -coverprofile=rssPkg.out
go tool cover -html=rssPkg.out
