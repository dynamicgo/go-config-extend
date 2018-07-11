package extend

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"

	config "github.com/dynamicgo/go-config"
	"github.com/dynamicgo/go-config/source/memory"
)

var configdata = `{
    "mq":{
		"leveldb":{
			"broker":"../../.test",
			"topic":"levemq",
			"consumer":"test"
		},
		"wallet":[
			{
				"broker":"../../.test",
				"topic":"levemq",
				"consumer":"test"
			},
			{
				"broker":"../../.test",
				"topic":"levemq",
				"consumer":"test"
			}
		]
	}
}`

var conf config.Config

func init() {
	conf = config.NewConfig()

	err := conf.Load(memory.NewSource(
		memory.WithData([]byte(configdata)),
	))

	if err != nil {
		panic(err)
	}
}

func TestSubconfig(t *testing.T) {
	subconfig, err := SubConfig(conf, "mq", "leveldb")

	require.NoError(t, err)

	require.Equal(t, subconfig.Get("broker").String(""), "../../.test")
}

func TestSubConfigSlice(t *testing.T) {

	subConfigSlice, err := SubConfigSlice(conf, "mq", "wallet")

	require.NoError(t, err)

	require.Equal(t, 2, len(subConfigSlice))

	require.Equal(t, subConfigSlice[0].Get("broker").String(""), "../../.test")
}

func TestHex(t *testing.T) {
	_, err := hex.DecodeString("")
	require.NoError(t, err)
}
