#!/bin/bash

DATADIR="./data"

if [ ! -d $DATADIR ]; then
    mkdir $DATADIR
    mkdir $DATADIR/ride
    mkdir $DATADIR/station
else
    echo "./data found"
fi



DOWNLOADYN=''
while [ "$DOWNLOADYN" != "no" ];
do
    echo "Do you want to download the station/journey data using curl? (yes/no/license)"
    read DOWNLOADYN

    if [ $DOWNLOADYN == "yes" ]; then
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
