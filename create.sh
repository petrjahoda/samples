#!/usr/bin/env bash
cd linux
upx sample_linux
cd ..
docker rmi -f petrjahoda/sample:latest
docker build -t petrjahoda/sample:latest .
docker push petrjahoda/sample:latest