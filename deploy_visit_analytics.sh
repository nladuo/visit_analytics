#!/bin/bash

GIN_MODE=release nohup ./visit_analytics config.yaml 2>&1 > web.log &
