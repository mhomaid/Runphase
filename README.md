# Runphase

**A durable release manager for self-hosted models.**

Serving engines and routers are solved: vLLM, SGLang, llm-d, NVIDIA Dynamo, KServe. Safely shipping a new model version is not. Teams still benchmark, eval, canary, watch dashboards, and roll back by hand.

Runphase turns that into one durable, auditable workflow:

```text
benchmark → offline eval → deploy candidate → canary via InferencePool
→ watch latency, errors, quality, cluster health → promote or roll back
```

It runs the same on a laptop (MLX, llama.cpp, Ollama, vLLM endpoints) and on Kubernetes.

**License:** Apache-2.0  
**Stage:** pre-alpha (V0 in progress)

## How it works

- **Temporal** runs release workflows that survive crashes, wait for approvals, and roll back cleanly.
- **A Kubernetes operator** reconciles `ModelDeployment` resources (`runphase.dev`).
- **Gateway API Inference Extension** `InferencePool` backends carry canary traffic; Runphase shifts weights, the gateway routes.
- **Eval suites** are mandatory gates, not dashboards.
- **OpenTelemetry GenAI** metrics feed the gates.

Runphase does not serve inference, route requests, autoscale, or store model bytes. It orchestrates releases across the tools that do.

## Layout

```text
Runphase/
├── apps/web/                 # Next.js UI (bun)
├── cmd/
│   ├── apiserver/            # Go API
│   ├── worker/               # Temporal worker
│   ├── operator/             # Kubernetes operator
│   ├── cli/                  # Go CLI
│   └── demo-inference/       # tiny runtime for local demos
├── internal/                 # domain, store, temporal, traffic, evals, benchmark, ...
├── api/{openapi,crds}/
├── deploy/{compose,kind,helm,examples}/
├── docs/{architecture,adrs,runbooks,benchmarks,incidents}/
├── e2e/
└── scripts/
```

## Roadmap

See [ROADMAP.md](ROADMAP.md).

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md). Security issues: [SECURITY.md](SECURITY.md).
