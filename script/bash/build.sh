#!/bin/bash

docker build -f Dockerfile -t machengcheng/12306:v0.1 .
docker push machengcheng/12306
