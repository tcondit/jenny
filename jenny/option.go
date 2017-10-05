package jenny

import (
	"github.com/c-bata/go-prompt"
	"math/rand"
	"strings"
	"time"
)

func optionCompleter(args []string, long bool) []prompt.Suggest {
	l := len(args)
	if l <= 1 {
		if long {
			return prompt.FilterHasPrefix(optionHelp, "--", false)
		}
		return optionHelp
	}

	var suggests []prompt.Suggest
	commandArgs := excludeOptions(args)
	if commandArgs[1] == "save" {
		suggests = flagProfile
	}

	if long {
		return prompt.FilterContains(
			prompt.FilterHasPrefix(suggests, "--", false),
			strings.TrimLeft(args[l-1], "--"),
			true,
		)
	}
	return prompt.FilterContains(suggests, strings.TrimLeft(args[l-1], "-"), true)
}

func Banner() string {
	rand.Seed(time.Now().UnixNano())
	return banner[rand.Int()%len(banner)]
}

var banner = []string{`
    ___       ___       ___       ___       ___   
   /\  \     /\  \     /\__\     /\__\     /\__\  
  _\:\  \   /::\  \   /:| _|_   /:| _|_   |::L__L 
 /\/::\__\ /::\:\__\ /::|/\__\ /::|/\__\  |:::\__\
 \::/\/__/ \:\:\/  / \/|::/  / \/|::/  /  /:;;/__/
  \/__/     \:\/  /    |:/  /    |:/  /   \/__/   
             \/__/     \/__/     \/__/            
`,
	`
   __     ______     __   __     __   __     __  __    
  /\ \   /\  ___\   /\ "-.\ \   /\ "-.\ \   /\ \_\ \   
 _\_\ \  \ \  __\   \ \ \-.  \  \ \ \-.  \  \ \____ \  
/\_____\  \ \_____\  \ \_\\"\_\  \ \_\\"\_\  \/\_____\ 
\/_____/   \/_____/   \/_/ \/_/   \/_/ \/_/   \/_____/ 
`,
	`
     ██╗███████╗███╗   ██╗███╗   ██╗██╗   ██╗
     ██║██╔════╝████╗  ██║████╗  ██║╚██╗ ██╔╝
     ██║█████╗  ██╔██╗ ██║██╔██╗ ██║ ╚████╔╝ 
██   ██║██╔══╝  ██║╚██╗██║██║╚██╗██║  ╚██╔╝  
╚█████╔╝███████╗██║ ╚████║██║ ╚████║   ██║   
 ╚════╝ ╚══════╝╚═╝  ╚═══╝╚═╝  ╚═══╝   ╚═╝   
`,
	`
::::::'##:'########:'##::: ##:'##::: ##:'##:::'##:
:::::: ##: ##.....:: ###:: ##: ###:: ##:. ##:'##::
:::::: ##: ##::::::: ####: ##: ####: ##::. ####:::
:::::: ##: ######::: ## ## ##: ## ## ##:::. ##::::
'##::: ##: ##...:::: ##. ####: ##. ####:::: ##::::
 ##::: ##: ##::::::: ##:. ###: ##:. ###:::: ##::::
. ######:: ########: ##::. ##: ##::. ##:::: ##::::
:......:::........::..::::..::..::::..:::::..:::::
`,
	`
     @@@  @@@@@@@@  @@@  @@@  @@@  @@@  @@@ @@@  
     @@@  @@@@@@@@  @@@@ @@@  @@@@ @@@  @@@ @@@  
     @@!  @@!       @@!@!@@@  @@!@!@@@  @@! !@@  
     !@!  !@!       !@!!@!@!  !@!!@!@!  !@! @!!  
     !!@  @!!!:!    @!@ !!@!  @!@ !!@!   !@!@!   
     !!!  !!!!!:    !@!  !!!  !@!  !!!    @!!!   
     !!:  !!:       !!:  !!!  !!:  !!!    !!:    
!!:  :!:  :!:       :!:  !:!  :!:  !:!    :!:    
::: : ::   :: ::::   ::   ::   ::   ::     ::    
 : :::    : :: ::   ::    :   ::    :      :     
`}

var profile = []prompt.Suggest{
	{Text: "uri", Description: "Location of the Jenkins server."},
	{Text: "user", Description: "Username credential."},
	{Text: "pwd", Description: "Password credential."},
	{Text: "name", Description: "If you have multiple Jenkins profiles, default: Default."},
	{Text: "use", Description: "Use different Jenkins credentials."},

	{Text: "cancel", Description: "Close and discard configuration."},
	{Text: "save", Description: "Save and close configuration."},
}

var actions = []prompt.Suggest{
	{Text: "get", Description: "Display one or many resources"},
	{Text: "describe", Description: "Show details of a specific resource or group of resources"},
	{Text: "create", Description: "Create a resource by filename or stdin"},
	{Text: "delete", Description: "Delete resources by filenames, stdin, resources and names, or by resources and label selector."},
	{Text: "edit", Description: "Edit a resource on the server"},
	{Text: "apply", Description: "Apply a configuration to a resource by filename or stdin"},
	{Text: "logs", Description: "Print the logs for a job or an event."},
	{Text: "run", Description: "Run a particular job."},
	{Text: "version", Description: "Print the Jenkins version information."},
	{Text: "explain", Description: "Documentation of resources."},

	{Text: "profile", Description: "Add or use a profile."},

	{Text: "exit", Description: "Exit this program."},
	{Text: "quit", Description: "Exit this program."},
}

var optionHelp = []prompt.Suggest{
	{Text: "help"},
}

var flagProfile = []prompt.Suggest{
	{Text: "--force-save", Description: "Save current configuration in .jenny.yml file."},
	// aliases
	{Text: "-f", Description: "Save current configuration in .jenny.yml file."},
}