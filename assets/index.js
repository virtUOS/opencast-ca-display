var config = null;
var is_active = null;
var calendar = [];
var network_status = {
    interfaces: [],
    connected: false,
    hostename: ""
};

/**
 * Load configuration and initialize timer
 */
fetch("/config")
    .then((response) => response.json())
    .then((cfg) => {
        config = cfg;
        console.info("config", config);
        updateNetworkStatus();
        setInterval(updateTimer, 1000);
        setInterval(updateNetworkStatus, 1000);
        updateCalendar();
        setInterval(updateCalendar, 60000);
    });

/**
 * Update the view
 */
function updateView(active) {
    // Update text
    document.getElementById("text").innerText = active.text;

    // Update colors
    const body = document.getElementsByTagName("body")[0];
    body.style.backgroundColor = active.background;
    body.style.color = active.color;

    // Update logo
    document.getElementById("logo").src = active.image.replace(/\s/g, "");
    if (network_status.connected) {
        document.getElementById("info").innerText = parseCalendar(active);
    }
}

/**
 * Check for capture agent status
 */
function updateTimer() {
    if (network_status.connected == true) {
        fetch("/status")
            .then((response) => {
                if (!response.ok) {
                    const active = config.unknown;
                    active.text = response.statusText;
                    updateView(active);
                    throw Error(response.statusText);
                }
                return response.json();
            })
            .then((capturing) => {
                console.debug("capturing", capturing);
                is_active = capturing;
                updateView(capturing ? config.capturing : config.idle);
            });
    } else {
        console.log("Network status is not connected");
        const active = config.unknown;
        active.text = "No network connection";
        updateView(active);
    }
}

function parseCalendar(active) {
    console.debug("Is Active? ", is_active);
    if (calendar.length == 0) {
        console.debug("Calendar is empty");
        if (!is_active) {
            return active.none;
        }
    }
    let time_remaining = 0;

    const event_time = is_active ? calendar[0].end : calendar[0].start;

    now = Date.now();
    // TODO Maybe switch 'is_active' to 'capturing' here?
    //t = is_active ? calendar[0].End : calendar[0].Start;
    time_remaining = event_time - now;
    console.debug("Time Remaining: ", time_remaining, event_time, now);

    hours =
        time_remaining > 0 ? Math.floor(time_remaining / (1000 * 60 * 60)) : 0;
    minutes =
        time_remaining > 0 ? Math.floor((time_remaining / (1000 * 60)) % 60) : 0;
    seconds = time_remaining > 0 ? Math.floor((time_remaining / 1000) % 60) : 0;

    hours = hours < 10 ? "0" + hours : hours;
    minutes = minutes < 10 ? "0" + minutes : minutes;
    seconds = seconds < 10 ? "0" + seconds : seconds;

    time_remaining = `${hours}:${minutes}:${seconds}`;
    console.debug("Compare ", is_active, now, event_time);
    return is_active && now < calendar[0].Start
        ? ""
        : active.info + " " + time_remaining;
}

function updateCalendar() {
    if (network_status.connected) {
        fetch("/calendar")
            .then((response) => {
                console.debug("Calendar fetched; Status ", response.status);
                return response.json();
            })
            .then((json) => {
                console.log("Calendar ", json);
                calendar = json;
            });
    }
}

function updateNetworkStatus() {
    fetch("/network_info")
        .then((response) => {
            return response.json();
        })
        .then((json) => {
            network_status = json;
            if (!network_status.connected) {
                console.log("No network");
                const active = config.unknown;
                active.text = "No network connection";
                updateView(active);
            }
        });
}

document.addEventListener('DOMContentLoaded', function () {
    // Netzwerk-Button und Popup initialisieren
    const networkButton = document.getElementById('network-button');
    const networkPopup = document.getElementById('network-popup');
    const closePopup = document.getElementById('close-popup');

    // Event-Listener für den versteckten Button
    networkButton.addEventListener('click', function () {
        fetchNetworkInfo();
        networkPopup.style.display = 'block';
    });

    // Event-Listener für den Schließen-Button
    closePopup.addEventListener('click', function () {
        networkPopup.style.display = 'none';
    });

    // Popup schließen, wenn außerhalb geklickt wird
    window.addEventListener('click', function (event) {
        if (event.target === networkPopup) {
            networkPopup.style.display = 'none';
        }
    });

    // Netzwerkinformationen abrufen
    function fetchNetworkInfo() {
        fetch('/network_info')
            .then(response => response.json())
            .then(data => {
                updateNetworkInfo(data);
            })
            .catch(error => {
                console.error('Fehler beim Abrufen der Netzwerkinformationen:', error);
                document.getElementById('interfaces-list').innerHTML = 'Fehler beim Laden der Netzwerkinformationen.';
            });
    }

    // Netzwerkinformationen aktualisieren
    function updateNetworkInfo(data) {
        // Verbindungsstatus anzeigen
        const connectionStatus = document.getElementById('connection-status');
        connectionStatus.textContent = data.connected ? 'Verbunden' : 'Nicht verbunden';
        connectionStatus.className = data.connected ? 'connected' : 'disconnected';

        // Hostname anzeigen
        const hostnameElement = document.getElementById('hostname');
        if (data.hostname) {
            hostnameElement.textContent = data.hostname;
            hostnameElement.parentElement.style.display = 'block';
        } else {
            hostnameElement.parentElement.style.display = 'none';
        }

        // Netzwerkschnittstellen anzeigen
        const interfacesList = document.getElementById('interfaces-list');
        interfacesList.innerHTML = '';

        if (data.interfaces && data.interfaces.length > 0) {
            data.interfaces.forEach(net_interface => {
                const interfaceDiv = document.createElement('div');
                console.log(net_interface);
                interfaceDiv.className = 'interface-item';

                let interfaceHtml = `<div class="interface-name">${net_interface.name}</div>`;

                if (net_interface.mac && net_interface.mac !== "") {
                    interfaceHtml += `<div>MAC: ${net_interface.mac}</div>`;
                }

                if (net_interface.addr && net_interface.addr.length > 0) {
                    console.log(net_interface.addr);
                    console.log(net_interface.addr.join(', '));
                    interfaceHtml += `<div>IP-Adressen: ${net_interface.addr.join(', ')}</div>`;
                }

                if (net_interface.flags) {
                    interfaceHtml += `<div>Flags: ${net_interface.flags}</div>`;
                }

                interfaceDiv.innerHTML = interfaceHtml;
                interfacesList.appendChild(interfaceDiv);
            });
        } else {
            interfacesList.innerHTML = 'Keine Netzwerkschnittstellen gefunden.';
        }
    }
});