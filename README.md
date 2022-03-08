# discord_rpc
Discord rich presence made easy

## Customization
1. Fork/Clone the repo
2. [Make a new discord application](https://discord.com/developers/applications)
3. Give the application a name, this will be what the "game" you are "Playing" will be called
4. Scroll down until you see the "Application ID" and click the copy button, for later use
5. Go to the "Rich Presence" tab and scroll down to the "Rich Presence Assets" section
6. Upload up to 2 images you want to use for your presence and name them (must be alphanumerical)
7. Run `go mod tidy`
8. Run `go build` (or `go run rpc.go` if you dont want to build)
9. Every time you want the presence, run `rpc` (or `go run rpc.go` if you didnt built it)

## Configuration
The config file is `config.json`.

### Format
The config file is a JSON file, so contains JSON data. The fields are as follows:

> `clientid`: The ID of your discord application

* Needed
* Type: `string`

(See [Customization](#customization) step 4)
<br>

> `state`: Text displayed at the bottom of your rich presence.

* Optional
* Type: `string`
<br>

> `details`: Text displayed on your rich presence.

* Optional
* Type: `string`
<br>

> `largeid`: The name you gave to the image you uploaded to the discord application, that you want to be the larger image.

* Optional
* Type: `string`

(See [Customization](#customization) step 6)
<br>

> `largeimgtext`: The text you want to be displayed when hovering over the larger image.
<br>

* Optional, even if you have a larger image
* Type: `string`
<br>

> `smallid`: The name you gave to the image you uploaded to the discord application, that you want to be the smaller image.

* Optional
* Type: `string`

(See [Customization](#customization) step 6)
<br>

> `smallimgtext`: The text you want to be displayed when hovering over the smaller image.

* Optional, even if you have a smaller image
* Type: `string`
<br>

> `buttons`: The amount of buttons you want to have on your presence.

* Optional
* Type: `int`
* Max value: `2`
<br>

> `button1`: The text of the first button.

* Needed if `buttons` is greater than `0`
* Type: `string`
<br>

> `button1url`: The URL of the first button.

* Needed if `buttons` is greater than `0`
* Type: `string`
<br>

> `button2`: The text of the second button.

* Needed if `buttons` is `2`
* Type: `string`
<br>

> `button2url`: The URL of the second button.

* Needed if `buttons` is `2`
* Type: `string`

### Example
```json
{
    "clientid": "950420436020240415",
    "state": "Powered by rich-go",
    "details": "A very difficult game",
    "largeid": "pfpimg",
    "largeimgtext": "pfp",
    "smallid": "verifiedimg",
    "smallimgtext": "verified",
    "buttons": 2,
    "button1": "My Website",
    "button1url": "https://j1233.minetest.land/",
    "button2": "My GitHub",
    "button2url": "https://github.com/Minetest-j45"
}
```
<br>

Powered by [rich-go](https://github.com/hugolgst/rich-go)