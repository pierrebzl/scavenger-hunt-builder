# scavenger-hunt-builder

This app is meant to ease the process of creating homemade scavenger hunts.

## Basic workflow:
- Set game environment parameters:
    - Players:
        - Number and names of the participants
        - Number, names and color/logo of the teams
    - Place:
        - List of the rooms/main location (ideally number of room > number of teams)

- Add clues to find spots in each of the rooms/places
    - `room A`
        - `spot 1`
            - `clue 1`
                - Short name
                - The clue (could be text, rebus, picture) the player will need to find `spot 1` in `room A`
                - Picture of `spot 1` (optional)
                - Relative difficulty level (from 1 to 10)

- Compute optimal `spot` path for each team
    - Avoid teams to be in the same `room` at the same time
    - Use difficulty levels

- Generate usable outputs to setup the game IRL
    - Printable and cuttable layout of the clues to be hidden at each of the `spot`
        - Each `clue ticket` will show:
            - The color/logo of which team the clue is meant for
            - The rank of the clue relative to the team is meant for (ex: spot 1 will be the 3rd spot for team A but the 1st spot to visit for team B)
            - The clue to get the team on the path of the next spot
            - The id/location of where this `clue ticket` should be hidden during the game setup

    - A table / story sum up of each team hunt (spot order, etc.)


## Dev
`curl -H "Content-Type: application/json" -X POST -d '{"name":"room 1"}' http://localhost:8080/room/ | jq`

`curl -H "Content-Type: application/json" -X GET http://localhost:8080/room/2 | jq`

`curl -H "Content-Type: application/json" -X PUT -d '{"name":"new room name"}' http://localhost:8080/room/2`

`curl -H "Content-Type: application/json" -X POST -d '{"name":"clue 1", "clue":"indice 1"}' http://localhost:8080/room/2/spot/ | jq`

`curl -H "Content-Type: application/json" -X PUT -d '{"name":"test-new-clue", "clue": "changed"}' http://localhost:8080/spot/18`
