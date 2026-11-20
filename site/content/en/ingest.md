---
title: Ingest
weight: 4
eyebrow: ingestion
budget: ≤ 1 page
badge: idempotent
---

Code structure is data. Config says what to read; the result is `code_*` rows in the same JSONL.

## Model

<div class="model-line" aria-label="code_repo to repo_branch to code_file, optionally to scip_symbol and scip_link">
  <span class="ent">code_repo</span>
  <span class="arrow">→</span>
  <span class="ent">repo_branch</span>
  <span class="arrow">→</span>
  <span class="ent pk">code_file</span>
  <span class="arrow">⇢</span>
  <span class="ent">scip_symbol</span>
  <span class="arrow">⇢</span>
  <span class="ent">scip_link</span>
</div>

Dotted lines are reserved, they do not block the slice. AC↔code links work at file level there.

## Config

{{% panel title="coherence.ingest.yaml" tag="tool config, not proto" %}}
```yaml
repos:
  - path: .
    branches: [main]
    include:
      - "src/**"
      - "tests/**"
      - "e2e/**"
```
{{% /panel %}}

## Messages

{{% panel title="core.proto" tag="code ingestion" %}}
```proto
message CodeRepo {
  coherence.kernel.v1.RecordMeta meta = 1;
  string path = 2;
}
message RepoBranch {
  coherence.kernel.v1.RecordMeta meta = 1;
  string code_repo_id = 2;
  string name     = 3;
  string head_sha = 4;
}
message CodeFile {
  coherence.kernel.v1.RecordMeta meta = 1;
  string repo_branch_id = 2;
  string file_path = 3;
  string sha       = 4;
}
```
{{% /panel %}}

## Result

| Event | Action |
|---|---|
| File appeared | insert `code_file` with status `active` |
| sha changed | update `sha`, `updated_at_unix_s` |
| File disappeared | `status=archived`, keep the row |
| Repeated ingest | zero new rows if nothing changed |

<div class="plan-out">
  <div class="sum">ingest report · <b>main @ 4f2a1c</b></div>
<pre>code_file  +128  ~3  -2 (archived)
elapsed    0.41s
exit       0</pre>
</div>

File identity key: `(repo.path, branch.name, file_path)`. Id is derived, so repeated ingest creates no duplicates.
