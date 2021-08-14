package config

// Configurations exported
type Configurations struct {
	Server       ServerConfigurations
	Database     DatabaseConfigurations
	// EXAMPLE_PATH string
	// EXAMPLE_VAR  string
}

//NOTE:
// The configuration file contains, several variables,
// namely the server port, database name, credentials 
// to access the database and also "EXAMPLE_PATH" and
// "EXAMPLE_VAR" variables, where "EXAMPLE_VAR" is
// also available in the environment variables.


// ServerConfigurations exported
type ServerConfigurations struct {
	Port int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
}