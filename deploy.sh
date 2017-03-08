#!/bin/bash

GIN_MODE=release ./visit_analytics config.yaml 2>&1  > web.log
