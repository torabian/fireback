import Form from "@rjsf/core";
import { type RJSFSchema } from "@rjsf/utils";
import validator from "@rjsf/validator-ajv8";

const schema: RJSFSchema = {
  title: "Project Configuration",
  description: "Configure a project, its team, deployment and integrations.",
  type: "object",

  required: ["name", "owner", "projectType", "team", "deployment"],

  properties: {
    name: {
      type: "string",
      title: "Project name",
      minLength: 3,
      maxLength: 100,
    },

    description: {
      type: "string",
      title: "Description",
      maxLength: 1000,
    },

    projectType: {
      type: "string",
      title: "Project type",
      enum: ["web", "mobile", "api", "data"],
      default: "web",
    },

    owner: {
      type: "object",
      title: "Project owner",
      required: ["name", "email"],
      properties: {
        name: {
          type: "string",
          title: "Name",
        },
        email: {
          type: "string",
          title: "Email",
          format: "email",
        },
        phone: {
          type: "string",
          title: "Phone",
        },
      },
    },

    team: {
      type: "array",
      title: "Team members",
      minItems: 1,
      items: {
        type: "object",
        required: ["name", "email", "role"],
        properties: {
          name: {
            type: "string",
            title: "Name",
          },
          email: {
            type: "string",
            title: "Email",
            format: "email",
          },
          role: {
            type: "string",
            title: "Role",
            enum: ["developer", "designer", "product_manager", "qa", "devops"],
          },
          allocation: {
            type: "integer",
            title: "Allocation %",
            minimum: 0,
            maximum: 100,
            default: 100,
          },
        },
      },
    },

    repository: {
      type: "object",
      title: "Repository",
      properties: {
        provider: {
          type: "string",
          title: "Provider",
          enum: ["github", "gitlab", "bitbucket", "other"],
          default: "github",
        },

        url: {
          type: "string",
          title: "Repository URL",
          format: "uri",
        },

        private: {
          type: "boolean",
          title: "Private repository",
          default: true,
        },

        branch: {
          type: "string",
          title: "Default branch",
          default: "main",
        },
      },
    },

    deployment: {
      type: "object",
      title: "Deployment",
      required: ["environment", "region"],
      properties: {
        environment: {
          type: "string",
          title: "Environment",
          enum: ["development", "staging", "production"],
          default: "development",
        },

        region: {
          type: "string",
          title: "Region",
          enum: [
            "eu-central-1",
            "eu-west-1",
            "us-east-1",
            "us-west-2",
            "ap-southeast-1",
          ],
        },

        autoDeploy: {
          type: "boolean",
          title: "Enable automatic deployment",
          default: true,
        },

        replicas: {
          type: "integer",
          title: "Replicas",
          minimum: 1,
          maximum: 100,
          default: 2,
        },

        resources: {
          type: "object",
          title: "Resources",
          properties: {
            cpu: {
              type: "number",
              title: "CPU (cores)",
              minimum: 0.1,
              maximum: 64,
              default: 1,
            },
            memory: {
              type: "number",
              title: "Memory (GB)",
              minimum: 0.25,
              maximum: 256,
              default: 2,
            },
          },
        },
      },
    },

    features: {
      type: "object",
      title: "Features",
      properties: {
        authentication: {
          type: "boolean",
          title: "Authentication",
          default: true,
        },

        payments: {
          type: "boolean",
          title: "Payments",
          default: false,
        },

        analytics: {
          type: "boolean",
          title: "Analytics",
          default: true,
        },

        notifications: {
          type: "boolean",
          title: "Notifications",
          default: false,
        },
      },
    },

    database: {
      type: "object",
      title: "Database",
      required: ["engine"],
      properties: {
        engine: {
          type: "string",
          title: "Database engine",
          enum: ["postgresql", "mysql", "mongodb", "none"],
          default: "postgresql",
        },

        host: {
          type: "string",
          title: "Host",
        },

        port: {
          type: "integer",
          title: "Port",
          minimum: 1,
          maximum: 65535,
          default: 5432,
        },

        databaseName: {
          type: "string",
          title: "Database name",
        },

        ssl: {
          type: "boolean",
          title: "SSL",
          default: true,
        },
      },
    },

    integrations: {
      type: "array",
      title: "Integrations",
      items: {
        type: "object",
        required: ["type", "name"],
        properties: {
          type: {
            type: "string",
            title: "Integration",
            enum: [
              "slack",
              "stripe",
              "sentry",
              "datadog",
              "sendgrid",
              "custom",
            ],
          },

          name: {
            type: "string",
            title: "Name",
          },

          enabled: {
            type: "boolean",
            title: "Enabled",
            default: true,
          },

          config: {
            type: "object",
            title: "Configuration",
            properties: {
              endpoint: {
                type: "string",
                title: "Endpoint",
                format: "uri",
              },
              apiKey: {
                type: "string",
                title: "API key",
              },
            },
          },
        },
      },
    },

    security: {
      type: "object",
      title: "Security",
      properties: {
        requireMfa: {
          type: "boolean",
          title: "Require MFA",
          default: true,
        },

        allowedIpRanges: {
          type: "array",
          title: "Allowed IP ranges",
          items: {
            type: "string",
          },
        },

        sessionTimeout: {
          type: "integer",
          title: "Session timeout (minutes)",
          minimum: 5,
          maximum: 1440,
          default: 60,
        },
      },
    },

    tags: {
      type: "array",
      title: "Tags",
      uniqueItems: true,
      items: {
        type: "string",
      },
    },
  },
};

const log = (type) => console.log.bind(console, type);

const uiSchema: UiSchema = {
  description: {
    "ui:widget": "textarea",
    "ui:options": {
      rows: 5,
    },
  },

  owner: {
    "ui:order": ["name", "email", "phone"],
  },

  team: {
    items: {
      allocation: {
        "ui:widget": "range",
      },
    },
  },

  repository: {
    url: {
      "ui:placeholder": "https://github.com/company/project",
    },
  },

  deployment: {
    resources: {
      cpu: {
        "ui:widget": "updown",
      },
      memory: {
        "ui:widget": "updown",
      },
    },
  },

  database: {
    "ui:order": ["engine", "host", "port", "databaseName", "ssl"],
    host: {
      "ui:placeholder": "db.example.com",
    },
  },

  integrations: {
    items: {
      config: {
        apiKey: {
          "ui:widget": "password",
        },
      },
    },
  },

  security: {
    allowedIpRanges: {
      "ui:options": {
        orderable: false,
      },
    },
  },

  tags: {
    "ui:options": {
      orderable: false,
    },
  },
};

export const SampleForm = () => {
  return (
    <Form
      schema={schema}
      uiSchema={uiSchema}
      validator={validator}
      showErrorList="top"
      liveValidate
      noHtml5Validate
      onChange={log("changed")}
      onSubmit={log("submitted")}
      onError={log("errors")}
    />
  );
};
