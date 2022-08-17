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
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	defaultSchemaNamespace  = "th2-schema"
	defaultVhost            = "th2"
	defaultServiceNamespace = "service"
	monitoringNamespace     = "monitoring"
	rabbitmqPod             = "rabbitmq-0"
	rabbitmqSvc             = "rabbitmq-discovery"
	rabbitmqUser            = "th2"
	rabbitmqPassword        = "test"
	cassandraPod            = "cassandra-0"
	dataProviderSvc         = "rpt-data-provider"
	reportViewerSvc         = "rpt-data-viewer"
	infraMgrSvc             = "infra-mgr"
	infraEditorSvc          = "infra-editor"
	dashboardSvc            = "th2-infra-dashboard"
	retries                 = 10
	timeout                 = 5 * time.Second
)

var (
	serviceNamespace, schemaNamespace string
	exists                            bool
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setUp() {
	if schemaNamespace, exists = os.LookupEnv("SCHEMA_NAMESPACE"); !exists {
		schemaNamespace = defaultSchemaNamespace
	}
	if serviceNamespace, exists = os.LookupEnv("INFRA_NAMESPACE"); !exists {
		serviceNamespace = defaultServiceNamespace
	}
}

func tearDown() {}

func validFunc(t *testing.T, testCode int, substr string) func(int, string) bool {
	return func(code int, body string) bool {
		if testCode != code {
			logger.Log(t, "Incorrect response code")
			return false
		}
		if !strings.Contains(body, substr) {
			logger.Log(t, "Body does not contain "+substr)
			return false
		}
		return true
	}
}

func TestDashboardEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/dashboard/"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilServiceAvailable(t, options, dashboardSvc, retries, timeout)
	validator := validFunc(t, 200, "<title>Kubernetes Dashboard</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}

func TestDashboardRedirectEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/dashboard"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilServiceAvailable(t, options, dashboardSvc, retries, timeout)
	validator := validFunc(t, 200, "<title>Kubernetes Dashboard</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}

func TestInfraEditorEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/editor/"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilServiceAvailable(t, options, infraEditorSvc, retries, timeout)

	validator := validFunc(t, 200, "<title>Infra editor</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}

func TestInfraEditorRedirectEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/editor"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilServiceAvailable(t, options, infraEditorSvc, retries, timeout)

	validator := validFunc(t, 200, "<title>Infra editor</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}

func TestRabbitMQEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/rabbitmq/"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilPodAvailable(t, options, rabbitmqPod, retries, timeout)

	validator := validFunc(t, 200, "<title>RabbitMQ Management</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}

func TestRabbitMQRedirectEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/rabbitmq"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilPodAvailable(t, options, rabbitmqPod, retries, timeout)

	validator := validFunc(t, 200, "<title>RabbitMQ Management</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}

func TestInfraMgrEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/editor/backend/actuator/health"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilPodAvailable(t, options, cassandraPod, retries, timeout)

	validator := validFunc(t, 200, "{\"status\":\"UP\",\"groups\":[\"liveness\",\"readiness\"]}")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}

func TestNamespaceReportEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := fmt.Sprintf("http://localhost:30000/%s/", schemaNamespace)
	options := k8s.NewKubectlOptions("", "", schemaNamespace)
	k8s.WaitUntilServiceAvailable(t, options, reportViewerSvc, retries, timeout)

	validator := validFunc(t, 200, "<title>TH2 Report</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}

func TestNamespaceDataProviderEndpoint(t *testing.T) {
	// t.Parallel()
	endpoint := fmt.Sprintf("http://localhost:30000/%s/backend/filters/sse-events", schemaNamespace)
	options := k8s.NewKubectlOptions("", "", schemaNamespace)
	k8s.WaitUntilServiceAvailable(t, options, dataProviderSvc, retries, timeout)

	validator := validFunc(t, 200, "[\"attachedMessageId\",\"type\",\"name\",\"body\",\"status\"]")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}

func TestRabbitMQQueues(t *testing.T) {
	// t.Parallel()
	endpoint := fmt.Sprintf("http://%s:%s@localhost:30000/rabbitmq/api/queues/%s/link%%5B%s%%3Arpt-data-provider%%3Afrom_codec%%5D",
		rabbitmqUser, rabbitmqPassword, defaultVhost, schemaNamespace,
	)
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilPodAvailable(t, options, rabbitmqPod, retries, timeout)
	expectedString := fmt.Sprintf("\"name\":\"link[%s:rpt-data-provider:from_codec]\"", schemaNamespace)
	validator := validFunc(t, 200, expectedString)
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}

func TestPodAnnotations(t *testing.T) {
	options := k8s.NewKubectlOptions("", "", schemaNamespace)
	filters := metav1.ListOptions{
		LabelSelector: "app=rpt-data-viewer",
	}
	pods := k8s.ListPods(t, options, filters)
	assert.Equal(t, "test-annotation", pods[0].ObjectMeta.Annotations["e2e"])
	assert.Equal(t, "test-common-annotation", pods[0].ObjectMeta.Annotations["e2ecommon"])
}

func TestPodCommonAnnotationsOnly(t *testing.T) {
	options := k8s.NewKubectlOptions("", "", schemaNamespace)
	filters := metav1.ListOptions{
		LabelSelector: "app=rpt-data-provider",
	}
	pods := k8s.ListPods(t, options, filters)
	assert.Empty(t, pods[0].ObjectMeta.Annotations["e2e"])
	assert.Equal(t, "test-common-annotation", pods[0].ObjectMeta.Annotations["e2ecommon"])
}

func TestCodecFixPod(t *testing.T) {
	options := k8s.NewKubectlOptions("", "", schemaNamespace)
	filters := metav1.ListOptions{
		LabelSelector: "app=codec-fix",
	}
	pods := k8s.ListPods(t, options, filters)
	codecFixPodName := pods[0].ObjectMeta.Name
	k8s.WaitUntilPodAvailable(t, options, codecFixPodName, retries, timeout)
}

func TestTH2Main(t *testing.T) {
	// t.Parallel()
	endpoint := "http://localhost:30000/"
	options := k8s.NewKubectlOptions("", "", serviceNamespace)
	k8s.WaitUntilPodAvailable(t, options, rabbitmqPod, retries, timeout)
	validator := validFunc(t, 200, "<title>Welcome to th2</title>")
	http_helper.HttpGetWithRetryWithCustomValidation(t, endpoint, nil, retries, timeout, validator)
}
