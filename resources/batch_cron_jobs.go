package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func BatchCronJobs() *schema.Table {
	return &schema.Table{
		Name:         "k8s_batch_cron_jobs",
		Description:  "CronJob represents the configuration of a single cron job.",
		Multiplex:    client.ContextMultiplex,
		DeleteFilter: client.DeleteContextFilter,
		Resolver:     fetchBatchCronJobs,
		Columns: []schema.Column{
			client.CommonContextField,
			{
				Name:        "kind",
				Description: "Kind is a string value representing the REST resource this object represents.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TypeMeta.Kind"),
			},
			{
				Name:        "api_version",
				Description: "Defines the versioned schema of this representation of an object.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TypeMeta.APIVersion"),
			},
			{
				Name:        "name",
				Description: "Unique name within a namespace.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.Name"),
			},
			{
				Name:        "namespace",
				Description: "Namespace defines the space within which each name must be unique.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.Namespace"),
			},
			{
				Name:        "uid",
				Description: "UID is the unique in time and space value for this object.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.UID"),
			},
			{
				Name:        "resource_version",
				Description: "An opaque value that represents the internal version of this object.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.ResourceVersion"),
			},
			{
				Name:        "generation",
				Description: "A sequence number representing a specific generation of the desired state.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ObjectMeta.Generation"),
			},
			{
				Name:        "deletion_grace_period_seconds",
				Description: "Number of seconds allowed for this object to gracefully terminate.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("ObjectMeta.DeletionGracePeriodSeconds"),
			},
			{
				Name:        "labels",
				Description: "Map of string keys and values that can be used to organize and categorize objects.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ObjectMeta.Labels"),
			},
			{
				Name:        "annotations",
				Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ObjectMeta.Annotations"),
			},
			{
				Name:        "owner_references",
				Description: "List of objects depended by this object.",
				Type:        schema.TypeJSON,
				Resolver:    resolveBatchCronJobOwnerReferences,
			},
			{
				Name:        "finalizers",
				Description: "List of finalizers",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ObjectMeta.Finalizers"),
			},
			{
				Name:        "cluster_name",
				Description: "The name of the cluster which the object belongs to.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ObjectMeta.ClusterName"),
			},
			{
				Name:        "schedule",
				Description: "The schedule in Cron format.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.Schedule"),
			},
			{
				Name:        "starting_deadline_seconds",
				Description: "Optional deadline in seconds for starting the job if it misses scheduled time for any reason",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Spec.StartingDeadlineSeconds"),
			},
			{
				Name:        "concurrency_policy",
				Description: "Specifies how to treat concurrent executions of a Job.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Spec.ConcurrencyPolicy"),
			},
			{
				Name:        "suspend",
				Description: "This flag tells the controller to suspend subsequent executions, it does not apply to already started executions",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Spec.Suspend"),
			},
			{
				Name:        "successful_jobs_history_limit",
				Description: "The number of successful finished jobs to retain",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.SuccessfulJobsHistoryLimit"),
			},
			{
				Name:        "failed_jobs_history_limit",
				Description: "The number of failed finished jobs to retain",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Spec.FailedJobsHistoryLimit"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "k8s_batch_cron_job_templates",
				Description: "JobTemplateSpec describes the data a Job should have when created from a template",
				Resolver:    fetchBatchCronJobTemplates,
				Columns: []schema.Column{
					{
						Name:        "cron_job_cq_id",
						Description: "Unique CloudQuery ID of k8s_batch_cron_jobs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "Unique name within a namespace.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ObjectMeta.Name"),
					},
					{
						Name:        "namespace",
						Description: "Namespace defines the space within which each name must be unique.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ObjectMeta.Namespace"),
					},
					{
						Name:        "uid",
						Description: "UID is the unique in time and space value for this object.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ObjectMeta.UID"),
					},
					{
						Name:        "resource_version",
						Description: "An opaque value that represents the internal version of this object.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ObjectMeta.ResourceVersion"),
					},
					{
						Name:        "generation",
						Description: "A sequence number representing a specific generation of the desired state.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ObjectMeta.Generation"),
					},
					{
						Name:        "deletion_grace_period_seconds",
						Description: "Number of seconds allowed for this object to gracefully terminate.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ObjectMeta.DeletionGracePeriodSeconds"),
					},
					{
						Name:        "labels",
						Description: "Map of string keys and values that can be used to organize and categorize objects.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("ObjectMeta.Labels"),
					},
					{
						Name:        "annotations",
						Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("ObjectMeta.Annotations"),
					},
					{
						Name:        "owner_references",
						Description: "List of objects depended by this object.",
						Type:        schema.TypeJSON,
						Resolver:    resolveBatchCronJobTemplateOwnerReferences,
					},
					{
						Name:        "finalizers",
						Description: "List of finalizers",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("ObjectMeta.Finalizers"),
					},
					{
						Name:        "cluster_name",
						Description: "The name of the cluster which the object belongs to.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ObjectMeta.ClusterName"),
					},
					{
						Name:        "parallelism",
						Description: "Specifies the maximum desired number of pods the job should run at any given time",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Spec.Parallelism"),
					},
					{
						Name:        "completions",
						Description: "Specifies the desired number of successfully finished pods the job should be run with",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Spec.Completions"),
					},
					{
						Name:        "active_deadline_seconds",
						Description: "Specifies the duration in seconds relative to the startTime that the job may be continuously active before the system tries to terminate it.",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("Spec.ActiveDeadlineSeconds"),
					},
					{
						Name:        "backoff_limit",
						Description: "Specifies the number of retries before marking this job failed.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Spec.BackoffLimit"),
					},
					{
						Name:        "selector",
						Description: "A label query over pods that should match the pod count.",
						Type:        schema.TypeJSON,
						Resolver:    resolveBatchCronJobTemplateSelector,
					},
					{
						Name:        "manual_selector",
						Description: "manualSelector controls generation of pod labels and pod selectors.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Spec.ManualSelector"),
					},
					{
						Name:        "ttl_seconds_after_finished",
						Description: "Limits the lifetime of a Job that has finished execution.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Spec.TTLSecondsAfterFinished"),
					},
					{
						Name:        "completion_mode",
						Description: "CompletionMode specifies how Pod completions are tracked",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Spec.CompletionMode"),
					},
					{
						Name:        "suspend",
						Description: "Suspend specifies whether the Job controller should create Pods or not",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Spec.Suspend"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "k8s_batch_cron_job_template_pod",
						Description: "Describes the pod that will be created when executing a job.",
						Resolver:    fetchBatchCronJobTemplatePod,
						Columns: []schema.Column{
							{
								Name:        "cron_job_template_cq_id",
								Description: "Unique CloudQuery ID of k8s_batch_cron_job_templates table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "Unique name within a namespace.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ObjectMeta.Name"),
							},
							{
								Name:        "namespace",
								Description: "Namespace defines the space within which each name must be unique.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ObjectMeta.Namespace"),
							},
							{
								Name:        "uid",
								Description: "UID is the unique in time and space value for this object.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ObjectMeta.UID"),
							},
							{
								Name:        "resource_version",
								Description: "An opaque value that represents the internal version of this object.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ObjectMeta.ResourceVersion"),
							},
							{
								Name:        "generation",
								Description: "A sequence number representing a specific generation of the desired state.",
								Type:        schema.TypeBigInt,
								Resolver:    schema.PathResolver("ObjectMeta.Generation"),
							},
							{
								Name:        "deletion_grace_period_seconds",
								Description: "Number of seconds allowed for this object to gracefully terminate.",
								Type:        schema.TypeBigInt,
								Resolver:    schema.PathResolver("ObjectMeta.DeletionGracePeriodSeconds"),
							},
							{
								Name:        "labels",
								Description: "Map of string keys and values that can be used to organize and categorize objects.",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("ObjectMeta.Labels"),
							},
							{
								Name:        "annotations",
								Description: "Annotations is an unstructured key value map stored with a resource that may be set by external tools.",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("ObjectMeta.Annotations"),
							},
							{
								Name:        "owner_references",
								Description: "List of objects depended by this object.",
								Type:        schema.TypeJSON,
								Resolver:    resolveBatchCronJobTemplatePodJSON(func(podSpec corev1.PodTemplateSpec) interface{} { return podSpec.ObjectMeta.OwnerReferences }),
							},
							{
								Name:        "finalizers",
								Description: "List of finalizers",
								Type:        schema.TypeStringArray,
								Resolver:    schema.PathResolver("ObjectMeta.Finalizers"),
							},
							{
								Name:        "cluster_name",
								Description: "The name of the cluster which the object belongs to.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ObjectMeta.ClusterName"),
							},
							{
								Name:        "restart_policy",
								Description: "Restart policy for all containers within the pod.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Spec.RestartPolicy"),
							},
							{
								Name:        "termination_grace_period_seconds",
								Description: "Optional duration in seconds the pod needs to terminate gracefully",
								Type:        schema.TypeBigInt,
								Resolver:    schema.PathResolver("Spec.TerminationGracePeriodSeconds"),
							},
							{
								Name:        "active_deadline_seconds",
								Description: "Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers.",
								Type:        schema.TypeBigInt,
								Resolver:    schema.PathResolver("Spec.ActiveDeadlineSeconds"),
							},
							{
								Name:        "dns_policy",
								Description: "Sets DNS policy for the pod.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Spec.DNSPolicy"),
							},
							{
								Name:        "node_selector",
								Description: "Selector which must be true for the pod to fit on a node.",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("Spec.NodeSelector"),
							},
							{
								Name:        "service_account_name",
								Description: "Name of the ServiceAccount to use to run this pod.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Spec.ServiceAccountName"),
							},
							{
								Name:        "automount_service_account_token",
								Description: "Indicates whether a service account token should be automatically mounted.",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("Spec.AutomountServiceAccountToken"),
							},
							{
								Name:        "node_name",
								Description: "Requests to schedule this pod onto a specific node.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Spec.NodeName"),
							},
							{
								Name:        "host_network",
								Description: "Host networking requested for this pod.",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("Spec.HostNetwork"),
							},
							{
								Name:        "host_pid",
								Description: "Use the host's pid namespace.",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("Spec.HostPID"),
							},
							{
								Name:        "host_ipc",
								Description: "Use the host's ipc namespace.",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("Spec.HostIPC"),
							},
							{
								Name:        "share_process_namespace",
								Description: "Share a single process namespace between all of the containers in a pod.",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("Spec.ShareProcessNamespace"),
							},
							{
								Name:        "security_context",
								Description: "Holds pod-level security attributes and common container settings.",
								Type:        schema.TypeJSON,
								Resolver:    resolveBatchCronJobTemplatePodJSON(func(podSpec corev1.PodTemplateSpec) interface{} { return podSpec.Spec.SecurityContext }),
							},
							{
								Name:        "image_pull_secrets",
								Description: "Optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.",
								Type:        schema.TypeJSON,
								Resolver:    resolveBatchCronJobTemplatePodJSON(func(podSpec corev1.PodTemplateSpec) interface{} { return podSpec.Spec.ImagePullSecrets }),
							},
							{
								Name:        "hostname",
								Description: "Specifies the hostname of the Pod.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Spec.Hostname"),
							},
							{
								Name:        "subdomain",
								Description: "Specifies the subdomain of the Pod.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Spec.Subdomain"),
							},
							{
								Name:        "affinity",
								Description: "If specified, the pod's scheduling constraints.",
								Type:        schema.TypeJSON,
								Resolver:    resolveBatchCronJobTemplatePodJSON(func(podSpec corev1.PodTemplateSpec) interface{} { return podSpec.Spec.Affinity }),
							},
							{
								Name:        "scheduler_name",
								Description: "If specified, the pod will be dispatched by specified scheduler.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Spec.SchedulerName"),
							},
							{
								Name:        "tolerations",
								Description: "If specified, the pod's tolerations.",
								Type:        schema.TypeJSON,
								Resolver:    resolveBatchCronJobTemplatePodJSON(func(podSpec corev1.PodTemplateSpec) interface{} { return podSpec.Spec.Tolerations }),
							},
							{
								Name:        "host_aliases",
								Description: "Optional list of hosts and IPs that will be injected into the pod's hosts file if specified.",
								Type:        schema.TypeJSON,
								Resolver:    resolveBatchCronJobTemplatePodJSON(func(podSpec corev1.PodTemplateSpec) interface{} { return podSpec.Spec.HostAliases }),
							},
							{
								Name:        "priority_class_name",
								Description: "If specified, indicates the pod's priority",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Spec.PriorityClassName"),
							},
							{
								Name:        "priority",
								Description: "The priority value",
								Type:        schema.TypeInt,
								Resolver:    schema.PathResolver("Spec.Priority"),
							},
							{
								Name:        "dns_config",
								Description: "Specifies the DNS parameters of a pod.",
								Type:        schema.TypeJSON,
								Resolver:    resolveBatchCronJobTemplatePodJSON(func(podSpec corev1.PodTemplateSpec) interface{} { return podSpec.Spec.DNSConfig }),
							},
							{
								Name:        "readiness_gates",
								Description: "If specified, all readiness gates will be evaluated for pod readiness.",
								Type:        schema.TypeJSON,
								Resolver:    resolveBatchCronJobTemplatePodJSON(func(podSpec corev1.PodTemplateSpec) interface{} { return podSpec.Spec.ReadinessGates }),
							},
							{
								Name:        "runtime_class_name",
								Description: "Refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Spec.RuntimeClassName"),
							},
							{
								Name:        "enable_service_links",
								Description: "Indicates whether information about services should be injected into pod's environment variables.",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("Spec.EnableServiceLinks"),
							},
							{
								Name:        "preemption_policy",
								Description: "Policy for preempting pods with lower priority.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Spec.PreemptionPolicy"),
							},
							{
								Name:        "overhead",
								Description: "Represents the resource overhead associated with running a pod for a given RuntimeClass.",
								Type:        schema.TypeJSON,
								Resolver:    schema.PathResolver("Spec.Overhead"),
							},
							{
								Name:        "topology_spread_constraints",
								Description: "Describes how a group of pods ought to spread across topology domains",
								Type:        schema.TypeJSON,
								Resolver:    resolveBatchCronJobTemplatePodJSON(func(podSpec corev1.PodTemplateSpec) interface{} { return podSpec.Spec.TopologySpreadConstraints }),
							},
							{
								Name:        "set_hostname_as_fqdn",
								Description: "If true the pod's hostname will be configured as the pod's FQDN, rather than the leaf name.",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("Spec.SetHostnameAsFQDN"),
							},
						},
						Relations: []*schema.Table{
							{
								Name:        "k8s_batch_cron_job_template_pod_init_containers",
								Description: "A single application container that you want to run within a pod.",
								Resolver:    fetchBatchCronJobTemplatePodInitContainers,
								Columns: []schema.Column{
									{
										Name:        "pod_cq_id",
										Description: "Unique CloudQuery ID of k8s_batch_cron_job_template_pod table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "name",
										Description: "Name of the container specified as a DNS_LABEL.",
										Type:        schema.TypeString,
									},
									{
										Name:        "image",
										Description: "Docker image name.",
										Type:        schema.TypeString,
									},
									{
										Name:        "command",
										Description: "Entrypoint array",
										Type:        schema.TypeStringArray,
									},
									{
										Name:        "args",
										Description: "Arguments to the entrypoint.",
										Type:        schema.TypeStringArray,
									},
									{
										Name:        "working_dir",
										Description: "Container's working directory.",
										Type:        schema.TypeString,
									},
									{
										Name:        "env_from",
										Description: "List of sources to populate environment variables in the container.",
										Type:        schema.TypeJSON,
										Resolver:    resolveContainerJSONField(func(c corev1.Container) interface{} { return c.EnvFrom }),
									},
									{
										Name:        "resources_limits",
										Description: "Limits describes the maximum amount of compute resources allowed.",
										Type:        schema.TypeJSON,
										Resolver:    schema.PathResolver("Resources.Limits"),
									},
									{
										Name:        "resources_requests",
										Description: "Requests describes the minimum amount of compute resources required.",
										Type:        schema.TypeJSON,
										Resolver:    schema.PathResolver("Resources.Requests"),
									},
									{
										Name:        "liveness_probe",
										Description: "Periodic probe of container liveness.",
										Type:        schema.TypeJSON,
										Resolver:    resolveContainerJSONField(func(c corev1.Container) interface{} { return c.LivenessProbe }),
									},
									{
										Name:        "readiness_probe",
										Description: "Periodic probe of container service readiness.",
										Type:        schema.TypeJSON,
										Resolver:    resolveContainerJSONField(func(c corev1.Container) interface{} { return c.ReadinessProbe }),
									},
									{
										Name:        "startup_probe",
										Description: "Startup probe indicates that the Pod has successfully initialized.",
										Type:        schema.TypeJSON,
										Resolver:    resolveContainerJSONField(func(c corev1.Container) interface{} { return c.ReadinessProbe }),
									},
									{
										Name:        "lifecycle",
										Description: "Actions that the management system should take in response to container lifecycle events.",
										Type:        schema.TypeJSON,
										Resolver:    resolveContainerJSONField(func(c corev1.Container) interface{} { return c.Lifecycle }),
									},
									{
										Name:        "termination_message_path",
										Description: "Path at which the file to which the container's termination message will be written is mounted into the container's filesystem.",
										Type:        schema.TypeString,
									},
									{
										Name:        "termination_message_policy",
										Description: "Indicate how the termination message should be populated.",
										Type:        schema.TypeString,
									},
									{
										Name:        "image_pull_policy",
										Description: "Image pull policy.",
										Type:        schema.TypeString,
									},
									{
										Name:        "security_context",
										Description: "security options the container should be run with.",
										Type:        schema.TypeJSON,
										Resolver:    resolveContainerJSONField(func(c corev1.Container) interface{} { return c.SecurityContext }),
									},
									{
										Name:        "stdin",
										Description: "Whether this container should allocate a buffer for stdin in the container runtime",
										Type:        schema.TypeBool,
									},
									{
										Name:        "stdin_once",
										Description: "Whether the container runtime should close the stdin channel after it has been opened by a single attach",
										Type:        schema.TypeBool,
									},
									{
										Name:        "tty",
										Description: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",
										Type:        schema.TypeBool,
										Resolver:    schema.PathResolver("TTY"),
									},
								},
								Relations: []*schema.Table{
									{
										Name:        "k8s_batch_cron_job_template_pod_init_container_ports",
										Description: "ContainerPort represents a network port in a single container.",
										Resolver:    fetchCorePodContainerPorts,
										Columns: []schema.Column{
											{
												Name:        "container_cq_id",
												Description: "Unique CloudQuery ID of k8s_batch_cron_job_template_pod_init_containers table (FK)",
												Type:        schema.TypeUUID,
												Resolver:    schema.ParentIdResolver,
											},
											{
												Name:        "name",
												Description: "If specified, this must be an IANA_SVC_NAME and unique within the pod",
												Type:        schema.TypeString,
											},
											{
												Name:        "host_port",
												Description: "Number of port to expose on the host.",
												Type:        schema.TypeInt,
											},
											{
												Name:        "container_port",
												Description: "Number of port to expose on the pod's IP address.",
												Type:        schema.TypeInt,
											},
											{
												Name:        "protocol",
												Description: "Protocol for port",
												Type:        schema.TypeString,
											},
											{
												Name:        "host_ip",
												Description: "What host IP to bind the external port to.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("HostIP"),
											},
										},
									},
									{
										Name:        "k8s_batch_cron_job_template_pod_init_container_envs",
										Description: "EnvVar represents an environment variable present in a Container.",
										Resolver:    fetchCorePodContainerEnvs,
										Columns: []schema.Column{
											{
												Name:        "container_cq_id",
												Description: "Unique CloudQuery ID of k8s_batch_cron_job_template_pod_init_containers table (FK)",
												Type:        schema.TypeUUID,
												Resolver:    schema.ParentIdResolver,
											},
											{
												Name:        "name",
												Description: "Name of the environment variable",
												Type:        schema.TypeString,
											},
											{
												Name:        "value",
												Description: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables",
												Type:        schema.TypeString,
											},
											{
												Name:        "value_from_field_ref_api_version",
												Description: "Version of the schema the FieldPath is written in terms of.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.FieldRef.APIVersion"),
											},
											{
												Name:        "value_from_field_ref_field_path",
												Description: "Path of the field to select in the specified API version.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.FieldRef.FieldPath"),
											},
											{
												Name:        "value_from_resource_field_ref_container_name",
												Description: "Container name: required for volumes, optional for env vars.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.ResourceFieldRef.ContainerName"),
											},
											{
												Name:        "value_from_resource_field_ref_resource",
												Description: "Required: resource to select",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.ResourceFieldRef.Resource"),
											},
											{
												Name:     "value_from_resource_field_ref_divisor_format",
												Type:     schema.TypeString,
												Resolver: schema.PathResolver("ValueFrom.ResourceFieldRef.Divisor.Format"),
											},
											{
												Name:        "value_from_config_map_key_ref_local_object_reference_name",
												Description: "Name of the referent.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name"),
											},
											{
												Name:        "value_from_config_map_key_ref_key",
												Description: "The key to select.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.Key"),
											},
											{
												Name:        "value_from_config_map_key_ref_optional",
												Description: "Specify whether the ConfigMap or its key must be defined.",
												Type:        schema.TypeBool,
												Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.Optional"),
											},
											{
												Name:        "value_from_secret_key_ref_local_object_reference_name",
												Description: "Name of the referent.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.LocalObjectReference.Name"),
											},
											{
												Name:        "value_from_secret_key_ref_key",
												Description: "The key of the secret to select from",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.Key"),
											},
											{
												Name:        "value_from_secret_key_ref_optional",
												Description: "Specify whether the Secret or its key must be defined",
												Type:        schema.TypeBool,
												Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.Optional"),
											},
										},
									},
									{
										Name:        "k8s_batch_cron_job_template_pod_init_container_volume_mounts",
										Description: "VolumeMount describes a mounting of a Volume within a container.",
										Resolver:    fetchCorePodContainerVolumeMounts,
										Columns: []schema.Column{
											{
												Name:        "container_cq_id",
												Description: "Unique CloudQuery ID of k8s_batch_cron_job_template_pod_init_containers table (FK)",
												Type:        schema.TypeUUID,
												Resolver:    schema.ParentIdResolver,
											},
											{
												Name:        "name",
												Description: "This must match the Name of a Volume.",
												Type:        schema.TypeString,
											},
											{
												Name:        "read_only",
												Description: "Mounted read-only if true, read-write otherwise (false or unspecified).",
												Type:        schema.TypeBool,
											},
											{
												Name:        "mount_path",
												Description: "Path within the container at which the volume should be mounted",
												Type:        schema.TypeString,
											},
											{
												Name:        "sub_path",
												Description: "Path within the volume from which the container's volume should be mounted.",
												Type:        schema.TypeString,
											},
											{
												Name:        "mount_propagation",
												Description: "Determines how mounts are propagated from the host to container and the other way around.",
												Type:        schema.TypeString,
											},
											{
												Name:        "sub_path_expr",
												Description: "Expanded path within the volume from which the container's volume should be mounted.",
												Type:        schema.TypeString,
											},
										},
									},
									{
										Name:        "k8s_batch_cron_job_template_pod_init_container_volume_devices",
										Description: "volumeDevice describes a mapping of a raw block device within a container.",
										Resolver:    fetchCorePodContainerVolumeDevices,
										Columns: []schema.Column{
											{
												Name:        "container_cq_id",
												Description: "Unique CloudQuery ID of k8s_batch_cron_job_template_pod_init_containers table (FK)",
												Type:        schema.TypeUUID,
												Resolver:    schema.ParentIdResolver,
											},
											{
												Name:        "name",
												Description: "name must match the name of a persistentVolumeClaim in the pod",
												Type:        schema.TypeString,
											},
											{
												Name:        "device_path",
												Description: "devicePath is the path inside of the container that the device will be mapped to.",
												Type:        schema.TypeString,
											},
										},
									},
								},
							},
							{
								Name:        "k8s_batch_cron_job_template_pod_containers",
								Description: "A single application container that you want to run within a pod.",
								Resolver:    fetchBatchCronJobTemplatePodContainers,
								Columns: []schema.Column{
									{
										Name:        "pod_cq_id",
										Description: "Unique CloudQuery ID of k8s_batch_cron_job_template_pod table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "name",
										Description: "Name of the container specified as a DNS_LABEL.",
										Type:        schema.TypeString,
									},
									{
										Name:        "image",
										Description: "Docker image name.",
										Type:        schema.TypeString,
									},
									{
										Name:        "command",
										Description: "Entrypoint array",
										Type:        schema.TypeStringArray,
									},
									{
										Name:        "args",
										Description: "Arguments to the entrypoint.",
										Type:        schema.TypeStringArray,
									},
									{
										Name:        "working_dir",
										Description: "Container's working directory.",
										Type:        schema.TypeString,
									},
									{
										Name:        "env_from",
										Description: "List of sources to populate environment variables in the container.",
										Type:        schema.TypeJSON,
										Resolver:    resolveContainerJSONField(func(c corev1.Container) interface{} { return c.EnvFrom }),
									},
									{
										Name:        "resources_limits",
										Description: "Limits describes the maximum amount of compute resources allowed.",
										Type:        schema.TypeJSON,
										Resolver:    schema.PathResolver("Resources.Limits"),
									},
									{
										Name:        "resources_requests",
										Description: "Requests describes the minimum amount of compute resources required.",
										Type:        schema.TypeJSON,
										Resolver:    schema.PathResolver("Resources.Requests"),
									},
									{
										Name:        "liveness_probe",
										Description: "Periodic probe of container liveness.",
										Type:        schema.TypeJSON,
										Resolver:    resolveContainerJSONField(func(c corev1.Container) interface{} { return c.LivenessProbe }),
									},
									{
										Name:        "readiness_probe",
										Description: "Periodic probe of container service readiness.",
										Type:        schema.TypeJSON,
										Resolver:    resolveContainerJSONField(func(c corev1.Container) interface{} { return c.ReadinessProbe }),
									},
									{
										Name:        "startup_probe",
										Description: "Startup probe indicates that the Pod has successfully initialized.",
										Type:        schema.TypeJSON,
										Resolver:    resolveContainerJSONField(func(c corev1.Container) interface{} { return c.StartupProbe }),
									},
									{
										Name:        "lifecycle",
										Description: "Actions that the management system should take in response to container lifecycle events.",
										Type:        schema.TypeJSON,
										Resolver:    resolveContainerJSONField(func(c corev1.Container) interface{} { return c.Lifecycle }),
									},
									{
										Name:        "termination_message_path",
										Description: "Path at which the file to which the container's termination message will be written is mounted into the container's filesystem.",
										Type:        schema.TypeString,
									},
									{
										Name:        "termination_message_policy",
										Description: "Indicate how the termination message should be populated.",
										Type:        schema.TypeString,
									},
									{
										Name:        "image_pull_policy",
										Description: "Image pull policy.",
										Type:        schema.TypeString,
									},
									{
										Name:        "security_context",
										Description: "security options the container should be run with.",
										Type:        schema.TypeJSON,
										Resolver:    resolveContainerJSONField(func(c corev1.Container) interface{} { return c.SecurityContext }),
									},
									{
										Name:        "stdin",
										Description: "Whether this container should allocate a buffer for stdin in the container runtime",
										Type:        schema.TypeBool,
									},
									{
										Name:        "stdin_once",
										Description: "Whether the container runtime should close the stdin channel after it has been opened by a single attach",
										Type:        schema.TypeBool,
									},
									{
										Name:        "tty",
										Description: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true.",
										Type:        schema.TypeBool,
										Resolver:    schema.PathResolver("TTY"),
									},
								},
								Relations: []*schema.Table{
									{
										Name:        "k8s_batch_cron_job_template_pod_container_ports",
										Description: "ContainerPort represents a network port in a single container.",
										Resolver:    fetchCorePodContainerPorts,
										Columns: []schema.Column{
											{
												Name:        "pod_container_cq_id",
												Description: "Unique CloudQuery ID of k8s_batch_cron_job_template_pod_containers table (FK)",
												Type:        schema.TypeUUID,
												Resolver:    schema.ParentIdResolver,
											},
											{
												Name:        "name",
												Description: "If specified, this must be an IANA_SVC_NAME and unique within the pod",
												Type:        schema.TypeString,
											},
											{
												Name:        "host_port",
												Description: "Number of port to expose on the host.",
												Type:        schema.TypeInt,
											},
											{
												Name:        "container_port",
												Description: "Number of port to expose on the pod's IP address.",
												Type:        schema.TypeInt,
											},
											{
												Name:        "protocol",
												Description: "Protocol for port",
												Type:        schema.TypeString,
											},
											{
												Name:        "host_ip",
												Description: "What host IP to bind the external port to.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("HostIP"),
											},
										},
									},
									{
										Name:        "k8s_batch_cron_job_template_pod_container_envs",
										Description: "EnvVar represents an environment variable present in a Container.",
										Resolver:    fetchCorePodContainerEnvs,
										Columns: []schema.Column{
											{
												Name:        "pod_container_cq_id",
												Description: "Unique CloudQuery ID of k8s_batch_cron_job_template_pod_containers table (FK)",
												Type:        schema.TypeUUID,
												Resolver:    schema.ParentIdResolver,
											},
											{
												Name:        "name",
												Description: "Name of the environment variable",
												Type:        schema.TypeString,
											},
											{
												Name:        "value",
												Description: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables",
												Type:        schema.TypeString,
											},
											{
												Name:        "value_from_field_ref_api_version",
												Description: "Version of the schema the FieldPath is written in terms of, defaults to \"v1\".",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.FieldRef.APIVersion"),
											},
											{
												Name:        "value_from_field_ref_field_path",
												Description: "Path of the field to select in the specified API version.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.FieldRef.FieldPath"),
											},
											{
												Name:        "value_from_resource_field_ref_container_name",
												Description: "Container name: required for volumes, optional for env vars",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.ResourceFieldRef.ContainerName"),
											},
											{
												Name:        "value_from_resource_field_ref_resource",
												Description: "Required: resource to select",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.ResourceFieldRef.Resource"),
											},
											{
												Name:     "value_from_resource_field_ref_divisor_format",
												Type:     schema.TypeString,
												Resolver: schema.PathResolver("ValueFrom.ResourceFieldRef.Divisor.Format"),
											},
											{
												Name:        "value_from_config_map_key_ref_local_object_reference_name",
												Description: "Name of the referent.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name"),
											},
											{
												Name:        "value_from_config_map_key_ref_key",
												Description: "The key to select.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.Key"),
											},
											{
												Name:        "value_from_config_map_key_ref_optional",
												Description: "Specify whether the ConfigMap or its key must be defined",
												Type:        schema.TypeBool,
												Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.Optional"),
											},
											{
												Name:        "value_from_secret_key_ref_local_object_reference_name",
												Description: "Name of the referent.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.LocalObjectReference.Name"),
											},
											{
												Name:        "value_from_secret_key_ref_key",
												Description: "The key of the secret to select from",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.Key"),
											},
											{
												Name:        "value_from_secret_key_ref_optional",
												Description: "Specify whether the Secret or its key must be defined.",
												Type:        schema.TypeBool,
												Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.Optional"),
											},
										},
									},
									{
										Name:        "k8s_batch_cron_job_template_pod_container_volume_mounts",
										Description: "VolumeMount describes a mounting of a Volume within a container.",
										Resolver:    fetchCorePodContainerVolumeMounts,
										Columns: []schema.Column{
											{
												Name:        "pod_container_cq_id",
												Description: "Unique CloudQuery ID of k8s_batch_cron_job_template_pod_containers table (FK)",
												Type:        schema.TypeUUID,
												Resolver:    schema.ParentIdResolver,
											},
											{
												Name:        "name",
												Description: "This must match the Name of a Volume.",
												Type:        schema.TypeString,
											},
											{
												Name:        "read_only",
												Description: "Mounted read-only if true, read-write otherwise (false or unspecified).",
												Type:        schema.TypeBool,
											},
											{
												Name:        "mount_path",
												Description: "Path within the container at which the volume should be mounted",
												Type:        schema.TypeString,
											},
											{
												Name:        "sub_path",
												Description: "Path within the volume from which the container's volume should be mounted.",
												Type:        schema.TypeString,
											},
											{
												Name:        "mount_propagation",
												Description: "Determines how mounts are propagated from the host to container and the other way around.",
												Type:        schema.TypeString,
											},
											{
												Name:        "sub_path_expr",
												Description: "Expanded path within the volume from which the container's volume should be mounted.",
												Type:        schema.TypeString,
											},
										},
									},
									{
										Name:        "k8s_batch_cron_job_template_pod_container_volume_devices",
										Description: "volumeDevice describes a mapping of a raw block device within a container.",
										Resolver:    fetchCorePodContainerVolumeDevices,
										Columns: []schema.Column{
											{
												Name:        "pod_container_cq_id",
												Description: "Unique CloudQuery ID of k8s_batch_cron_job_template_pod_containers table (FK)",
												Type:        schema.TypeUUID,
												Resolver:    schema.ParentIdResolver,
											},
											{
												Name:        "name",
												Description: "name must match the name of a persistentVolumeClaim in the pod",
												Type:        schema.TypeString,
											},
											{
												Name:        "device_path",
												Description: "devicePath is the path inside of the container that the device will be mapped to.",
												Type:        schema.TypeString,
											},
										},
									},
								},
							},
							{
								Name:        "k8s_batch_cron_job_template_pod_ephemeral_containers",
								Description: "An EphemeralContainer is a container that may be added temporarily to an existing pod for user-initiated activities such as debugging",
								Resolver:    fetchBatchCronJobTemplatePodEphemeralContainers,
								Columns: []schema.Column{
									{
										Name:        "target_container_name",
										Description: "If set, the name of the container from PodSpec that this ephemeral container targets.",
										Type:        schema.TypeString,
									},
									{
										Name:        "pod_cq_id",
										Description: "Unique CloudQuery ID of k8s_batch_cron_job_template_pod table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "name",
										Description: "Name of the container specified as a DNS_LABEL.",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("EphemeralContainerCommon.Name"),
									},
									{
										Name:        "image",
										Description: "Docker image name.",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("EphemeralContainerCommon.Image"),
									},
									{
										Name:        "command",
										Description: "Entrypoint array",
										Type:        schema.TypeStringArray,
										Resolver:    schema.PathResolver("EphemeralContainerCommon.Command"),
									},
									{
										Name:        "args",
										Description: "Arguments to the entrypoint.",
										Type:        schema.TypeStringArray,
										Resolver:    schema.PathResolver("EphemeralContainerCommon.Args"),
									},
									{
										Name:        "working_dir",
										Description: "Container's working directory.",
										Resolver:    schema.PathResolver("EphemeralContainerCommon.WorkingDir"),
										Type:        schema.TypeString,
									},
									{
										Name:        "env_from",
										Description: "List of sources to populate environment variables in the container.",
										Type:        schema.TypeJSON,
										Resolver:    resolveEphemeralContainerJSONField(func(c corev1.EphemeralContainer) interface{} { return c.EphemeralContainerCommon.EnvFrom }),
									},
									{
										Name:        "resources_limits",
										Description: "Limits describes the maximum amount of compute resources allowed.",
										Type:        schema.TypeJSON,
										Resolver:    schema.PathResolver("EphemeralContainerCommon.Resources.Limits"),
									},
									{
										Name:        "resources_requests",
										Description: "Requests describes the minimum amount of compute resources required.",
										Type:        schema.TypeJSON,
										Resolver:    schema.PathResolver("EphemeralContainerCommon.Resources.Requests"),
									},
									{
										Name:        "liveness_probe",
										Description: "Periodic probe of container liveness.",
										Type:        schema.TypeJSON,
										Resolver:    resolveEphemeralContainerJSONField(func(c corev1.EphemeralContainer) interface{} { return c.EphemeralContainerCommon.LivenessProbe }),
									},
									{
										Name:        "readiness_probe",
										Description: "Periodic probe of container service readiness.",
										Type:        schema.TypeJSON,
										Resolver:    resolveEphemeralContainerJSONField(func(c corev1.EphemeralContainer) interface{} { return c.EphemeralContainerCommon.ReadinessProbe }),
									},
									{
										Name:        "startup_probe",
										Description: "Startup probe indicates that the Pod has successfully initialized.",
										Type:        schema.TypeJSON,
										Resolver:    resolveEphemeralContainerJSONField(func(c corev1.EphemeralContainer) interface{} { return c.EphemeralContainerCommon.StartupProbe }),
									},
									{
										Name:        "lifecycle",
										Description: "Actions that the management system should take in response to container lifecycle events.",
										Type:        schema.TypeJSON,
										Resolver:    resolveEphemeralContainerJSONField(func(c corev1.EphemeralContainer) interface{} { return c.EphemeralContainerCommon.Lifecycle }),
									},
									{
										Name:        "termination_message_path",
										Description: "Path at which the file to which the container's termination message will be written is mounted into the container's filesystem.",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("EphemeralContainerCommon.TerminationMessagePath"),
									},
									{
										Name:        "termination_message_policy",
										Description: "Indicate how the termination message should be populated.",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("EphemeralContainerCommon.TerminationMessagePolicy"),
									},
									{
										Name:        "image_pull_policy",
										Description: "Image pull policy.",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("EphemeralContainerCommon.ImagePullPolicy"),
									},
									{
										Name:        "security_context",
										Description: "security options the container should be run with.",
										Type:        schema.TypeJSON,
										Resolver:    resolveEphemeralContainerJSONField(func(c corev1.EphemeralContainer) interface{} { return c.EphemeralContainerCommon.SecurityContext }),
									},
									{
										Name:        "stdin",
										Description: "Whether this container should allocate a buffer for stdin in the container runtime",
										Type:        schema.TypeBool,
										Resolver:    schema.PathResolver("EphemeralContainerCommon.Stdin"),
									},
									{
										Name:        "stdin_once",
										Description: "Whether the container runtime should close the stdin channel after it has been opened by a single attach",
										Type:        schema.TypeBool,
										Resolver:    schema.PathResolver("EphemeralContainerCommon.StdinOnce"),
									},
									{
										Name:        "tty",
										Description: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true.",
										Type:        schema.TypeBool,
										Resolver:    schema.PathResolver("EphemeralContainerCommon.TTY"),
									},
								},
								Relations: []*schema.Table{
									{
										Name:        "k8s_batch_cron_job_template_pod_ephemeral_container_ports",
										Description: "ContainerPort represents a network port in a single container.",
										Resolver:    fetchCorePodEphemeralContainerPorts,
										Columns: []schema.Column{
											{
												Name:        "pod_ephemeral_container_cq_id",
												Description: "Unique CloudQuery ID of k8s_batch_cron_job_template_pod_ephemeral_containers table (FK)",
												Type:        schema.TypeUUID,
												Resolver:    schema.ParentIdResolver,
											},
											{
												Name:        "name",
												Description: "If specified, this must be an IANA_SVC_NAME and unique within the pod",
												Type:        schema.TypeString,
											},
											{
												Name:        "host_port",
												Description: "Number of port to expose on the host.",
												Type:        schema.TypeInt,
											},
											{
												Name:        "container_port",
												Description: "Number of port to expose on the pod's IP address.",
												Type:        schema.TypeInt,
											},
											{
												Name:        "protocol",
												Description: "Protocol for port",
												Type:        schema.TypeString,
											},
											{
												Name:        "host_ip",
												Description: "What host IP to bind the external port to.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("HostIP"),
											},
										},
									},
									{
										Name:        "k8s_batch_cron_job_template_pod_ephemeral_container_envs",
										Description: "EnvVar represents an environment variable present in a Container.",
										Resolver:    fetchCorePodEphemeralContainerEnvs,
										Columns: []schema.Column{
											{
												Name:        "pod_ephemeral_container_cq_id",
												Description: "Unique CloudQuery ID of k8s_batch_cron_job_template_pod_ephemeral_containers table (FK)",
												Type:        schema.TypeUUID,
												Resolver:    schema.ParentIdResolver,
											},
											{
												Name:        "name",
												Description: "Name of the environment variable",
												Type:        schema.TypeString,
											},
											{
												Name:        "value",
												Description: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables",
												Type:        schema.TypeString,
											},
											{
												Name:        "value_from_field_ref_api_version",
												Description: "Version of the schema the FieldPath is written in terms of, defaults to \"v1\".",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.FieldRef.APIVersion"),
											},
											{
												Name:        "value_from_field_ref_field_path",
												Description: "Path of the field to select in the specified API version.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.FieldRef.FieldPath"),
											},
											{
												Name:        "value_from_resource_field_ref_container_name",
												Description: "Container name: required for volumes, optional for env vars.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.ResourceFieldRef.ContainerName"),
											},
											{
												Name:        "value_from_resource_field_ref_resource",
												Description: "Required: resource to select",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.ResourceFieldRef.Resource"),
											},
											{
												Name:     "value_from_resource_field_ref_divisor_format",
												Type:     schema.TypeString,
												Resolver: schema.PathResolver("ValueFrom.ResourceFieldRef.Divisor.Format"),
											},
											{
												Name:        "value_from_config_map_key_ref_local_object_reference_name",
												Description: "Name of the referent.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name"),
											},
											{
												Name:        "value_from_config_map_key_ref_key",
												Description: "The key to select.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.Key"),
											},
											{
												Name:        "value_from_config_map_key_ref_optional",
												Description: "Specify whether the ConfigMap or its key must be defined",
												Type:        schema.TypeBool,
												Resolver:    schema.PathResolver("ValueFrom.ConfigMapKeyRef.Optional"),
											},
											{
												Name:        "value_from_secret_key_ref_local_object_reference_name",
												Description: "Name of the referent.",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.LocalObjectReference.Name"),
											},
											{
												Name:        "value_from_secret_key_ref_key",
												Description: "The key of the secret to select from",
												Type:        schema.TypeString,
												Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.Key"),
											},
											{
												Name:        "value_from_secret_key_ref_optional",
												Description: "Specify whether the Secret or its key must be defined.",
												Type:        schema.TypeBool,
												Resolver:    schema.PathResolver("ValueFrom.SecretKeyRef.Optional"),
											},
										},
									},
									{
										Name:        "k8s_batch_cron_job_template_pod_ephemeral_container_volume_mounts",
										Description: "VolumeMount describes a mounting of a Volume within a container.",
										Resolver:    fetchCorePodEphemeralContainerVolumeMounts,
										Columns: []schema.Column{
											{
												Name:        "pod_ephemeral_container_cq_id",
												Description: "Unique CloudQuery ID of k8s_batch_cron_job_template_pod_ephemeral_containers table (FK)",
												Type:        schema.TypeUUID,
												Resolver:    schema.ParentIdResolver,
											},
											{
												Name:        "name",
												Description: "This must match the Name of a Volume.",
												Type:        schema.TypeString,
											},
											{
												Name:        "read_only",
												Description: "Mounted read-only if true, read-write otherwise (false or unspecified).",
												Type:        schema.TypeBool,
											},
											{
												Name:        "mount_path",
												Description: "Path within the container at which the volume should be mounted",
												Type:        schema.TypeString,
											},
											{
												Name:        "sub_path",
												Description: "Path within the volume from which the container's volume should be mounted.",
												Type:        schema.TypeString,
											},
											{
												Name:        "mount_propagation",
												Description: "mountPropagation determines how mounts are propagated from the host to container and the other way around.",
												Type:        schema.TypeString,
											},
											{
												Name:        "sub_path_expr",
												Description: "Expanded path within the volume from which the container's volume should be mounted.",
												Type:        schema.TypeString,
											},
										},
									},
									{
										Name:        "k8s_batch_cron_job_template_pod_ephemeral_containers_vol_devices",
										Description: "volumeDevice describes a mapping of a raw block device within a container.",
										Resolver:    fetchCorePodEphemeralContainerVolumeDevices,
										Columns: []schema.Column{
											{
												Name:        "pod_ephemeral_container_cq_id",
												Description: "Unique CloudQuery ID of k8s_batch_cron_job_template_pod_ephemeral_containers table (FK)",
												Type:        schema.TypeUUID,
												Resolver:    schema.ParentIdResolver,
											},
											{
												Name:        "name",
												Description: "name must match the name of a persistentVolumeClaim in the pod",
												Type:        schema.TypeString,
											},
											{
												Name:        "device_path",
												Description: "devicePath is the path inside of the container that the device will be mapped to.",
												Type:        schema.TypeString,
											},
										},
									},
								},
							},
							{
								Name:        "k8s_batch_cron_job_template_pod_volumes",
								Description: "Volume represents a named volume in a pod that may be accessed by any container in the pod.",
								Resolver:    fetchBatchCronJobTemplatePodVolumes,
								Columns: []schema.Column{
									{
										Name:        "pod_cq_id",
										Description: "Unique CloudQuery ID of k8s_batch_cron_job_template_pod table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "name",
										Description: "Volume's name. Must be a DNS_LABEL and unique within the pod.",
										Type:        schema.TypeString,
									},
									{
										Name:        "host_path",
										Description: "Pre-existing file or directory on the host machine that is directly exposed to the container.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.HostPath }),
									},
									{
										Name:        "empty_dir",
										Description: "Temporary directory that shares a pod's lifetime.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.EmptyDir }),
									},
									{
										Name:        "gce_persistent_disk",
										Description: "GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.GCEPersistentDisk }),
									},
									{
										Name:        "aws_elastic_block_store",
										Description: "AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.AWSElasticBlockStore }),
									},
									{
										Name:        "secret",
										Description: "A secret that should populate this volume.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.Secret }),
									},
									{
										Name:        "nfs",
										Description: "NFS mount on the host that shares a pod's lifetime",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.NFS }),
									},
									{
										Name:        "iscsi",
										Description: "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.ISCSI }),
									},
									{
										Name:        "glusterfs",
										Description: "Glusterfs mount on the host that shares a pod's lifetime.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.Glusterfs }),
									},
									{
										Name:        "persistent_volume_claim",
										Description: "Persistent volume claim.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.PersistentVolumeClaim }),
									},
									{
										Name:        "rbd",
										Description: "Rados Block Device mount on the host that shares a pod's lifetime.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.RBD }),
									},
									{
										Name:        "flex_volume",
										Description: "Generic volume resource that is provisioned/attached using an exec based plugin.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.FlexVolume }),
									},
									{
										Name:        "cinder",
										Description: "Cinder volume attached and mounted on kubelets host machine.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.Cinder }),
									},
									{
										Name:        "ceph_fs",
										Description: "Ceph FS mount on the host that shares a pod's lifetime.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.CephFS }),
									},
									{
										Name:        "flocker",
										Description: "Flocker volume attached to a kubelet's host machine.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.Flocker }),
									},
									{
										Name:        "downward_api",
										Description: "Optional: mode bits to use on created files by default",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.DownwardAPI }),
									},
									{
										Name:        "fc",
										Description: "Fibre Channel resource that is attached to a kubelet's host machine.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.FC }),
									},
									{
										Name:        "azure_file",
										Description: "Azure File Service mount on the host and bind mount to the pod.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.AzureFile }),
									},
									{
										Name:        "config_map",
										Description: "configMap that should populate this volume",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.ConfigMap }),
									},
									{
										Name:        "vsphere_volume",
										Description: "vSphere volume attached and mounted on kubelets host machine.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.VsphereVolume }),
									},
									{
										Name:        "quobyte",
										Description: "Quobyte mount on the host that shares a pod's lifetime.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.Quobyte }),
									},
									{
										Name:        "azure_disk",
										Description: "The Name of the data disk in the blob storage",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.AzureDisk }),
									},
									{
										Name:        "photon_persistent_disk",
										Description: "PhotonController persistent disk attached and mounted on kubelets host machine.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.PhotonPersistentDisk }),
									},
									{
										Name:        "projected",
										Description: "Items for all in one resources secrets, configmaps, and downward API.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.Projected }),
									},
									{
										Name:        "portworx_volume",
										Description: "Portworx volume attached and mounted on kubelets host machine.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.PortworxVolume }),
									},
									{
										Name:        "scale_io",
										Description: "ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.ScaleIO }),
									},
									{
										Name:        "storage_os",
										Description: "StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.StorageOS }),
									},
									{
										Name:        "csi",
										Description: "CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.CSI }),
									},
									{
										Name:        "ephemeral",
										Description: "Ephemeral represents a volume that is handled by a cluster storage driver.",
										Type:        schema.TypeJSON,
										Resolver:    resolveVolumeJSONField(func(v corev1.Volume) interface{} { return v.Ephemeral }),
									},
								},
							},
						},
					}},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchBatchCronJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	jobs := meta.(*client.Client).Services.CronJobs
	opts := metav1.ListOptions{}
	for {
		result, err := jobs.List(ctx, opts)
		if err != nil {
			return err
		}
		res <- result.Items
		next := result.GetContinue()
		if next == "" {
			return nil
		}
		opts.Continue = next
	}
}

