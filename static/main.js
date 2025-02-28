function doSomething() {
	console.log("Doing something")
}

function createTeamButton() {
	const teamForm = document.getElementById("createTeamForm")
	formVis = teamForm.style.display
	if(formVis === 'none' || formVis === '') {
		teamForm.style.display = "block"
	} else {
		teamForm.style.display = "none"
	}
}
