#!/bin/bash

echo "====================="
echo "Bike App Setup Script"
echo "====================="

if [ ! -d "./frontend" ] || [ ! -d "./backend" ]; then
    echo "The script uses relative paths and should be run in the bike_app folder."
    exit 1
fi

# Required software installed check
type mysql >/dev/null 2>&1 || { echo "mariadb not found. Install mariadb and run the script again."; exit 1; }
type go >/dev/null 2>&1 || { echo "go not found. Install go and run the script again."; exit 1; }
type npm >/dev/null 2>&1 || { echo "npm not found. Install npm and run the script again."; exit 1; }

# Environment variable check
if [[ -z "${BIKE_APP_PW}" ]]; then
    echo "Environment variable BIKE_APP_PW not found"
    echo "Please input the password for database user bike_app_user:"
    read -s BIKE_APP_PW
    export BIKE_APP_PW
fi
if [[ -z "${BIKE_APP_TEST_PW}" ]]; then
    echo "Environment variable BIKE_APP_TEST_PW not found"
    echo "Please input the password for database user bike_app_test_user:"
    read -s BIKE_APP_TEST_PW
    export BIKE_APP_TEST_PW
fi

# Database check
echo
echo "-- Setting up the database"
DBYN=''
while [ "$DBYN" != "skip" ];
do
    echo "The script can handle setting up the MariaDB users and databases if provided with an user with"
    echo " -ALL PRIVILEGES ON *.* WITH GRANT OPTION"
    echo " -password identification"
    echo
    echo "If this is not an option for you, set up the databases and users manually as described in the README.md and ./*.sql files and select 'skip'."
    echo
    echo "Do you want to let the script set up the users and databases? (yes/skip)"
    read DBYN

    if [ $DBYN == "yes" ]; then
        echo "Input database user with ALL PRIVILEGES ON *.* WITH GRANT OPTION:"
        read DBU
        echo "Input $DBU password"
        read -s DBPW

        echo "Creating the databases and users" 
        sed "s/BIKE_APP_PW/${BIKE_APP_PW}/g; s/BIKE_APP_TEST_PW/${BIKE_APP_TEST_PW}/g; s/DBU/${DBU}/g" ./users_databases.sql | mysql -u "$DBU" -p"$DBPW"

        if [ $? -eq 0 ]; then
            echo "Database operations successful"
        else
            echo "Database operations failed."
            exit 1
        fi
        break
    fi
done


# Data check and acquisition
echo
DATADIR="./data"

if [ ! -d $DATADIR ]; then
    mkdir $DATADIR
    mkdir $DATADIR/ride
    mkdir $DATADIR/station
else
    echo "Listing files in ./data/station and ./data/ride"
    FILECOUNT=0
    STATIONCOUNT=0
    RIDECOUNT=0
    for f in ./data/station/*
    do
        echo "$f"
        ((++FILECOUNT))
        ((++STATIONCOUNT))
    done
    for f in ./data/ride/*
    do
        echo "$f"
        ((++FILECOUNT))
        ((++RIDECOUNT))
    done
    echo "Found ${FILECOUNT} files."
    if [[ $STATIONCOUNT == 0 || $RIDECOUNT == 0 ]]; then
        echo
        echo "Missing data files. Downloading the data files is recommended."
        echo "You may also manually copy the station .csv files into ./data/station and the ride .csv files into ./data/ride now before proceeding."
    fi
fi

echo
DOWNLOADYN=''
while [ "$DOWNLOADYN" != "no" ];
do
    echo "Do you want to download the station/journey data using curl? (yes/no/license)"
    read DOWNLOADYN

    if [ $DOWNLOADYN == "yes" ]; then
        type curl >/dev/null 2>&1 || { echo "curl not found. Install curl and run the script again."; exit 1; }
        echo "https://opendata.arcgis.com/datasets/726277c507ef4914b0aec3cbcfcbfafc_0.csv"
        curl -o "$DATADIR/station/stations.csv" https://opendata.arcgis.com/datasets/726277c507ef4914b0aec3cbcfcbfafc_0.csv
        echo "https://dev.hsl.fi/citybikes/od-trips-2021/2021-{05,06,07}.csv"
        curl -L "https://dev.hsl.fi/citybikes/od-trips-2021/2021-{05,06,07}.csv" -o "$DATADIR/ride/2021-#1.csv"
        break
    elif [ $DOWNLOADYN == "license" ]; then
        echo "Bike journey data is owned by City Bike Finland."
        echo "City Bicycle Station data license and information can be found at https://www.avoindata.fi/data/en/dataset/hsl-n-kaupunkipyoraasemat/resource/a23eef3a-cc40-4608-8aa2-c730d17e8902"
    fi
done

# Build the frontend
echo
echo "-- Build the frontend and copy it to the backend static folder"
cd frontend
npm install
npm run buildandcopy

# Run the csv_reader tests and generate a coverage HTML file
echo
echo "-- Running csv_reader tests and generating a coverage HTML file"
cd ../csv_reader
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
echo
echo "csv_reader test coverage report available: csv_reader/coverage.html"

# Import the CSV data
echo
echo "-- Import data from CSV"
cd ../backend
go run ./cmd/ImportCSV

# Run the backend tests and generate a coverage HTML file
echo
echo "-- Running backend tests and generating a coverage HTML file"
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
echo
echo "Backend test coverage report available: backend/coverage.html"

# Run the backend
echo
echo "-- Starting the backend"
go run ./cmd/backend
