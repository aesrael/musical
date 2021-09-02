# musical 🎸🎧

[![.github/workflows/send-track.yml](https://github.com/aesrael/musical/actions/workflows/send-track.yml/badge.svg)](https://github.com/aesrael/musical/actions/workflows/send-track.yml)

[![.github/workflows/backup-db.yml](https://github.com/aesrael/musical/actions/workflows/backup-db.yml/badge.svg)](https://github.com/aesrael/musical/actions/workflows/backup-db.yml)

Building a job runner and exec environment around https://github.com/spotDL/spotify-downloader

## how this works (optional flow)
simply create an [issue comment](https://github.com/aesrael/musical/issues) with the url of a spotify track or playlist, and it is automatically uploaded to a drive folder.

**NOTE:** GitHub issues/actions is not required for this to work, just any http request to said endpoints will do, using github actions is merely a matter of convenience and personal choice

## prerequisites
### running locally
* Install go
* Install redis (brew install redis)

* Install the spotify-downloader executable (https://github.com/spotDL/spotify-downloader)
note that this requires ffmpeg in order to work, so be sure to install ffmpeg also.

## start project
```bash
make build-run
```

### or using docker

docker
```bash
docker-compose build
docker-compose up
```

## how it works
This service consists of an http server, a redis database (used also as a queue & cache) and a set of workers consuming jobs published to said queue.

### Role of redis
Using [asynq](https://github.com/koddr/tutorial-go-asynq), Redis is utilized in this project as a durable queue.


### Role of github actions
comments on any [issue](https://github.com/aesrael/musical/issues) in this repo, trigger an action workflow, which enqueues jobs(tracks to download in the redis database), once these tracks are downloaded, they are uploaded to google drive, after which a cleanup occurs.

A backup job also happens at given time intervals to backup the redis db to google drive.

### Role of google drive
Using the [google drive API](https://developers.google.com/drive/api/v3/reference), authorized by means of a service account, tracks are uploaded to a google drive bucket, it also acts as a backup for the redis db in the event of a failure or crash.

connections are authenticated via [google service accounts](https://console.cloud.google.com/projectselector2/iam-admin/serviceaccounts)

for a more visual guide to service accounts see these

• [What are Service Accounts?](https://www.youtube.com/watch?v=xXk1YlkKW_k)

• [Uploading files to Google Drive API with a service account](https://www.youtube.com/watch?v=Q5b0ivBYqeQ)

http server is exposed on port `8999`
the endpoints are
| endpoint     | method | body                          | description                     |
| ----------- | --------|-------------------------------|---------------------------------|
| /api/job    | POST    | {"track": track\|playlist url} | enqueue a new track for download
| /api/backup | GET     | -                             | backup db                 |

both endpoint requires an `Authorization` header with any jwt token generated from the secret key.

### Role of asynqmon
using [asynqmon](https://github.com/hibiken/asynqmon) you can monitor and manage the jobs, install the binary and allow it access to your redis db.
