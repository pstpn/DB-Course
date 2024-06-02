package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/lib/pq"
)

func insertPostgres(count int, testCount int64) (map[string]int64, map[string]int, map[string]int64) {
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

	times := make(map[string]int64)
	sizes := make(map[string]int, count)
	ids := make(map[string]int64, count)
	for num := range count {
		file := fmt.Sprintf("data/t%d.jpg", num+1)
		data, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		sizes[file] = len(data)

		var id int64
		var summ int64
		for range testCount {
			start := time.Now()
			row := db.QueryRow("INSERT INTO images (data) VALUES ($1) RETURNING id", data)
			summ += time.Since(start).Microseconds()
			err = row.Scan(&id)
			if err != nil {
				log.Fatal(err)
			}
			_, err = db.Exec("DELETE FROM images WHERE id = $1", id)
			if err != nil {
				log.Fatal(err)
			}
		}

		times[file] = summ / testCount

		row := db.QueryRow("INSERT INTO images (data) VALUES ($1) RETURNING id", data)
		err = row.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		ids[file] = id

		fmt.Printf("Inserted %s to PostgreSQL\n", file)
	}

	return ids, sizes, times
}

func uploadMongo(count int, testCount int64) (map[string]primitive.ObjectID, map[string]int64) {
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

	times := make(map[string]int64)
	ids := make(map[string]primitive.ObjectID, count)
	for num := range count {
		file := fmt.Sprintf("data/t%d.jpg", num+1)
		data, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		uploadOpts := options.GridFSUpload().SetMetadata(bson.D{{
			Key:   file,
			Value: len(data)},
		})

		var id primitive.ObjectID
		var summ int64
		for range testCount {
			start := time.Now()
			id, err = bucket.UploadFromStream(
				file,
				bytes.NewReader(data),
				uploadOpts,
			)
			if err != nil {
				log.Fatal(err)
			}
			summ += time.Since(start).Microseconds()

			err = bucket.Delete(id)
			if err != nil {
				log.Fatal(err)
			}
		}

		times[file] = summ / testCount

		id, err = bucket.UploadFromStream(
			file,
			bytes.NewReader(data),
			uploadOpts,
		)
		if err != nil {
			log.Fatal(err)
		}
		ids[file] = id

		fmt.Printf("Uploaded %s to MongoDB\n", file)
	}

	return ids, times
}

func selectPostgres(count int, testCount int64, ids map[string]int64) map[string]int64 {
	connStr := "user=postgres dbname=testdb password='admin' sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_, err = db.Exec("DELETE FROM images")
		db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	times := make(map[string]int64, count)
	for num := range count {
		file := fmt.Sprintf("data/t%d.jpg", num+1)

		var summ int64
		for range testCount {
			start := time.Now()
			rows, err := db.Query("SELECT data FROM images WHERE id = $1", ids[file])
			if err != nil {
				log.Fatal(err)
			}
			summ += time.Since(start).Microseconds()
			rows.Close()
		}

		times[file] = summ / testCount
	}

	return times
}

func findMongo(count int, testCount int64, ids map[string]primitive.ObjectID) map[string]int64 {
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

	times := make(map[string]int64, count)
	for num := range count {
		file := fmt.Sprintf("data/t%d.jpg", num+1)

		var summ int64
		for range testCount {
			start := time.Now()
			_, err = bucket.Find(bson.D{{"_id", ids[file]}})
			if err != nil {
				log.Fatal(err)
			}
			summ += time.Since(start).Microseconds()
		}

		times[file] = summ / testCount
	}

	err = db.Drop(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	return times
}

func toCSV(postgres, mongo map[string]int64, sizes map[string]int, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(f)
	defer func() {
		w.Flush()
		f.Close()
	}()
	w.Comma = ';'

	for num := range len(sizes) {
		file := fmt.Sprintf("data/t%d.jpg", num+1)
		err = w.Write([]string{
			strconv.Itoa(sizes[file] / 1024),
			strconv.Itoa(int(postgres[file])),
			strconv.Itoa(int(mongo[file])),
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	postgresIDs, sizes, insertTimes := insertPostgres(14, 100)
	mongoIDs, uploadTimes := uploadMongo(14, 100)

	selectTimes := selectPostgres(14, 100, postgresIDs)
	findTimes := findMongo(14, 100, mongoIDs)

	fmt.Println()
	fmt.Println("Sizes: ", sizes)
	fmt.Println()
	fmt.Println("Insert times Postgres: ", selectTimes)
	fmt.Println()
	fmt.Println("Upload times Mongo: ", selectTimes)
	fmt.Println()
	fmt.Println("Select times Postgres: ", selectTimes)
	fmt.Println()
	fmt.Println("Find times Mongo: ", findTimes)
	fmt.Println()

	toCSV(selectTimes, findTimes, sizes, "result1.csv")
	toCSV(insertTimes, uploadTimes, sizes, "result2.csv")
}
