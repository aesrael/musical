# musical

[![.github/workflows/send-track.yml](https://github.com/aesrael/musical/actions/workflows/send-track.yml/badge.svg)](https://github.com/aesrael/musical/actions/workflows/send-track.yml)

[![.github/workflows/backup-db.yml](https://github.com/aesrael/musical/actions/workflows/backup-db.yml/badge.svg)](https://github.com/aesrael/musical/actions/workflows/backup-db.yml)

Building a job runner and exec environment around https://github.com/spotDL/spotify-downloader

## prerequisites
* install redis (brew install redis)

* install the spotify-downloader executable (https://github.com/spotDL/spotify-downloader)
note that this requies ffmpeg in other to work, so be sure to install ffmpeg also

## start server
```bash
make run
```

