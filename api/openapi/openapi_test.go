package openapi

import (
	"context"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestSpec(t *testing.T) {
	doc := load(t)

	if doc.OpenAPI != "3.0.3" {
		t.Fatalf("openapi = %q", doc.OpenAPI)
	}
	if doc.Info == nil || doc.Info.Title != "Runphase API" || doc.Info.Version != "0.0.1" {
		t.Fatalf("info = %#v", doc.Info)
	}

	paths := doc.Paths.Map()
	if len(paths) != 2 {
		t.Fatalf("paths = %d, want only /healthz and /readyz", len(paths))
	}
	assertGETJSON(t, doc, "/healthz")
	assertGETJSON(t, doc, "/readyz")
}

func load(t *testing.T) *openapi3.T {
	t.Helper()
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile("openapi.yaml")
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if err := doc.Validate(context.Background()); err != nil {
		t.Fatalf("validate: %v", err)
	}
	return doc
}

func assertGETJSON(t *testing.T, doc *openapi3.T, path string) {
	t.Helper()
	item := doc.Paths.Find(path)
	if item == nil || item.Get == nil {
		t.Fatalf("%s GET is missing", path)
	}
	response := item.Get.Responses.Status(200)
	if response == nil || response.Value == nil {
		t.Fatalf("%s 200 response is missing", path)
	}
	media := response.Value.Content.Get("application/json")
	if media == nil || media.Schema == nil || media.Schema.Ref != "#/components/schemas/StatusResponse" {
		t.Fatalf("%s 200 schema = %#v", path, media)
	}
}
