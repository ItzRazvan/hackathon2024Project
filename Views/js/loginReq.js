document.getElementById("loginForm").addEventListener("submit", (event) => {
	event.preventDefault();

	fetch("/login", {
		method: "POST",
		credentials: "include",
		body: new FormData(document.getElementById("loginForm")),
	}).then((response) => {
		if (response.ok) {
			window.location.href = response.url;
		} else {
			document.getElementById("loginError").innerText =
				"Email sau parola incorecta";
		}
	});
});
