package config

import "gorm.io/gorm"

var (
	Key = "ab9312a52781f4b7c7edf4341ef940daff94c567ffa503c3db8125fec68c4225"
)

type BodyStructure map[string]string

const (
	ERRORS_NULL = iota
)

const SERVER_HOST = "localhost"
const SERVER_PORT = "8088"

const UpdatePath = "./updates/"

var DB *gorm.DB
