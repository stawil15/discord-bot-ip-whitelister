package commands

type Command struct {
	Name        string
	Description string
	OptionName  string
	OptionDesc  string
}

var banCmd = Command{
	Name:        "ban",
	Description: "Ban a user",
	OptionName:  "user",
	OptionDesc:  "Discord user ID",
}

var whitelistCmd = Command{
	Name:        "whitelist",
	Description: "Whitelist your IP",
	OptionName:  "ip",
	OptionDesc:  "Your IP address",
}
