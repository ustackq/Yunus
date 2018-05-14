# Yunus

Yunus is a fully opensoure product, which built on pure-upstream Kubernetes but has an opinion on the best way to install and manager a Kubernetes cluster. This project not only helps you install a Kubernetes cluster using kubeadm which has supported multi-host automation install  ,but It also allows users to manage applications running in the cluster and troubleshoot them, as well as manage the cluster itself.

Goals of the project:

- Install Kubernetes clusters automation
- Secure by default (uses TLS, RBAC by default, OIDC AuthN, etcd)
- Automatable install process for scripts and CI/CD
- Run on such OS: Centos, Ubuntu
- Customizable and modular: Change DNS providers, security settings, authentication providers, Prometheus tool kits,and ELK log kits.
- Highly Available by default: Deploy all Kubernetes components HA.

## Getting Started

**To use a tested release** on a supported platform, follow the links below.

**To hack or modify** the templates or add a new platform, use the scripts in this repo to boot and tear down clusters.

### Architecture overview
![](dcos/architecture.png)
See the architecture below:



### Install

buger/jsonparser

## Documentation

Dashboard documentation can be found on [Wiki](https://github.com/kubernetes/dashboard/wiki) pages, it includes:

* Common: Entry-level overview

* Install Guide: [Installation](https://github.com/ustackq/yunus/docs/Installation), [Istaller Dashboard](
https://github.com/ustackq/yunus/docs/Accessing-dashboard) and more for users

* Manager Guide: [Management](https://github.com/ustackq/yunus/docs/management), [Management Dashboard](
https://github.com/ustackq/yunus/docs/Accessing-dashboard) and more for users

* Developer Guide: [Getting Started](https://github.com/ustackq/yunus/docs/Getting-started), [Dependency
Management](https://github.com/ustackq/yunus/docs/Dependency-management) and more for anyone interested in contributing

## License

The work done has been licensed under Apache License 2.0. The license file can be found [here](LICENSE). You can find
out more about the license at [www.apache.org/licenses/LICENSE-2.0](//www.apache.org/licenses/LICENSE-2.0).



