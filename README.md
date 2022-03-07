# discord_rpc
My discord rich presence

## Customization
1. Fork/Clone the repo
2. [Make a new discord application](https://discord.com/developers/applications)
3. Give the application a name, this will be what the "game" you are "Playing" will be called
4. Scroll down until you see the "Application ID" and click the copy button, for later use
5. Go to the "Rich Presence" tab and scroll down to the "Rich Presence Assets" section
6. Upload your images you want to use for your presence and name them (must be alphanumerical)
7. Change line `11` to your application ID, from step 4
8. Change lines `18` and `19` to what you want them to say (`State` will be the small one on the bottom and `Details` will be the small one under the name of the "game")
9. Chane lines `20` and `22` to the names of the images you uploaded
10. Change lines `21` and `23` to the descriptions of the images
11. Customize the buttons accordingly, if you dont want them, delete lines `27-36`
12. Change the number on line `43` to how many hours you want the presence to be there for when you run it
13. Run `go mod tidy`
14. Run `go build` (or `go run rpc.go` if you dont want to build)
15. Every time you want the presence, run `rpc` (or `go run rpc.go` if you didnt built it)