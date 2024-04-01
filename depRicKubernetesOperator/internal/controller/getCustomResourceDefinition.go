package controller

import (
	"context"
	"fmt"
)

func GetCustomResourceDefinition() []*unstructured.Unstructured {

	customResourceDefinition1 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/instance":   "release-name",
					"app.kubernetes.io/managed-by": "Helm",
					"app.kubernetes.io/name":       "kong",
					"app.kubernetes.io/version":    "1.4",
					"helm.sh/chart":                "kong-0.36.6",
				},
				"name": "kongconsumers.configuration.konghq.com",
			},
			"spec": map[string]interface{}{
				"group": "configuration.konghq.com",
				"names": map[string]interface{}{
					"kind":   "KongConsumer",
					"plural": "kongconsumers",
					"shortNames": []interface{}{
						"kc",
					},
				},
				"scope": "Namespaced",
				"versions": map[string]interface{}{
					"- name": "v1",
					"served": true,
					"storage": true,
					"schema": map[string]interface{}{
						"openAPIV3Schema": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"credentials": map[string]interface{}{
									"items": map[string]interface{}{
										"type": "string",
									},
									"type": "array",
								},
								"custom_id": map[string]interface{}{
									"type": "string",
								},
								"username": map[string]interface{}{
									"type": "string",
								},
							},
						},
					},
					"additionalPrinterColumns": []interface{}{
						map[string]interface{}{
							"jsonPath":    ".username",
							"description": "Username of a Kong Consumer",
							"name":        "Username",
							"type":        "string",
						},
						map[string]interface{}{
							"type":        "date",
							"jsonPath":    ".metadata.creationTimestamp",
							"description": "Age",
							"name":        "Age",
						},
					},
				},
			},
		},
	}

	customResourceDefinition2 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/instance":   "release-name",
					"app.kubernetes.io/managed-by": "Helm",
					"app.kubernetes.io/name":       "kong",
					"app.kubernetes.io/version":    "1.4",
					"helm.sh/chart":                "kong-0.36.6",
				},
				"name": "kongcredentials.configuration.konghq.com",
			},
			"spec": map[string]interface{}{
				"group": "configuration.konghq.com",
				"names": map[string]interface{}{
					"kind":   "KongCredential",
					"plural": "kongcredentials",
				},
				"scope": "Namespaced",
				"versions": map[string]interface{}{
					"- name": "v1",
					"served": true,
					"storage": true,
					"additionalPrinterColumns": []interface{}{
						map[string]interface{}{
							"jsonPath":    ".type",
							"description": "Type of credential",
							"name":        "Credential-type",
							"type":        "string",
						},
						map[string]interface{}{
							"jsonPath":    ".metadata.creationTimestamp",
							"description": "Age",
							"name":        "Age",
							"type":        "date",
						},
						map[string]interface{}{
							"type":        "string",
							"jsonPath":    ".consumerRef",
							"description": "Owner of the credential",
							"name":        "Consumer-Ref",
						},
					},
					"schema": map[string]interface{}{
						"openAPIV3Schema": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"consumerRef": map[string]interface{}{
									"type": "string",
								},
								"type": map[string]interface{}{
									"type": "string",
								},
							},
							"required": []interface{}{
								"consumerRef",
								"type",
							},
						},
					},
				},
			},
		},
	}

	customResourceDefinition3 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/instance":   "release-name",
					"app.kubernetes.io/managed-by": "Helm",
					"app.kubernetes.io/name":       "kong",
					"app.kubernetes.io/version":    "1.4",
					"helm.sh/chart":                "kong-0.36.6",
				},
				"name": "kongplugins.configuration.konghq.com",
			},
			"spec": map[string]interface{}{
				"names": map[string]interface{}{
					"kind":   "KongPlugin",
					"plural": "kongplugins",
					"shortNames": []interface{}{
						"kp",
					},
				},
				"scope": "Namespaced",
				"group": "configuration.konghq.com",
				"versions": map[string]interface{}{
					"- name": "v1",
					"served": true,
					"storage": true,
					"schema": map[string]interface{}{
						"openAPIV3Schema": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"protocols": map[string]interface{}{
									"items": map[string]interface{}{
										"type": "string",
										"enum": []interface{}{
											"http",
											"https",
											"tcp",
											"tls",
											"grpc",
											"grpcs",
										},
									},
									"type": "array",
								},
								"run_on": map[string]interface{}{
									"enum": []interface{}{
										"first",
										"second",
										"all",
									},
									"type": "string",
								},
								"config": map[string]interface{}{
									"type": "object",
								},
								"disabled": map[string]interface{}{
									"type": "boolean",
								},
								"plugin": map[string]interface{}{
									"type": "string",
								},
							},
							"required": []interface{}{
								"plugin",
							},
						},
					},
					"additionalPrinterColumns": []interface{}{
						map[string]interface{}{
							"jsonPath":    ".plugin",
							"description": "Name of the plugin",
							"name":        "Plugin-Type",
							"type":        "string",
						},
						map[string]interface{}{
							"description": "Age",
							"name":        "Age",
							"type":        "date",
							"jsonPath":    ".metadata.creationTimestamp",
						},
						map[string]interface{}{
							"jsonPath":    ".disabled",
							"description": "Indicates if the plugin is disabled",
							"name":        "Disabled",
							"priority":    1,
							"type":        "boolean",
						},
						map[string]interface{}{
							"jsonPath":    ".config",
							"description": "Configuration of the plugin",
							"name":        "Config",
							"priority":    1,
							"type":        "string",
						},
					},
			},
		},
	}

	customResourceDefinition4 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"helm.sh/chart":                "kong-0.36.6",
					"app.kubernetes.io/instance":   "release-name",
					"app.kubernetes.io/managed-by": "Helm",
					"app.kubernetes.io/name":       "kong",
					"app.kubernetes.io/version":    "1.4",
				},
				"name": "kongingresses.configuration.konghq.com",
			},
			"spec": map[string]interface{}{
				"group":   "configuration.konghq.com",
				"names": map[string]interface{}{
					"plural": "kongingresses",
					"shortNames": []interface{}{
						"ki",
					},
					"kind": "KongIngress",
				},
				"scope": "Namespaced",
				"versions": map[string]interface{}{
					"- name": "v1",
					"served": true,
					"storage": true,
					"schema": map[string]interface{}{
						"openAPIV3Schema": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"upstream": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"hash_fallback_header": map[string]interface{}{
											"type": "string",
										},
										"hash_on": map[string]interface{}{
											"type": "string",
										},
										"hash_on_cookie": map[string]interface{}{
											"type": "string",
										},
										"hash_on_cookie_path": map[string]interface{}{
											"type": "string",
										},
										"healthchecks": map[string]interface{}{
											"type": "object",
											"properties": map[string]interface{}{
												"passive": map[string]interface{}{
													"type": "object",
													"properties": map[string]interface{}{
														"healthy": map[string]interface{}{
															"type": "object",
															"properties": map[string]interface{}{
																"http_statuses": map[string]interface{}{
																	"items": map[string]interface{}{
																		"type": "integer",
																	},
																	"type": "array",
																},
																"interval": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
																"successes": map[string]interface{}{
																	"type":    "integer",
																	"minimum": 0,
																},
															},
															"type": "object",
														},
														"unhealthy": map[string]interface{}{
															"type": "object",
															"properties": map[string]interface{}{
																"timeout": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
																"http_failures": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
																"http_statuses": map[string]interface{}{
																	"items": map[string]interface{}{
																		"type": "integer",
																	},
																	"type": "array",
																},
																"interval": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
																"tcp_failures": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
															},
														},
													}
												},
												"active": map[string]interface{}{
													"type": "object",
													"properties": map[string]interface{}{
														"concurrency": map[string]interface{}{
															"minimum": 1,
															"type":    "integer",
														},
														"healthy": map[string]interface{}{
															"properties": map[string]interface{}{
																"http_statuses": map[string]interface{}{
																	"items": map[string]interface{}{
																		"type": "integer",
																	},
																	"type": "array",
																},
																"interval": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
																"successes": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
															},
															"type": "object",
														},
														"http_path": map[string]interface{}{
															"pattern": "^/.*$",
															"type":    "string",
														},
														"timeout": map[string]interface{}{
															"minimum": 0,
															"type":    "integer",
														},
														"unhealthy": map[string]interface{}{
															"type": "object",
															"properties": map[string]interface{}{
																"timeout": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
																"http_failures": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
																"http_statuses": map[string]interface{}{
																	"items": map[string]interface{}{
																		"type": "integer",
																	},
																	"type": "array",
																},
																"interval": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
																"tcp_failures": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
															},
														},
													},
												},
											},
										},
										"slots": map[string]interface{}{
											"minimum": 10,
											"type":    "integer",
										},
										"algorithm": map[string]interface{}{
											"enum": []interface{}{
												"round-robin",
												"consistent-hashing",
												"least-connections",
											},
											"type": "string",
										},
										"hash_fallback": map[string]interface{}{
											"type": "string",
										},
										"hash_on_header": map[string]interface{}{
											"type": "string",
										},
										"host_header": map[string]interface{}{
											"type": "string",
										},
									},
									"type": "object",
								},
								"proxy": map[string]interface{}{
									"properties": map[string]interface{}{
										"connect_timeout": map[string]interface{}{
											"minimum": 0,
											"type":    "integer",
										},
										"path": map[string]interface{}{
											"pattern": "^/.*$",
											"type":    "string",
										},
										"protocol": map[string]interface{}{
											"enum": []interface{}{
												"http",
												"https",
												"grpc",
												"grpcs",
											},
											"type": "string",
										},
										"read_timeout": map[string]interface{}{
											"minimum": 0,
											"type":    "integer",
										},
										"retries": map[string]interface{}{
											"minimum": 0,
											"type":    "integer",
										},
										"write_timeout": map[string]interface{}{
											"minimum": 0,
											"type":    "integer",
										},
									},
									"type": "object",
								},
								"route": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"protocols": map[string]interface{}{
											"items": map[string]interface{}{
												"type": "string",
												"enum": []interface{}{
													"http",
													"https",
													"grpc",
													"grpcs",
												},
											},
											"type": "array",
										},
										"regex_priority": map[string]interface{}{
											"type": "integer",
										},
										"strip_path": map[string]interface{}{
											"type": "boolean",
										},
										"headers": map[string]interface{}{
											"additionalProperties": map[string]interface{}{
												"items": map[string]interface{}{
													"type": "string",
												},
												"type": "array",
											},
											"type": "object",
										},
										"https_redirect_status_code": map[string]interface{}{
											"type": "integer",
										},
										"methods": map[string]interface{}{
											"items": map[string]interface{}{
												"type": "string",
											},
											"type": "array",
										},
										"preserve_host": map[string]interface{}{
											"type": "boolean",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	customResourceDefinition5 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/instance":   "release-name",
					"app.kubernetes.io/managed-by": "Helm",
					"app.kubernetes.io/name":       "kong",
					"app.kubernetes.io/version":    "1.4",
					"helm.sh/chart":                "kong-0.36.6",
				},
				"name": "kongconsumers.configuration.konghq.com",
			},
			"spec": map[string]interface{}{
				"group": "configuration.konghq.com",
				"names": map[string]interface{}{
					"kind":   "KongConsumer",
					"plural": "kongconsumers",
					"shortNames": []interface{}{
						"kc",
					},
				},
				"scope": "Namespaced",
				"versions": map[string]interface{}{
					"- name": "v1",
					"served": true,
					"storage": true,
					"schema": map[string]interface{}{
						"openAPIV3Schema": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"credentials": map[string]interface{}{
									"type": "array",
									"items": map[string]interface{}{
										"type": "string",
									},
								},
								"custom_id": map[string]interface{}{
									"type": "string",
								},
								"username": map[string]interface{}{
									"type": "string",
								},
							},
						},
					},
					"additionalPrinterColumns": []interface{}{
						map[string]interface{}{
							"jsonPath":    ".username",
							"description": "Username of a Kong Consumer",
							"name":        "Username",
							"type":        "string",
						},
						map[string]interface{}{
							"jsonPath":    ".metadata.creationTimestamp",
							"description": "Age",
							"name":        "Age",
							"type":        "date",
						},
					},
				},
			},
		},
	}

	customResourceDefinition6 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/managed-by": "Helm",
					"app.kubernetes.io/name":       "kong",
					"app.kubernetes.io/version":    "1.4",
					"helm.sh/chart":                "kong-0.36.6",
					"app.kubernetes.io/instance":   "release-name",
				},
				"name": "kongcredentials.configuration.konghq.com",
			},
			"spec": map[string]interface{}{
				"group": "configuration.konghq.com",
				"names": map[string]interface{}{
					"kind":   "KongCredential",
					"plural": "kongcredentials",
				},
				"scope": "Namespaced",
				"versions": map[string]interface{}{
					"- name": "v1",
					"served": true,
					"storage": true,
					"schema": map[string]interface{}{
						"openAPIV3Schema": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"consumerRef": map[string]interface{}{
									"type": "string",
								},
								"type": map[string]interface{}{
									"type": "string",
								},
							},
							"required": []interface{}{
								"consumerRef",
								"type",
							},
						},
					},
					"additionalPrinterColumns": []interface{}{
						map[string]interface{}{
							"jsonPath":    ".type",
							"description": "Type of credential",
							"name":        "Credential-type",
							"type":        "string",
						},
						map[string]interface{}{
							"description": "Age",
							"name":        "Age",
							"type":        "date",
							"jsonPath":    ".metadata.creationTimestamp",
						},
						map[string]interface{}{
							"jsonPath":    ".consumerRef",
							"description": "Owner of the credential",
							"name":        "Consumer-Ref",
							"type":        "string",
						},
					},
				},
			},
		},
	}

	customResourceDefinition7 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app.kubernetes.io/instance":   "release-name",
					"app.kubernetes.io/managed-by": "Helm",
					"app.kubernetes.io/name":       "kong",
					"app.kubernetes.io/version":    "1.4",
					"helm.sh/chart":                "kong-0.36.6",
				},
				"name": "kongplugins.configuration.konghq.com",
			},
			"spec": map[string]interface{}{
				"group": "configuration.konghq.com",
				"scope": "Namespaced",
				"names": map[string]interface{}{
					"kind":   "KongPlugin",
					"plural": "kongplugins",
					"shortNames": []interface{}{
						"kp",
					},
				},
				"versions": map[string]interface{}{
					"- name": "v1",
					"served": true,
					"storage": true,
					"schema": map[string]interface{}{
						"openAPIV3Schema": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"config": map[string]interface{}{
									"type": "object",
								},
								"disabled": map[string]interface{}{
									"type": "boolean",
								},
								"plugin": map[string]interface{}{
									"type": "string",
								},
								"protocols": map[string]interface{}{
									"items": map[string]interface{}{
										"enum": []interface{}{
											"http",
											"https",
											"tcp",
											"tls",
											"grpc",
											"grpcs",
										},
										"type": "string",
									},
									"type": "array",
								},
								"run_on": map[string]interface{}{
									"type": "string",
									"enum": []interface{}{
										"first",
										"second",
										"all",
									},
								},
							},
							"required": []interface{}{
								"plugin",
							},
						},
					},
					"additionalPrinterColumns": []interface{}{
						map[string]interface{}{
							"description": "Name of the plugin",
							"name":        "Plugin-Type",
							"type":        "string",
							"jsonPath":    ".plugin",
						},
						map[string]interface{}{
							"description": "Age",
							"name":        "Age",
							"type":        "date",
							"jsonPath":    ".metadata.creationTimestamp",
						},
						map[string]interface{}{
							"priority":    1,
							"type":        "boolean",
							"jsonPath":    ".disabled",
							"description": "Indicates if the plugin is disabled",
							"name":        "Disabled",
						},
						map[string]interface{}{
							"jsonPath":    ".config",
							"description": "Configuration of the plugin",
							"name":        "Config",
							"priority":    1,
							"type":        "string",
						},
					},
				},
			},
		},
	}

	customResourceDefinition8 := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"helm.sh/chart":                "kong-0.36.6",
					"app.kubernetes.io/instance":   "release-name",
					"app.kubernetes.io/managed-by": "Helm",
					"app.kubernetes.io/name":       "kong",
					"app.kubernetes.io/version":    "1.4",
				},
				"name": "kongingresses.configuration.konghq.com",
			},
			"spec": map[string]interface{}{
				"group": "configuration.konghq.com",
				"names": map[string]interface{}{
					"plural": "kongingresses",
					"shortNames": []interface{}{
						"ki",
					},
					"kind": "KongIngress",
				},
				"scope": "Namespaced",
				"versions": map[string]interface{}{
					"- name": "v1",
					"served": true,
					"storage": true,
					"schema": map[string]interface{}{
						"openAPIV3Schema": map[string]interface{}{
							"type": "object",
							"properties": map[string]interface{}{
								"proxy": map[string]interface{}{
									"properties": map[string]interface{}{
										"path": map[string]interface{}{
											"pattern": "^/.*$",
											"type":    "string",
										},
										"protocol": map[string]interface{}{
											"type": "string",
											"enum": []interface{}{
												"http",
												"https",
												"grpc",
												"grpcs",
											},
										},
										"read_timeout": map[string]interface{}{
											"minimum": 0,
											"type":    "integer",
										},
										"retries": map[string]interface{}{
											"minimum": 0,
											"type":    "integer",
										},
										"write_timeout": map[string]interface{}{
											"minimum": 0,
											"type":    "integer",
										},
										"connect_timeout": map[string]interface{}{
											"minimum": 0,
											"type":    "integer",
										},
									},
									"type": "object",
								},
								"route": map[string]interface{}{
									"type": "object",
									"properties": map[string]interface{}{
										"https_redirect_status_code": map[string]interface{}{
											"type": "integer",
										},
										"methods": map[string]interface{}{
											"items": map[string]interface{}{
												"type": "string",
											},
											"type": "array",
										},
										"preserve_host": map[string]interface{}{
											"type": "boolean",
										},
										"protocols": map[string]interface{}{
											"items": map[string]interface{}{
												"type": "string",
												"enum": []interface{}{
													"http",
													"https",
													"grpc",
													"grpcs",
												},
											},
											"type": "array",
										},
										"regex_priority": map[string]interface{}{
											"type": "integer",
										},
										"strip_path": map[string]interface{}{
											"type": "boolean",
										},
										"headers": map[string]interface{}{
											"additionalProperties": map[string]interface{}{
												"items": map[string]interface{}{
													"type": "string",
												},
												"type": "array",
											},
											"type": "object",
										},
									},
								},
								"upstream": map[string]interface{}{
									"properties": map[string]interface{}{
										"algorithm": map[string]interface{}{
											"enum": []interface{}{
												"round-robin",
												"consistent-hashing",
												"least-connections",
											},
											"type": "string",
										},
										"hash_fallback_header": map[string]interface{}{
											"type": "string",
										},
										"hash_on_cookie": map[string]interface{}{
											"type": "string",
										},
										"hash_on_cookie_path": map[string]interface{}{
											"type": "string",
										},
										"host_header": map[string]interface{}{
											"type": "string",
										},
										"hash_fallback": map[string]interface{}{
											"type": "string",
										},
										"hash_on": map[string]interface{}{
											"type": "string",
										},
										"hash_on_header": map[string]interface{}{
											"type": "string",
										},
										"healthchecks": map[string]interface{}{
											"type": "object",
											"properties": map[string]interface{}{
												"active": map[string]interface{}{
													"properties": map[string]interface{}{
														"timeout": map[string]interface{}{
															"minimum": 0,
															"type":    "integer",
														},
														"unhealthy": map[string]interface{}{
															"properties": map[string]interface{}{
																"http_failures": map[string]interface{}{
																	"type":    "integer",
																	"minimum": 0,
																},
																"http_statuses": map[string]interface{}{
																	"items": map[string]interface{}{
																		"type": "integer",
																	},
																	"type": "array",
																},
																"interval": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
																"tcp_failures": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
																"timeout": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
															},
															"type": "object",
														},
														"concurrency": map[string]interface{}{
															"minimum": 1,
															"type":    "integer",
														},
														"healthy": map[string]interface{}{
															"properties": map[string]interface{}{
																"http_statuses": map[string]interface{}{
																	"items": map[string]interface{}{
																		"type": "integer",
																	},
																	"type": "array",
																},
																"interval": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
																"successes": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
															},
															"type": "object",
														},
														"http_path": map[string]interface{}{
															"pattern": "^/.*$",
															"type":    "string",
														},
													},
													"type": "object",
												},
												"passive": map[string]interface{}{
													"properties": map[string]interface{}{
														"healthy": map[string]interface{}{
															"type": "object",
															"properties": map[string]interface{}{
																"http_statuses": map[string]interface{}{
																	"items": map[string]interface{}{
																		"type": "integer",
																	},
																	"type": "array",
																},
																"interval": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
																"successes": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
															},
														},
														"unhealthy": map[string]interface{}{
															"type": "object",
															"properties": map[string]interface{}{
																"http_statuses": map[string]interface{}{
																	"type": "array",
																	"items": map[string]interface{}{
																		"type": "integer",
																	},
																},
																"interval": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
																"tcp_failures": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
																"timeout": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
																"http_failures": map[string]interface{}{
																	"minimum": 0,
																	"type":    "integer",
																},
															},
														},
													},
													"type": "object",
												},
											},
										},
										"slots": map[string]interface{}{
											"minimum": 10,
											"type":    "integer",
										},
									},
									"type": "object",
								},
							},
						},
					},
				},
			},
		},
	}

	return []*unstructured.Unstructured{customResourceDefinition1, customResourceDefinition2, customResourceDefinition3, customResourceDefinition4, customResourceDefinition5, customResourceDefinition6, customResourceDefinition7, customResourceDefinition8}
}
