#!/usr/bin/env python3
"""Measure recall/precision of TelemetryQA rules against fixtures/demo-*.

Usage:
    semgrep --config rules/ fixtures/demo-go fixtures/demo-python --json > out/results.json
    python3 scripts/measure.py out/results.json fixtures/ground-truth.json
"""
import json
import sys

results_path, truth_path = sys.argv[1], sys.argv[2]

with open(results_path) as f:
    findings = json.load(f)["results"]
with open(truth_path) as f:
    truth = json.load(f)["defects"]

# Normalize semgrep paths (may be relative or absolute) down to the
# "demo-go/..." / "demo-python/..." form used in ground-truth.json.
def normalize(path):
    for marker in ("demo-go", "demo-python"):
        idx = path.find(marker)
        if idx != -1:
            return path[idx:]
    return path

found = {(normalize(r["path"]), r["start"]["line"]) for r in findings}
expected = {(d["file"], d["line"]) for d in truth}

tp = found & expected
fn = expected - found
fp = found - expected

print(f"defects planted : {len(expected)}")
print(f"true positives   : {len(tp)}")
print(f"false negatives  : {len(fn)}")
print(f"false positives  : {len(fp)}")

if expected:
    print(f"recall           : {len(tp) / len(expected) * 100:.1f}%")
if found:
    print(f"precision        : {len(tp) / len(found) * 100:.1f}%")

if fn:
    print("\n[missed]")
    for path, line in sorted(fn):
        print(f"  {path}:{line}")

if fp:
    print("\n[false positives]")
    for path, line in sorted(fp):
        print(f"  {path}:{line}")