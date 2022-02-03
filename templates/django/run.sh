#!/bin/bash

IMAGE=_username/_appname

# check if is running from the root directory
# if yes cancekl execution
if [ -d "src" ]; then 
    echo "-------------------"
    echo "you can not run docker from this folder"
    echo "go the built folder to run your project"
    echo 
    exit 0
fi

# running locally
# shows that the image has been built locally
if [[ $1 == --local ]]; then 
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then 
        sudo docker-compose up -d --remove-orphans 
    else 
        docker-compose up -d --remove-orphans 
    fi

    exit 0
elif [[ $1 == --server ]]; then

    # check operating system
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then 

        #  check command
        if [[ $2 == stop ]]; then 
            # stopping services
            sudo docker-compose down
            exit 0
        elif [[ $2 == restart ]]; then
            # retart services
            sudo docker-compose down
            sudo docker-compose up -d
            exit 0
        fi

        # if on sever we need to remove the image and run the commmand
        # when we are developing we do not need to remove the image. 

        # check if there is a image exists to remove it
        if [[ ! $(sudo docker inspect $IMAGE | grep -F []) == '[]' ]]; then 
            echo .... removing image : $IMAGE, 
            sudo docker-compose down --rmi all
        fi

        sudo docker-compose build --force-rm --remove-orphans 
        sudo docker-compose up -d 

    else

        # check command
        if [[ $2 == stop ]]; then 
            # stopping services
            docker-compose down
            exit 0
        elif [[ $2 == restart ]]; then
            # retart services
            docker-compose down
            docker-compose up -d
            exit 0
        fi

        # check if there is a image exists to remove it
        if [[ ! $(docker inspect $IMAGE | grep -F []) == '[]' ]]; then 
            echo .... removing image : $IMAGE, 
            docker-compose down --rmi all
        fi
        
        # start docker services
        docker-compose build --force-rm --remove-orphans 
        docker-compose up -d

    fi

else 
    echo "------------------"
    echo "Specify the environment you are running this project from. "
    echo "Syntax ./run.sh [environmet]"
    echo "environment - is either --local or --server"
fi