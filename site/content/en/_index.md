---
title: Overview
linkTitle: Overview
weight: 1
eyebrow: summary
budget: 1 page
heading: "One spec. All ACs. Code. Run. Report."
---

The minimal contract that covers the vertical path end to end and nothing beyond it. Tables are JSONL rows in git. Protobuf describes a row, not an ORM. Everything this path does not need stays outside the slice boundary.

<div class="meta-row">
  <span class="chip accent">kernel ⊥ core</span>
  <span class="chip">~150 lines of proto</span>
  <span class="chip">28 → 14 messages</span>
  <span class="chip teal">link — the only generic join</span>
</div>

{{< pipeline >}}

{{< spine >}}

## In the slice

<div class="cols">
  <div>
    <ul class="rule-list">
      <li><span class="mk">+</span><span>One spec with all its ACs and an owner.</span></li>
      <li><span class="mk">+</span><span>Ingest a list of repositories and branches from config.</span></li>
      <li><span class="mk">+</span><span>Manual or agent link: AC —<code>verified_by</code>→ code_file.</span></li>
      <li><span class="mk">+</span><span>Deterministic selection: links → run plan.</span></li>
      <li><span class="mk">+</span><span>One behavior: <code>run_tests(files)</code>; artifacts in <code>run/</code>.</span></li>
      <li><span class="mk">+</span><span>Report per RunItem; <code>--persist</code> decides what gets published.</span></li>
      <li><span class="mk">+</span><span>Extension example: a Confluence page via <code>persisted_to</code>.</span></li>
    </ul>
  </div>
  <div>
    <h2>Deliberately deferred</h2>
    <ul class="rule-list">
      <li><span class="mk no">−</span><span><code>ObjectKindDefinition</code> and <code>Any</code>-envelope: not needed until third-party objects appear at runtime.</span></li>
      <li><span class="mk no">−</span><span>Validation protocols: in the slice, validation is tool output plus an exit code.</span></li>
      <li><span class="mk no">−</span><span>Entity / Attribute / State / Transition / Invariant — outside the path.</span></li>
      <li><span class="mk no">−</span><span>Deep SCIP ingest: tables reserved, links stay at file level.</span></li>
      <li><span class="mk no">−</span><span>Multiple behavior engines: one, with a visible extension point.</span></li>
      <li><span class="mk no">−</span><span>Jira/GitHub publishers: pattern proven by one Confluence example.</span></li>
    </ul>
  </div>
</div>

## Package map

{{% panel title="imports" tag="rule: kernel does not import core" %}}
```proto
package coherence.kernel.v1   // RecordMeta, Link
   ↑
package coherence.core.v1     // Spec, AcceptanceCriterion, Owner, Code*, Scip*, Run*
   ↑
package coherence.ext.*       // publisher tables: confluence.v1.Page, ...
```
{{% /panel %}}

<p class="footnote">
  proposal · not the final contract · proto files not rewritten yet<br>
  all sizes are budgets, not promises
</p>
