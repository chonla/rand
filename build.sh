#!/bin/bash

set -eux

go mod download
go install .
