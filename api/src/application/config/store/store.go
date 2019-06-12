package store

import "app/application/config"

// StorePath Define a pasta onde será guardados os dados históricos
var StorePath = config.GetEnv("STORE_PATH", "/tmp/")
