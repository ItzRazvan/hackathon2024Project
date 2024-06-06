document.getElementById("signupForm").addEventListener("submit", () => {
	event.preventDefault();

	fetch("/signup", {
		method: "POST",
		credentials: "include",
		body: new FormData(document.getElementById("signupForm")),
	}).then((response) => {
		if (response.ok) {
			document.getElementById("name").value = "";
			document.getElementById("email").value = "";
			document.getElementById("password").value = "";
			document.getElementById("signupSuccess").innerText =
				"Cont creat cu succes!";
		} else {
			document.getElementById("signupError").innerText =
				"Sintaxa emailului este gresita sau parola e pre scurta";
		}
	});
});

document.getElementById("goBack").addEventListener("click", () => {
	window.location.href = "/";
});
