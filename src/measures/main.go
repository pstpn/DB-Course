package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/lib/pq"
)

func uploadPostgres() {
	connStr := "user=postgres dbname=testdb password='admin' sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS images (id SERIAL PRIMARY KEY, data BYTEA)")
	if err != nil {
		log.Fatal(err)
	}

	imageFiles := []string{"data/1.jpg", "data/2.jpg", "data/3.jpg", "data/4.jpg", "data/5.jpg", "data/6.jpg", "data/7.jpg"}
	for _, file := range imageFiles {
		data, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec("INSERT INTO images (data) VALUES ($1)", data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Uploaded %s to PostgreSQL\n", file)
	}
}

func uploadMongo() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	var result bson.M
	if err = client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		log.Fatal(err)
	}

	db := client.Database("testdb")
	bucketOpts := options.GridFSBucket().SetName("images")
	bucket, err := gridfs.NewBucket(db, bucketOpts)
	if err != nil {
		log.Fatal(err)
	}

	imageFiles := []string{"data/1.jpg", "data/2.jpg", "data/3.jpg", "data/4.jpg", "data/5.jpg", "data/6.jpg", "data/7.jpg"}
	for _, file := range imageFiles {
		data, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		uploadOpts := options.GridFSUpload().SetMetadata(bson.D{{
			Key:   file,
			Value: len(data)},
		})
		_, err = bucket.UploadFromStream(
			file,
			bytes.NewReader(data),
			uploadOpts,
		)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Uploaded %s to MongoDB\n", file)
	}
}

func measurePostgres() {
	connStr := "user=postgres dbname=testdb password='admin' sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	start := time.Now()
	rows, err := db.Query("SELECT data FROM images")
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Printf("Time to select: %v\n", elapsed)
	defer rows.Close()

	var times []time.Duration
	for rows.Next() {
		var data []byte
		start := time.Now()
		err := rows.Scan(&data)
		if err != nil {
			log.Fatal(err)
		}
		elapsed := time.Since(start)
		times = append(times, elapsed)
	}

	for i, t := range times {
		fmt.Printf("Time to fetch image %d from PostgreSQL: %v\n", i+1, t)
	}
}

func measureMongo() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	var result bson.M
	if err = client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		log.Fatal(err)
	}

	db := client.Database("testdb")
	bucketOpts := options.GridFSBucket().SetName("images")
	bucket, err := gridfs.NewBucket(db, bucketOpts)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	cursor, err := bucket.Find(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Printf("Time to find: %v\n", elapsed)
	defer cursor.Close(context.Background())

	var times []time.Duration
	for cursor.Next(context.Background()) {
		var result bson.M
		start := time.Now()
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		elapsed := time.Since(start)
		times = append(times, elapsed)
	}

	for i, t := range times {
		fmt.Printf("Time to fetch image %d from MongoDB: %v\n", i+1, t)
	}
}

func main() {
	uploadPostgres()
	uploadMongo()

	measurePostgres()
	measureMongo()
}
