# TODO
- .csv reader
- database
    - creation scripts and env variables
    - foreign keys
    - change `ride` primary to an autoincremented int to account for very similar rides
- back end
    - database access
        - csv reader data should use this db access to insert
    - REST API
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
| name | string |

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
