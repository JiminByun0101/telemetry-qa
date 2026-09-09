# TelemetryQA

Static analysis that catches observability configuration defects in CI, before deployment.

Three defect categories:
1. Silent Failure - instrumentation emits no signal when things go wrong
2. Cardinality Explorsion - high-cardinality metric labels cause time series blowup
3. PII Logging - personally identifiable information leaks into logs 

Target languages: Go, Python
Status: in development (private)