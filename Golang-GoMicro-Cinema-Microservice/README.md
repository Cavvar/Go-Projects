# Movie Theater Management - Microservices
Die gesamte Anwendung kann mittels Docker-Compose gestartet werden.
```
docker-compose -f docker-compose up
```
Mit Docker-Compose werden alle Services gestartet inklusive dem Client.

Der Client beinhaltet zwei Szenarien:
- Ein Kinosaal in dem Vorstellungen geplant sind, für die es Reservierung gibt, wird gelöscht
- Zwei Benutzer haben überlappende Reservierungsanfragen, die zunächst jede für sich möglich sind, und beide wollen die Reservierung durchführen
(aber nur einer gewinnt;-))

Szenario 1 wird wie folgt belegt:
- Existiert der Kinosaal?
- Löschen des Kinosaals
- Existieren Vorstellungen im Kinosaal?
- Existiert der Kinosaal?

Szenario 2 wird wie folgt belegt:
- Beantrage Reservierung von n-Sitzen in Vorstellung X
- Beantrage Reservierung von m-Sitzen in Vorstellung X

Eine Reservierung ohne dessen Bestätigung belegt bereits die Anzahl (n, m) an Sitze. Wenn aber nicht ausreichend Sitze vorhanden sind,
dann schlägt bereits die Beantragung fehl.

![Protocols](PROTOCOL.md)