package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
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

{app_name} EPRINT_ID_LIST

# DESCRIPTION

The program will generate a SQL statements to produce a list of
eprint records containing eprintid, title, doi, publication type
and pub date. It is suitable to process via the mysql client and
rendering a tab delimited file result.

The EPRINT_ID_LIST should have only eprintid one per line.

# OPTIONS

-help
: display this help page.

# EXAMPLE

Process the eprint id list and produce a report of eprintid containing
their title, doi, publication type and publication date.

~~~
{app_name} eprint_ids.txt \
    > eprintid_pub_type.sql

mysql caltechauthors --batch --skip-column-names \
    < eprintid_pub_type.sql >eprints_pub_type.tsv

grep '\tarticle\t' <eprint_pub_type_report.tsv| wc -l
~~~

`
)

func usage(appName string) string {
	return strings.ReplaceAll(helpText, "{app_name}", appName)
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
	if len(args) < 1 || len(args) > 1 {
		fmt.Fprintf(os.Stderr, "Expected a filename containing eprint id one per row\n")
		os.Exit(1)
	}
	fName := args[0]
	src, err := ioutil.ReadFile(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read %q, %s\n", fName, err)
		os.Exit(1)
	}
	txt := fmt.Sprintf("%s", src)
	for i, line := range strings.Split(txt, "\n") {
		eprintid := strings.TrimSpace(line)
		if eprintid == "" {
			fmt.Fprintf(os.Stderr, "WARNING: skipping line %d, not data.\n", i)
		} else {
			fmt.Printf("SELECT eprintid, title, IFNULL(doi, '') AS doi, IFNULL(type, '') AS publication_type, IF(date_type = 'pub', CONCAT(date_year, '-', date_month, '-', date_day), '') AS publication_date FROM eprint WHERE eprintid = %q;\n", eprintid)
		}
	}
}