func resolveBatchCronJobOwnerReferences(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cronJob, ok := resource.Item.(batchv1.CronJob)
	if !ok {
		return fmt.Errorf("not a batchv1.CronJob instance: %T", resource.Item)
	}
	b, err := json.Marshal(cronJob.ObjectMeta.OwnerReferences)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}

func fetchBatchCronJobTemplates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	cronJob, ok := parent.Item.(batchv1.CronJob)
	if !ok {
		return fmt.Errorf("not a batchv1.CronJob instance: %T", parent.Item)
	}
	res <- cronJob.Spec.JobTemplate
	return nil
}

func resolveBatchCronJobTemplateOwnerReferences(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	template, ok := resource.Item.(batchv1.JobTemplateSpec)
	if !ok {
		return fmt.Errorf("not a batchv1.JobTemplateSpec instance: %T", resource.Item)
	}
	b, err := json.Marshal(template.ObjectMeta.OwnerReferences)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}

func resolveBatchCronJobTemplateSelector(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	template, ok := resource.Item.(batchv1.JobTemplateSpec)
	if !ok {
		return fmt.Errorf("not a batchv1.JobTemplateSpec instance: %T", resource.Item)
	}
	b, err := json.Marshal(template.Spec.Selector)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}

func fetchBatchCronJobTemplatePod(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	template, ok := parent.Item.(batchv1.JobTemplateSpec)
	if !ok {
		return fmt.Errorf("not a batchv1.JobTemplateSpec instance: %T", parent.Item)
	}
	res <- template.Spec.Template
	return nil
}

