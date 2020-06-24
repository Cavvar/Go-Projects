# Kommunikation der Services
Die Services untereinander kommunizieren über Clients zu dem jeweiligen Service.
Da z.B. bei der Nicht-Verfügbarkeit eines Kinosaals, die in dem Kinosaal stattfindenen Vorstellungen gelöscht werden müssen, wird im Handler des `room-services` der Client zum `showing-service` initialisiert.

Die andere Möglichkeit wäre die Kommunikation untereinander mittels `Publish & Subscribe` zu realisieren.
