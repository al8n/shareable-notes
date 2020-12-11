package common

import "errors"

var (
	// Identity
	ErrorInvalidEmail = errors.New("email address is invalid")
	ErrorEmailIdentity = errors.New("email or password is incorrect")
	ErrorPhoneIdentity = errors.New("phone number or password is incorrect")

	// Config
	ErrorConfigurationFileType = errors.New("config file type is not supported")
	ErrorNoServicesConfig = errors.New("no services configurations are provided")
	ErrorInvalidPort = errors.New("invalid port number")

	// Mongo config
	ErrorNoMongoDBConfig = errors.New("no MongoDB configuration is provided")
	ErrorNoDatabaseName = errors.New("MongoDB database name cannot be null")
	ErrorNoCollectionName = errors.New("MongoDB collection name cannot be null")

)
