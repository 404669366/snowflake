package snowflake

import "github.com/bwmarrin/snowflake"

var snow *snowflake.Node

func Init(node int64) {
	var err error
	snow, err = snowflake.NewNode(node)
	if err != nil {
		panic("init snowflake error : " + err.Error())
	}
}

func CreateId() int64 {
	return snow.Generate().Int64()
}
