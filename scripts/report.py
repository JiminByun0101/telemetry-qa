#!/usr/bin/env python3
"""Turn a Semgrep JSON scan into a PR comment and a CI job summary.

Usage:
    semgrep --config rules/ <targets> --json --output out/results.json
    python3 scripts/report.py out/results.json

Writes out/comment.md and appends the same content to the GitHub
Actions job summary when running in CI.

This script does not read ground-truth.json. It only sees what a real
user would see: the scan output for a codebase it knows nothing about.
"""
import json
import os
import sys
from collections import defaultdict

CATEGORY_NAMES = {
    "cardinality": "Cardinality",
    "pii": "PII logging",
    "silent-failure": "Silent failure",
}


def load_findings(path):
    with open(path) as f:
        return json.load(f).get("results", [])


def build_report(findings):
    by_severity = defaultdict(list)
    by_category = defaultdict(list)

    for r in findings:
        extra = r.get("extra", {})
        severity = extra.get("severity", "INFO")
        category = extra.get("metadata", {}).get("subcategory", "other")
        by_severity[severity].append(r)
        by_category[category].append(r)

    errors = by_severity.get("ERROR", [])
    warnings = by_severity.get("WARNING", [])

    lines = []

    if errors:
        lines.append(f"## Observability check failed — {len(errors)} blocking issue(s)")
        lines.append("")
        lines.append("These must be resolved before this branch can be merged.")
    else:
        lines.append("## Observability check passed")
        lines.append("")
        if warnings:
            lines.append(f"No blocking issues. {len(warnings)} warning(s) below are informational.")
        else:
            lines.append("No issues found.")
    lines.append("")

    if by_category:
        lines.append("| Category | Blocking | Warnings |")
        lines.append("|---|---|---|")
        for cat in sorted(by_category):
            items = by_category[cat]
            e = sum(1 for i in items if i.get("extra", {}).get("severity") == "ERROR")
            w = sum(1 for i in items if i.get("extra", {}).get("severity") == "WARNING")
            lines.append(f"| {CATEGORY_NAMES.get(cat, cat)} | {e} | {w} |")
        lines.append("")

    for severity, heading in (("ERROR", "Blocking"), ("WARNING", "Warnings")):
        items = by_severity.get(severity, [])
        if not items:
            continue
        lines.append(f"### {heading}")
        lines.append("")
        for r in items:
            path = r.get("path", "?")
            line = r.get("start", {}).get("line", "?")
            extra = r.get("extra", {})
            message = " ".join(extra.get("message", "").split())
            rule_id = r.get("check_id", "").split(".")[-1]
            lines.append(f"**`{path}:{line}`**")
            lines.append("")
            lines.append(message)
            lines.append("")
            lines.append(f"<sub>Rule: `{rule_id}`</sub>")
            lines.append("")

    return "\n".join(lines), len(errors), len(warnings)


def main():
    if len(sys.argv) < 2:
        print("usage: report.py <semgrep-results.json>", file=sys.stderr)
        sys.exit(2)

    findings = load_findings(sys.argv[1])
    markdown, n_errors, n_warnings = build_report(findings)

    os.makedirs("out", exist_ok=True)
    with open("out/comment.md", "w") as f:
        f.write(markdown)

    summary_path = os.environ.get("GITHUB_STEP_SUMMARY")
    if summary_path:
        with open(summary_path, "a") as f:
            f.write(markdown)

    print(f"{n_errors} blocking, {n_warnings} warnings")


if __name__ == "__main__":
    main()