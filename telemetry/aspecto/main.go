package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
	"log"
	"net/http"
)

var client *mongo.Client
var Tracer opentracing.Tracer

func JaegerTraceProvider() (*sdktrace.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("todo-service"),
			semconv.DeploymentEnvironmentKey.String("production"),
		)),
	)
	return tp, nil
}
func main() { //Export traces to Jaeger
	tp, tpErr := JaegerTraceProvider()
	if tpErr != nil {
		log.Fatal(tpErr)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	connectMongo()
	startWebServer()
}
func connectMongo() {
	opts := options.Client()
	//Mongo OpenTelemetry instrumentation
	opts.Monitor = otelmongo.NewMonitor()
	opts.ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(context.Background(), opts)
	//Seed the database with some todo's
	docs := []interface{}{
		bson.D{{"id", "1"}, {"title", "Buy groceries"}},
		bson.D{{"id", "2"}, {"title", "install Aspecto.io"}},
		bson.D{{"id", "3"}, {"title", "Buy dogz.io domain"}},
	}
	client.Database("todo").Collection("todos").InsertMany(context.Background(), docs)
}
func startWebServer() {
	r := gin.Default()
	//gin OpenTelemetry instrumentation
	r.Use(otelgin.Middleware("todo-service"))
	r.GET("/todo", func(c *gin.Context) {
		collection := client.Database("todo").Collection("todos")
		//Make sure to pass c.Request.Context() as the context and not c itself
		cur, findErr := collection.Find(c.Request.Context(), bson.D{})
		if findErr != nil {
			c.AbortWithError(500, findErr)
			return
		}
		results := make([]interface{}, 0)
		curErr := cur.All(c, &results)
		if curErr != nil {
			c.AbortWithError(500, curErr)
			return
		}
		c.JSON(http.StatusOK, results)
	})
	_ = r.Run(":8080")
}
