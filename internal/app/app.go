package app

func IsDev() bool {
	return Config.GetString("APP_ENV") == "development"
}

func IsProd() bool {
	return Config.GetString("APP_ENV") == "production"
}
