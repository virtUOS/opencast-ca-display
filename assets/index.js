var config = null;
updateCalendar();
var is_active = null;
var calendar = [];

/**
 * Load configuration and initialize timer
 */
fetch("/config")
	.then((response) => response.json())
	.then((cfg) => {
		config = cfg;
		console.info("config", config);
		setInterval(updateTimer, 1000);
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

	document.getElementById("info").innerText = parseCalendar(active);
}

/**
 * Check for capture agent status
 */
function updateTimer() {
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
}

function parseCalendar(active) {
	let time_remaining = 0;
	console.debug("Is Active? ", is_active);
	const event_time = is_active ? calendar[0].end : calendar[0].start;

	now = Date.now();
	if (calendar.length > 0) {
		// TODO Maybe switch 'is_active' to 'capturing' here?
		//t = is_active ? calendar[0].End : calendar[0].Start;
		time_remaining = event_time - now;
		console.debug("Time Remaining: ", time_remaining, event_time, now);
	} else {
		console.debug("Calendar is empty");
		if (!is_active) {
			return active.none;
		}
	}

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
