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
 * Check for capture agent status
 */
function updateTimer() {
	fetch('/status')
	.then(response => response.json())
	.then(capturing => {
		console.debug('capturing', capturing)
		const active = capturing ? config.capturing : config.idle;
		document.getElementById('text').innerText = active.text;
		const body = document.getElementsByTagName('body')[0];
		body.style.backgroundColor = active.background;
		body.style.color = active.color;
	})
}
