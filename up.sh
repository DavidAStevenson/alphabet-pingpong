#!/bin/bash

for LETTER in A B C D E F G H I J K L M N O P Q R S T U V W X Y Z; 
do
	docker run -d --rm --name=alphabet-pingpong-${LETTER} ds35472/alphabet-pingpong:latest --url=nats://demo.nats.io:4222 --letter=${LETTER};
# If running gnatsd locally (e.g. via docker run gnatds) or demo.nats.io is unavailable, use below one instead
# 	docker run -d --rm --name=alphabet-pingpong-${LETTER} ds35472/alphabet-pingpong:latest --url=nats://192.168.99.100:4222 --letter=${LETTER};
done
