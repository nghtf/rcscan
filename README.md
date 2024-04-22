# rcscan
Minimalistic parser for rc-files

== Overview

Package provides minimalistic access to parameters stored in old-fashined .rc-like configuration files (similar to .toml, but not)

Example for getmailrc cnfiguration file:

	[retriever]
	type = SimplePOP3SSLRetriever
	server = pop.domain.example
	username = user@domain.example
	password = P@$$w0rd

	[destination]
	type = Maildir
	path = ./getmail/maildir/

	[options]
	delete = false
	message_log = ./getmail/log

With rscan you can easily access any parameter inside the configuration.
Let's pull "path" parameter from "destination" section:

	rc, err := rcscan.New("./path-to/getmailrc")
	if err != nil {
		fmt.Println(err)
	} else {
		path, err := rc.GetParam("destination", "path")
		if err != nil {
			log.Debug("problem")
		}
		fmt.Println(path)
	}