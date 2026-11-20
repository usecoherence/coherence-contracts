# coherence-contracts

Protobuf contracts for Coherence kernel + first-party core extension.

```bash
make check
make inventory
```

Repository shape:

```txt
proto/coherence/kernel/v1   # generic object registry, records, links, validation
proto/coherence/core/v1     # bundled core extension: spec, AC, behavior, domain model
proto/examples/jira/v1      # example third-party extension object kind
```

Design notes:

- [Vertical slice: spec → AC → code → run → report](design/vertical-slice.html) — proposed minimal contract for the end-to-end path. [Русская версия](design/vertical-slice.ru.html).

The kernel only knows `ObjectKindDefinition`, `ObjectRef`, `ObjectRecord`, `Link`, and validation/manifest contracts.
Core concepts like `Spec`, `AcceptanceCriterion`, and `Behavior` are just preinstalled extension payloads.
