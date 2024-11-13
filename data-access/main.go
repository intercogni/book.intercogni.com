package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

var db *sql.DB

func getAlbumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("select * from album where artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb Album
		err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
		if err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

func getAlbumByID(id int64) (Album, error) {
	var alb Album

	row := db.QueryRow("select * from album where id = ?", id)
	err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec(
		"insert into album (title, artist, price) values (?, ?, ?)",
		alb.Title, alb.Artist, alb.Price,
	)

	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil
}

func main() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	ping_err := db.Ping()
	if ping_err != nil {
		log.Fatal(ping_err)
	}
	fmt.Println("connected to recordings mega-library!!! (also called database)")

	albums, err := getAlbumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("albums found by artist: %v\n", albums)

	_alb, _err := getAlbumByID(2)
	if _err != nil {
		log.Fatal(_err)
	}
	fmt.Printf("album found by id: %v\n", _alb)

	__alb, __err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})

	if __err != nil {
		log.Fatal(__err)
	}
	fmt.Printf("id of added album: %v\n", __alb)
}
