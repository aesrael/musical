# musical

[![.github/workflows/send-track.yml](https://github.com/aesrael/musical/actions/workflows/send-track.yml/badge.svg)](https://github.com/aesrael/musical/actions/workflows/send-track.yml)

[![.github/workflows/backup-db.yml](https://github.com/aesrael/musical/actions/workflows/backup-db.yml/badge.svg)](https://github.com/aesrael/musical/actions/workflows/backup-db.yml)

Building a job runner and exec environment around https://github.com/spotDL/spotify-downloader

## prerequisites
* Install go
* install redis (brew install redis)

* install the spotify-downloader executable (https://github.com/spotDL/spotify-downloader)
<<<<<<< HEAD
note that this requies ffmpeg in order to work, so be sure to install ffmpeg also
=======
note that this requires ffmpeg in other to work, so be sure to install ffmpeg also
>>>>>>> 2da2ceb (dockerize app)

## how it works
This service consists of an http server, a redis database (used also as a queue & cache) and a set of workers consuming jobs published to said queue.

### Role of redis
Using [asynq](https://github.com/koddr/tutorial-go-asynq), Redis is utilized in this project as a durable queue.


### Role of github actions
comments on any [issue](https://github.com/aesrael/musical/issues) in this repo, trigger an action workflow, which enqueues jobs(tracks to download in the redis database), once these tracks are downloaded, they are uploaded to google drive, after which a cleanup occurs.

A backup job also happens at given time intervals to backup the redis db to google drive.


### Role of google drive
Using the [google drive API](https://developers.google.com/drive/api/v3/reference), authorized by means of a service account, tracks are uploaded to a google drive bucket, it also acts as a backup for the redis db in the event of a failure or crash.


## start project
```bash
make build-run
```

docker
```bash
docker-compose build
docker-compose up
```

http server is exposed on port `8999`
the endpoints are

* /api/job [POST]
```js
       body {
         track String,
       }

       headers {
         Authorization: ${TOKEN},
       }
```

* /api/backup [GET]

```js
    headers {
       Authorization: ${TOKEN},
      }
```