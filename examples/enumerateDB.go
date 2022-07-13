package main

// Simple program to list databases and the tables

import (
	"context"
	"database/sql"
	"log"

	impala "github.com/MicePilot/go-impala"
)

func main() {

	opts := impala.DefaultOptions

	opts.Host = ""
	opts.Port = "21050"

	// enable LDAP authentication:
	opts.UseLDAP = false
	opts.Username = "<ldap username>"
	opts.Password = "<ldap password>"
	// enable TLS
	opts.UseTLS = false
	opts.CACertPath = "/path/to/cacert"

	opts.Database = ""

	connector := impala.NewConnector(&opts)
	db := sql.OpenDB(connector)
	defer db.Close()

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "SHOW TABLES")
	if err != nil {
		log.Fatal(err)
	}

	var name string
	var tables []string
	for rows.Next() {
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		tables = append(tables, name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println("All table names of the current connected database", tables)

	rows, err = db.QueryContext(ctx, "SHOW DATABASES")
	if err != nil {
		log.Fatal(err)
	}

	r := struct {
		name    string
		comment string
	}{}

	databases := make([]string, 0) // databases will contain all the DBs to enumerate later
	for rows.Next() {
		if err := rows.Scan(&r.name, &r.comment); err != nil {
			log.Fatal(err)
		}
		databases = append(databases, r.name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println("List of Databases", databases)

	stmt, err := db.PrepareContext(ctx, "SHOW TABLES IN @p1")
	if err != nil {
		log.Fatal(err)
	}

	tbl := struct {
		name string
	}{}

	for _, d := range databases {
		rows, err := stmt.QueryContext(ctx, d)
		if err != nil {
			log.Printf("error in querying database %s: %s", d, err.Error())
			continue
		}

		tables := make([]string, 0)
		for rows.Next() {
			if err := rows.Scan(&tbl.name); err != nil {
				log.Println(err)
				continue
			}
			tables = append(tables, tbl.name)
		}
		log.Printf("List of Tables in Database %s: %v\n", d, tables)
	}
}
