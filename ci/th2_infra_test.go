package test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/logger"
)

const (
	schemaNamespace  = "schema-e2e-v101"
	serviceNamespace = "service"
	rabbitmqSvc      = "rabbitmq-discovery"
	dataProviderSvc  = "rpt-data-provider"
	reportViewerSvc  = "rpt-data-viewer"
	infraMgrSvc      = "infra-mgr"
	infraEditorSvc   = "infra-editor"
	dashboardSvc     = "dashboard-kubernetes-dashboard"
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
	endpoint := "http://localhost:30000/dashboard/"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilServiceAvailable(t, options, dashboardSvc, 10, 3*time.Second)
	validator := validFunc(t, 200, "<title>Kubernetes Dashboard</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, 0, time.Second, validator)
}

func TestDashboardRedirectEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/dashboard"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilServiceAvailable(t, options, dashboardSvc, 10, 3*time.Second)
	validator := validFunc(t, 301, "<title>Kubernetes Dashboard</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, 0, time.Second, validator)
}

func TestRabbitMQEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/rabbitmq/"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilServiceAvailable(t, options, rabbitmqSvc, 10, 3*time.Second)

	validator := validFunc(t, 200, "<title>RabbitMQ Management</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, 0, time.Second, validator)
}

func TestInfraEditorEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/editor/"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilServiceAvailable(t, options, infraEditorSvc, 10, 1*time.Second)

	validator := validFunc(t, 200, "<title>Infra editor</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, 0, time.Second, validator)
}

func TestInfraMgrEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/editor/backend/"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilServiceAvailable(t, options, infraMgrSvc, 10, 1*time.Second)

	validator := validFunc(t, 404, "{\"status_code\":404,\"error_code\":\"NOT_FOUND\",\"message\":\"Not Found\"}")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, 0, time.Second, validator)
}

func TestNamespaceReportEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := fmt.Sprintf("http://localhost:30000/%s/", schemaNamespace)
	options := k8s.NewKubectlOptions("", "", schemaNamespace)
	k8s.WaitUntilServiceAvailable(t, options, reportViewerSvc, 30, 10*time.Second)

	validator := validFunc(t, 200, "<title>TH2 Report</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, 10, 3*time.Second, validator)
}

func TestNamespaceDataProviderEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := fmt.Sprintf("http://localhost:30000/%s/backend/search/events?timestampFrom=0&timestampTo=0", schemaNamespace)
	options := k8s.NewKubectlOptions("", "", schemaNamespace)
	k8s.WaitUntilServiceAvailable(t, options, dataProviderSvc, 30, 10*time.Second)

	validator := validFunc(t, 200, "[]")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, 10, 3*time.Second, validator)
}
