---
title: Execute
weight: 6
eyebrow: execution
budget: ≤ 1 page
badge: behavior engine
badgeClass: teal
---

The plan becomes artifacts. A behavior is a pure function of (link kind, target kind) and environment context.

## Model

| link kind | target | context | behavior |
|---|---|---|---|
| `verified_by` | `code_file` | `local` / `ci` | `run_tests(files)` |

One behavior in the slice. The behavior registry is an extension point, not part of the contract.

## Messages

{{% panel title="core.proto" tag="run / item / artifact" %}}
```proto
message Run {
  string id = 1;
  string behavior = 2;  // run_tests
  string context  = 3;  // local | ci
  string status   = 4;  // ok | failed | partial
  int64 started_at_unix_s  = 5;
  int64 finished_at_unix_s = 6;
  repeated RunItem items = 10;
}

message RunItem {
  string id = 1;
  string link_id = 2;
  string acceptance_criterion_id = 3;
  string code_file_id = 4;
  string expected = 5;
  string actual   = 6;
  string reason   = 7;
  string status   = 8;  // ok | failed | skipped
  repeated Artifact artifacts = 10;
}

message Artifact {
  string id = 1;
  string kind = 2;  // log | screenshot | recording | trace
  string uri  = 3;  // run/{ts}-{behavior}/...
  int64  size_bytes = 4;
}
```
{{% /panel %}}

## Result

<div class="panel-head" style="border:1px solid var(--line-strong);border-bottom:0;background:var(--panel)"><span>artifact bundle</span><span class="tag">gitignored, inspectable</span></div>
<div class="path-tree"><span class="dir">run/</span>1763550000-run_tests/
  plan.json
  <span class="dir">link-001/</span>
    stdout.log
    stderr.log
    junit.xml
    screenshot.png      <span class="note"># playwright</span>
    trace.zip           <span class="note"># step-by-step repro</span>
  <span class="dir">link-002/</span>
    stdout.log
  report.json</div>

<ul class="rule-list">
  <li><span class="mk">→</span><span>Artifacts are always written while a run is in progress: you can open them by hand.</span></li>
  <li><span class="mk">→</span><span><code>--persist</code> controls publication only: report rows and external evidence copies.</span></li>
  <li><span class="mk">→</span><span>Exit code: <code>0</code> — all ok/skipped, <code>1</code> — any failed.</span></li>
  <li><span class="mk">→</span><span>Each RunItem references the original <code>link_id</code> — provenance is not lost.</span></li>
</ul>
