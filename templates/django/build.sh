#!/bin/bash

build () {

    # check for the type of OS 
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then 

        # remove production file if exists
        if [ -d "$1-_appname" ]; then 
            sudo rm -r $1-_appname
        fi

        # create production folder contents
        sudo mkdir $1-_appname
        sudo cp -r nginx $1-_appname/
        sudo cp docker-compose.yml $1-_appname/
        sudo cp run.sh $1-_appname/

        # its linux
        if [[ $1 == test ]]; then 
            _rev="_username/_appname"
            _put="_username/$1-_appname"
            _name="_appname"
            _newname="$1-_appname"
            
            sudo sed -i "s|$_rev|$_put|g" $1-_appname/docker-compose.yml
            sudo sed -i "s|$_rev|$_put|g" $1-_appname/run.sh
            sudo sed -i "s|$_name|$_newname|g" $1-_appname/nginx/nginx.conf
            sudo sed -i "s|$_name|$_newname|g" $1-_appname/docker-compose.yml
            sudo sed -i "s|$_name|$_newname|g" .github/workflows/deploy-test.yml
            sudo sed -i "s|$_name|$_newname|g" .github/workflows/deploy-prod.yml

            # since it created double $1-$1 on the image, remove that
            _double_name="$1-$1-_appname"
            _singlename="$1-_appname"

            sudo sed -i "s|$_double_name|$_singlename|g" $1-_appname/docker-compose.yml

            # build
            sudo docker build -t _username/$1-_appname src/
        else 

            # build
            sudo docker build -t _username/_appname src/
        fi
    else

        # remove production file if exists
        if [ -d "$1-_appname" ]; then 
            rm -r $1-_appname
        fi

        mkdir $1-_appname
        cp -r nginx $1-_appname/
        cp docker-compose.yml $1-_appname/
        cp run.sh $1-_appname/

        # its windows
        if [[ $1 == test ]]; then 
            _rev="_username/_appname"
            _put="_username/$1-_appname"
            _name="_appname"
            _newname="$1-_appname"

            sed -i "s|$_rev|$_put|g" $1-_appname/docker-compose.yml
            sed -i "s|$_rev|$_put|g" $1-_appname/run.sh
            sed -i "s|$_name|$_newname|g" $1-_appname/nginx/nginx.conf
            sed -i "s|$_name|$_newname|g" $1-_appname/docker-compose.yml
            sed -i "s|$_name|$_newname|g" .github/workflows/deploy-test.yml
            sed -i "s|$_name|$_newname|g" .github/workflows/deploy-prod.yml

            # since it created double $1-$1 on the image, remove that
            _double_name="$1-$1-_appname"
            _singlename="$1-_appname"

            sed -i "s|$_double_name|$_singlename|g" $1-_appname/docker-compose.yml

            # build
            docker build -t _username/$1-_appname src/
        else 
            
            # build
            docker build -t _username/_appname src/
        fi

    fi

}


# running tasks
if [[ $1 == "prod" ]]; then
    echo "Building for production"
    build "prod"
elif [[ $1 == "test" ]]; then
    echo "Building for testing"
    build "test"
else 
    echo
    echo "-------------"
    echo "specify if you are building for production or test"
    echo "syntax: ./build.sh [type]"
    echo "[type] - prod or test"
fi

