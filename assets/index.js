var config = null;
var calendar = updateCalendar();
var is_active = null;

/**
 * [ ] Task A: Updating the textbox 
 * 		[ ] A1: Fetch the calendar from the correct Endpoint
 * 		[ ] A2: Parse the calendar according to the capturing status (either start or end date)
 *  	[ ] A3: Set the correct time remaining until end/next
 * [ ] Task B: Retrieving the authentication details, the agent and the url --> Backend
 * 	 	[ ] B1: read & parse yaml (?) Or is there a better way to do things? 
 * 		[ ] B2: make a request with fetch 
 * [x] Task C: Update the Calendar every two minutes (or something like that), not every 2 seconds
 * [ ] Task D: 
 * 
 * 
 * Kommentare von Lars:
 * 	- Falls in der nächsten (Viertel-)Stunde startet, so und so viel Minuten sonst einfach mit Datum
 *  - YAML im Backend auslesen und Opencast anfragen (alle 5 Minuten oder so)
 *  - Aus boolean ein json objekt machen ? Sonst zweiten Endpunkt machen
 *  - Statusfeld etwas nach oben verschieben, damit Display besser aussieht --> bessere Verteilung
 */

/**
 * Load configuration and initialize timer
 */
fetch('/config')
.then(response => response.json())
.then(cfg => {
	config = cfg;
	console.info('config', config)
	setInterval(updateTimer, 1000);
	setInterval(updateCalendar, 60000);
})

/**
 * Update the view
 */
function updateView(active) {
	// Update text
	document.getElementById('text').innerText = active.text;

	// Update colors
	const body = document.getElementsByTagName('body')[0];
	body.style.backgroundColor = active.background;
	body.style.color = active.color;

	// Update logo
	document.getElementById('logo').src = active.image.replace(/\s/g, '');

	document.getElementById('info').innerText = parseCalendar(active);
}


/**
 * Check for capture agent status
 */
function updateTimer() {
	fetch('/status')
	.then(response => {
		if (!response.ok) {
			const active = config.unknown;
			active.text = response.statusText;
			updateView(active);
			throw Error(response.statusText);
		}
		return response.json()
	}).then(capturing => {
		console.log('capturing', capturing);
		// the second condition is not used for debugging
		if(is_active != capturing && !capturing){
			updateCalendar()
		}
		is_active = capturing;
		updateView(capturing ? config.capturing : config.idle);
	})
}

function parseCalendar(active){
	// Do we want 'Startet/Endet in' or 'Startet/Endet um'?
	let diff = 0;
	let t = 0;

	console.debug('Lenght ', calendar.length);

	now = Date.now();
	if (calendar.length > 0){
		t = is_active ? calendar[0].End : calendar[0].Start;
		diff = t - now;
		console.debug('Diff ', diff, t, now);
	} else {
		console.debug('Calendar is empty');
	}

	hours = Math.floor(diff / (1000 * 60 * 60));
	minutes = Math.floor((diff / (1000 * 60)) % 60);
	seconds = Math.floor((diff / 1000) % 60);

	hours = (hours < 10) ? '0' + hours : hours;
	minutes = (minutes < 10) ? '0' + minutes : minutes;
	seconds = (seconds < 10) ? '0' + seconds : seconds;

	console.debug('Remaining ', diff/1000);
	return calendar.length == 0 && !is_active ? 'Keine Aufzeichnung geplant' : active.info + ' ' + hours + ':' + minutes + ':' + seconds;
}

function updateCalendar() {
	fetch("/calendar")
	.then(response => {
		console.debug('Status ', response.status)
		return response.json()})
	.then(json => {
		console.log('Calendar ', json);
		calendar = json;
	});
}
