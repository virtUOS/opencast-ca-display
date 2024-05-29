var config = null;

/**
 * Load configuration and initialize timer
 */
fetch('/config')
.then(response => response.json())
.then(cfg => {
	config = cfg;
	console.info('config', config)
	setInterval(updateTimer, 2000);
})

/**
 * Update the view
 */
function updateView(active) {
	// Update text
	document.getElementById('text').innerText = active ? config.capturing.text : config.idle.text;

	// Update colors
	const body = document.getElementsByTagName('body')[0];
	body.style.backgroundColor = active ? config.capturing.background : config.idle.background;
	body.style.color = active? config.capturing.color : config.idle.color;

	// Update logo
	document.getElementById('logo').src = active ? config.capturing.image.replace(/\s/g, '') : config.idle.image.replace(/\s/g, '');

	document.getElementById('info').innerText = getAdditionalInfo(active);
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
		console.debug('capturing', capturing)
		//updateView(capturing ? config.capturing : config.idle);
		updateView(capturing);
	})
}

function formatSeconds(s){
	return (new Date(s * 1000)).toUTCString().match(/(\d\d:\d\d:\d\d)/)[0];
}

function parseCalendar(calendar, active){
	//console.log('In parseCalendar ', calendar, active);
	// Do we want 'Startet/Endet in' or 'Startet/Endet um'?
	let diff = 0;
	let t = 0;

	console.log('Active ', active);
	console.log('Lenght ', calendar.length);

	now = Date.now();
	try{
		t = active ? Date.parse(calendar[0].data.endDate) : Date.parse(calendar[0].data.startDate);
		diff = t - now;
		console.log('Diff ', diff, t, now);
	} catch (TypeError) {
		console.debug('Calendar is empty');
	}

	let remaining = null;
	console.log('Remaining ', new Date(diff/1000).toISOString());
	remaining = formatSeconds(diff/1000)
	if(calendar.length > 0){
		return (active ? config.capturing.info : config.idle.info) + ' ' + remaining;
	} else {
		return (active ? config.capturing.info + ' ' + remaining : 'Keine Aufzeichnung geplant');
	}
}

function getAdditionalInfo(active) {
	// TODO find a better cutoff (24h ?)
		/**
		 * [ ] Task A: Updating the textbox 
		 * 		[x] A1: Fetch the calendar
		 * 		[x] A2: Parse the calendar according to the capturing status (either start or end date)
		 *  	[ ] A3: Set the correct time remaining until end/next 
		 * 			--> Check for capturing state, not changing for some reason
		 * [ ] Task B: Retrieving the authentication details, the agent and the url
		 * 	 	[ ] B1: read & parse yaml
		 * 		[ ] B2: make a request with fetch 
		 */

	time = Date.now();
	cutoff = time + 86400000; // Cutoff is set to 24h from now
	
	user = 'admin';
	pw = 'opencast';
	url = 'https://develop.opencast.org/recordings/calendar.json';
	agent = 'test';

	calendar = null;

	let headers = new Headers();

	headers.set('Authorization', 'Basic ' + btoa(user + ":" + pw));

	url = url + '?' + new URLSearchParams({
		agentid: agent,
		cutoff: cutoff,
	});

	console.log('Request URL ', url);

	fetch(url, {method:'GET',
		headers:headers,
	})
	.then(response => {
		console.debug('Status ', response.status)
		return response.json()})
	.then(json => {
		console.debug('Calendar ', json);
		calendar = json;
		ans = parseCalendar(calendar, active);
		console.debug('Ans ', ans);
		return ans;
	});
	console.debug('Answer ', ans);
	return String(ans);
}
