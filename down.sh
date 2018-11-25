#!/bin/bash

for LETTER in {A..Z}
do   
    echo Stopping alphabet-pingpong-${LETTER};
	docker stop alphabet-pingpong-${LETTER}; 
done
