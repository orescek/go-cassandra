package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gocql/gocql"
)

var (
	cluster      = "cassandra"
	keyspacename = "projekt1"
	tablename    = "test"
	filename     = "index.html"
)

func createkyespace() {
	c := gocql.NewCluster(cluster)
	c.Keyspace = "system"
	//c.Timeout = 1 * time.Second
	session, err := c.CreateSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	table := fmt.Sprintf(`CREATE KEYSPACE %s WITH replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1}`, keyspacename)
	if err := session.Query(table).RetryPolicy(nil).Exec(); err == nil {
		//log.Printf("error creating table table=%q err=%v\n", table, err)
		os.Remove(filename)
	} else {
		log.Printf("Already created keyspace")
	}

}

func createtable() {
	c := gocql.NewCluster(cluster)
	c.Keyspace = keyspacename
	//c.Timeout = 30 * time.Second
	session, err := c.CreateSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	table := fmt.Sprintf(`CREATE TABLE %s.%s (cas text, PRIMARY KEY(cas))`, keyspacename, tablename)
	if err := session.Query(table).RetryPolicy(nil).Exec(); err != nil {
		//log.Printf("error creating table table=%q err=%v\n", table, err)
		log.Printf("Already created table")
	}

}
func write_to_file(str string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err = f.WriteString(str); err != nil {
		panic(err)
	}
}

func mojprojekt(w http.ResponseWriter, r *http.Request) {

	createkyespace()
	createtable()
	cluster := gocql.NewCluster(cluster)
	cluster.Keyspace = keyspacename
	cluster.Consistency = gocql.Quorum

	session, _ := cluster.CreateSession()
	defer session.Close()

	t := time.Now()
	if err := session.Query(`INSERT INTO test (cas) VALUES (?)`, t.Format("2006-01-02 15:04:05")).Exec(); err != nil {
		log.Fatal(err)
	}

	var cas string

	// here magic happens
	os.Remove(filename)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	write_to_file("<table>")

	iter := session.Query(`SELECT cas FROM test`).Iter()
	for iter.Scan(&cas) {
		//fmt.Println("casovi:", cas)
		write_to_file("<tr><td>" + cas + "</td></tr>")
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	write_to_file("</table>")
	file.Close()

	b, err := ioutil.ReadFile(filename) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'

	//fmt.Println(str) // print the content as a 'string'
	fmt.Fprintf(w, str)

}

func main() {
	http.HandleFunc("/test", mojprojekt)
	http.ListenAndServe(":8898", nil)
}
