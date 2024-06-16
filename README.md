# rcscan
Minimalistic parser for .rc and .ini files

## Overview

Package provides minimalistic access to parameters stored in old-fashined .rc and .ini -like configuration files:

	[Section 1]
	paramA = 2023
	paramB = 2024

	[Section 2]
	paramC = some string value

RCScan allows you to easily extract any parameter __as a string__ value:

Code example:

	rc, _ := rcscan.New("./example/example.rc")
	param, _ := rc.Get("Section 1", "paramA")
	fmt.Println("paramA:", param)

Output:

	paramA: 2023

RCScan supports files with no sections and also comment blocks started from `;`, `#` or `/`:

	# parameter outside a section:
	paramA = my value

	[Section 1]
	paramA = 2023
	paramB = 2024

Code example:

	rc, _ := rcscan.New("./example/example.rc")
	
	param, _ := rc.Get("", "paramA")
	fmt.Println("paramA:", param)
	
	param, _ := rc.Get("Section 1", "paramA")
	fmt.Println("paramA:", param)

Output:

	paramA: my value
	paramA: 2023

Some real world example for popular getmail.rc configuration file:

	[retriever]
	type = SimplePOP3SSLRetriever
	server = pop.domain.example

	[destination]
	type = Maildir
	path = ./getmail/maildir/

	[options]
	delete = false
	message_log = ./getmail/log

With rcscan you can easily access parameters from it:

	rc, err := rcscan.New("./path-to-your/getmail.rc")
	if err == nil {
		path, _ := rc.Get("destination", "path")
		fmt.Println("path:", path)

		// sections can be specified with brackets as well:

		path, _ := rc.Get("[options]", "delete")
		fmt.Println("delete:", path)
	}

Output:

    path: ./getmail/maildir/
	delete: false

