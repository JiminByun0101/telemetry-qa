# Defect Catalog

Detection scope for TelemetryQA v1: 3 defect categories, 2 languages (Go, Python).

---

## C001 — High-cardinality metric label (name)

| | |
|---|---|
| Language | Go, Python |
| Severity | ERROR |
| Confidence | HIGH |

**Blocked label vocabulary**
```
user_id, userid, uid, customer_id, account_id
request_id, req_id, trace_id, span_id, correlation_id
session_id, sid
email, phone, ip, ip_address, remote_addr
uuid, guid, hash
timestamp, ts, time, created_at
url, path, full_path, query, referer
```

**Allowed label vocabulary**
```
method, status, status_code, code
route, endpoint, handler, operation
service, namespace, pod, instance, job
result, error_type, cache_hit
```

**Rule of thumb**: a label is safe if the set of possible values is finite and determined by code (e.g. `route` — bounded by registered routes). A label is unsafe if its values are determined by request data (e.g. `path` — varies per request).

**Operational impact**: label cardinality multiplies time series count. One high-cardinality label can turn a bounded metric into millions of series, causing Prometheus TSDB memory growth and OOM.

**Detection difficulty**: Low. Label names are typically string literals in metric constructors, directly visible to AST-based matching.

---

## C002 — Dynamic label value (timestamp / UUID)

| | |
|---|---|
| Language | Go, Python |
| Severity | ERROR |
| Confidence | HIGH |

Detects label values constructed from `time.Now()`, `uuid.New()`, etc. at `WithLabelValues()` / `.labels()` call sites.

**Operational impact**: same as C001, but at the call site instead of declaration site — every observation creates a new series.

**Detection difficulty**: Medium. Only catches inline expressions; a value assigned to a variable first is missed (would need dataflow / taint analysis).

---

## C003 — Raw URL path as attribute

| | |
|---|---|
| Language | Go, Python |
| Severity | WARNING |
| Confidence | MEDIUM |

Detects raw request path (`r.URL.Path`, `r.RequestURI`) used as an attribute value instead of a normalized route pattern (`/users/{id}`).

**Operational impact**: same as C001, but the impact depends on whether the attribute is consumed as a metric dimension or a trace span attribute — the latter is normal per OTel semantic conventions. Static analysis alone cannot distinguish these without following the value to its sink, so this is kept at WARNING rather than ERROR.

**Detection difficulty**: Medium.

---

## P001 — Sensitive field name in structured log

| | |
|---|---|
| Language | Go, Python |
| Severity | ERROR (credentials/PII), WARNING (general PII) |
| Confidence | HIGH |

**Blocked field names — ERROR tier (credentials, unique identifiers)**
```
password, passwd, pwd, secret, api_key, apikey
token, access_token, refresh_token, id_token
authorization, auth_header, bearer
private_key, secret_key, credential, credentials
ssn, resident_number, rrn, jumin
card_number, cardno, pan, cvv, cvc
account_number, passport_number, driver_license
```

**Blocked field names — WARNING tier (general PII)**
```
email, phone, mobile, address, birthdate, dob
full_name, real_name, ip_address
```

**Allowed field names (must not match)**
```
password_hash, hashed_password, password_policy
has_password, password_length, password_changed_at
token_type, token_expiry, token_count, csrf_token
email_verified, email_domain, email_hash
account_id, account_type, account_status
card_brand, card_last4, card_type
auth_method, auth_provider, is_authenticated
```

**Operational impact**: logs are shipped to a central collector, viewed by multiple people, and retained long-term. Once a sensitive value lands in a log, it is effectively out of the team's control.

**Detection difficulty**: Low for key-name detection; requires negative lookahead in the regex to avoid matching safe compound names like `password_hash`.

---

## P002 — Full object/request dump in log

| | |
|---|---|
| Language | Go, Python |
| Severity | WARNING |
| Confidence | MEDIUM |

| Language | Detected pattern |
|---|---|
| Go | `%+v`, `%#v` format verbs |
| Python | `.__dict__`, `vars()`, `locals()`, `.headers`, `.body` |

**Operational impact**: what gets logged is not controlled by the logging line itself but by whatever fields exist on the struct/object at call time. Adding a new field to a struct can silently start leaking it through logging code that was never touched.

**Detection difficulty**: Medium.

---

## P003 — Exception context dump

| | |
|---|---|
| Language | Python |
| Severity | WARNING |
| Confidence | MEDIUM |

Detects `locals()` or f-string interpolation of `locals()` inside an `except` block.

