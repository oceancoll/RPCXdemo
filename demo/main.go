package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)


//docker run -d -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p5775:5775/udp -p6831:6831/udp -p6832:6832/udp   -p5778:5778 -p16686:16686 -p14268:14268 -p9411:9411 jaegertracing/all-in-one:latest
func NewJaegerTracer(service string) (opentracing.Tracer, io.Closer) {
	cfg := jaegerCfg.Configuration{
		Sampler: &jaegerCfg.SamplerConfig{
			Type:jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter:&jaegerCfg.ReporterConfig{
			LogSpans:true,
			CollectorEndpoint:"http://106.14.160.84:14268/api/traces",
			//LocalAgentHostPort:"106.14.160.84:6831",
		},
	}
	cfg.ServiceName = service
	tracer, closer, err := cfg.NewTracer(
		jaegerCfg.Logger(jaeger.StdLogger),
	)

	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}

func foo3(req string, ctx context.Context) (reply string){
	//1.创建子span
	span, _ := opentracing.StartSpanFromContext(ctx, "span_foo3")
	defer func() {
		//4.接口调用完，在tag中设置request和reply
		span.SetTag("request", req)
		span.SetTag("reply", reply)
		span.Finish()
	}()

	println(req)
	//2.模拟处理耗时
	time.Sleep(time.Second/2)
	//3.返回reply
	reply = "foo3Reply"
	return
}
//跟foo3一样逻辑
func foo4(req string, ctx context.Context) (reply string){
	span, _ := opentracing.StartSpanFromContext(ctx, "span_foo4")
	defer func() {
		span.SetTag("request", req)
		span.SetTag("reply", reply)
		span.Finish()
	}()

	println(req)
	time.Sleep(time.Second/2)
	reply = "foo4Reply"
	return
}

func main() {
	tracer, closer := NewJaegerTracer("jaeger-demo")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)//StartspanFromContext创建新span时会用到

	span := tracer.StartSpan("span_root")
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	r1 := foo3("Hello foo3", ctx)
	r2 := foo4("Hello foo4", ctx)
	fmt.Println(r1, r2)
	span.Finish()
}
