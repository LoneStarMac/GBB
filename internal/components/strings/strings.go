package strings

const (
	AppTitle           = "GBB"
	AppVersion         = "v0.1"
	ScanStart          = "Scanning for CMake dependencies...\n"
	CMakeNotFound      = "CMakeLists.txt not found or unreadable."
	NoDepsFound        = "No dependencies found in CMakeLists.txt."
	DepsFoundIntro     = "Found dependencies:"
	BrewPrompt         = "Would you like to search Homebrew for related formulas? [y/N] "
	ProceedPrompt      = "Proceed with CMake build anyway? [Y/n] "
	ExitMessage        = "Exiting."
	ErrorSearchingBrew = "Error searching Homebrew:"
	NoMatchesFound     = "No matches found."
	FoundFormulasIntro = "Found Homebrew formulas:"
	HeaderTemplate     = "%s %s — [%s] — %s\nBuild: %s | Method: %s | Dir: %s\n"
	MenuCheckOnline    = "Check online for %s build plugins"
	MenuChangeInstall  = "Change install location"
	MenuChangeFolder   = "Change/specify build folder"
	MenuChangeBuilder  = "Change builder (currently: %s)"
	MenuOptions        = "Options"
	MenuBuildNow       = "🚀 Build %s now"
	MainMenuPrompt     = "Choose an action:"
)
