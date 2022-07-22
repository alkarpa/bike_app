# TODO
- database
    - creation scripts and env variables
    - foreign keys
    - change station names to a station_name table with (ref to station id, language key, value)
- back end
    - database access
        - get pagination for rides
    - REST API
        - fix the unreliable server tests
- front end
    - fetching data
    - displaying data
 
===

# Parts

## React Front end

The web application that provides the user with a view into the data.

===
## Go Back end

REST API that provides the data to the front end.

===
## MariaDB Database

Stores the related data.

db: bike_app

station

| column | type |
|--------|----|
| id | int |
| operator | varchar |
| capacity | int |
| x | double|
| y | double|

station_lang_field

| column | type |
|--------|----|
| id | int |
| lang | varchar |
| key | varchar |
| value | varchar |

trip

| column | type |
|--------|------|
| id | int |
| from_station | int |
| to_station | int | 
| distance | int |
| duration | int |


===
## Go .csv reader

Reads .csv files and inserts the related data into the database.
