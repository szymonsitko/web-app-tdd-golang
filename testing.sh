#!/bin/bash
#  Script running tests in right orser

go test functional_test.go -v
go test webapp_test.go -v
