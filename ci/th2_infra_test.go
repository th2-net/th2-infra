package test

import (
	"os"
	"strings"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/logger"
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setUp() {}

func tearDown() {}

func validFunc(t *testing.T, code int, substr string) func(int, string) bool {
	return func(code int, body string) bool {
		if code != code {
			logger.Logf(t, "Incorrect response code")
			return false
		}
		if !strings.Contains(body, substr) {
			logger.Logf(t, "Body does not contain "+substr)
			return false
		}
		return true
	}
}

func TestDashboardEndpoint(t *testing.T) {
	// t.Parallel()
	options := k8s.NewKubectlOptions("", "", "service")
	k8s.WaitUntilServiceAvailable(t, options, "ingress-ingress-nginx-controller", 10, 3*time.Second)
	validator := validFunc(t, 200, "<title>Kubernetes Dashboard</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, "http://localhost:30000/dashboard/", nil, 0, time.Second, validator)
}

func TestRabbitMQEndpoint(t *testing.T) {
	// t.Parallel()
	options := k8s.NewKubectlOptions("", "", "service")
	k8s.WaitUntilServiceAvailable(t, options, "rabbitmq-discovery", 10, 3*time.Second)

	validator := validFunc(t, 200, "<title>RabbitMQ Management</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, "http://localhost:30000/rabbitmq/", nil, 0, time.Second, validator)
}

func TestInfraEditorEndpoint(t *testing.T) {
	// t.Parallel()
	options := k8s.NewKubectlOptions("", "", "service")
	k8s.WaitUntilServiceAvailable(t, options, "infra-editor", 10, 1*time.Second)

	validator := validFunc(t, 200, "<title>Schema editor</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, "http://localhost:30000/editor/", nil, 0, time.Second, validator)
}

func TestInfraMgrEndpoint(t *testing.T) {
	// t.Parallel()
	options := k8s.NewKubectlOptions("", "", "service")
	k8s.WaitUntilServiceAvailable(t, options, "infra-mgr", 10, 1*time.Second)

	validator := validFunc(t, 404, "{\"status_code\":404,\"error_code\":\"NOT_FOUND\",\"message\":\"Not Found\"}")
	http_helper.HttpGetWithRetryWithCustomValidation(t, "http://localhost:30000/editor/backend/", nil, 0, time.Second, validator)
}

func TestNamespaceReportEndpoint(t *testing.T) {
	// t.Parallel()
	options := k8s.NewKubectlOptions("", "", "schema-e2e")
	k8s.WaitUntilServiceAvailable(t, options, "rpt-data-viewer", 30, 5*time.Second)

	validator := validFunc(t, 200, "<title>TH2 Report</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, "http://localhost:30000/schema-e2e/", nil, 10, 5*time.Second, validator)
}

func TestNamespaceDataProviderEndpoint(t *testing.T) {
	// t.Parallel()
	options := k8s.NewKubectlOptions("", "", "schema-e2e")
	k8s.WaitUntilServiceAvailable(t, options, "rpt-data-provider", 30, 5*time.Second)

	validator := validFunc(t, 200, "[]")
	http_helper.HttpGetWithRetryWithCustomValidation(t, "http://localhost:30000/schema-e2e/backend/messageStreams", nil, 10, 5*time.Second, validator)
}
