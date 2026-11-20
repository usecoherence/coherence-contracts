---
title: Схема
weight: 2
eyebrow: protobuf
budget: ≤ 1 страница
---

Три пакета, четырнадцать сообщений. Строка таблицы — это message с общим `RecordMeta`; связь любых двух строк — один `Link`.

{{% panel title="coherence/kernel/v1/kernel.proto" tag="~30 строк" %}}
```proto
syntax = "proto3";
package coherence.kernel.v1;

// Общая идентичность и отображение для любой строки.
message RecordMeta {
  string id    = 1;  // стабильный, не переиспользуется
  string slug  = 2;  // человеческий handle
  string title = 3;
  string description = 4;
  string status = 5; // draft | active | archived
  int64 created_at_unix_s = 10;
  int64 updated_at_unix_s = 11;
  map<string, string> annotations = 20;
}

// Одно направленное семантическое ребро между строками.
message Link {
  string id = 1;
  string type_from = 2;  // "acceptance_criterion"
  string from_id   = 3;
  string type_to   = 4;  // "code_file"
  string to_id     = 5;
  string relation_kind = 6; // "verified_by"
  string note = 7;
  double confidence = 8;    // 1.0 authored, <1 inferred
  int64 created_at_unix_s = 10;
  int64 updated_at_unix_s = 11;
}
```
{{% /panel %}}

{{% panel title="coherence/core/v1/core.proto" tag="~90 строк" %}}
```proto
syntax = "proto3";
package coherence.core.v1;

import "coherence/kernel/v1/kernel.proto";

message Owner {
  coherence.kernel.v1.RecordMeta meta = 1;
  string name = 2;
}

message Spec {
  coherence.kernel.v1.RecordMeta meta = 1;
  string level    = 2; // product | system | module | ...
  string owner_id = 3;
}

message AcceptanceCriterion {
  coherence.kernel.v1.RecordMeta meta = 1;
  string spec_id = 2;
}

message CodeRepo {
  coherence.kernel.v1.RecordMeta meta = 1;
  string path = 2;
}

message RepoBranch {
  coherence.kernel.v1.RecordMeta meta = 1;
  string code_repo_id = 2;
  string name         = 3;
  string head_sha     = 4;
}

message CodeFile {
  coherence.kernel.v1.RecordMeta meta = 1;
  string repo_branch_id = 2;
  string file_path      = 3;
  string sha            = 4;
}

// Зарезервировано под SCIP; в слайсе не заполняется.
message ScipSymbol {
  coherence.kernel.v1.RecordMeta meta = 1;
  string code_file_id = 2;
  string symbol       = 3;
  string kind         = 4;
  int32  start_line   = 5;
  int32  end_line     = 6;
}

message ScipLink {
  coherence.kernel.v1.RecordMeta meta = 1;
  string from_symbol_id = 2;
  string to_symbol_id   = 3;
  string relation_kind  = 4;
}

// Run — это факт исполнения, а не объект каталога.
message Run {
  string id = 1;
  string behavior = 2; // "run_tests"
  string context  = 3; // "local" | "ci"
  string status   = 4; // ok | failed | partial
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
  string status   = 8; // ok | failed | skipped
  repeated Artifact artifacts = 10;
}

message Artifact {
  string id   = 1;
  string kind = 2; // log | screenshot | recording | trace
  string uri  = 3; // run/{ts}-{behavior}/...
  int64  size_bytes = 4;
  map<string, string> annotations = 5;
}
```
{{% /panel %}}

{{% panel title="coherence/ext/confluence/v1/page.proto" tag="~12 строк" %}}
```proto
syntax = "proto3";
package coherence.ext.confluence.v1;

import "coherence/kernel/v1/kernel.proto";

// Composition over inheritance: core не знает об этой таблице.
message Page {
  coherence.kernel.v1.RecordMeta meta = 1;
  string space_key   = 2;
  string external_id = 3;
  string url         = 4;
}
```
{{% /panel %}}

## Решения, которые делают схему маленькой

| Решение | Почему |
|---|---|
| Один `RecordMeta` на все таблицы | id / slug / title / description / status / время описываются один раз, а не в каждом message. |
| Один `Link` вместо edge-таблиц | `verified_by`, `depends_on`, `persisted_to` — строки одного вида, без отдельных message на каждую связь. |
| Канонические имена — singular | `type_from` / `type_to` совпадают с именами таблиц и message: `spec`, `acceptance_criterion`, `code_file`. Один словарь вместо двух. |
| Никакого `Any` и реестров | JSONL уже типизирован именем таблицы. Реестр kinds вернётся, когда появятся runtime-объекты третьих сторон. |
| Расширения — отдельные пакеты | Confluence/Jira добавляют таблицы и линкуются к core через `Link`. Core не меняется. |

{{% callout type="teal" %}}
Граница пакетов проверяема: `coherence.kernel.v1` не импортирует `coherence.core.v1`. Это правило уже зафиксировано в AGENTS.md и покрывается buf-плагином.
{{% /callout %}}
