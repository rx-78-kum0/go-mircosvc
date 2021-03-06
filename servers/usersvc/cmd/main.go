package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/opentracing/opentracing-go"

	"github.com/openzipkin/zipkin-go"

	"github.com/kum0/go-mircosvc/shared/db"
	"github.com/kum0/go-mircosvc/shared/email"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/kum0/go-mircosvc/servers/usersvc/config"
	"github.com/kum0/go-mircosvc/servers/usersvc/endpoints"
	"github.com/kum0/go-mircosvc/servers/usersvc/middleware"
	"github.com/kum0/go-mircosvc/servers/usersvc/transport"
	sharedEtcd "github.com/kum0/go-mircosvc/shared/etcd"
	"github.com/kum0/go-mircosvc/shared/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/reflection"

	kitGrpc "github.com/go-kit/kit/transport/grpc"
	userPb "github.com/kum0/go-mircosvc/pb/user"
	sharedZipkin "github.com/kum0/go-mircosvc/shared/zipkin"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	zipkinGrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	conf := config.GetConfig()
	log, f := logger.NewLogger(conf.LogPath)
	defer f.Close()

	zipkinTracer, reporter := sharedZipkin.NewZipkin(log, conf.ZipkinAddr, "localhost:"+conf.GrpcPort,
		conf.ServiceName)
	defer reporter.Close()

	opentracing.SetGlobalTracer(zipkinot.Wrap(zipkinTracer))
	tracer := opentracing.GlobalTracer()

	{
		etcdClient := sharedEtcd.NewEtcd(conf.EtcdAddr)
		register := sharedEtcd.Register("/usersvc", "localhost:"+conf.GrpcPort, etcdClient, log)
		defer register.Register()
	}

	var svc endpoints.UserSerivcer
	{
		mdb := db.NewMysql(conf.MysqlUsername, conf.MysqlPassword, conf.MysqlAddr, conf.MysqlAuthsource)
		rd := db.NewRedis(conf.RedisAddr, conf.RedisPassword, conf.RedisMaxIdle, conf.RedisMaxActive)
		email := email.NewEmail(conf.EmailFrom, conf.EmailAuthCode, conf.EmailHost, conf.EmailSender, conf.EmailPort)
		svc = endpoints.NewUserService(mdb, rd, email)
		svc = middleware.MakeServiceMiddleware(svc)
	}
	eps := endpoints.NewEndpoints(svc, log, tracer, zipkinTracer)

	hs := health.NewServer()
	hs.SetServingStatus(conf.ServiceName, healthgrpc.HealthCheckResponse_SERVING)

	errs := make(chan error, 1)
	go grpcServer(transport.MakeGRPCServer(eps, tracer, zipkinTracer, log), conf.GrpcPort, zipkinTracer, hs, log, errs)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	level.Info(log).Log("serviceName", conf.ServiceName, "terminated", <-errs)
}

func grpcServer(grpcsvc userPb.UsersvcServer, port string, zipkinTracer *zipkin.Tracer, hs *health.Server,
	logger log.Logger,
	errs chan error) {
	p := fmt.Sprintf(":%s", port)
	listener, err := net.Listen("tcp", p)
	if err != nil {
		level.Error(logger).Log("protocol", "GRPC", "listen", port, "err", err)
		os.Exit(1)
	}
	level.Info(logger).Log("protocol", "GRPC", "protocol", "GRPC", "exposed", port)

	server := grpc.NewServer(grpc.UnaryInterceptor(kitGrpc.Interceptor),
		grpc.StatsHandler(zipkinGrpc.NewServerHandler(zipkinTracer)),
	)
	userPb.RegisterUsersvcServer(server, grpcsvc)
	healthgrpc.RegisterHealthServer(server, hs)
	reflection.Register(server)
	errs <- server.Serve(listener)
}

// func httpServer(logger log.Logger, port string, handler http.Handler, errs chan error) {
// 	http.Handle("/", accessControl(handler))
//
// 	p := fmt.Sprintf(":%s", port)
// 	svr := &http.Server{
// 		Addr:    p,
// 		Handler: handler,
// 	}
// 	err := svr.ListenAndServe()
// 	if err != nil {
// 		logger.Log("listen: %s\n", err)
// 	}
// 	errs <- err
// }
//
// func accessControl(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
//
// 		if r.Method == "OPTIONS" {
// 			return
// 		}
//
// 		h.ServeHTTP(w, r)
// 	})
// }
