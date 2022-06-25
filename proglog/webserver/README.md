# ProgLog

Go example commit log service.

A commit log is a sequence of records where each record has its own unique identifier. Records can only be appended at the end of a commit log. Records are immutable once written.

Reading a commit log happens from left to right. There's no querying so offsets are used to specify the start and end points of the read.

E.g. 
Kafka's individual records are identified using <topic, partition, offset>.

![Example Kafka Topic][ExampleKafkaTopic]

Commit logs provide a source of truth for what occurred in a system and in what order.

[ExampleKafkaTopic]: https://quarkus.io/assets/images/posts/kafka-commit-strategies/topics-partitions.png