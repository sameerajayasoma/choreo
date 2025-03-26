# OpenChoreo Concepts

OpenChoreo defines a set of application and platform-level abstractions that simplify how teams design, deploy, and manage cloud-native applications and their environments. These abstractions provide a unified model for defining, deploying, and managing cloud-native applications and platform infrastructure. While some are more commonly used by developers and others by platform engineers, they are designed to work cohesively in a shared platform model.

---

## Platform Concepts
These abstractions are primarily used to define infrastructure, environments, and deployment flow.

### DataPlane
Represents a Kubernetes cluster managed by OpenChoreo. It maintains health status, capabilities, and is referenced by environments.

### Environment
Defines a target deployment stage (e.g., `dev`, `staging`, `prod`). Each environment is bound to a specific `DataPlane` and can enforce policies and overlays.

### DeploymentPipeline
Defines the ordered progression of deployments across multiple environments. It allows teams to enforce promotion policies and lifecycle consistency.

### Secret
Represents securely managed configuration values or credentials. Secrets can be environment-specific and support integrations with systems like Vault, GitHub, Bitbucket, and DockerHub.

---

## Application Concepts
These abstractions describe how applications are structured, how they communicate with each other, and how they are deployed across environments.

### Project
Represents a cloud-native application composed of multiple components. It acts as the unit of isolation and maps directly to a Kubernetes namespace.

### Component
Defines a deployable unit within a project, such as a web service, API, worker, or scheduled task. Each component is translated into Kubernetes workload resources like `Deployment`, `Job`, or `StatefulSet`.

### Endpoint
Describes a network-accessible interface exposed by a component. It defines routing, protocol, and visibility (`public`, `organization`, `project`), and maps to `Service` and `Ingress` resources.

### Connection
Specifies outbound dependencies for a component. A connection can target either an external service (southbound) or a service in another project (westbound). Connections are enforced via Cilium `NetworkPolicies` and optionally routed through egress gateways.

### Endpoint
Describes a network-accessible interface exposed by a component. It defines routing rules, protocols, and visibility scopes such as `public`, `organization`, or `project`. Endpoints are mapped to Kubernetes `Service` and `Ingress` resources and are configured through shared ingress gateways.

### Connection
Defines outbound service dependencies for a component. These can represent connections to services in other projects (westbound traffic) or to external systems (southbound traffic). Connections are translated into network access policies and optionally routed through egress gateways.

### DeploymentTrack
Manages the progression of a component across environments. It supports automatic promotion and build management based on deployment status.

### Build
Represents the transformation of source code into a deployable artifact. This includes build configurations and metadata tracked by the control plane.

### DeployableArtifact
A build output with environment-agnostic configuration, ready to be deployed to any environment.

### Deployment
Represents the actual rollout of a `DeployableArtifact` to a specific environment. Manages status, versioning, and rollout policies.

### DeploymentRevision
Tracks the history of a deployment, allowing rollback and inspection of prior deployment states.

---

These abstractions allow teams to focus on application logic and infrastructure governance while OpenChoreo handles the orchestration, security, and operational complexity required to run modern cloud-native applications across multiple environments and clusters.

