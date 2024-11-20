package app

import (
	"common/config"
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

// Run 做一些项目主包的加载 初始化 grpc http redis mysql 等等
func Run(ctx context.Context) error {

	// 启动grpc服务
	server := grpc.NewServer()

	// 因为Listen之后会阻塞 这里启用一个协程去处理Listen
	go func() {
		//
		lis, err := net.Listen("tcp", config.Conf.Grpc.Addr)

		if err != nil {
			log.Fatal("grpc listen err:", err)
		}

		if err := server.Serve(lis); err != nil {
			log.Fatal("grpc Serve err:", err)
		}
	}()

	// 停止函数
	stop := func() {
		// 假设给予2秒处理释放资源的时间
		log.Println("app stop finish")
		time.Sleep(time.Second * 2)
		server.Stop()
	}

	// 控制启停 希望程序停止或者中断时 需要处理一下关闭其他服务或者释放程序的操作 需要一个优雅的启停
	// 使用一个缓存的channel来处理程序状态
	c := make(chan os.Signal, 1)

	// 监听程序信号 Notify用于接收信号 SIGTERM 终止 SIGQUIT 退出 SIGINT 中断   SIGHUP 挂断
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGHUP)

	log.Println("grpc run status ok")
	for {
		select {
		case <-ctx.Done():
			stop()
			return nil
		// 当接收到信号channel时
		case s := <-c:
			switch s {
			case syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT:
				stop()
				log.Println("app exit")
				return nil
			// 挂断信号  当你在liunx中登录的账户启动了这些服务 然后你退出登录的时候会收到这个信号 结束程序
			case syscall.SIGHUP:
				// TODO reload
				stop()
				log.Println("user hang up app exit")
				return nil
			default:
				return nil
			}

		}
	}
}
