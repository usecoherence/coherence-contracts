---
title: Обзор
linkTitle: Обзор
weight: 1
eyebrow: сводка
budget: 1 страница
heading: "Один spec. Все AC. Код. Прогон. Отчёт."
---

Минимальный контракт, который покрывает вертикальный путь целиком и ничего сверх него. Таблицы — это JSONL-строки в git. Protobuf описывает строку, а не ORM. Всё, что не нужно этому пути, вынесено за границу слайса.

<div class="meta-row">
  <span class="chip accent">kernel ⊥ core</span>
  <span class="chip">~150 строк proto</span>
  <span class="chip">28 → 14 messages</span>
  <span class="chip teal">link — единственный generic join</span>
</div>

{{< pipeline >}}

{{< spine >}}

## В слайсе

<div class="cols">
  <div>
    <ul class="rule-list">
      <li><span class="mk">+</span><span>Один spec со всеми AC и владельцем.</span></li>
      <li><span class="mk">+</span><span>Ingest списка репозиториев и веток из конфига.</span></li>
      <li><span class="mk">+</span><span>Ручная или агентская связка AC —<code>verified_by</code>→ code_file.</span></li>
      <li><span class="mk">+</span><span>Детерминированный выбор: links → run plan.</span></li>
      <li><span class="mk">+</span><span>Один behavior: <code>run_tests(files)</code>; artifacts в <code>run/</code>.</span></li>
      <li><span class="mk">+</span><span>Отчёт по каждому RunItem; <code>--persist</code> решает, что публикуется.</span></li>
      <li><span class="mk">+</span><span>Пример расширения: Confluence-страница через <code>persisted_to</code>.</span></li>
    </ul>
  </div>
  <div>
    <h2>Отложено осознанно</h2>
    <ul class="rule-list">
      <li><span class="mk no">−</span><span><code>ObjectKindDefinition</code> и <code>Any</code>-envelope: не нужны, пока объекты третьих сторон не появляются в runtime.</span></li>
      <li><span class="mk no">−</span><span>Протоколы валидации: в слайсе валидация — это вывод тула и exit code.</span></li>
      <li><span class="mk no">−</span><span>Entity / Attribute / State / Transition / Invariant — вне пути.</span></li>
      <li><span class="mk no">−</span><span>Глубокий SCIP-ingest: таблицы зарезервированы, связи — на уровне файлов.</span></li>
      <li><span class="mk no">−</span><span>Много behavior-движков: один, но точка расширения видна.</span></li>
      <li><span class="mk no">−</span><span>Jira/GitHub-публикаторы: паттерн доказан одним Confluence-примером.</span></li>
    </ul>
  </div>
</div>

## Карта пакетов

{{% panel title="imports" tag="правило: kernel не импортирует core" %}}
```proto
package coherence.kernel.v1   // RecordMeta, Link
   ↑
package coherence.core.v1     // Spec, AcceptanceCriterion, Owner, Code*, Scip*, Run*
   ↑
package coherence.ext.*       // publisher-таблицы: confluence.v1.Page, ...
```
{{% /panel %}}

<p class="footnote">
  proposal · не финальный контракт · proto-файлы пока не переписаны<br>
  все размеры — бюджеты, а не обещания
</p>
