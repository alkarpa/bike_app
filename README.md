# Bike App
 
A web application for displaying bike data from a backend, the said backend, and a data importer.


## Setup

### CSV Data

Bike App is designed to process very particular biking and bike station datasets. You can download the files manually below or by using the Bash setup script (requires curl for download).

| Bike ride data owned by City Bike Finland |
|----------------|
| https://dev.hsl.fi/citybikes/od-trips-2021/2021-05.csv |
| https://dev.hsl.fi/citybikes/od-trips-2021/2021-06.csv |
| https://dev.hsl.fi/citybikes/od-trips-2021/2021-07.csv |

| Station data link | Contains |
|--------------|-----|
| https://opendata.arcgis.com/datasets/726277c507ef4914b0aec3cbcfcbfafc_0.csv | Data file |
| https://www.avoindata.fi/data/en_GB/dataset/hsl-n-kaupunkipyoraasemat/resource/a23eef3a-cc40-4608-8aa2-c730d17e8902 | License |

The importer expects the .csv files inside the data folder.
Station and bike ride .csv files should be placed in the `data/station` and `data/ride` folders respectively.


### Database

Bike App uses `MariaDB` to store and process the imported biking data.

You will need

1. mariadb installed and running
2. mariadb user and database creation privileges

Set up the users with password identification and according to the table below so that

* the production user `bike_app_user` has INSERT and SELECT privileges in the production database `bike_app`
* the testing user `bike_app_test_user` has all privileges (create table, drop table, etc.) in the testing database `bike_app_test`
* the user passwords are something secure chosen by you

| Purpose | Database name | Database user | User password environment variable |
|---------|---------------|---------------|------------------------------------|
|Production| `bike_app`     | `bike_app_user` | `BIKE_APP_PW` |
|Testing | `bike_app_test` | `bike_app_test_user` | `BIKE_APP_TEST_PW` |


The backend uses BIKE_APP_PW environment variable to log into the database. On Linux it can be applied while starting the backend through the command line,
eg.

```BIKE_APP_PW=bad_security go run ./backend/cmd/backend```

though this obviously shows the password in plaintext and likely stores it in your command history.

### Backend

The backend is made using Go and requires a Go compiler to build and run.

There are two backend commands

| command | purpose |
|---------|:---------|
| `cmd/ImportCSV` | Imports data from the .csv files in ./data/ . Only needs to be run once. |
| `cmd/backend` | Serves the http backend |

### Frontend

The frontend is a npm React project and requires npm to build and run.
