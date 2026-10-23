# 2026-10-23 - Starter Bootstrap

## Decisions

1. The starter is Go-native: `go.mod`, `Makefile`, `src/`;
   no JVM build system is used.
2. The reference starter's useful separation is retained: backend entry,
   implementation tests, plugin manifest/resources, optional UI, CI, and local
   agent knowledge.
3. The starter depends on `modality-gopn-sdk-go/service`; local development uses
   a module replacement until a public SDK release exists.
4. The initial service is a gRPC health service, avoiding a fictional business
   proto while proving lifecycle serving.
5. UI route declarations use the modality static host contract rather than Halo
   runtime packages.
