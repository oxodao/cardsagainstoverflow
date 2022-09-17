# Cards Against Overflow

This is the game Cards Against Humanity but you can play it in your browser.

It's a bit prettier than the existing one, although they are functional too (looking at you Pretend you're xyzzy :p)

It's a rewrite of a private one that was used during the lockdown in France in 2020 to be a bit better (ditching Vue.JS, cleaner Go, ...).

## Features

- [ ] Playing Cards Against Humanity
- [X] Importing decks
- [X] Managing decks
- [ ] Replay game (Watching it back)
- [ ] Display mode (Just showing the game, if you are playing in the same room for example)
- [ ] Admin tools (Room creator)
- [ ] Re-joining a game (Lost connection, changing device mid-game, ...)
- [ ] MSN Messenger's / WLM's Wizz

Technical:
- [ ] Auto-creating the database
- [ ] Auto-applying migrations

## Setup & Usage

You need to have a postgresql server initialized with the `sql/init.sql`.

Create a `config.yml` (current folder) or a `/etc/config.yml` containing the basic setup:
```yaml
server:
  # Database configuration
  database:
    host: 127.0.0.1
    port: 5432
    username: cao
    password: cao
    database: cao
  
  # Web-server configuration
  web:
    host: 127.0.0.1
    port: 2834

client:
  enable_wizz: true # Whether players can use the "wizz" from MSN
```

Then setup your reverse-proxy (nginx example):
```
server {
        listen      80;
        server_name game.example.tld;
        return 301 https://$server_name$request_uri;
}

server {
        listen 443 ssl http2;
        server_name game.example.tld;

        # Your SSL config (Lets Encrypt + https://ssl-config.mozilla.org/)
        include snippets/ssl.conf;

        location / {
                proxy_pass http://127.0.0.1:2834;
                proxy_http_version 1.1;
                proxy_set_header Upgrade $http_upgrade;
                proxy_set_header Connection "Upgrade";
                proxy_set_header Host $host;
        }
}
```

Now you can import your decks. As we are playing with the french `Blanc Manger Coco` version, we cannot share the files (Go buy the real one if you don't have it yet !). Later we'll probably provide Cards Against Humanity ones though as they are free.

In the meantime, if you are willing to commit to a painful task (Done it by hand on 1300+ cards), you can build the deckfile yourself:
```json
{
    "name": "Cards against humanity: Scary Pack",
    "author": "Cards against humanity",
    "selected_by_default": false,
    "white_cards": [
        { "text": "My white card", "is_black_card": false },
        { "text": "My white card 2", "is_black_card": false }
    ],
    "black_cards": [
        { "text": "My ____ black card", "is_black_card": true },
        { "text": "My black ____ card 2", "is_black_card": true }
    ]
}
```

The format may seems a bit bizarre but there are two reasons it's done like this:
- Re-using internal structure (lazy)
- Future upgrades (ex: double score cards) while keeping the format compatible

Finally, you can import them:
```sh
$ cao deck import [--selected-by-default=true|false] deckfile.json
```

You can manage them from there:
```sh
$ cao deck list
$ cao deck delete [id]
$ cao deck set-default [id] (true|false) # Make it selected by default in the web-ui
```

You are now ready to play ! Three players minimum, or the game won't start (That would be weird to play)

## Contributing

There will be a CLA added to this project, as I haven't took the time yet to do this properly, please don't create PR or contribute to the project yet.

Issues are still welcome of course !


## License

>    Cards Against Overflow
>    Copyright (C) 2022 - Nathan <Oxodao> JANCZEWSKI
>
>    This program is free software: you can redistribute it and/or modify
>    it under the terms of the GNU Affero General Public License as
>    published by the Free Software Foundation, either version 3 of the
>    License, or (at your option) any later version.
>
>    This program is distributed in the hope that it will be useful,
>    but WITHOUT ANY WARRANTY; without even the implied warranty of
>    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
>    GNU Affero General Public License for more details.
>
>    You should have received a copy of the GNU Affero General Public License
>    along with this program.  If not, see <https://www.gnu.org/licenses/>.

