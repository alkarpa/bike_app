CREATE USER IF NOT EXISTS  bike_app_user@localhost IDENTIFIED BY 'BIKE_APP_PW';
CREATE USER IF NOT EXISTS  bike_app_test_user@localhost IDENTIFIED BY 'BIKE_APP_TEST_PW';

CREATE DATABASE IF NOT EXISTS bike_app;
CREATE DATABASE IF NOT EXISTS bike_app_test;

GRANT ALL PRIVILEGES ON `bike_app`.* TO 'bike_app_user'@'localhost';
GRANT ALL PRIVILEGES ON `bike_app_test`.* TO 'bike_app_test_user'@'localhost';

USE bike_app;
SOURCE bike_app.sql;
