package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"

	sc "entgo.io/ent/dialect/sqlcommenter"
	_ "github.com/mattn/go-sqlite3"

	otelCommenter "entgo.io/ent/dialect/sqlcommenter/commenters/otel"
	"github.com/go-chi/chi/v5"

	"entgo.io/ent/examples/sqlcommenter/ent"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type routeKey struct{}

func initTracer() *sdktrace.TracerProvider {
	// For the demonstration, use sdktrace.AlwaysSample sampler to sample all traces.
	// In a production application, use sdktrace.ProbabilitySampler with a desired probability.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String("ExampleService"))),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp
}

func main() {
	tp := initTracer()
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	// create and configure ent client
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer client.Close()
	// add debug logs
	client = client.Debug()
	client = client.SqlComments(sc.WithCommenter(
		otelCommenter.HttpCommenter{},
		sc.NewContextMapper(sc.RouteCommentKey, routeKey{}),
	), sc.WithComments(sc.SqlComments{
		sc.ApplicationCommentKey: "bootcamp",
		sc.FrameworkCommentKey:   "go-chi",
	}))
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	client.User.Create().SetName("hedwigz").SaveX(context.Background())

	middleware := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			c := context.WithValue(r.Context(), routeKey{}, r.URL.Path)
			next.ServeHTTP(w, r.WithContext(c))
		}
		return http.HandlerFunc(fn)
	}
	r := chi.NewRouter()
	r.Use(middleware)
	r.Get("/users", func(rw http.ResponseWriter, r *http.Request) {
		users := client.User.Query().AllX(r.Context())
		b, _ := json.Marshal(users)
		rw.WriteHeader(http.StatusOK)
		rw.Write(b)
	})

	backend := otelhttp.NewHandler(r, "app")
	testRequest(backend)
}

func testRequest(handler http.Handler) {
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()

	// debug printer should print sql statement with comments
	handler.ServeHTTP(w, req)
}
