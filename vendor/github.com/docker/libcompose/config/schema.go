package config

var schemaV1 = `{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "id": "config_schema_v1.json",

  "type": "object",

  "patternProperties": {
    "^[a-zA-Z0-9._-]+$": {
      "$ref": "#/definitions/service"
    }
  },

  "additionalProperties": false,

  "definitions": {
    "service": {
      "id": "#/definitions/service",
      "type": "object",

      "properties": {
        "build": {"type": "string"},
        "cap_add": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "cap_drop": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "certs": {"$ref": "#/definitions/list_of_strings"},
        "cgroup_parent": {"type": "string"},
        "command": {
          "oneOf": [
            {"type": "string"},
            {"type": "array", "items": {"type": "string"}}
          ]
        },
        "container_name": {"type": "string"},
        "cpu_shares": {"type": ["number", "string"]},
        "cpu_quota": {"type": ["number", "string"]},
        "cpuset": {"type": "string"},
        "device_read_iops": {"$ref": "#/definitions/list_or_dict"},
        "devices": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "device_write_iops": {"$ref": "#/definitions/list_or_dict"},
        "default_cert": {"type": "string"},
        "disks": {"type": "array"},
        "dns": {"$ref": "#/definitions/string_or_list"},
        "dns_search": {"$ref": "#/definitions/string_or_list"},
        "dockerfile": {"type": "string"},
        "domainname": {"type": "string"},
        "entrypoint": {
          "oneOf": [
            {"type": "string"},
            {"type": "array", "items": {"type": "string"}}
          ]
        },
        "env_file": {"$ref": "#/definitions/string_or_list"},
        "environment": {"$ref": "#/definitions/list_or_dict"},

        "expose": {
          "type": "array",
          "items": {
            "type": ["string", "number"],
            "format": "expose"
          },
          "uniqueItems": true
        },

        "extends": {
          "oneOf": [
            {
              "type": "string"
            },
            {
              "type": "object",

              "properties": {
                "service": {"type": "string"},
                "file": {"type": "string"}
              },
              "required": ["service"],
              "additionalProperties": false
            }
          ]
        },

        "extra_hosts": {"$ref": "#/definitions/list_or_dict"},
        "external_ips": {"$ref": "#/definitions/list_of_strings"},
        "external_links": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "health_check": {"type": "object"},
        "hostname": {"type": "string"},
        "image": {"type": "string"},
        "ipc": {"type": "string"},
        "labels": {"$ref": "#/definitions/list_or_dict"},
        "links": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "load_balancer_config": {"type": "object"},
        "log_driver": {"type": "string"},
        "log_opt": {"type": "object"},
        "mac_address": {"type": "string"},
        "memory": {"type": ["number", "string"]},
        "mem_limit": {"type": ["number", "string"]},
        "memswap_limit": {"type": ["number", "string"]},
        "metadata": {"type": "object"},
        "net": {"type": "string"},
        "pid": {"type": ["string", "null"]},

        "ports": {
          "type": "array",
          "items": {
            "type": ["string", "number"],
            "format": "ports"
          },
          "uniqueItems": true
        },

        "privileged": {"type": "boolean"},
        "read_only": {"type": "boolean"},
        "restart": {"type": "string"},
        "retain_ip": {"type": "boolean"},
        "scale": {"type": ["number", "string"]},
        "scale_policy": {"type": "object"},
        "security_opt": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "shm_size": {"type": ["number", "string"]},
        "stdin_open": {"type": "boolean"},
        "stop_signal": {"type": "string"},
        "tty": {"type": "boolean"},
        "type": {"type": "string"},
        "update_strategy": {"type": "object"},
        "ulimits": {
          "type": "object",
          "patternProperties": {
            "^[a-z]+$": {
              "oneOf": [
                {"type": "integer"},
                {
                  "type":"object",
                  "properties": {
                    "hard": {"type": "integer"},
                    "soft": {"type": "integer"}
                  },
                  "required": ["soft", "hard"],
                  "additionalProperties": false
                }
              ]
            }
          }
        },
        "user": {"type": "string"},
        "userdata": {"type": "string"},
        "vcpu": {"type": ["number", "string"]},
        "volumes": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "volume_driver": {"type": "string"},
        "volumes_from": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "working_dir": {"type": "string"}
      },

      "dependencies": {
        "memswap_limit": ["mem_limit"]
      },
      "additionalProperties": false
    },

    "string_or_list": {
      "oneOf": [
        {"type": "string"},
        {"$ref": "#/definitions/list_of_strings"}
      ]
    },

    "list_of_strings": {
      "type": "array",
      "items": {"type": "string"},
      "uniqueItems": true
    },

    "list_or_dict": {
      "oneOf": [
        {
          "type": "object",
          "patternProperties": {
            ".+": {
              "type": ["string", "number", "null", "boolean"]
            }
          },
          "additionalProperties": false
        },
        {"type": "array", "items": {"type": "string"}, "uniqueItems": true}
      ]
    },

    "constraints": {
      "service": {
        "id": "#/definitions/constraints/service",
        "anyOf": [
          {
            "required": ["build"],
            "not": {"required": ["image"]}
          },
          {
            "required": ["image"],
            "not": {"anyOf": [
              {"required": ["build"]},
              {"required": ["dockerfile"]}
            ]}
          }
        ]
      }
    }
  }
}
`

