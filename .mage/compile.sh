#!/bin/sh

set -e

cd .mage
exec go run mage.go -compile ../mage

