# songdb
Song Database written in Go with Echo framework

I created this repository to practice my Go skills. 
I am trying to create REST API running with Echo framework.

## Requirements
* [gorm.io/gorm](https://gorm.io/)
* [github.com/labstack/echo/v4](https://echo.labstack.com/)
* [github.com/aidarkhanov/nanoid](https://github.com/aidarkhanov/nanoid)

## Routes
* `GET /songs` to get all songs in database
* `GET /songs/:id` to get one of the song in database
* `GET /albums` to get all albums in database
* `GET /albums/:id` to get one of the album in database
* `GET /albums/:id/songs` to get all songs in specified album
* `POST /songs` to create a new song to database
* `POST /albums` to create a new album to database
* `POST /albums/:id/songs` to create a new song in a specified album
* `PUT /songs/:id` to update specific song
* `PUT /albums/:id` to update specific albums
* `DELETE /songs/:id` to delete specific song
* `DELETE /albums/:id` to delete specific album

## Directory tree
```
├── cmd
│   └── server
│       └── main.go
├── go.mod
├── go.sum
└── pkg
    ├── config
    │   └── db.go
    ├── controller
    │   ├── album.go
    │   ├── song.go
    │   └── songrel.go
    ├── db
    │   ├── album.go
    │   ├── consts.go
    │   ├── song.go
    │   └── songrel.go
    ├── errors
    │   └── nodata.go
    ├── models
    │   ├── album.go
    │   └── song.go
    └── routes
        ├── album.go
        ├── song.go
        └── songrel.go
```

Here is also [the ERD](https://user-images.githubusercontent.com/73155095/236813865-0c89bce1-940a-4cd9-83c1-402849972b88.png)
[mysqldump.txt](https://github.com/dibalikpohon/songdb/files/11420669/mysqldump.txt)