func resolveBatchCronJobTemplatePodJSON(valueGetter func(podSpec corev1.PodTemplateSpec) interface{}) func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return func(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		pod, ok := resource.Item.(corev1.PodTemplateSpec)
		if !ok {
			return fmt.Errorf("not a batchv1.PodTemplateSpec instance: %T", resource.Item)
		}
		b, err := json.Marshal(valueGetter(pod))
		if err != nil {
			return err
		}
		return resource.Set(c.Name, b)
	}
}

func fetchBatchCronJobTemplatePodInitContainers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	podSpec, ok := parent.Item.(corev1.PodTemplateSpec)
	if !ok {
		return fmt.Errorf("not a corev1.PodTemplateSpec instance: %T", parent.Item)
	}
	res <- podSpec.Spec.InitContainers
	return nil
}

func fetchBatchCronJobTemplatePodContainers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	podSpec, ok := parent.Item.(corev1.PodTemplateSpec)
	if !ok {
		return fmt.Errorf("not a corev1.PodTemplateSpec instance: %T", parent.Item)
	}
	res <- podSpec.Spec.Containers
	return nil
}

func fetchBatchCronJobTemplatePodEphemeralContainers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	podSpec, ok := parent.Item.(corev1.PodTemplateSpec)
	if !ok {
		return fmt.Errorf("not a corev1.PodTemplateSpec instance: %T", parent.Item)
	}
	res <- podSpec.Spec.EphemeralContainers
	return nil
}

func fetchBatchCronJobTemplatePodVolumes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	podSpec, ok := parent.Item.(corev1.PodTemplateSpec)
	if !ok {
		return fmt.Errorf("not a corev1.PodTemplateSpec instance: %T", parent.Item)
	}
	res <- podSpec.Spec.Volumes
	return nil
}
