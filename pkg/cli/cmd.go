//==================================================================
//创建时间：2017-4-27 首次创建
//功能描述：Flag实现配置
//创建人：郭世江
//修改记录：若要修改请记录
//==================================================================
package cli

import (
	"t180web/pkg/db"
	"t180web/pkg/handler"
	"t180web/pkg/model"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

func HeartDataCmd(name string) *cobra.Command {
	t180web := &cobra.Command{
		Use:   name,
		Short: "t180web services",
	}
	t180web.AddCommand(startCmd())
	return t180web
}

func startCmd() *cobra.Command {
	var addr string
	var user, passwd, ip, port, dbname string
	start := &cobra.Command{
		Use:   "start",
		Short: "Start t180web system service",
		PreRun: func(cmd *cobra.Command, args []string) {
			if err := db.Init(user, passwd, ip, port, dbname); err != nil {
				glog.Errorf("reporter init db err %v\n", err)
				os.Exit(1)
			}
			if err := model.Init(db.GetDB()); err != nil {
				glog.Errorf("reporter init model err %v\n", err)
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			router := handler.CreateHttpRouter()
			if err := http.ListenAndServe(addr, router); err != nil {
				glog.Errorf("reporter start err %v\n", err)
				os.Exit(1)
			}
		},
	}
	flags := start.Flags()
	flags.StringVar(&addr, "listen", ":8899", " http request listen port")
	flags.StringVar(&user, "db_user", "root", "App Database User")
	flags.StringVar(&passwd, "db_passwd", "guoshijiang858103", "App Database Passwd")
	flags.StringVar(&ip, "db_ip", "39.108.103.32", "App Database IP")
	flags.StringVar(&port, "db_port", "3306", "App Database Port")
	flags.StringVar(&dbname, "db_name", "t180", "App Database Name")
	return start
}
