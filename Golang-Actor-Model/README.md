## Ausführen mit Docker

-   Images bauen

    ```
    make docker
    ```

-   ein (Docker)-Netzwerk `actors` erzeugen

    ```
    docker network create actors
    ```

-   Starten des Tree-Services und binden an den Port 1860 des Containers mit dem DNS-Namen
    `treeservice` (entspricht dem Argument von `--name`) im Netzwerk `actors`:

    ```
    docker run --rm --net actors --name treeservice treeservice \
      --host="treeservice.actors:1860"
    ```

-   Starten des Tree-CLI, Binden an `treecli.actors:1338` und nutzen des Services unter
    dem Namen und Port `treeservice.actors:1860`:

    ```
    docker run --rm --net actors --name treecli treecli --bind="treecli.actors:1338" \
      --remote="treeservice.actors:1860"
    ```

    Hier sind wieder die beiden Flags `--host` und `--serviceName` beliebig gewählt und
    in der Datei `treeservice/main.go` implementiert.

-   Zum Beenden, killen Sie einfach den Tree-Service-Container mit `Ctrl-C` und löschen
    Sie das Netzwerk mit

    ```
    docker network rm actors
    ```

## Ausführen mit Docker ohne vorher die Docker-Images zu bauen

Nach einem Commit baut der Jenkins, wenn alles durch gelaufen ist, die beiden
Docker-Images. Sie können diese dann mit `docker pull` herunter laden. Schauen Sie für die
genaue Bezeichnung in die Consolenausgabe des Jenkins-Jobs.

Wenn Sie die Imagenamen oben (`treeservice` und `treecli`) durch die Namen aus der
Registry ersetzen, können Sie Ihre Lösung mit den selben Kommandos wie oben beschrieben,
ausprobieren.

## CLI-Commands TreeCli
* ```bind (string)```: 
        bind address to (default "localhost:1338")
* ```deletekey (bool)```: 
        flag for deleting a key/value in the tree
* ```deletetree (bool)```:
        flag for deleting a tree
* ```find (bool)```:
        flag for find a value in the tree
* ```id (int)```:
        flag for id. necessary for all operations
* ```insert (bool)```:
        flag for inserting a value to the tree
* ```key (int)```:
        key when inserting/deleting/finding values
* ```leafSize (int)```:
        leafSize
* ```newtree (bool)```:
        create tree. default not creating tree
* ```remote (string)```:
        remote address (default "localhost:1860")
* ```token (string)```:
        flag for token. necessary for all operations
* ```traverse (bool)```:
        flag for traversing the tree
* ```value (string)```:
        value when inserting a key
        
## CLI-Commands TreeService
* ```host (string)```: 
    flag for the address of the treeservice
* ```serviceName (string)```: 
    flag for the serviceName

## Usage without docker
- First, start the treeservice:

```
    cd treeservice && go run main.go
```

- Second, start the treecli:

```
    cd treecli && go run main.go
```
Without any of the available flags, it will print out the currently available trees!

- Show available trees: ```go run main.go```
- Create Tree: ```go run main.go --newtree --leafSize=3```
- Insert Key-Value: ```go run main.go --insert --token="AAAAA" --id=1337 --key=5 --value="golang"```
- Find: ```go run main.go --find --token="AAAAA" --id=1337 --key=5```
- Delete Key: ```go run main.go --deletekey --token="AAAAA" --id=1337 --key=5```
- Delete Tree: ```go run main.go--deletetree --token="AAAAA" --id=1337```
- Traverse: ```go run main.go --traverse --token="AAAAA" --id=1337```
