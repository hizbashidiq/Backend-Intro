package main

import (
	// this is a postgres driver. and most commonly used. you can also use pq or go-pg
	// go get github.com/jackc/pgx/v5
	// btw, it's better to use native pgx instead of pgx+database/sql
	// use pgx+database/sql if you plan to swap driver to mysql for example
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)


func main(){
	// in RDBMS, sometimes table called relation, rows called tuples, columns called attributes,
	// and data type of a columns called domain, a set of allowable values
	// sometimes also table = files, rows = records, columns = fields

	// condition for RDBMS:
	// All rows in a relation must be unique
	// All columns in a relation must have a unique Name
	// 		Columns with same name in different relation (table) is allowed

	// primary key must satisfy:
	// 1. unique
	// 2. not null
	// 3. stable (should not change)
	// 4. minimal (no unnecessary column)
	// 5. meaningless (ideally, no business meaning)

	// database only enforced 1 and 2
	// e.g. Students have ID and email -> email as primary is fine technically but risky design wise
	// email can be changed, length can be larged, and has real world meaning/usage

	// superkey are just a concept of every possible attributes combination that are unique
	// candidate key are a minimal superkey
	// so the hierarchy is superkey->candidatekey->primarykey
	// surrogate key -> artificial key which aim to uniquely identify every rows i.e. ID
	// Entity Relationship Diagram (ERD) has 3 concepts: entities, attributes, relationship

	// general guideline is to first create a well-normalized database first and denormalize if needed
	// normalization mostly only happen up to 3NF but there's higher normal form such as BCNF, 4NF, 5NF

	// Categories of SQL
	// Data Definition Language(DDL)-design database structure i.e. tables, views, indexes
	// Data Query Language(DQL)-retrieving data from the database
	// Data Manipulation Language(DML)-modify or write new information to the database
	// Data Control Language(DCL)-manage or authorize database access for other users
	// Transaction Control Language(TCL)

	// PostgreSQL Conventions
	// database name should be written as snake_case with suffix _db
	// Length of database name <42 char
	// Length of table name <42 char
	// field should be written as snake_case
	// foreign key should be suffixed with _id
	// primary key with int value should use BIGSERIAL as data type
	// prefer to use a numeric value that doesn't lose precision (INT, DECIMAL)
	// do not allow NULL unless necessary

	// Database in Golang
	// sql.DB isn't a database connection
	// sql.DB perform important task behind the scenes: open and close connection to the actual underlaying database
	// via the driver. It also manages a pool of connections as needed.
	// driver basically is a translator. you code something in go, driver translate it into something
	// postgreSQL can understand. it's adapter to something external


	// \l or \list to list database in postgres
	// \l+ or \list+ for more detailed information

	// [\c or \connect] database_name -> to connect to a database

	// don't use log.fatal unless it's in the startup since it's gonna immediately stop program
	// defer can't run, connections aren't closed, goroutines abruptly killed, HTTP server stop immediately
	// Backend rule: errors are data. crashes are decision. only crash when the process cannot safely continue
	// startup failure->log.fatal, recoverable error->return err, unexpected but survivable log.Printf

	// log for recording events, fmt for formatting output

	// rule: the function that opens a resource should not usually defer it's close,
	// the function that owns the lifetime shoulr

	conn, err := connect()
	if err!=nil{
		log.Fatal(err)
	}else{
		// db is a long-live object means, you should do open() and close() very rarely
		defer conn.Close(context.Background())
	}

  // iterate rows.Next() if you need control or large database, otherwise just use pgx.CollectRows()

  dbs, err := listOfDB(conn)
  if err!=nil{
    fmt.Println(err)
  }else{
    fmt.Println("LIST OF DATABASES")
    fmt.Println(dbs)
  }

  tables, err := listOfTable(conn)
  if err!=nil{
    fmt.Println(err)
  }else{
    fmt.Println("LIST OF TABLES")
    fmt.Println(tables)
  }

  table_name := "person"
	attributes, err := listOfAttributes(conn, table_name)
  if err!=nil{
    fmt.Println(err)
  }else{
    fmt.Println("LIST OF ATTRIBUTES IN TABLE",table_name)
    fmt.Println(attributes)
  }

  currentDB, err := currentDB(conn)
  if err!=nil{
    fmt.Println(err)
  }else{
    fmt.Println("Current Database: ", currentDB)
  }


  var p1 = person{}

  query := "SELECT * FROM person where id =$1"
  err = conn.QueryRow(context.Background(), query, 1).
  Scan(&p1.Id, &p1.FirstName, &p1.LastName, &p1.Gender, &p1.BirthDate, &p1.Email, &p1.BirthCountry)
  // btw queryrow only return pgx.row, .Scan that return error

  // interesting, so queryrow does store error internally, but it doesn't show up until we call scan
  // i guess this in some way is an wrap error thing?

  if err!=nil{
    fmt.Println("ERROR")
    fmt.Println(err)
  }else{
    // fmt.Printf("ID: %d\nFirst Name: %s\nLast Name: %s\n", id, first_name, last_name)
    fmt.Println(p1)
    fmt.Printf("%#v %#v %#v\n", *p1.Gender, *p1.BirthDate, p1.Email)
  }
  b, _ := json.Marshal(p1)
  fmt.Println(string(b))

  c := p1.BirthDate
  d := *p1.BirthDate
  e := &p1.BirthDate
  f := time.Now().Format("02-01-2006")
  g := time.Now()
  fmt.Printf("%T %T %T %T %T\n", c,d,e,f,g)

  // so time.Time.Format() is fine and I understand that means *time.Time.Format() also fine
  // but why I can't do *p1.BirthDate.Format since it's change *time.Time -> time.Time
  // why I need to write it as p1.BirthDate.Format which *time.Time?
  // lol, so *p1.BirthDate.Format = *(p1.BirthDate.Format) not (*p.BirthDate).Format

  query = `SELECT * FROM person WHERE id=$1`
  _, err = conn.Prepare(context.Background(), "get_user_by_id", query)
  if err!=nil{
    fmt.Println(err)
  }
  var p100 person
  err = conn.QueryRow(context.Background(), "get_user_by_id",100).
        Scan(&p100.Id, &p100.FirstName, &p100.LastName, &p100.Gender, &p100.BirthDate, &p100.Email, &p100.BirthCountry)
  if err!=nil{
    fmt.Println(err)
  }else{
    fmt.Println(p100)
  }
  var p500 person
  err = conn.QueryRow(context.Background(), "get_user_by_id",500).
        Scan(&p500.Id, &p500.FirstName, &p500.LastName, &p500.Gender, &p500.BirthDate, &p500.Email, &p500.BirthCountry)
  if err!=nil{
    fmt.Println(err)
  }else{
    fmt.Println(p500)
  }

  // in pgx, prepared statement is connection-scoped, not pool-scoped.
  // preparing on one connection does not prepare it on others.
  // pgx encourage to use implicit prepared statement. it's more advanced.
  // you need to understand statement caching and control it via pgx.ConnConfig or pgxpool.Config
  // only use explicit prepare when it's extremely high QPS, you managing long-lived connection (not generic pool),
  // and you want explicit control over statement naming
  // typical case: low-level data access layers, long-running bg workers, specialized performance-critical path

  // query = `INSERT INTO PERSON(first_name, last_name, email, gender, date_of_birth, country_of_birth)
  // query = `INSERT INTO PERSON VALUES
  // (1000,'Sup', 'Brik', 'Male', '2000-01-01', 'supbrik@gmail.com', 'Indonesia')`
  // commandTag, err := conn.Exec(context.Background(), query)
  // if err!=nil{
  //   fmt.Println(err)
  // }

  // fmt.Println("isInsert:", commandTag.Insert(), "rowsAffected:", commandTag.RowsAffected())
  // in pgx it's still the same, .query / .queryRow for retrieve data,
  // .exec for data manipulation (delete, insert, update)

  // transaction is when you do more than one modify data command in sql but you need all of them
  // to success or nothing happen at all (all or nothing), like you transfer money from A to B
  // you need to substract from A AND add to B, you can't just substract from A but then the add to B failed
  // in this case, use transaction (Tx)
  // begin are normal transaction while beginTx provide more control and more advance (isolation level, etc)

  // so, prepared statement != prepared queries? idk

  // normal prepared statement under the hood tied with certain connection, if that connection busy, it'll
  // automatically reprepared using other connection available and will be tied with that connection, and if later
  // that connection also busy when the prepare want to use it, it's back to reprepared to other connection,
  // and so on, so on. (can cause performance overhead)
  // while prepared statement that bound to Tx will only used that certain connection it tied with, and will not
  // reconnect to other connection if the connection aren't available.

  // handling error in database related code also aren't straightforward, there's some way to do it based on
  // the code logic and condition. basically you need to understand how the code work or interact with database
  // under the hood i guess.

  // oh you can use COALESCE() so it'll return zero-value to your variable instead of nil
  // so you don't need to make some of your variable as a pointer like i did in person struct (?)
  // wait so actually there's no best way to handle nullable field? at some case it can be the pointer (like
  // my person struct), some case using coalesce(), some case using sql.nullstring(which is a bit tricky
  // since there's no sql.nullUint64 and some other data type, but you can define your own types to handle null,
  // just copy the sql.nullstring design)

  // in general, coalesce/sql.nullstring(?) when null is noise, use pointer when null is information
  // but it's better to design a clean architecture (database, repository layer, domain layer)
  // like allow null only when meaningful, use default when possible (database)
  // use sql.Nullx or raw values, use COALESCE only for view-like queries (repository)
  // avoid sql.Nullx, use value types for required fields, pointer for optional ones (domain layer)
  // so you know how to handle the null value right way

  // connection pooling is normal concept, but Golang make it more controllable therefore explicit

  // result almost always scanned row by row in low-level. but ORM usually have feature to return bulk result
  // but under the hood it still row by row, but ORM store the result in a slice for us and return bulk result

  // btw I'm still torn between using pgx directly or use it through database/sql
}

