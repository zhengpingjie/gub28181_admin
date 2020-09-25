package util
import "flag"

type Option struct {
	Debug    bool
	MysqlDsn string
	RedisDsn string
}
var Opt *Option= new(Option)
func init(){
	flag.BoolVar(&Opt.Debug, "debug", false, "open debug mode")
}