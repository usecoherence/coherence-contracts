---
title: Author
weight: 3
eyebrow: authoring
budget: ≤ 1 страница
badge: source of truth
badgeClass: teal
---

Спека и её критерии живут в версионированном JSONL. Это не база и не API — это diff-friendly файл в git.

## Модель

<div class="model-line" aria-label="owner 1 к N spec 1 к N acceptance criterion">
  <span class="ent pk">owner</span>
  <span class="rel">1<br>──<br>N</span>
  <span class="ent pk">spec</span>
  <span class="rel">1<br>──<br>N</span>
  <span class="ent pk">acceptance_criterion</span>
</div>

{{% panel title="таблицы" tag="columns" %}}
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

## Результат

{{% panel title="contracts.jsonl" tag="одна строка = одна запись" %}}
```json
{"table":"spec","id":"spec-001","slug":"specs-as-code","level":"module","status":"active","title":"Specs as code","description":"...","owner_id":"owner-001"}
{"table":"acceptance_criterion","id":"ac-001","spec_id":"spec-001","slug":"ac-jsonl-roundtrip","status":"active","title":"JSONL roundtrip"}
{"table":"owner","id":"owner-001","name":"core"}
```
{{% /panel %}}

<ul class="rule-list">
  <li><span class="mk">1</span><span>Детерминированный порядок: <code>(table, id)</code>. Diff читается, merge предсказуем.</span></li>
  <li><span class="mk">2</span><span>Уникальность: <code>id</code> глобально, <code>slug</code> — в пределах таблицы.</span></li>
  <li><span class="mk">3</span><span>Ссылки проверяются при чтении: <code>owner_id</code> и <code>spec_id</code> должны существовать.</span></li>
  <li><span class="mk">4</span><span>Ошибки — список <code>{line, table, id, problem}</code> и exit code 1.</span></li>
</ul>

{{% callout %}}
Один файл — осознанный старт. Когда объём вырастет, разбиение на `contracts/spec.jsonl` и т.д. — механическая операция, схема не меняется.
{{% /callout %}}