func (p person) String() string{
  gender, birthDate, email, birthCountry := "<empty>", "<empty>", "<empty>", "<empty>"

  if p.Gender!=nil{
    gender = *p.Gender
  }
  if p.BirthDate!=nil{
    birthDate = p.BirthDate.Format("02-01-2006")
    // birthDate = (*p.BirthDate).Format("02-01-2006")
  }
  if p.Email!=nil{
    email = *p.Email
  }
  if p.BirthCountry!=nil{
    birthCountry = *p.BirthCountry
  }
  return fmt.Sprintf("Person\nID: %d\nFirst Name: %s\nLast Name: %s\nGender: %s\nBirthdate: %s\nEmail: %s\nCountry: %s\n",
                      p.Id, p.FirstName, p.LastName, gender, birthDate, email, birthCountry)
}

type person struct{
  Id int
  FirstName string
  LastName string
  Gender *string
  BirthDate *time.Time //since it's nullable, use pointer
  Email *string
  BirthCountry *string
}

func currentDB(conn *pgx.Conn) (string, error){
  query := `SELECT current_database()`
  var current_db string
  err := conn.QueryRow(context.Background(), query).Scan(&current_db)
  if err!=nil{
    return "", err
  }
  return current_db, nil
}

func listOfAttributes(conn *pgx.Conn, tableName string) ([]string, error){
	query := `SELECT column_name FROM information_schema.columns
            WHERE table_name = $1 ORDER BY ordinal_position`

  rows, err := conn.Query(context.Background(), query, tableName)
  if err!=nil{
    return nil, err
  }
  // defer rows.Close()

  // for rows.Next(){
  //   var column_name string
  //   err := rows.Scan(&column_name)
  //   if err!=nil{
  //     return err
  //   }
  //   fmt.Printf("column_name: %s\n", column_name)
  // }
  // return nil

  attributes, err := pgx.CollectRows(rows, pgx.RowTo[string])
  if err!=nil{
    return nil, err
  }
  return attributes, err
}



