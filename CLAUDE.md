# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build and Test

- **Build all**: `make build` (builds all binaries in `cmd/` to `_output/bin/`)
- **Build specific component**: `make cmd/<component_name>` (e.g., `make cmd/host`, `make cmd/region`)
- **Build Docker Image**: `make image` (requires `REGISTRY`, `VERSION`, `ARCH` env vars)
  - Example: `REGISTRY=myreg VERSION=dev ARCH=all make image`
- **Test all**: `make test` (runs unit tests for most packages)
- **Run single test**: `go test -v ./pkg/path/to/package -run TestName`
- **Lint**: `make vet`
- **Format**: `make fmt` (formats Go code using `gofmt`)
- **Mod tidy**: `make mod` (updates `go.mod` and `go.sum`)
- **Generate API code**: `make gen-model-api` (requires `model-api-gen` tool)
- **Generate Swagger**: `make gen-swagger` (requires `swagger`, `swagger-gen` tools)

## Architecture

Cloudpods (OneCloud) is a cloud-native unified multi-cloud management platform (CMP). It manages resources across public clouds (AWS, Azure, GCP, Alibaba, etc.) and private clouds (OpenStack, KVM, VMware).

### Directory Structure
- `cmd/`: Entry points for services and tools.
  - `climc`: Unified CLI tool for interacting with the platform.
  - `region`: Core controller service (region controller).
  - `host`: Host agent (hostman) running on compute nodes (KVM/baremetal).
  - `keystone`: Identity service (compatible with OpenStack Keystone).
  - `glance`: Image service (compatible with OpenStack Glance).
  - `apigateway`: API Gateway.
  - `scheduler`: Resource scheduler.
  - `monitor`: Monitoring service.
- `pkg/`: Core library code and business logic.
  - `apis/`: API definitions, models, and structs.
  - `mcclient/`: Client library for internal services and API calls.
  - `cloudprovider/`: Interfaces and implementations for different cloud providers.
  - `compute/`: Compute resource management logic (VMs, disks, etc.).
  - `hostman/`: Logic for the host agent.
  - `util/`: Shared utility functions.
- `build/`: Build scripts, packaging logic, and output directory (`_output/`).
- `locales/`: Localization files (generated with `y18n`).

### Key Concepts
- **Region**: The central control plane that manages cloud accounts and resources.
- **Host**: The agent that runs on hypervisors to manage local VMs and storage.
- **Cloudprovider**: The abstraction layer for interacting with external clouds (AWS, Azure, etc.).
- **Models**: Defines the resource schemas and business logic (often found in `pkg/<service>/models`).
- **Dispatcher**: Requests are routed via a Model Dispatcher to specific resource models using reflection.

## Development Environment

### Prerequisites
- **OS**: Linux recommended (CentOS 7 VM for Windows/Mac users).
- **Go Version**: 1.24 (Documentation mentions 1.18+, but repo uses 1.24).
- **Tools**: Docker (enable "experimental": true in daemon.json), Kubernetes cluster.
- **GOPATH**: Code must be in `$GOPATH/src/yunion.io/x/cloudpods`.

### Setup Tips
- **Local Run**: To run `region` locally, set `fetch_etcd_service_info_and_use_etcd_lock: false` in config.
- **Host Build**: Compiling `host` locally requires Ceph dev libraries (`libcephfs-devel`, `librbd-devel`, etc.).

## Coding Standards

### Naming Conventions
- **Files**: Use kebab-case (e.g., `dns_zones.go`).
- **Structs**:
  - Public: Start with `S` (e.g., `SDnsZone`).
  - Private: Start with `s` (e.g., `sDnsZone`).
- **Interfaces**:
  - Public: Start with `I` (e.g., `ICloudProvider`).
  - Private: Start with `i` (e.g., `iCloudProvider`).
- **Functions**: Verb-noun, CamelCase (e.g., `GetDetails`). Booleans start with `Is`, `Has`, `Can`, `Allow`.
- **Constants**: UPPER_SNAKE_CASE.
- **Imports**: Grouped into Std, 3rd-party, Internal (separated by blank lines).

### Style
- Use `gofmt` for formatting.
- Avoid exposing internal IDs in APIs; use names/keywords.
- Functions should prioritize returning `error` over nil/bools.

## Contribution Guidelines

### Git Commit Messages
Format: `<type>(<scope>): <subject>`

- **Types**:
  - `feat`: New feature
  - `fix`: Bug fix
  - `refactor`: Code change that neither fixes a bug nor adds a feature
  - `test`: Adding missing tests or correcting existing tests
  - `chore`: Changes to build process or auxiliary tools
- **Scope**: Affected component (e.g., `region`, `scheduler`, `compute`).
- **Subject**: Short description (< 50 chars).
- **Body**: Detailed description (optional, wrap at 72 chars).

### Adding a New API
1. **Model**: Add `GetDetails<Action>` method to the model struct in `pkg/<service>/models/`.
   - Example: `func (self *SMyRes) GetDetailsStatus(...)`
2. **Logic**: Implement business logic in the method.
3. **CLI**: Register the command in `cmd/climc/shell/<service>/` using `cmd.GetWithCustomShow("<action>", ...)`.
4. **Test**: Recompile `climc` and the service, then test the new command.
