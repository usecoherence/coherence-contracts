---
title: Author
weight: 3
eyebrow: authoring
budget: ≤ 1 page
badge: source of truth
badgeClass: teal
---

The spec and its criteria live in versioned JSONL. Not a database, not an API — a diff-friendly file in git.

## Model

<div class="model-line" aria-label="owner 1 to N spec 1 to N acceptance criterion">
  <span class="ent pk">owner</span>
  <span class="rel">1<br>──<br>N</span>
  <span class="ent pk">spec</span>
  <span class="rel">1<br>──<br>N</span>
  <span class="ent pk">acceptance_criterion</span>
</div>

{{% panel title="tables" tag="columns" %}}
| table | columns |
|---|---|
| `owner` | id · name |
| `spec` | id · slug · level · status · title · description · owner_id |
| `acceptance_criterion` | id · spec_id · slug · status · title · description |
{{% /panel %}}

{{% panel title="messages" tag="core.proto" %}}
```proto
message Spec {
  coherence.kernel.v1.RecordMeta meta = 1;
  string level    = 2;
  string owner_id = 3;
}

message AcceptanceCriterion {
  coherence.kernel.v1.RecordMeta meta = 1;
  string spec_id = 2;
}

message Owner {
  coherence.kernel.v1.RecordMeta meta = 1;
  string name = 2;
}
```
{{% /panel %}}

## Result

{{% panel title="contracts.jsonl" tag="one line = one record" %}}
```json
{"table":"spec","id":"spec-001","slug":"specs-as-code","level":"module","status":"active","title":"Specs as code","description":"...","owner_id":"owner-001"}
{"table":"acceptance_criterion","id":"ac-001","spec_id":"spec-001","slug":"ac-jsonl-roundtrip","status":"active","title":"JSONL roundtrip"}
{"table":"owner","id":"owner-001","name":"core"}
```
{{% /panel %}}

<ul class="rule-list">
  <li><span class="mk">1</span><span>Deterministic order: <code>(table, id)</code>. Diffs are readable, merges are predictable.</span></li>
  <li><span class="mk">2</span><span>Uniqueness: <code>id</code> globally, <code>slug</code> — within a table.</span></li>
  <li><span class="mk">3</span><span>References are checked on read: <code>owner_id</code> and <code>spec_id</code> must exist.</span></li>
  <li><span class="mk">4</span><span>Errors are a list of <code>{line, table, id, problem}</code> and exit code 1.</span></li>
</ul>

{{% callout %}}
One file is a deliberate start. As volume grows, splitting into `contracts/spec.jsonl` etc. is a mechanical operation; the schema does not change.
{{% /callout %}}
