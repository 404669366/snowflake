package snowflake

var snow *Node

func Init(node int64) {
	var err error
	snow, err = NewNode(node)
	if err != nil {
		panic("init snowflake error : " + err.Error())
	}
}

func CreateId() int64 {
	return snow.Generate().Int64()
}
