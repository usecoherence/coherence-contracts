---
title: Ingest
weight: 4
eyebrow: ingestion
budget: ≤ 1 страница
badge: idempotent
---

Структура кода — это данные. Конфиг говорит, что читать; результат — строки `code_*` в том же JSONL.

## Модель

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

Пунктир — зарезервированная часть, не блокирует слайс. Связи AC↔код в нём работают на уровне файлов.

## Конфиг

{{% panel title="coherence.ingest.yaml" tag="tool config, не proto" %}}
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

## Результат

| Событие | Действие |
|---|---|
| Файл появился | insert `code_file` со статусом `active` |
| sha изменился | update `sha`, `updated_at_unix_s` |
| Файл исчез | `status=archived`, строку не удаляем |
| Повторный ingest | ноль новых строк, если ничего не менялось |

<div class="plan-out">
  <div class="sum">ingest report · <b>main @ 4f2a1c</b></div>
<pre>code_file  +128  ~3  -2 (archived)
elapsed    0.41s
exit       0</pre>
</div>

Ключ идентичности файла: `(repo.path, branch.name, file_path)`. Id — производный, поэтому повторный ingest не создаёт дублей.
