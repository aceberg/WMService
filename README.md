[![Docker](https://github.com/aceberg/WMService/actions/workflows/main-docker-all.yml/badge.svg)](https://github.com/aceberg/WMService/actions/workflows/main-docker-all.yml)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/aceberg/wmservice)

# WMService

Washing machines service ticket system    
https://github.com/aceberg/WMService


## Config

Configuration can be done through config file or environment variables

| Variable  | Description | Default |
| --------  | ----------- | ------- |
| DB_PATH    | Path to Database | /data/wmservice/rs1.db |
| GUI_PORT   | Port for web GUI | 8843 |
| THEME | Any theme name from https://bootswatch.com in lowcase | united |
| SHOW | Tickets to show on index page | 10 |
| TZ | Set your timezone for correct time | "" |

## Config file

Config file path is `/data/wmservice/config`. All variables could be set there.