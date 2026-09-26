# Runphase Roadmap

LLM-first. Prove the release lifecycle on one workload type before adding breadth.

| Milestone | Scope |
|---|---|
| **V0** | Go API, Postgres, Temporal, `ModelDeployment` operator, OpenAI-compatible and demo runtimes, durable deploy → smoke → benchmark, rollback on failure, worker-crash recovery, read-only run graph |
| **V0.1** | vLLM and SGLang profiles, hardware profiles and cost per million tokens, benchmark provenance, `InferencePool` on kind, idle-resource teardown |
| **V0.2** | model versions and aliases, eval gates, `SafeReleaseWorkflow` with canary, online quality and cluster-health gates, approvals, release reports |
| **V0.3** | release on a real GPU cluster (llm-d target), cross-hardware equivalence evals, multi-LoRA releases |
| **V1 alpha** | editable runbook canvas, Helm chart, basic auth/RBAC, docs, CI, tagged release |
| **Later** | KServe and NVIDIA Dynamo targets, deprecation schedules, OIDC and audit, multi-cluster, residency-aware releases, vision and classic-ML profiles |

## Non-goals

A serving engine, request router, autoscaler, GPU allocator, model registry, training pipeline system, or agent builder.
