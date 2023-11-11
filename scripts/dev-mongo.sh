#!/bin/zsh

docker run -dit --rm -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME="root" -e MONGO_INITDB_ROOT_PASSWORD="mongopw" -v ./build/mongodb/init.js:/docker-entrypoint-initdb.d/init.js --name lizard-db lizard-db mongod
