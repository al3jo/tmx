# TMR Segment Sorter

Reads files in the [TMX format](https://www.gala-global.org/tmx-14b)
and sorts its [Translation Units](https://www.gala-global.org/tmx-14b#tu)
alphabetically by its [Segment](https://www.gala-global.org/tmx-14b#seg)
in the default language.

It will try to preserve all other properties and parameters intact.

To build run:
```bash
$ docker-compose run --rm buildbox
$ make
```

This will generate a binary at `./tmx`. To execute it run:
```bash
$ ./tmx > sorted.tmx
```

This will read the contents of `test/tm.tmx`, sort them and print them to `stdout`.

