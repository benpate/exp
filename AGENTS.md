# exp — Notes for AI Agents

Database-agnostic query expressions; see [README.md](README.md) for the type overview. This package describes only the SHAPE of a logical expression — every comparison semantic lives in an adapter (data-mongo's `ExpressionToBSON`, data-mock's `MatcherFunc`) or in whatever `MatcherFunc` a caller hands to `Match`.

- **Empty And and empty Or match OPPOSITE ways.** An empty `AndExpression` is vacuously true — that is what makes `All()` match everything — while an empty `OrExpression` matches nothing, mirroring SQL's `x IN ()`. Building an `Or(...)` dynamically and ending up with zero clauses yields match-NOTHING, not match-everything.

- **An empty-string value is a legitimate predicate value.** `Equal("field", "")` OR'd with a real value is the only way to express "unset or X", so nothing here rejects, drops, or normalizes empty values — or validates `Field`/`Operator`/`Value` at all. Guards against downstream hazards like `CONTAINS ""` (which data-mongo compiles to an empty `$regex` matching every record) belong in the producer — see the exp-builder AGENTS.md — never in this package.

- **A `nil` Expression is absorbed at every combining site, and the nil check stays inline.** Every `And`/`Or` method returns its receiver when handed nil, so a nil can never enter the tree — where it would panic much later, inside `Match()` or `Fields()`, far from the call that produced it. The check is repeated at each site instead of factored into a helper because nilaway cannot see a nil check through a function call.

- **`EmptyExpression` is absorbing glue.** Combining anything with `Empty()` returns the other side unchanged, and Empty's own `And*`/`Or*` shortcuts return a bare `Predicate`. It exists so "no constraint" is a value rather than a nil.

- **The three-index appends are important.** `append(e[:len(e):len(e)], ...)` in [and.go](and.go) and [or.go](or.go) caps the receiver so append always allocates a fresh backing array; without the cap, two expressions derived from the same parent would share spare capacity and clobber each other. Do not simplify to a plain append.

- **Flattening depends on the receiver, so don't rely on tree shape.** `AndExpression.And(AndExpression)` splices the children in (AND is associative), but `Predicate.And(AndExpression)` nests the argument as a single element. Adapters must translate recursively, not assume a flat list.

- **`And()`/`Or()` constructors return empty, non-nil slices.** A no-argument expression marshals to JSON as `[]` rather than `null`, and downstream code (data-mongo's empty-criteria handling) relies on getting a real value. Keep the `make(..., 0)`.

- **`Parse` fails OPEN.** A string missing any of its three space-separated parts, or naming an unrecognized operator, returns `Empty()` — which matches EVERYTHING. Parse input usually arrives from a URL, so if the result gates access, AND it with your own access rules; a garbage filter must not silently become "no constraint" on protected data.

- **`Operator()` falls back to `=` for unrecognized input.** `OperatorOk` reports recognition; the `Operator` wrapper discards it, so a typo silently becomes an equality comparison. Matching is case-insensitive and accepts HTML entity spellings (`&gt;`, `&le;`) so operators survive a trip through a URL or an HTML attribute.

- **Operator constants outrun adapter support.** `IN`, `NOT IN`, `IN ALL`, `EXISTS`, and the GEO operators are defined here and implemented by data-mongo, but data-mock's matcher cannot evaluate them (they never match there) — check the adapter's AGENTS.md before relying on an operator. The GEO constructors call `GeoJSON()` at build time and store the resulting map, which adapters pass through untouched.

- **`Exists` and `NotExists` share one operator.** Both emit `OperatorExists`, differing only in the boolean `Value`; there is no "NOT EXISTS" token, so adapters must branch on the value.

- **`InAll` wraps its variadic args as one `[]any` Value.** Passing an existing slice as a single argument wraps it a second time (`Value = []any{yourSlice}`); spread it with `values...` instead.
