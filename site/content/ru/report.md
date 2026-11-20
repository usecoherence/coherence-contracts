---
title: Report
weight: 7
eyebrow: reporting
budget: ≤ 1 страница
badge: human + agent signal
badgeClass: accent
---

Отчёт — это не дамп логов, а минимальный набор сигналов: какой AC сломан, где код, что ожидалось, как повторить.

## Модель

| field | meaning |
|---|---|
| `run_id` | какой прогон |
| `link_id` | какое ребро проверялось |
| `acceptance_criterion_id` | какой AC |
| `code_file_id` | какой файл исполнялся |
| `expected` / `actual` | ожидание и факт |
| `reason` | почему failed/skipped |
| `status` | `ok` · `failed` · `skipped` |

## Результат

{{% panel title="report.json" tag="строка на RunItem" %}}
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
  <li><span class="mk">1</span><span>Сломанный AC: текст и id. Самое важное — первым.</span></li>
  <li><span class="mk">2</span><span>Файл и диапазон строк; сниппет под спойлером.</span></li>
  <li><span class="mk">3</span><span>Контекст: соседние символы, другие файлы, соседние AC — свёрнуто.</span></li>
  <li><span class="mk">4</span><span>Воспроизведение: screenshot / recording / trace + команда re-run.</span></li>
</ol>

{{% callout type="amber" %}}
**Принцип:** отчёт должен читаться за 30 секунд, а раскрываться до полного debug-контекста только по запросу. Evidence пре-процессится под эти четыре уровня.
{{% /callout %}}

## Persist policy

Один и тот же прогон ведёт себя по-разному в зависимости от контекста и флага. Переключи:

{{< persist >}}

## Publish — composition

<div class="model-line">
  <span class="ent">spec</span>
  <span class="rel">persisted_to →</span>
  <span class="ent">confluence_page</span>
  <span class="arrow">⇒</span>
  <span style="color:var(--muted)">upsert_page(spec)</span>
</div>

Публикация не добавляет полей в `Spec`. Она добавляет таблицу расширения и ребро. Тот же паттерн — Jira, GitHub, HTML-сайт.
