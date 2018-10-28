#!/bin/bash

for LETTER in A B C D E F G H I J K L M N O P Q R S T U V W X Y Z; 
do
	docker run -d --rm --name=alphabet-pingpong-${LETTER} alphabet-pingpong:slices --url=nats://192.168.99.100:4222 --letter=${LETTER}; 
done
