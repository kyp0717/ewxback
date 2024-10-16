#!/bin/bash

docker run --name postgres-ew -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres

