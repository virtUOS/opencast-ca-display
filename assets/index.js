const myInterval = setInterval(myTimer, 2000);

function myTimer() {
	fetch('/status')
	.then(response => response.json())
	.then(capturing => {
		console.info(capturing)
		document.getElementById('text').innerText = capturing ? 'Aufzeichnung läuft' : 'Keine Aufzeichnung';
		const body = document.getElementsByTagName('body')[0];
		if (capturing) {
			body.style.backgroundColor = '#ac0634';
			body.style.color = '#fff';
		} else {
			body.style.backgroundColor = '#fff';
			body.style.color = '#000';
		}
	})
}
