package ya

import (
	//	"io/ioutil"
	logger "log"
	"os"
)

var Log *logger.Logger

func init() {
	Log = logger.New(os.Stdout, ":: ", logger.LstdFlags|logger.Lshortfile)
	//	Log.SetOutput(ioutil.Discard) // выключаем

}
