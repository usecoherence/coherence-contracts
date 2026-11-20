---
title: Link
weight: 5
eyebrow: linking
budget: ≤ 1 страница
badge: generic join
badgeClass: accent
---

AC становится проверяемым критерием, когда указывает на код. Это ребро — строка таблицы `link`.

## Модель

<div class="model-line">
  <span class="ent pk">acceptance_criterion</span>
  <span class="rel">verified_by →</span>
  <span class="ent pk">code_file</span>
</div>

{{% panel title="link columns" tag="один shape на все связи" %}}
| field | example |
|---|---|
| `id` | `link-001` |
| `type_from` / `from_id` | `acceptance_criterion` / `ac-001` |
| `type_to` / `to_id` | `code_file` / `file-017` |
| `relation_kind` | `verified_by` |
| `note` | `"e2e: jsonl roundtrip"` |
| `confidence` | `1.0` |
{{% /panel %}}

{{% panel title="relation vocabulary" tag="extensible by data" %}}
| relation | from → to | behavior |
|---|---|---|
| `verified_by` | `ac → code_file` | `run_tests(files)` |
| `depends_on` | `spec → spec` | impact / ordering, не исполняется |
| `persisted_to` | `spec → confluence_page` | `upsert_page` |
{{% /panel %}}

## Messages

{{% panel title="coherence/kernel/v1/kernel.proto" tag="одна форма для всех рёбер" %}}
```proto
message Link {
  string id = 1;
  string type_from = 2;  // "acceptance_criterion"
  string from_id   = 3;
  string type_to   = 4;  // "code_file"
  string to_id     = 5;
  string relation_kind = 6; // "verified_by"
  string note = 7;
  double confidence = 8;    // 1.0 authored, <1 inferred
}
```
{{% /panel %}}

## Selection → run plan

Выбор полностью детерминирован: никакого LLM в момент прогона. Агент может предлагать links, но выполняются только записанные рёбра.

{{% panel title="правило" tag="verified_by" %}}
```
links
  .where(relation_kind == "verified_by")
  .where(source.status == "active")
  .join(acceptance_criterion on from_id)
  .join(code_file on to_id)
  .group_by(ac.id)
```
{{% /panel %}}

<div class="plan-out">
  <div class="sum">plan · <b>1 spec · 3 AC · 4 files</b></div>

| AC | files |
|---|---|
| `ac-001` | `e2e/spec.spec.ts` · `e2e/api.spec.ts` |
| `ac-002` | `tests/jsonl_test.go` |
| `ac-003` | `tests/report_test.go` |

</div>

{{% callout %}}
AC с `status=archived` или link с `confidence<1` не попадают в план автоматически. Их можно включить явным scope-фильтром.
{{% /callout %}}