var schemaV2 = `{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "id": "config_schema_v2.0.json",
  "type": "object",

  "properties": {
    "version": {
      "type": "string"
    },

    "services": {
      "id": "#/properties/services",
      "type": "object",
      "patternProperties": {
        "^[a-zA-Z0-9._-]+$": {
          "$ref": "#/definitions/service"
        }
      },
      "additionalProperties": false
    },

    "networks": {
      "id": "#/properties/networks",
      "type": "object",
      "patternProperties": {
        "^[a-zA-Z0-9._-]+$": {
          "$ref": "#/definitions/network"
        }
      }
    },

    "volumes": {
      "id": "#/properties/volumes",
      "type": "object",
      "patternProperties": {
        "^[a-zA-Z0-9._-]+$": {
          "$ref": "#/definitions/volume"
        }
      },
      "additionalProperties": false
    }
  },

  "additionalProperties": false,

  "definitions": {

    "service": {
      "id": "#/definitions/service",
      "type": "object",

      "properties": {
        "build": {
          "oneOf": [
            {"type": "string"},
            {
              "type": "object",
              "properties": {
                "context": {"type": "string"},
                "dockerfile": {"type": "string"},
                "args": {"$ref": "#/definitions/list_or_dict"}
              },
              "additionalProperties": false
            }
          ]
        },
        "cap_add": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "cap_drop": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "certs": {"$ref": "#/definitions/list_of_strings"},
        "cgroup_parent": {"type": "string"},
        "command": {
          "oneOf": [
            {"type": "string"},
            {"type": "array", "items": {"type": "string"}}
          ]
        },
        "container_name": {"type": "string"},
        "cpu_shares": {"type": ["number", "string"]},
        "cpu_quota": {"type": ["number", "string"]},
        "cpuset": {"type": "string"},
        "default_cert": {"type": "string"},
        "depends_on": {"$ref": "#/definitions/list_of_strings"},
        "device_read_iops": {"$ref": "#/definitions/list_or_dict"},
        "devices": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "device_write_iops": {"$ref": "#/definitions/list_or_dict"},
        "disks": {"type": "array"},
        "dns": {"$ref": "#/definitions/string_or_list"},
        "dns_search": {"$ref": "#/definitions/string_or_list"},
        "domainname": {"type": "string"},
        "entrypoint": {
          "oneOf": [
            {"type": "string"},
            {"type": "array", "items": {"type": "string"}}
          ]
        },
        "env_file": {"$ref": "#/definitions/string_or_list"},
        "environment": {"$ref": "#/definitions/list_or_dict"},

        "expose": {
          "type": "array",
          "items": {
            "type": ["string", "number"],
            "format": "expose"
          },
          "uniqueItems": true
        },

        "extends": {
          "oneOf": [
            {
              "type": "string"
            },
            {
              "type": "object",

              "properties": {
                "service": {"type": "string"},
                "file": {"type": "string"}
              },
              "required": ["service"],
              "additionalProperties": false
            }
          ]
        },

        "external_ips": {"$ref": "#/definitions/list_of_strings"},
        "external_links": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "extra_hosts": {"$ref": "#/definitions/list_or_dict"},
        "health_check": {"type": "object"},
        "hostname": {"type": "string"},
        "image": {"type": "string"},
        "ipc": {"type": "string"},
        "labels": {"$ref": "#/definitions/list_or_dict"},
        "links": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "load_balancer_config": {"type": "object"},

        "logging": {
            "type": "object",

            "properties": {
                "driver": {"type": "string"},
                "options": {"type": "object"}
            },
            "additionalProperties": false
        },

        "mac_address": {"type": "string"},
        "memory": {"type": ["number", "string"]},
        "mem_limit": {"type": ["number", "string"]},
        "memswap_limit": {"type": ["number", "string"]},
        "metadata": {"type": "object"},
        "network_mode": {"type": "string"},

        "networks": {
          "oneOf": [
            {"$ref": "#/definitions/list_of_strings"},
            {
              "type": "object",
              "patternProperties": {
                "^[a-zA-Z0-9._-]+$": {
                  "oneOf": [
                    {
                      "type": "object",
                      "properties": {
                        "aliases": {"$ref": "#/definitions/list_of_strings"},
                        "ipv4_address": {"type": "string"},
                        "ipv6_address": {"type": "string"}
                      },
                      "additionalProperties": false
                    },
                    {"type": "null"}
                  ]
                }
              },
              "additionalProperties": false
            }
          ]
        },
        "pid": {"type": ["string", "null"]},

        "ports": {
          "type": "array",
          "items": {
            "type": ["string", "number"],
            "format": "ports"
          },
          "uniqueItems": true
        },

        "privileged": {"type": "boolean"},
        "read_only": {"type": "boolean"},
        "restart": {"type": "string"},
        "retain_ip": {"type": "boolean"},
        "scale": {"type": "number"},
        "scale_policy": {"type": "object"},
        "security_opt": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "shm_size": {"type": ["number", "string"]},
        "stdin_open": {"type": "boolean"},
        "stop_signal": {"type": "string"},
        "tmpfs": {"$ref": "#/definitions/string_or_list"},
        "tty": {"type": "boolean"},
        "type": {"type": "string"},
        "update_strategy": {"type": "object"},
        "ulimits": {
          "type": "object",
          "patternProperties": {
            "^[a-z]+$": {
              "oneOf": [
                {"type": "integer"},
                {
                  "type":"object",
                  "properties": {
                    "hard": {"type": "integer"},
                    "soft": {"type": "integer"}
                  },
                  "required": ["soft", "hard"],
                  "additionalProperties": false
                }
              ]
            }
          }
        },
        "user": {"type": "string"},
        "userdata": {"type": "string"},
        "vcpu": {"type": ["number", "string"]},
        "volumes": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "volume_driver": {"type": "string"},
        "volumes_from": {"type": "array", "items": {"type": "string"}, "uniqueItems": true},
        "working_dir": {"type": "string"}
      },

      "dependencies": {
        "memswap_limit": ["mem_limit"]
      },
      "additionalProperties": false
    },

    "network": {
      "id": "#/definitions/network",
      "type": "object",
      "properties": {
        "driver": {"type": "string"},
        "driver_opts": {
          "type": "object",
          "patternProperties": {
            "^.+$": {"type": ["string", "number"]}
          }
        },
        "ipam": {
            "type": "object",
            "properties": {
                "driver": {"type": "string"},
                "config": {
                    "type": "array"
                }
            },
            "additionalProperties": false
        },
        "external": {
          "type": ["boolean", "object"],
          "properties": {
            "name": {"type": "string"}
          },
          "additionalProperties": false
        }
      },
      "additionalProperties": false
    },

    "volume": {
      "id": "#/definitions/volume",
      "type": ["object", "null"],
      "properties": {
        "driver": {"type": "string"},
        "driver_opts": {
          "type": "object",
          "patternProperties": {
            "^.+$": {"type": ["string", "number"]}
          }
        },
        "external": {
          "type": ["boolean", "object"],
          "properties": {
            "name": {"type": "string"}
          }
        },
        "additionalProperties": false
      },
      "additionalProperties": false
    },

    "string_or_list": {
      "oneOf": [
        {"type": "string"},
        {"$ref": "#/definitions/list_of_strings"}
      ]
    },

    "list_of_strings": {
      "type": "array",
      "items": {"type": "string"},
      "uniqueItems": true
    },

    "list_or_dict": {
      "oneOf": [
        {
          "type": "object",
          "patternProperties": {
            ".+": {
              "type": ["string", "number", "null", "boolean"]
            }
          },
          "additionalProperties": false
        },
        {"type": "array", "items": {"type": "string"}, "uniqueItems": true}
      ]
    },

    "constraints": {
      "service": {
        "id": "#/definitions/constraints/service",
        "anyOf": [
          {"required": ["build"]},
          {"required": ["image"]}
        ],
        "properties": {
          "build": {
            "required": ["context"]
          }
        }
      }
    }
  }
}
`