**Operational impact**: an exception handler that dumps the entire local scope may include request bodies or credentials that happened to be in scope at the point of failure.

**Detection difficulty**: Medium.

---

## P004 — Korean PII literal (resident registration number / card number)

| | |
|---|---|
| Language | Go, Python (text-level, language-agnostic pattern) |
| Severity | ERROR (RRN), WARNING (card number) |
| Confidence | MEDIUM (RRN), LOW (card number) |

Regex-based detection of literal values shaped like a Korean resident registration number or a 16-digit card number, anywhere in the source (comments, string literals, test fixtures).

**Known limitation**: no checksum validation (RRN weighted checksum, card Luhn algorithm). This is a precision/recall tradeoff — narrowing the regex further would reduce both false positives and true positive coverage. Listed as a v2 roadmap item (post-processing checksum validator).

---

## S001 — Span not ended (Go)

| | |
|---|---|
| Language | Go |
| Severity | ERROR |
| Confidence | MEDIUM |

Detects `tracer.Start()` without a corresponding `defer span.End()` or `span.End()` later in the same block.

**Operational impact**: an unterminated span is never flushed to the exporter — the trace for that request never appears in the backend, and the span object leaks in memory.

**Known limitations**:
- Absence-of-code detection is scoped to same-block matching; a `span.End()` inside a different block (e.g. nested `if`) may not be recognized, depending on Semgrep's `...` scoping behavior — needs fixture verification.
- Multi-return-path functions where `End()` is only called on the success path are NOT reliably caught (would need control-flow analysis) — tracked as `todoruleid` in fixtures, not claimed as covered.

---

## S002 — Error not recorded on owned span

| | |
|---|---|
| Language | Go |
| Severity | ERROR (pending fixture-measured precision ≥ 90%, else WARNING or excluded from v1) |
| Confidence | MEDIUM |

Detects `if err != nil` blocks, inside a function that owns a span (started via `tracer.Start()` in that function), where neither `span.RecordError(err)` nor `span.SetStatus(codes.Error, ...)` is called.

**Operational impact**: `RecordError` alone leaves the span status as `Unset` — the trace UI still shows success/green, error rate metrics stay at 0%, and no alert fires, even though the request failed. This is the canonical "silent failure" case for this product.

**Design decision — precision guardrail**: matching is scoped to the function that owns the span, not every `if err != nil` in the codebase. Error propagation in a function without its own span is normal Go idiom and must not be flagged. This single constraint is what keeps false positives in the low tens instead of the thousands.

**Known limitations**:
- Relies on variable name heuristic (`err` substring) rather than type information — cannot distinguish a real `error` from another nilable type named similarly.
- May false-positive on the `defer func() { if err != nil { ... } }()` pattern (deferred recording via named return value), which is actually a best-practice pattern — must be verified against a fixture and either excluded via `pattern-not`, downgraded to WARNING, or cut from v1 if precision doesn't clear 90%.

---

## S003 — Telemetry init error ignored

| | |
|---|---|
| Language | Go |
| Severity | ERROR |
| Confidence | HIGH |

Detects `_ := someOtelInitFn(...)` where the error return of an OTel exporter/provider/resource constructor is discarded via blank identifier.

**Operational impact**: the service starts normally even if telemetry initialization failed — no traces or metrics are ever emitted, and nothing signals this failure.

**Detection difficulty**: Low. Presence of `_` in place of the error return is directly visible to AST matching — this is presence detection, not absence detection, unlike the rest of the S category.

---

## S004 — Manual span not ended (Python)

| | |
|---|---|
| Language | Python |
| Severity | ERROR |
| Confidence | MEDIUM |

Detects `tracer.start_span()` (manual API) without a corresponding `.end()` call or being wrapped in a `with` statement.

**Operational impact**: same as S001, for the Python manual span API. Recommended fix is switching to `start_as_current_span()` context manager, which handles this automatically.

**Detection difficulty**: Medium. Same block-scoping caveat as S001.

---

## Out of scope for v1

- Cross-function (interprocedural) tracking of wrapped constructors
- Control-flow analysis of multi-return-path functions
- Taint tracking of values through variables
- Checksum validation for card numbers / resident registration numbers

---

## Measurement history

```
(append one line per day, after running scripts/measure.py)

YYYY-MM-DD: <what was added> - defects N / TP n / FP n (recall X%, precision Y%)
```

- 2026-09-07 C001 only: defects 2 / TP 1 / FP 0 (recall 50.0%, precision 100.0%)
- 2026-09-07 C001+C002: defects 4 / TP 3 / FP 0 (recall 75.0%, precision 100.0%)