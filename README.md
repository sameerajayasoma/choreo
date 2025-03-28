# OpenChoreo
Internal Developer Platform

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![GitHub last commit](https://img.shields.io/github/last-commit/choreo-idp/choreo.svg)](https://github.com/choreo-idp/choreo/commits/main)
[![Go Report Card](https://goreportcard.com/badge/github.com/choreo-idp/choreo)](https://goreportcard.com/report/github.com/choreo-idp/choreo)
[![GitHub issues](https://img.shields.io/github/issues/choreo-idp/choreo.svg)](https://github.com/choreo-idp/choreo/issues)
[![Twitter Follow](https://img.shields.io/twitter/follow/Choreo?style=social)](https://twitter.com/ChoreoDev)

OpenChoreo is an open-source internal developer platform (IDP) designed for both platform engineers and application developers.
- For platform engineers, it provides a complete yet configurable platform that integrates with CI/CD systems, security tools, and cloud infrastructure. It supports enforcement of organizational policies and reduces the need to manage individual toolchains.
- For developers, it offers a set of abstractions that simplify working with Kubernetes and other infrastructure components. These abstractions allow developers to design, deploy, and manage applications without direct interaction with the underlying systems.

OpenChoreo builds upon years of experience at [WSO2](https://wso2.com/) in building [WSO2 Choreo](https://choreo.dev/) — an IDP as a Service. WSO2 Choreo is designed not only to automate software delivery workflows but also to support software engineering practices by enforcing architectural standards, encouraging service reuse, and providing integrated capabilities for API management and observability.

With OpenChoreo, we're bringing these best ideas to the open-source community in a Kubernetes-native, GitOps-friendly architecture. It is intended to be a complete self-hosted internal developer platform for organizations, while remaining highly customizable to fit existing tools, workflows, and infrastructure requirements.

## How does OpenChoreo work?
OpenChoreo provides a set of developer and platform-focused [abstractions](./docs/choreo-concepts.md) that simplify how applications are built, deployed, and managed on Kubernetes. Developers define [Projects](./docs/choreo-concepts.md#project), [Components](./docs/choreo-concepts.md#component), [Endpoints](./docs/choreo-concepts.md#environment), and [Connections](./docs/choreo-concepts.md#connection)—while platform engineers configure [DataPlanes](./docs/choreo-concepts.md#dataplane), [Environments](./docs/choreo-concepts.md#environment), [DeploymentPipelines](./docs/choreo-concepts.md#deploymentpipeline), and shared infrastructure.

The OpenChoreo control plane interprets these high-level abstractions and orchestrates their deployment across one or more Kubernetes clusters. Each Project is mapped to a Kubernetes namespace, and each Component is deployed as a workload resource such as a Deployment, Job, or StatefulSet. Endpoints define how components are exposed within or outside the platform, and Connections declare outbound dependencies to internal or external services. OpenChoreo uses this information to configure ingress and egress gateways, apply network policies, and ensure secure, policy-driven service communication across projects and environments.

![overview](./docs/images/choreo_overview.svg)


## Getting Started

The easiest way to try OpenChoreo is by following the **[Quick Start Guide](./docs/quick-start-guide.md)**. It walks you through setting up Choreo using a Dev Container, so you can start experimenting without affecting your local environment.

For a deeper understanding of OpenChoreo’s architecture, see **[Choreo Concepts](./docs/choreo-concepts.md)**.

Visit **[Installation Guide](./docs/install-guide.md)** to learn more about installation methods.

## Samples

Explore hands-on examples to help you configure and deploy applications using OpenChoreo.

- **[Configuring Choreo](./samples/configuring-choreo/)** – Set up organizations, environments, data planes, and deployment pipelines.
- **[Deploying Applications](./samples/deploying-applications/)** – Deploy services, web apps, and scheduled tasks.

Check out the **[Samples Directory](./samples/)** for more details.

## Join the Community & Contribute

We’d love for you to be part of OpenChoreo’s journey! 
Whether you’re fixing a bug, improving documentation, or suggesting new features, every contribution counts.

- **[Contributor Guide](./docs/contributors/README.md)** – Learn how to get started.
- **[Report an Issue](https://github.com/choreo-idp/choreo/issues)** – Help us improve Choreo.
- **[Join our Discord](https://discord.gg/asqDFC8suT)** – Be part of the community.

We’re excited to have you onboard!

## Roadmap
Explore the OpenChoreo roadmap, including completed milestones and upcoming plans, in our **[Roadmap]( https://github.com/orgs/choreo-idp/projects/1)**.

## License
OpenChoreo is licensed under Apache 2.0. See the **[LICENSE](./LICENSE)** file for full details.

