# Runphase

Local-first, Kubernetes-ready control plane for deploying, benchmarking, rolling out, and operating heterogeneous AI/ML workloads through durable typed runbooks.

Runphase integrates specialized systems instead of replacing them: Temporal for durable operational workflows, Kubernetes controllers for desired-state reconciliation, and existing inference runtimes for the data plane.

**License:** Apache-2.0  
**Stage:** V0 skeleton

Private curriculum and labs live in the sibling workspace `AI-Infra`. This repository is the public product.

## Layout

```text
Runphase/
├── apps/web/                 # Next.js
├── cmd/
│   ├── apiserver/            # Go API
│   ├── worker/               # Temporal worker
│   ├── operator/             # Kubernetes controller
│   ├── cli/
│   └── demo-inference/
├── internal/
├── api/{openapi,crds}/
├── deploy/{compose,kind,helm,examples}/
├── docs/{architecture,adrs,runbooks,benchmarks,incidents}/
├── e2e/
└── scripts/
```