func listOfTable(conn *pgx.Conn) ([]string, error){
  // query := `SELECT * FROM pg_catalog.pg_tables`
  query := `SELECT tablename FROM pg_catalog.pg_tables
            WHERE schemaname!='pg_catalog' AND schemaname!='information_schema'`

  rows, err := conn.Query(context.Background(), query)
  if err!=nil{
    return nil, err
  }
  defer rows.Close()

  // for rows.Next(){
  //   var table_name string
  //   err := rows.Scan(&table_name)
  //   if err!=nil{
  //     return err
  //   }
  //   fmt.Println("table_name: ", table_name)
  // }
  // return nil
  tables, err := pgx.CollectRows(rows, pgx.RowTo[string])
  if err!=nil{
    return nil, err
  }
  return tables, err
}

func listOfDB(conn *pgx.Conn) ([]string, error){
  query := `SELECT datname FROM pg_database`
  rows, err := conn.Query(context.Background(), query)
  if err!=nil{
    return nil, err
  }
  defer rows.Close()

  // for rows.Next(){
  //   var datname string
  //   err := rows.Scan(&datname)
  //   if err!=nil{
  //     return err
  //   }
  //   fmt.Printf("database_name: %s\n", datname)
  // }
  // return nil
  dbs, err := pgx.CollectRows(rows, pgx.RowTo[string])
  if err!=nil{
    return nil, err
  }
  return dbs, err
}

func connect() (*pgx.Conn, error){
	conn, err := pgx.Connect(context.Background(),
	"postgres://postgres:Strikeboy007@localhost:5432/backend_intro")
	// open one single Postgres connection, no pooling, one TCP connection, one backend session
	// for experimenting use pgx.Connect()
	// pool, err := pgxpool.New()
	// manages multiple connections, production-grade default
	if err!=nil{
		return nil, err
	}
	return conn, nil
}