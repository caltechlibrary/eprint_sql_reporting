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

{app_name} EPRINT_ID_LIST GROUP_NAME

# DESCRIPTION

The program will generate a SQL statements to add a local_group
for a set of eprintid.

The EPRINT_ID_LIST should have only eprintid one per line. The
local group name applied is provided in the command line (i.e.
GROU_NAME above)

# OPTIONS

-help
: display this help page.

-license
: display software license

-version
: display version


# EXAMPLE

Process the eprint id list and the group name "Astronomy Department"
updated the eprint table appropriately.

~~~
{app_name} eprint_ids.txt "Astronomy Department" \
    > eprintid_add_group.sql

mysql caltechauthors < eprintid_add_group.sql
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
	if len(args) < 2 || len(args) > 2 {
		fmt.Fprintf(os.Stderr, "Expected a filename containing eprint id one per row as a local group name in quotes.\n")
		os.Exit(1)
	}
	fName := args[0]
	localGroup := args[1]
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
			// NOTE: We need to prevent duplicates so we preemtively
			// remove the row because adding a new one.
			fmt.Printf("DELETE FROM eprint_local_group WHERE eprintid = %q AND local_group = %q;\n", eprintid, localGroup)
			fmt.Printf("SELECT pos+1 FROM eprint_local_group WHERE eprintid = %q ORDER BY pos DESC LIMIT 1 INTO @local_group_pos;\n", eprintid)
			fmt.Printf("INSERT INTO eprint_local_group (eprintid, local_group, pos) VALUES (%q, %q, IFNULL(@local_group_pos, 0));\n\n", eprintid, localGroup)
		}
	}
}
