# Bike App
 
A web application for displaying bike data from a backend, the said backend, and a data importer.

=== 

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


===

## Use

Run the following command in the backend folder to start the backend
  
```
go run ./cmd/backend/
```
> The backend requires the `BIKE_APP_PW` environment variable to be the database user's password.

Run the following command in the frontend folder to start the frontend

```
npm start
```

> The frontend requires that both the frontend and the backend are running to show data.

### Frontend manual

#### Header navigation

The header at the top of the page contains interactive buttons to alter the page and data presented.

* `Data language` buttons affect the language of station data such as names, addresses, and cities.

* `Stations` and `Bike rides` tabs open views for viewing station and bike ride data respectively.

#### Stations view

Stations view displays lists and details of the station data retrieved from the backend.


##### List of stations

The default view lists all the stations.

* Filters
  * Text filter
    * Filters the listed stations by text inclusion
    * Uses the stringified JSON of the station object and as such doesn't filter out some words eg. object keys like `city`
  * Min ID
    * Sets the minimum ID. Any station with an ID smaller than the given value will be filtered out
  * Max ID
    * Sets the Maximum ID. Any station with an ID greater than the given value will be filtered out
    
* Station list
  * Page buttons
    * For your convenience, the listed stations are split into pages
    * Click the buttons to change the page
  * List of stations
    * The header columns can be clicked to change the ordering of the list
    * The station rows can be clicked to view details about the station

* Relative position map
  * Uses the station coordinates to draw the station locations relative to each other
  * North is up. West is left.
  * The station dots change based on user actions
      * The default station representation is a gray dot
      * Stations on the current page of the list are red dots
      * Stations filtered out by the filters are light pink dots
      
      
##### Station details

Can be opened by clicking a station in the stations list.

The return button is at the very top and can be identified by its dark red color and informative labeling.

* Details
  * The first part of the page is a unlabeled and noninteractive details table
* Statistics
  * Retrieves and displays information about the station from the backend
    * Can be filtered by all data or by month using the `Filter statistics` buttons
    * The information is grouped by whether the bike ride started or ended at the station
* Relative position map
  * The station being looked at is shown as a red dot

#### Bike rides view

Lists bike rides.

The data is retrieved from the backend for each and every list modifying action in this view.

* Station Search
  * Search for rides by departure or return station name
  * Uses wildcard search in the database
* Ride list
  * Uses pagination and has page controls at the top
    * LIMIT OFFSET pagination for simplicity
  * Ordering can be changed by clicking the table headers
