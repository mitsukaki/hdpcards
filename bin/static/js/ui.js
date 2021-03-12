function scrape() {
    // get the request parameters
    let mapname = document.getElementById('mapname').value;
    let playerstring = document.getElementById('players').value;

    // get the player list
    let players = playerstring.replace(/\s/g, '').split(",");

    // make the request
    fetch("http://localhost:8883/api/scrape", {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "mapName": mapname,
            "players": players
        })
    });
}