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
	document.getElementById('text').innerText = active.text;

	// Update colors
	const body = document.getElementsByTagName('body')[0];
	body.style.backgroundColor = active.background;
	body.style.color = active.color;

	// Update logo
	document.getElementById('logo').src = active.image.replace(/\s/g, '');
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
		updateView(capturing ? config.capturing : config.idle);
	})
}
