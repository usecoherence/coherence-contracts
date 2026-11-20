---
title: Report
weight: 7
eyebrow: reporting
budget: ≤ 1 page
badge: human + agent signal
badgeClass: accent
---

A report is not a log dump but a minimal set of signals: which AC broke, where the code is, what was expected, how to reproduce.

## Model

| field | meaning |
|---|---|
| `run_id` | which run |
| `link_id` | which edge was checked |
| `acceptance_criterion_id` | which AC |
| `code_file_id` | which file was executed |
| `expected` / `actual` | expected and actual |
| `reason` | why failed/skipped |
| `status` | `ok` · `failed` · `skipped` |

## Result

{{% panel title="report.json" tag="one row per RunItem" %}}
```json
{
  "run_id": "run-1763550000",
  "link_id": "link-001",
  "acceptance_criterion_id": "ac-001",
  "code_file_id": "file-017",
  "expected": "JSONL roundtrip preserves all ACs",
  "actual": "1 AC lost its spec_id after re-import",
  "reason": "assertion failed at e2e/spec.spec.ts:42",
  "status": "failed"
}
```
{{% /panel %}}

## Progressive disclosure

<ol class="rule-list">
  <li><span class="mk">1</span><span>Broken AC: text and id. Most important first.</span></li>
  <li><span class="mk">2</span><span>File and line range; snippet behind a spoiler.</span></li>
  <li><span class="mk">3</span><span>Context: neighboring symbols, other files, neighboring ACs — collapsed.</span></li>
  <li><span class="mk">4</span><span>Reproduction: screenshot / recording / trace + re-run command.</span></li>
</ol>

{{% callout type="amber" %}}
**Principle:** the report should read in 30 seconds and expand to full debug context only on demand. Evidence is pre-processed into these four levels.
{{% /callout %}}

## Persist policy

The same run behaves differently depending on context and flag. Toggle:

{{< persist >}}

## Publish — composition

<div class="model-line">
  <span class="ent">spec</span>
  <span class="rel">persisted_to →</span>
  <span class="ent">confluence_page</span>
  <span class="arrow">⇒</span>
  <span style="color:var(--muted)">upsert_page(spec)</span>
</div>

Publishing does not add fields to `Spec`. It adds an extension table and an edge. Same pattern: Jira, GitHub, HTML site.
