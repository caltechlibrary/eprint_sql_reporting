package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/caltechlibrary/eprint_sql_reporting"
)

var (
	helpText = `% {app_name}(1) user-manual
% R. S. Doiel
% 2022-10-26

# NAME

{app_name}

# SYNOPSIS

{app_name} CSV_FILENAME COL_NUMBER_OF_ID

# DESCRIPTION

The program will generate a SQL statements for matching a spreadsheet
with a "Name" column to the author id in CaltechAUTHOR's eprint_creators_id
table.  This can be then passed via mysql client to generate a tab
delimited file if eprintid that next be passeed through the Unix sort
command to produce a list of unique eprint id for a group of authors.

The CSV_FILENAME should be the name of the file containing the author
ids.  COL_NUMBER_OF_ID is the column number (1 is initial column)
holding the author_id.

# OPTIONS

-help
: display this help page.

# EXAMPLE

CSV file called "GPS_Faculty_ORCIDS - Sheet1.csv", the author id is
in the column called "Name". Here's the steps to produce the report.

~~~
{app_name} "GPS_Faculty_ORCIDS - Sheet1.csv" 1 \
    > creator_ids.sql

mysql caltechauthors --batch --skip-column-names < creator_ids.sql | \
   sort -a >eprint_ids.txt
wc -l eprint_ids.txt
~~~

This leaves you author a list of eprint ids related to all the authors
in the CSV file.  You need to filter these one step further if you
are interested in publication type.

~~~
eprintid_pub_types eprint_ids.txt >eprint_pub_types.sql
mysql caltechauthors --batch --skip-column-names <eprintid_pub_types.sql |\
   >eprintid_pub_types.tsv
~~

This last step will produce a CSV file with one eprint record per line
with the eprintid, article title, doi, publicatin type, publication date.
`
)

func usage(appName string) string {
	return strings.ReplaceAll(helpText, "{app_name}", appName)
}

func getAuthorIDs(fName string, columnNo int) ([]string, error) {
	if columnNo < 0 {
		columnNo = 0
	}
	src, err := ioutil.ReadFile(fName)
	if err != nil {
		return nil, err
	}
	txt := fmt.Sprintf("%s", src)
	r := csv.NewReader(strings.NewReader(txt))
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	authorIDs := []string{}
	for i, record := range records {
		if i == 0 {
			if columnNo < len(record) {
				fmt.Fprintf(os.Stdout, "# column %d in %s is labeled %q\n", columnNo+1, fName, record[columnNo])
			} else {
				fmt.Fprintf(os.Stderr, "WARNING: row %d missing column %d\n", columnNo+1)
			}
		} else {
			//fmt.Fprintf(os.Stderr, "DEBUG record %+v\n", record)
			if columnNo < len(record) {
				authorID := record[columnNo]
				authorIDs = append(authorIDs, authorID)
			} else {
				fmt.Fprintf(os.Stderr, "WARNING: Row %d is missing column %d Name value\n", i, columnNo+1)
			}
		}
	}
	return authorIDs, nil
}

func main() {
	appName := path.Base(os.Args[0])
	showHelp, showVersion, showLicense := false, false, false
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.Parse()
	args := flag.Args()

	if showHelp {
		fmt.Fprintf(os.Stdout, "%s\n", usage(appName))
		fmt.Fprintf(os.Stdout, "%s %s\n", appName, eprint_sql_reporting.Version)
		os.Exit(0)
	}
	if showVersion {
		fmt.Fprintf(os.Stdout, "%s %s\n", appName, eprint_sql_reporting.Version)
		os.Exit(0)
	}
	if showLicense {
		fmt.Fprintf(os.Stdout, "%s %s\n", appName, eprint_sql_reporting.Version)
		fmt.Fprintf(os.Stdout, "%s\n", eprint_sql_reporting.LicenseText)
		os.Exit(0)
	}
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "%s\n", usage(appName))
		fmt.Fprintf(os.Stdout, "%s %s\n", appName, eprint_sql_reporting.Version)
		os.Exit(1)
	}
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Missing CSV filename or column number holding creator id\n")
		os.Exit(1)
	}
	fName := args[0]
	columnNo, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to convert %q to an integer, %s\n", args[1], err)
		os.Exit(1)
	}
	authorIDs, err := getAuthorIDs(fName, columnNo-1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	for _, authorID := range authorIDs {
		fmt.Printf("SELECT eprintid FROM eprint_creators_id WHERE creators_id = %q;\n", authorID)
	}
}
